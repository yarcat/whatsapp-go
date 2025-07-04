package whatsapp

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"io"
	"net/http"
	"strings"
)

// WebhookHandler is an interface that defines the methods that must be implemented
// by a webhook handler. It is used to handle incoming webhook requests from the WhatsApp Business API.
type WebhookHandler interface {
	HandleWebhook(context.Context, http.ResponseWriter, *WebhookRequest)
}

// WebhookHandlerFunc is a function type that implements the WebhookHandler interface.
type WebhookHandlerFunc func(context.Context, http.ResponseWriter, *WebhookRequest)

// HandleWebhook calls the function with the given parameters.
func (f WebhookHandlerFunc) HandleWebhook(ctx context.Context, w http.ResponseWriter, r *WebhookRequest) {
	f(ctx, w, r)
}

// WebhookErrHandler is an interface that defines the methods that must be implemented
// by a webhook error handler. It is used to handle errors that occur during the processing of webhook requests.
type WebhookErrHandler interface {
	HandleWebhookErr(context.Context, http.ResponseWriter, *WebhookRequest, error) bool
}

// WebhookErrHandlerFunc is a function type that implements the WebhookErrHandler interface.
type WebhookErrHandlerFunc func(context.Context, http.ResponseWriter, *WebhookRequest, error) bool

// HandleWebhookErr calls the function with the given parameters.
func (f WebhookErrHandlerFunc) HandleWebhookErr(ctx context.Context, w http.ResponseWriter, r *WebhookRequest, err error) bool {
	return f(ctx, w, r, err)
}

// Webhook is a struct that represents a WhatsApp webhook.
// It is used to handle incoming messages and events from the WhatsApp Business API.
// It implements the http.Handler interface to handle HTTP requests.
type Webhook struct {
	WebhookSecret string
	AppSecret     string
	Handler       WebhookHandler
	ErrHandler    WebhookErrHandler
}

// NewWebhook creates a new WhatsApp webhook with the given parameters.
func NewWebhook(webhookSecret, appSecret string, handler WebhookHandler) *Webhook {
	wh := &Webhook{
		WebhookSecret: webhookSecret,
		AppSecret:     appSecret,
		Handler:       handler,
	}
	wh.ErrHandler = wh
	return wh
}

// ServeHTTP handles incoming HTTP requests for the WhatsApp webhook.
func (wh *Webhook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		wh.verifyChallenge(w, r)
	case http.MethodPost:
		wh.handleWebhookPOST(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// HandleWebhookErr is called when an error occurs during the processing of a webhook request.
func (wh *Webhook) HandleWebhookErr(ctx context.Context, w http.ResponseWriter, r *WebhookRequest, err error) bool {
	if wh.ErrHandler != nil {
		return wh.ErrHandler.HandleWebhookErr(ctx, w, r, err)
	}
	return false
}

func (wh *Webhook) verifyChallenge(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("hub.mode")
	challenge := r.URL.Query().Get("hub.challenge")
	verifyToken := r.URL.Query().Get("hub.verify_token")

	if mode == "subscribe" && verifyToken == wh.WebhookSecret {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// verifySignature verifies the X-Hub-Signature or X-Hub-Signature-256 header
// against the request body using the webhook secret.
func (wh *Webhook) verifySignature(r *http.Request, body []byte) bool {
	if signature := r.Header.Get("X-Hub-Signature-256"); signature != "" {
		return wh.verifySignatureImpl(signature, "sha256=", body, sha256.New)
	}
	if signature := r.Header.Get("X-Hub-Signature"); signature != "" {
		return wh.verifySignatureImpl(signature, "sha1=", body, sha1.New)
	}
	return false
}

func (wh *Webhook) verifySignatureImpl(signature, prefix string, body []byte, hashFunc func() hash.Hash) bool {
	expectedSig, foundPrefix := strings.CutPrefix(signature, prefix)
	if !foundPrefix {
		return false
	}

	mac := hmac.New(hashFunc, []byte(wh.AppSecret))
	mac.Write(body)
	actualSig := hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(expectedSig), []byte(actualSig))
}

func (wh *Webhook) handleWebhookPOST(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("reading body: %w", err)
		if !wh.HandleWebhookErr(r.Context(), w, nil, err) {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
		}
		return
	}

	if !wh.verifySignature(r, body) {
		if !wh.HandleWebhookErr(r.Context(), w, nil, errors.New("invalid signature")) {
			http.Error(w, "Invalid signature", http.StatusForbidden)
		}
		return
	}

	var request WebhookRequest
	if err := json.Unmarshal(body, &request); err != nil {
		err = fmt.Errorf("unmarshalling request body: %w", err)
		if !wh.HandleWebhookErr(r.Context(), w, &request, err) {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		}
		return
	}

	wh.Handler.HandleWebhook(r.Context(), w, &request)
}
