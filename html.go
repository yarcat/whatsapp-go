package whatsapp

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Text, Link string
}

type Options struct {
	CollectLinks func([]Link)
}

type OptionFn func(*Options)

func FormatLinks(links []Link) string {
	if len(links) == 0 {
		return ""
	}

	var out strings.Builder
	for i, link := range links {
		if i > 0 {
			out.WriteString("\n")
		}
		out.WriteString(link.Text)
		out.WriteString(" - ")
		out.WriteString(link.Link)
	}
	return out.String()
}

func FromHTML(text string, opts ...OptionFn) string {
	simpleMappings := map[string]string{
		"b": "*",
		"i": "_",
		"s": "~",
	}

	var options Options
	for _, opt := range opts {
		opt(&options)
	}

	tokenizer := html.NewTokenizer(strings.NewReader(text))
	var (
		out   strings.Builder
		links []Link
	)

	func() {
		var currentLink *Link
		for {
			tokenType := tokenizer.Next()
			switch tokenType {
			case html.ErrorToken:
				return
			case html.TextToken:
				text := tokenizer.Token().Data
				if currentLink != nil {
					currentLink.Text += text
				}
				out.WriteString(text)
			case html.StartTagToken, html.EndTagToken:
				token := tokenizer.Token()
				switch token.Data {
				case "a":
					if tokenType == html.StartTagToken && options.CollectLinks != nil {
						for _, attr := range token.Attr {
							if attr.Key == "href" {
								currentLink = &Link{Link: attr.Val}
								break
							}
						}
					} else if tokenType == html.EndTagToken && currentLink != nil {
						links = append(links, *currentLink)
						currentLink = nil
					}
				default:
					if mapping, exists := simpleMappings[token.Data]; exists {
						out.WriteString(mapping)
					}
				}
			}
		}
	}()

	if options.CollectLinks != nil && len(links) > 0 {
		options.CollectLinks(links)
	}

	return out.String()
}

func FromHTMLWithLinks(text string, opts ...OptionFn) string {
	var links []Link
	opts = append(opts, func(opt *Options) {
		prev := opt.CollectLinks
		opt.CollectLinks = func(v []Link) {
			links = v
			if prev != nil {
				prev(v)
			}
		}
	})
	result := strings.TrimSpace(FromHTML(text, opts...))
	if len(links) > 0 {
		result += "\n\n" + FormatLinks(links)
	}
	return result
}
