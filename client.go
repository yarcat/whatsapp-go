package whatsapp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
)

const (
	// DefaultBaseURL is the base URL for the WhatsApp Business API.
	DefaultBaseURL = "https://graph.facebook.com"

	// DefaultAPIVersion is the version of the WhatsApp Business API.
	DefaultAPIVersion = "v22.0"
)

// Client is a Client client that provides methods to interact with the Client Business API.
type Client struct {
	AccessToken   string       // AccessToken is the access token for the WhatsApp Business API.
	BaseURL       string       // BaseURL is the base URL for the WhatsApp Business API.
	APIVersion    string       // APIVersion is the version of the WhatsApp Business API.
	PhoneNumberID string       // PhoneNumberID is the ID of the phone number associated with the WhatsApp Business account.
	Client        *http.Client // Client is the HTTP client used to make requests to the WhatsApp Business API.
}

// NewClient creates a new WhatsApp API client with the provided access token and phone number ID.
func NewClient(accessToken, phoneNumberID string) *Client {
	return &Client{
		AccessToken:   accessToken,
		BaseURL:       DefaultBaseURL,
		APIVersion:    DefaultAPIVersion,
		PhoneNumberID: phoneNumberID,
		Client:        http.DefaultClient,
	}
}

// SendText sends a text message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/text-messages
func (wa *Client) SendText(ctx context.Context, recipient string, params *SendTextParams) (*MessagesResponse, error) {
	request := &Request{
		MessagingProduct: MessagingProductWhatsApp,
		RecipientType:    RecipientTypeIndividual,
		To:               recipient,
		Type:             MessageTypeText,
		Text:             params,
	}
	var response MessagesResponse
	if err := sendRequest(ctx, wa, "messages", request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SendImage sends an image message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/image-messages
func (wa *Client) SendImage(ctx context.Context, recipient string, params *SendImageParams) (*MessagesResponse, error) {
	request := &Request{
		MessagingProduct: MessagingProductWhatsApp,
		RecipientType:    RecipientTypeIndividual,
		To:               recipient,
		Type:             MessageTypeImage,
		Image:            params,
	}
	var response MessagesResponse
	if err := sendRequest(ctx, wa, "messages", request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SendInteractiveButtons sends an interactive reply buttons message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-reply-buttons-messages
func (wa *Client) SendInteractiveButtons(ctx context.Context, recipient string, params *SendInteractiveButtonsParams) (*MessagesResponse, error) {
	interactive := &Interactive{
		Type:   InteractiveTypeButton,
		Header: params.Header,
		Body:   params.Body,
		Footer: params.Footer,
		Action: &Action{
			Buttons: params.Buttons,
		},
	}

	request := &Request{
		MessagingProduct: MessagingProductWhatsApp,
		RecipientType:    RecipientTypeIndividual,
		To:               recipient,
		Type:             MessageTypeInteractive,
		Interactive:      interactive,
	}

	var response MessagesResponse
	if err := sendRequest(ctx, wa, "messages", request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SendInteractiveList sends an interactive list message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-list-messages
func (wa *Client) SendInteractiveList(ctx context.Context, recipient string, params *SendInteractiveListParams) (*MessagesResponse, error) {
	interactive := &Interactive{
		Type:   InteractiveTypeList,
		Header: params.Header,
		Body:   params.Body,
		Footer: params.Footer,
		Action: &Action{
			Button:   params.Button,
			Sections: params.Sections,
		},
	}

	request := &Request{
		MessagingProduct: MessagingProductWhatsApp,
		RecipientType:    RecipientTypeIndividual,
		To:               recipient,
		Type:             MessageTypeInteractive,
		Interactive:      interactive,
	}

	var response MessagesResponse
	if err := sendRequest(ctx, wa, "messages", request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SendInteractiveFlow sends an interactive flow message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
func (wa *Client) SendInteractiveFlow(ctx context.Context, recipient string, params *SendInteractiveFlowParams) (*MessagesResponse, error) {
	action := &Action{
		Name:       "flow",
		Parameters: params.FlowParameters,
	}

	// Validate the action for type safety
	if err := ValidateAction(action); err != nil {
		return nil, fmt.Errorf("invalid flow action: %w", err)
	}

	interactive := &Interactive{
		Type:   InteractiveTypeFlow,
		Header: params.Header,
		Body:   params.Body,
		Footer: params.Footer,
		Action: action,
	}

	request := &Request{
		MessagingProduct: MessagingProductWhatsApp,
		RecipientType:    RecipientTypeIndividual,
		To:               recipient,
		Type:             MessageTypeInteractive,
		Interactive:      interactive,
	}

	var response MessagesResponse
	if err := sendRequest(ctx, wa, "messages", request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SendInteractiveCTAURL sends an interactive call-to-action URL message.
// This allows you to map any URL to a button so you don't have to include the raw URL in the message body.
//
// Example usage:
//
//	params := &SendInteractiveCTAURLParams{
//	    Header: &Header{
//	        Type: HeaderTypeText,
//	        Text: "Special Offer!",
//	    },
//	    Body: &Body{
//	        Text: "Tap the button below to view our latest deals.",
//	    },
//	    Footer: &Footer{
//	        Text: "Limited time offer",
//	    },
//	    DisplayText: "View Deals",
//	    URL: "https://example.com/deals?ref=whatsapp",
//	}
//
//	response, err := client.SendInteractiveCTAURL(ctx, "1234567890", params)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-cta-url-messages
func (wa *Client) SendInteractiveCTAURL(ctx context.Context, recipient string, params *SendInteractiveCTAURLParams) (*MessagesResponse, error) {
	action := &Action{
		Name: "cta_url",
		Parameters: &CTAURLParameters{
			DisplayText: params.DisplayText,
			URL:         params.URL,
		},
	}

	// Validate the action for type safety
	if err := ValidateAction(action); err != nil {
		return nil, fmt.Errorf("invalid CTA URL action: %w", err)
	}

	interactive := &Interactive{
		Type:   InteractiveTypeCTAURL,
		Header: params.Header,
		Body:   params.Body,
		Footer: params.Footer,
		Action: action,
	}

	request := &Request{
		MessagingProduct: MessagingProductWhatsApp,
		RecipientType:    RecipientTypeIndividual,
		To:               recipient,
		Type:             MessageTypeInteractive,
		Interactive:      interactive,
	}

	var response MessagesResponse
	if err := sendRequest(ctx, wa, "messages", request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetMedia retrieves media information including the download URL for a given media ID.
// The URL returned is valid for 5 minutes and can be used to download the media file.
//
// Example usage:
//
//	mediaInfo, err := client.GetMedia(ctx, "media_id_from_webhook")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Validate media size before downloading
//	if err := ValidateMediaSize(mediaInfo.MimeType, mediaInfo.FileSize); err != nil {
//	    log.Printf("Media validation failed: %v", err)
//	    return
//	}
//
//	// Download the media content (streaming)
//	reader, err := client.DownloadMedia(ctx, mediaInfo.URL)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer reader.Close()
//
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#retrieve-media-url
func (wa *Client) GetMedia(ctx context.Context, mediaID string) (*MediaResponse, error) {
	var response MediaResponse
	if err := sendGetRequest(ctx, wa, mediaID, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DownloadMedia downloads the actual media content using the URL obtained from GetMedia.
// The URL is only valid for 5 minutes after retrieval from GetMedia.
//
// Note: This method returns an io.ReadCloser for streaming the media content.
// The caller is responsible for closing the returned ReadCloser.
//
// Example usage:
//
//	reader, err := client.DownloadMedia(ctx, mediaInfo.URL)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer reader.Close()
//
//	// Save to file
//	file, err := os.Create("downloaded_media.jpg")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer file.Close()
//
//	_, err = io.Copy(file, reader)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#download-media
func (wa *Client) DownloadMedia(ctx context.Context, mediaURL string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, mediaURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+wa.AccessToken)

	resp, err := wa.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() // Close body since we're returning an error
		return nil, fmt.Errorf("failed to download media: want 200 OK, got %s", resp.Status)
	}

	// Return the response body as ReadCloser - caller must close it
	return resp.Body, nil
}

// DownloadMediaBytes downloads the actual media content and reads it into memory.
// This is a convenience method for cases where you need all the content in memory.
// For large files or streaming use cases, prefer DownloadMedia which returns io.ReadCloser.
//
// Example usage:
//
//	content, err := client.DownloadMediaBytes(ctx, mediaInfo.URL)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Save to file
//	err = os.WriteFile("downloaded_media.jpg", content, 0644)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#download-media
func (wa *Client) DownloadMediaBytes(ctx context.Context, mediaURL string) ([]byte, error) {
	reader, err := wa.DownloadMedia(ctx, mediaURL)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	// Read the entire content into memory
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read media content: %w", err)
	}

	return content, nil
}

// GetAndDownloadMedia is a convenience method that retrieves media information
// and downloads the content in a single call. This is useful when you need both
// the metadata and the actual media content.
//
// Note: The caller is responsible for closing the returned ReadCloser.
//
// Example usage:
//
//	mediaInfo, reader, err := client.GetAndDownloadMedia(ctx, mediaID)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer reader.Close()
//
//	// Process the media stream...
//	_, err = io.Copy(destination, reader)
//
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media
func (wa *Client) GetAndDownloadMedia(ctx context.Context, mediaID string) (*MediaResponse, io.ReadCloser, error) {
	// First, get the media information including the download URL
	mediaInfo, err := wa.GetMedia(ctx, mediaID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get media info: %w", err)
	}

	// Then download the actual media content
	content, err := wa.DownloadMedia(ctx, mediaInfo.URL)
	if err != nil {
		return mediaInfo, nil, fmt.Errorf("failed to download media: %w", err)
	}

	return mediaInfo, content, nil
}

// GetAndDownloadMediaBytes retrieves media information and downloads the media content into memory.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media
func (wa *Client) GetAndDownloadMediaBytes(ctx context.Context, mediaID string) (*MediaResponse, []byte, error) {
	// First, get the media information including the download URL
	mediaInfo, err := wa.GetMedia(ctx, mediaID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get media info: %w", err)
	}

	// Then download the actual media content into memory
	content, err := wa.DownloadMediaBytes(ctx, mediaInfo.URL)
	if err != nil {
		return mediaInfo, nil, fmt.Errorf("failed to download media: %w", err)
	}

	return mediaInfo, content, nil
}

// UploadMedia uploads media to WhatsApp and returns the media ID that can be used in messages.
// The media file is uploaded as multipart form data with the specified MIME type.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#upload-media
func (wa *Client) UploadMedia(ctx context.Context, params *UploadMediaParams) (*UploadMediaResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("invalid upload parameters: %w", err)
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, params.Filename))
	h.Set("Content-Type", params.MimeType)

	part, err := writer.CreatePart(h)
	if err != nil {
		return nil, fmt.Errorf("creating multipart part: %w", err)
	}
	if _, err := io.Copy(part, params.File); err != nil {
		return nil, fmt.Errorf("copying file data: %w", err)
	}

	if err := errors.Join(
		writer.WriteField("messaging_product", string(params.MessagingProduct)),
		writer.WriteField("type", params.MimeType),
		writer.Close(),
	); err != nil {
		return nil, fmt.Errorf("setting up multipart writer: %w", err)
	}

	u, err := url.JoinPath(wa.BaseURL, wa.APIVersion, wa.PhoneNumberID, "media")
	if err != nil {
		return nil, fmt.Errorf("build URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, &body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+wa.AccessToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := wa.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiError APIError
		if decodeErr := json.NewDecoder(resp.Body).Decode(&apiError); decodeErr != nil {
			return nil, fmt.Errorf("upload status %s", resp.Status)
		}
		return nil, fmt.Errorf("WhatsApp API error: %s (code: %d)", apiError.Error.Message, apiError.Error.Code)
	}

	var response UploadMediaResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return &response, nil
}

// UploadMediaFromFile is a convenience method that uploads media from a file path.
// This method automatically opens the file, detects basic MIME types, and uploads the media.
// For more control over the upload process, use UploadMedia directly.
//
// Example usage:
//
//	response, err := client.UploadMediaFromFile(ctx, "/path/to/image.jpg", "image/jpeg")
//	if err != nil {
//	    log.Printf("Failed to upload media: %v", err)
//	    return
//	}
//	// Use response.ID in SendImageParams or other message types
//
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#upload-media
func (wa *Client) UploadMediaFromFile(ctx context.Context, filePath, mimeType string) (*UploadMediaResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Extract filename from the path
	filename := filepath.Base(filePath)

	params, err := NewUploadMediaParams(file, filename, mimeType)
	if err != nil {
		return nil, fmt.Errorf("failed to create upload params: %w", err)
	}

	return wa.UploadMedia(ctx, params)
}

// DeleteMedia deletes media from WhatsApp servers.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#delete-media
func (wa *Client) DeleteMedia(ctx context.Context, mediaID string) (*DeleteMediaResponse, error) {
	if mediaID == "" {
		return nil, fmt.Errorf("media ID cannot be empty")
	}

	u, err := url.JoinPath(wa.BaseURL, wa.APIVersion, mediaID)
	if err != nil {
		return nil, fmt.Errorf("failed to build URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+wa.AccessToken)

	resp, err := wa.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiError APIError
		if decodeErr := json.NewDecoder(resp.Body).Decode(&apiError); decodeErr != nil {
			return nil, fmt.Errorf("delete failed with status %s", resp.Status)
		}
		return nil, fmt.Errorf("WhatsApp API error: %s (code: %d)", apiError.Error.Message, apiError.Error.Code)
	}

	var response DeleteMediaResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}

func sendRequest(ctx context.Context, wa *Client, endpoint string, request any, response any) error {
	u, err1 := url.JoinPath(wa.BaseURL, wa.APIVersion, wa.PhoneNumberID, endpoint)
	payloadBytes, err2 := json.Marshal(request)
	req, err3 := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewBuffer(payloadBytes))
	if err := errors.Join(err1, err2, err3); err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+wa.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := wa.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiError APIError
		if decodeErr := json.NewDecoder(resp.Body).Decode(&apiError); decodeErr != nil {
			return fmt.Errorf("want 200 OK, got %s", resp.Status)
		}
		return fmt.Errorf("WhatsApp API error: %s (code: %d)", apiError.Error.Message, apiError.Error.Code)
	}

	return json.NewDecoder(resp.Body).Decode(response)
}

func sendGetRequest(ctx context.Context, wa *Client, mediaID string, response any) error {
	u, err := url.JoinPath(wa.BaseURL, wa.APIVersion, mediaID)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+wa.AccessToken)

	resp, err := wa.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var mediaError MediaError
		if decodeErr := json.NewDecoder(resp.Body).Decode(&mediaError); decodeErr != nil {
			return fmt.Errorf("want 200 OK, got %s", resp.Status)
		}
		return fmt.Errorf("media API error: %s (code: %d)", mediaError.Error.Message, mediaError.Error.Code)
	}

	return json.NewDecoder(resp.Body).Decode(response)
}
