package whatsapp

import (
	"fmt"
	"io"
)

// MessagingProduct represents the type of messaging product used in the request.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages#messaging_product
type MessagingProduct string

const (
	// MessagingProductWhatsApp represents the WhatsApp messaging product.
	MessagingProductWhatsApp MessagingProduct = "whatsapp"
)

// RecipientType represents the type of recipient for the message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages#recipient_type
type RecipientType string

const (
	// RecipientTypeIndividual represents an individual recipient.
	// This is typically used for sending messages to a single user.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages#recipient_type
	RecipientTypeIndividual RecipientType = "individual"
)

// MessageType represents the type of message being sent.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages#type
type MessageType string

const (
	// MessageTypeText represents a text message.
	MessageTypeText MessageType = "text"
	// MessageTypeImage represents an image message.
	MessageTypeImage MessageType = "image"
	// MessageTypeAudio represents an audio message.
	MessageTypeAudio MessageType = "audio"
	// MessageTypeVideo represents a video message.
	MessageTypeVideo MessageType = "video"
	// MessageTypeDocument represents a document message.
	MessageTypeDocument MessageType = "document"
	// MessageTypeSticker represents a sticker message.
	MessageTypeSticker MessageType = "sticker"
	// MessageTypeLocation represents a location message.
	MessageTypeLocation MessageType = "location"
	// MessageTypeContacts represents a contacts message.
	MessageTypeContacts MessageType = "contacts"
	// MessageTypeButton represents a button message.
	MessageTypeButton MessageType = "button"
	// MessageTypeInteractive represents an interactive message.
	MessageTypeInteractive MessageType = "interactive"
	// MessageTypeOrder represents an order message.
	MessageTypeOrder MessageType = "order"
	// MessageTypeSystem represents a system message.
	MessageTypeSystem MessageType = "system"
	// MessageTypeReaction represents a reaction message.
	MessageTypeReaction MessageType = "reaction"
	// MessageTypeUnknown represents an unknown message type.
	MessageTypeUnknown MessageType = "unknown"
	// MessageTypeUnsupported represents an unsupported message type.
	MessageTypeUnsupported MessageType = "unsupported"
)

// InteractiveType represents the type of interactive message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
type InteractiveType string

const (
	// InteractiveTypeFlow represents a flow interactive message.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
	InteractiveTypeFlow InteractiveType = "flow"
	// InteractiveTypeButton represents a button interactive message.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-reply-buttons-messages
	InteractiveTypeButton InteractiveType = "button"
	// InteractiveTypeList represents a list interactive message.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-list-messages
	InteractiveTypeList InteractiveType = "list"
	// InteractiveTypeCTAURL represents a call-to-action URL interactive message.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-cta-url-messages
	InteractiveTypeCTAURL InteractiveType = "cta_url"
	// InteractiveTypeButtonReply represents a button reply interactive message.
	InteractiveTypeButtonReply InteractiveType = "button_reply"
	// InteractiveTypeListReply represents a list reply interactive message.
	InteractiveTypeListReply InteractiveType = "list_reply"
)

// HeaderType represents the type of header in an interactive message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
type HeaderType string

const (
	// HeaderTypeText represents a text header.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
	HeaderTypeText HeaderType = "text"
	// HeaderTypeImage represents an image header.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-reply-buttons-messages
	HeaderTypeImage HeaderType = "image"
	// HeaderTypeVideo represents a video header.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-reply-buttons-messages
	HeaderTypeVideo HeaderType = "video"
	// HeaderTypeDocument represents a document header.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-reply-buttons-messages
	HeaderTypeDocument HeaderType = "document"
)

// FlowAction represents the type of flow action.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
type FlowAction string

const (
	// FlowActionNavigate represents a navigate flow action.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
	FlowActionNavigate FlowAction = "navigate"
	// FlowActionDataExchange represents a data exchange flow action.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
	FlowActionDataExchange FlowAction = "data_exchange"
)

// FlowMode represents the mode of the flow.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
type FlowMode string

const (
	// FlowModeDraft represents a draft flow mode.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
	FlowModeDraft FlowMode = "draft"
	// FlowModePublished represents a published flow mode. This is the default mode for interactive flow messages.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
	FlowModePublished FlowMode = "published"
)

// Request represents a request to send a message via the WhatsApp Business API.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages
type Request struct {
	MessagingProduct MessagingProduct `json:"messaging_product"`
	RecipientType    RecipientType    `json:"recipient_type"`
	To               string           `json:"to"`
	Type             MessageType      `json:"type"`
	Text             *SendTextParams  `json:"text,omitempty"`
	Image            *SendImageParams `json:"image,omitempty"`
	Interactive      *Interactive     `json:"interactive,omitempty"`
}

// Interactive represents the interactive object for interactive messages.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
type Interactive struct {
	Type   InteractiveType `json:"type"`
	Header *Header         `json:"header,omitempty"`
	Body   *Body           `json:"body,omitempty"`
	Footer *Footer         `json:"footer,omitempty"`
	Action *Action         `json:"action,omitempty"`
}

// Header represents the header of an interactive message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-reply-buttons-messages
type Header struct {
	Type     HeaderType   `json:"type"`
	Text     string       `json:"text,omitempty"`
	Image    *MediaObject `json:"image,omitempty"`
	Video    *MediaObject `json:"video,omitempty"`
	Document *MediaObject `json:"document,omitempty"`
}

// Body represents the body of an interactive message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
type Body struct {
	Text string `json:"text"`
}

// Footer represents the footer of an interactive message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
type Footer struct {
	Text string `json:"text"`
}

// ActionParameters is an interface that all action parameters must implement.
// This ensures type safety while allowing different parameter types for different actions.
type ActionParameters interface {
	// ActionType returns the type of action these parameters are for
	ActionType() string
	// Validate performs validation on the parameters
	Validate() error
}

// Action represents the action of an interactive message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-reply-buttons-messages
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-list-messages
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-cta-url-messages
type Action struct {
	Name       string           `json:"name,omitempty"`
	Parameters ActionParameters `json:"parameters,omitempty"`
	Buttons    []Button         `json:"buttons,omitempty"`
	Button     string           `json:"button,omitempty"`
	Sections   []ListSection    `json:"sections,omitempty"`
}

// FlowParameters represents the parameters for a flow action.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
type FlowParameters struct {
	FlowMessageVersion string             `json:"flow_message_version"`
	FlowMode           FlowMode           `json:"flow_mode,omitempty"`
	FlowToken          string             `json:"flow_token"`
	FlowID             string             `json:"flow_id"`
	FlowCTA            string             `json:"flow_cta"`
	FlowAction         FlowAction         `json:"flow_action"`
	FlowActionPayload  *FlowActionPayload `json:"flow_action_payload,omitempty"`
}

// ActionType returns the action type for flow parameters
func (fp *FlowParameters) ActionType() string {
	return "flow"
}

// Validate validates the flow parameters
func (fp *FlowParameters) Validate() error {
	if fp == nil {
		return fmt.Errorf("flow parameters cannot be nil")
	}
	if fp.FlowMessageVersion == "" {
		return fmt.Errorf("flow_message_version is required")
	}
	if fp.FlowToken == "" {
		return fmt.Errorf("flow_token is required")
	}
	if fp.FlowID == "" {
		return fmt.Errorf("flow_id is required")
	}
	if fp.FlowCTA == "" {
		return fmt.Errorf("flow_cta is required")
	}
	if fp.FlowAction == "" {
		return fmt.Errorf("flow_action is required")
	}
	return nil
}

// NewFlowParameters creates a new FlowParameters instance with validation.
// This is a convenience constructor that ensures all required fields are provided.
func NewFlowParameters(flowMessageVersion, flowToken, flowID, flowCTA string, flowAction FlowAction) (*FlowParameters, error) {
	params := &FlowParameters{
		FlowMessageVersion: flowMessageVersion,
		FlowToken:          flowToken,
		FlowID:             flowID,
		FlowCTA:            flowCTA,
		FlowAction:         flowAction,
	}

	if err := params.Validate(); err != nil {
		return nil, err
	}

	return params, nil
}

// CTAURLParameters represents the parameters for a CTA URL action.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-cta-url-messages
type CTAURLParameters struct {
	DisplayText string `json:"display_text"`
	URL         string `json:"url"`
}

// ActionType returns the action type for CTA URL parameters
func (cp *CTAURLParameters) ActionType() string {
	return "cta_url"
}

// Validate validates the CTA URL parameters
func (cp *CTAURLParameters) Validate() error {
	if cp == nil {
		return fmt.Errorf("CTA URL parameters cannot be nil")
	}
	if cp.DisplayText == "" {
		return fmt.Errorf("display_text is required")
	}
	if cp.URL == "" {
		return fmt.Errorf("url is required")
	}
	return nil
}

// NewCTAURLParameters creates a new CTAURLParameters instance with validation.
// This is a convenience constructor that ensures all required fields are provided.
func NewCTAURLParameters(displayText, url string) (*CTAURLParameters, error) {
	params := &CTAURLParameters{
		DisplayText: displayText,
		URL:         url,
	}

	if err := params.Validate(); err != nil {
		return nil, err
	}

	return params, nil
}

// SendTextParams contains parameters for sending a text message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/text-messages
type SendTextParams struct {
	// PreviewURL should be set to true to have the WhatsApp client attempt
	// to render a link preview of any URL in the body text string.
	PreviewURL bool `json:"preview_url,omitempty"`
	// Body text. Required. URLs are automatically hyperlinked.
	// Maximum 1024 characters.
	Body string `json:"body"`
}

// SendImageParams contains parameters for sending an image message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/image-messages
type SendImageParams struct {
	// ID is the media object ID. Required when not using link.
	// Only one of ID or Link should be provided.
	ID string `json:"id,omitempty"`
	// Link is the URL of the image. Required when not using ID.
	// Only one of ID or Link should be provided.
	// The image must be 5MB or smaller.
	// Supported formats: JPEG, PNG.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#supported-media-types
	Link string `json:"link,omitempty"`
	// Caption is optional text that appears below the image.
	// Maximum 1024 characters.
	Caption string `json:"caption,omitempty"`
}

// Validate validates the image parameters
func (sip *SendImageParams) Validate() error {
	if sip == nil {
		return fmt.Errorf("image parameters cannot be nil")
	}
	if sip.ID == "" && sip.Link == "" {
		return fmt.Errorf("either ID or Link must be provided")
	}
	if sip.ID != "" && sip.Link != "" {
		return fmt.Errorf("only one of ID or Link should be provided")
	}
	if len(sip.Caption) > 1024 {
		return fmt.Errorf("caption exceeds maximum length of 1024 characters")
	}
	return nil
}

// NewSendImageParamsWithID creates a new SendImageParams instance using a media ID with validation.
// This is a convenience constructor for sending images using an existing media object.
func NewSendImageParamsWithID(id string, caption ...string) (*SendImageParams, error) {
	params := &SendImageParams{
		ID: id,
	}
	if len(caption) > 0 {
		params.Caption = caption[0]
	}

	if err := params.Validate(); err != nil {
		return nil, err
	}

	return params, nil
}

// NewSendImageParamsWithLink creates a new SendImageParams instance using a URL with validation.
// This is a convenience constructor for sending images using a direct URL.
func NewSendImageParamsWithLink(link string, caption ...string) (*SendImageParams, error) {
	params := &SendImageParams{
		Link: link,
	}
	if len(caption) > 0 {
		params.Caption = caption[0]
	}

	if err := params.Validate(); err != nil {
		return nil, err
	}

	return params, nil
}

// SendInteractiveFlowParams contains parameters for sending an interactive flow message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
type SendInteractiveFlowParams struct {
	// Header is optional header for the flow message
	Header *Header `json:"header,omitempty"`
	// Body is required body text for the flow message
	Body *Body `json:"body"`
	// Footer is optional footer for the flow message
	Footer *Footer `json:"footer,omitempty"`
	// FlowParameters contains the flow-specific parameters
	FlowParameters *FlowParameters `json:"flow_parameters"`
}

// SendInteractiveButtonsParams contains parameters for sending an interactive reply buttons message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-reply-buttons-messages
type SendInteractiveButtonsParams struct {
	// Header is optional header for the button message (text, image, video, or document)
	Header *Header `json:"header,omitempty"`
	// Body is required body text for the button message
	Body *Body `json:"body"`
	// Footer is optional footer for the button message
	Footer *Footer `json:"footer,omitempty"`
	// Buttons is the array of reply buttons (maximum 3)
	Buttons []Button `json:"buttons"`
}

// SendInteractiveListParams contains parameters for sending an interactive list message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-list-messages
type SendInteractiveListParams struct {
	// Header is optional header for the list message (text only)
	Header *Header `json:"header,omitempty"`
	// Body is required body text for the list message
	Body *Body `json:"body"`
	// Footer is optional footer for the list message
	Footer *Footer `json:"footer,omitempty"`
	// Button is the text displayed on the button that opens the list. Maximum 20 characters.
	Button string `json:"button"`
	// Sections is an array of sections. Maximum 10 sections. Maximum 10 rows across all sections.
	Sections []ListSection `json:"sections"`
}

// SendInteractiveCTAURLParams contains parameters for sending an interactive CTA URL message.
// CTA URL button messages allow you to map any URL to a button so you don't have to
// include the raw URL in the message body. This is useful when URLs contain lengthy
// or obscure strings that users may be hesitant to tap.
//
// Example usage:
//
//	params := &SendInteractiveCTAURLParams{
//	    Header: &Header{
//	        Type: HeaderTypeImage,
//	        Image: &MediaObject{
//	            Link: "https://example.com/banner.jpg",
//	        },
//	    },
//	    Body: &Body{
//	        Text: "Check out our amazing products and special offers!",
//	    },
//	    Footer: &Footer{
//	        Text: "Shop now and save big",
//	    },
//	    DisplayText: "Visit Store",
//	    URL: "https://example.com/store?utm_source=whatsapp&utm_campaign=promotion",
//	}
//
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-cta-url-messages
type SendInteractiveCTAURLParams struct {
	// Header is optional header for the CTA URL message (text, image, video, or document)
	Header *Header `json:"header,omitempty"`
	// Body is required body text for the CTA URL message
	Body *Body `json:"body"`
	// Footer is optional footer for the CTA URL message
	Footer *Footer `json:"footer,omitempty"`
	// DisplayText is the text displayed on the CTA button
	DisplayText string `json:"display_text"`
	// URL is the URL that will be opened when the button is tapped
	URL string `json:"url"`
}

// ListSection represents a section within an interactive list message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-list-messages
type ListSection struct {
	// Title is the title of the section. Maximum 24 characters.
	Title string `json:"title,omitempty"`
	// Rows is an array of rows within the section. Maximum 10 rows across all sections.
	Rows []ListRow `json:"rows"`
}

// ListRow represents a row within a section of an interactive list message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-list-messages
type ListRow struct {
	// ID is the unique identifier for the row. Maximum 200 characters.
	ID string `json:"id"`
	// Title is the title of the row. Maximum 24 characters.
	Title string `json:"title"`
	// Description is optional description for the row. Maximum 72 characters.
	Description string `json:"description,omitempty"`
}

// MessagesResponse represents the response from the WhatsApp Business API when sending a message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/guides/send-messages/
//
// The API will return the following JSON response if it successfully accepts your send
// message request without encountering any errors in the request itself. Note that this
// response only indicates that the API successfully accepted your request, it does not
// indicate successful delivery of your message. Message delivery status is communicated
// via messages webhooks instead.
type MessagesResponse struct {
	MessagingProduct MessagingProduct          `json:"messaging_product"`
	Contacts         []MessagesResponseContact `json:"contacts,omitempty"`
	Messages         []MessagesResponseMessage `json:"messages,omitempty"`
}

// MessagesResponseContact represents a contact in the response from the WhatsApp Business API.
type MessagesResponseContact struct {
	// Input is WhatsApp user's WhatsApp phone number. May not match wa_id value.
	Input string `json:"input"`
	// WaID is the WhatsApp user's WhatsApp ID. May not match input value.
	WaID string `json:"wa_id,omitempty"`
}

// MessagesResponseMessage represents a message in the response from the WhatsApp Business API.
type MessagesResponseMessage struct {
	// ID is a WhatsApp message ID. appears in associated messages webhooks, such
	// as sent, read, and delivered webhooks.
	ID string `json:"id"`
	// Indicates template pacing status. The message_status property is only included
	// in responses when sending a template message that uses a template that is being
	// paced.
	MessageStatus string `json:"status,omitempty"`
}

// WebhookRequest represents the top-level webhook notification payload from WhatsApp Business API.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components
type WebhookRequest struct {
	Object string         `json:"object"`
	Entry  []WebhookEntry `json:"entry"`
}

// WebhookEntry represents an entry in the webhook notification payload.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components
type WebhookEntry struct {
	ID      string          `json:"id"`
	Changes []WebhookChange `json:"changes"`
}

// WebhookChange represents a change in the webhook notification payload.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components
type WebhookChange struct {
	Value WebhookValue `json:"value"`
	Field string       `json:"field"`
}

// WebhookValue contains the details for the change that triggered the webhook.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components
type WebhookValue struct {
	MessagingProduct MessagingProduct `json:"messaging_product"`
	Metadata         WebhookMetadata  `json:"metadata"`
	Contacts         []WebhookContact `json:"contacts,omitempty"`
	Messages         []WebhookMessage `json:"messages,omitempty"`
	Statuses         []WebhookStatus  `json:"statuses,omitempty"`
	Errors           []WebhookError   `json:"errors,omitempty"`
}

// WebhookMetadata contains metadata about the webhook notification.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components
type WebhookMetadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

// WebhookContact represents a contact in the webhook notification.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookContact struct {
	Profile WebhookProfile `json:"profile"`
	WaID    string         `json:"wa_id"`
}

// WebhookProfile represents a user profile in the webhook notification.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookProfile struct {
	Name string `json:"name"`
}

// WebhookMessage represents a message in the webhook notification.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessage struct {
	From        string                     `json:"from"`
	ID          string                     `json:"id"`
	Timestamp   string                     `json:"timestamp"`
	Type        MessageType                `json:"type"`
	Context     *WebhookMessageContext     `json:"context,omitempty"`
	Text        *WebhookMessageText        `json:"text,omitempty"`
	Image       *WebhookMessageMedia       `json:"image,omitempty"`
	Audio       *WebhookMessageMedia       `json:"audio,omitempty"`
	Video       *WebhookMessageMedia       `json:"video,omitempty"`
	Document    *WebhookMessageMedia       `json:"document,omitempty"`
	Sticker     *WebhookMessageMedia       `json:"sticker,omitempty"`
	Location    *WebhookMessageLocation    `json:"location,omitempty"`
	Contacts    []WebhookMessageContact    `json:"contacts,omitempty"`
	Button      *WebhookMessageButton      `json:"button,omitempty"`
	Interactive *WebhookMessageInteractive `json:"interactive,omitempty"`
	Order       *WebhookMessageOrder       `json:"order,omitempty"`
	System      *WebhookMessageSystem      `json:"system,omitempty"`
	Reaction    *WebhookMessageReaction    `json:"reaction,omitempty"`
	Referral    *WebhookMessageReferral    `json:"referral,omitempty"`
	Errors      []WebhookError             `json:"errors,omitempty"`
}

// WebhookMessageContext represents the context of a message in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessageContext struct {
	From            string                  `json:"from,omitempty"`
	ID              string                  `json:"id,omitempty"`
	ReferredProduct *WebhookReferredProduct `json:"referred_product,omitempty"`
	Forwarded       bool                    `json:"forwarded,omitempty"`
}

// WebhookReferredProduct represents a referred product in message context.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookReferredProduct struct {
	CatalogID         string `json:"catalog_id"`
	ProductRetailerID string `json:"product_retailer_id"`
}

// WebhookMessageText represents a text message in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessageText struct {
	Body string `json:"body"`
}

// WebhookMessageMedia represents media (image, audio, video, document, sticker) in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessageMedia struct {
	Caption  string `json:"caption,omitempty"`
	Filename string `json:"filename,omitempty"`
	ID       string `json:"id"`
	MimeType string `json:"mime_type"`
	SHA256   string `json:"sha256"`
}

// WebhookMessageLocation represents a location message in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessageLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name,omitempty"`
	Address   string  `json:"address,omitempty"`
}

// WebhookMessageContact represents a contact in a contacts message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessageContact struct {
	Addresses []WebhookContactAddress `json:"addresses,omitempty"`
	Birthday  string                  `json:"birthday,omitempty"`
	Emails    []WebhookContactEmail   `json:"emails,omitempty"`
	Name      *WebhookContactName     `json:"name,omitempty"`
	Org       *WebhookContactOrg      `json:"org,omitempty"`
	Phones    []WebhookContactPhone   `json:"phones,omitempty"`
	URLs      []WebhookContactURL     `json:"urls,omitempty"`
}

// ContactAddressType represents the type of contact address.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type ContactAddressType string

const (
	// ContactAddressTypeHome represents a home address.
	ContactAddressTypeHome ContactAddressType = "HOME"
	// ContactAddressTypeWork represents a work address.
	ContactAddressTypeWork ContactAddressType = "WORK"
)

// WebhookContactAddress represents an address in a contact.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookContactAddress struct {
	Street      string             `json:"street,omitempty"`
	City        string             `json:"city,omitempty"`
	State       string             `json:"state,omitempty"`
	Zip         string             `json:"zip,omitempty"`
	Country     string             `json:"country,omitempty"`
	CountryCode string             `json:"country_code,omitempty"`
	Type        ContactAddressType `json:"type,omitempty"`
}

// ContactEmailType represents the type of contact email.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type ContactEmailType string

const (
	// ContactEmailTypeHome represents a home email.
	ContactEmailTypeHome ContactEmailType = "HOME"
	// ContactEmailTypeWork represents a work email.
	ContactEmailTypeWork ContactEmailType = "WORK"
)

// WebhookContactEmail represents an email in a contact.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookContactEmail struct {
	Email string           `json:"email"`
	Type  ContactEmailType `json:"type,omitempty"`
}

// WebhookContactName represents a name in a contact.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookContactName struct {
	FormattedName string `json:"formatted_name,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	MiddleName    string `json:"middle_name,omitempty"`
	Suffix        string `json:"suffix,omitempty"`
	Prefix        string `json:"prefix,omitempty"`
}

// WebhookContactOrg represents an organization in a contact.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookContactOrg struct {
	Company    string `json:"company,omitempty"`
	Department string `json:"department,omitempty"`
	Title      string `json:"title,omitempty"`
}

// ContactPhoneType represents the type of contact phone.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type ContactPhoneType string

const (
	// ContactPhoneTypeHome represents a home phone.
	ContactPhoneTypeHome ContactPhoneType = "HOME"
	// ContactPhoneTypeWork represents a work phone.
	ContactPhoneTypeWork ContactPhoneType = "WORK"
)

// WebhookContactPhone represents a phone number in a contact.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookContactPhone struct {
	Phone string           `json:"phone"`
	WaID  string           `json:"wa_id,omitempty"`
	Type  ContactPhoneType `json:"type,omitempty"`
}

// ContactURLType represents the type of contact URL.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type ContactURLType string

const (
	// ContactURLTypeHome represents a home URL.
	ContactURLTypeHome ContactURLType = "HOME"
	// ContactURLTypeWork represents a work URL.
	ContactURLTypeWork ContactURLType = "WORK"
)

// WebhookContactURL represents a URL in a contact.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookContactURL struct {
	URL  string         `json:"url"`
	Type ContactURLType `json:"type,omitempty"`
}

// WebhookMessageButton represents a button message in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessageButton struct {
	Text    string `json:"text"`
	Payload string `json:"payload"`
}

// WebhookMessageInteractive represents an interactive message in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessageInteractive struct {
	Type        InteractiveType                    `json:"type"`
	ButtonReply *WebhookMessageInteractiveButton   `json:"button_reply,omitempty"`
	ListReply   *WebhookMessageInteractiveListItem `json:"list_reply,omitempty"`
}

// WebhookMessageInteractiveButton represents a button reply in an interactive message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessageInteractiveButton struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// WebhookMessageInteractiveListItem represents a list reply in an interactive message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessageInteractiveListItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

// WebhookMessageOrder represents an order message in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessageOrder struct {
	CatalogID    string                    `json:"catalog_id"`
	ProductItems []WebhookMessageOrderItem `json:"product_items"`
	Text         string                    `json:"text,omitempty"`
}

// WebhookMessageOrderItem represents an item in an order message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessageOrderItem struct {
	ProductRetailerID string `json:"product_retailer_id"`
	Quantity          string `json:"quantity"`
	ItemPrice         string `json:"item_price"`
	Currency          string `json:"currency"`
}

// SystemMessageType represents the type of system message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type SystemMessageType string

const (
	// SystemMessageTypeUserChangedNumber represents a user changed number system message.
	SystemMessageTypeUserChangedNumber SystemMessageType = "user_changed_number"
)

// WebhookMessageSystem represents a system message in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessageSystem struct {
	Body    string            `json:"body"`
	NewWaID string            `json:"new_wa_id,omitempty"`
	Type    SystemMessageType `json:"type"`
}

// WebhookMessageReaction represents a reaction message in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessageReaction struct {
	MessageID string `json:"message_id"`
	Emoji     string `json:"emoji"`
}

// ReferralSourceType represents the source type of a referral.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type ReferralSourceType string

const (
	// ReferralSourceTypeAd represents an ad referral source.
	ReferralSourceTypeAd ReferralSourceType = "ad"
	// ReferralSourceTypePost represents a post referral source.
	ReferralSourceTypePost ReferralSourceType = "post"
)

// ReferralMediaType represents the media type of a referral.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type ReferralMediaType string

const (
	// ReferralMediaTypeImage represents an image referral media.
	ReferralMediaTypeImage ReferralMediaType = "image"
	// ReferralMediaTypeVideo represents a video referral media.
	ReferralMediaTypeVideo ReferralMediaType = "video"
)

// WebhookMessageReferral represents a referral message in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookMessageReferral struct {
	SourceURL    string             `json:"source_url"`
	SourceID     string             `json:"source_id"`
	SourceType   ReferralSourceType `json:"source_type"`
	Headline     string             `json:"headline,omitempty"`
	Body         string             `json:"body,omitempty"`
	MediaType    ReferralMediaType  `json:"media_type,omitempty"`
	ImageURL     string             `json:"image_url,omitempty"`
	VideoURL     string             `json:"video_url,omitempty"`
	ThumbnailURL string             `json:"thumbnail_url,omitempty"`
	CTWAClid     string             `json:"ctwa_clid,omitempty"`
}

// MessageStatus represents the status of a message in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type MessageStatus string

const (
	// MessageStatusSent represents a sent message status.
	MessageStatusSent MessageStatus = "sent"
	// MessageStatusDelivered represents a delivered message status.
	MessageStatusDelivered MessageStatus = "delivered"
	// MessageStatusRead represents a read message status.
	MessageStatusRead MessageStatus = "read"
	// MessageStatusFailed represents a failed message status.
	MessageStatusFailed MessageStatus = "failed"
)

// WebhookStatus represents a message status in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookStatus struct {
	ID           string                     `json:"id"`
	Status       MessageStatus              `json:"status"`
	Timestamp    string                     `json:"timestamp"`
	RecipientID  string                     `json:"recipient_id"`
	Conversation *WebhookStatusConversation `json:"conversation,omitempty"`
	Pricing      *WebhookStatusPricing      `json:"pricing,omitempty"`
	Errors       []WebhookError             `json:"errors,omitempty"`
}

// ConversationOriginType represents the origin type of a conversation.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type ConversationOriginType string

const (
	// ConversationOriginTypeReferralConversion represents a referral conversion origin.
	ConversationOriginTypeReferralConversion ConversationOriginType = "referral_conversion"
	// ConversationOriginTypeUserInitiated represents a user initiated origin.
	ConversationOriginTypeUserInitiated ConversationOriginType = "user_initiated"
	// ConversationOriginTypeBusinessInitiated represents a business initiated origin.
	ConversationOriginTypeBusinessInitiated ConversationOriginType = "business_initiated"
)

// WebhookStatusConversation represents conversation information in status notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookStatusConversation struct {
	ID                  string                     `json:"id"`
	ExpirationTimestamp string                     `json:"expiration_timestamp,omitempty"`
	Origin              *WebhookConversationOrigin `json:"origin,omitempty"`
}

// WebhookConversationOrigin represents the origin of a conversation.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookConversationOrigin struct {
	Type ConversationOriginType `json:"type"`
}

// PricingModel represents the pricing model in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type PricingModel string

const (
	// PricingModelCBP represents the CBP (Conversation-Based Pricing) model.
	PricingModelCBP PricingModel = "CBP"
)

// PricingCategory represents the pricing category in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type PricingCategory string

const (
	// PricingCategoryReferralConversion represents referral conversion pricing.
	PricingCategoryReferralConversion PricingCategory = "referral_conversion"
	// PricingCategoryUserInitiated represents user initiated pricing.
	PricingCategoryUserInitiated PricingCategory = "user_initiated"
	// PricingCategoryBusinessInitiated represents business initiated pricing.
	PricingCategoryBusinessInitiated PricingCategory = "business_initiated"
)

// WebhookStatusPricing represents pricing information in status notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookStatusPricing struct {
	Billable     bool            `json:"billable"`
	PricingModel PricingModel    `json:"pricing_model"`
	Category     PricingCategory `json:"category"`
}

// WebhookError represents an error in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookError struct {
	Code      int               `json:"code"`
	Title     string            `json:"title"`
	Message   string            `json:"message,omitempty"`
	Details   string            `json:"details,omitempty"`
	ErrorData *WebhookErrorData `json:"error_data,omitempty"`
	Href      string            `json:"href,omitempty"`
}

// WebhookErrorData represents additional error data in webhook notifications.
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type WebhookErrorData struct {
	Details string `json:"details"`
}

// MediaObject represents a media object (image, video, document) used in interactive messages.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-reply-buttons-messages
type MediaObject struct {
	ID       string `json:"id,omitempty"`
	Link     string `json:"link,omitempty"`
	Caption  string `json:"caption,omitempty"`
	Filename string `json:"filename,omitempty"`
}

// ButtonType represents the type of button in an interactive message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-reply-buttons-messages
type ButtonType string

const (
	// ButtonTypeReply represents a reply button.
	// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-reply-buttons-messages
	ButtonTypeReply ButtonType = "reply"
)

// ReplyButton represents a reply button in an interactive message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-reply-buttons-messages
type ReplyButton struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// Button represents a button in an interactive message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-reply-buttons-messages
type Button struct {
	Type  ButtonType   `json:"type"`
	Reply *ReplyButton `json:"reply,omitempty"`
}

// MediaResponse represents the response when retrieving media information.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#retrieve-media-url
type MediaResponse struct {
	URL              string `json:"url"`
	MimeType         string `json:"mime_type"`
	SHA256           string `json:"sha256"`
	FileSize         int64  `json:"file_size"`
	ID               string `json:"id"`
	MessagingProduct string `json:"messaging_product"`
}

// MediaError represents an error response when retrieving media.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#retrieve-media-url
type MediaError struct {
	Error struct {
		Message   string `json:"message"`
		Type      string `json:"type"`
		Code      int    `json:"code"`
		ErrorData struct {
			Details string `json:"details"`
		} `json:"error_data"`
		FBTraceID string `json:"fbtrace_id"`
	} `json:"error"`
}

// APIError represents an error response from the WhatsApp Business API.
// This structure is used for general API errors from endpoints like messages.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages#errors
type APIError struct {
	Error struct {
		Message   string `json:"message"`
		Type      string `json:"type"`
		Code      int    `json:"code"`
		ErrorData struct {
			Details string `json:"details"`
		} `json:"error_data"`
		FBTraceID string `json:"fbtrace_id"`
	} `json:"error"`
}

// MediaSizeLimit represents the maximum file size limits for different media types.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#supported-media-types
const (
	// MaxImageSize is the maximum size for image files (5MB)
	MaxImageSize = 5 * 1024 * 1024
	// MaxAudioSize is the maximum size for audio files (16MB)
	MaxAudioSize = 16 * 1024 * 1024
	// MaxVideoSize is the maximum size for video files (16MB)
	MaxVideoSize = 16 * 1024 * 1024
	// MaxDocumentSize is the maximum size for document files (100MB)
	MaxDocumentSize = 100 * 1024 * 1024
	// MaxStickerSize is the maximum size for sticker files (100KB)
	MaxStickerSize = 100 * 1024
)

// SupportedMimeType represents supported MIME types for WhatsApp media.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#supported-media-types
type SupportedMimeType string

const (
	// Image MIME types
	MimeTypeImageJPEG SupportedMimeType = "image/jpeg"
	MimeTypeImagePNG  SupportedMimeType = "image/png"
	MimeTypeImageWebP SupportedMimeType = "image/webp"

	// Audio MIME types
	MimeTypeAudioAAC  SupportedMimeType = "audio/aac"
	MimeTypeAudioMP4  SupportedMimeType = "audio/mp4"
	MimeTypeAudioMPEG SupportedMimeType = "audio/mpeg"
	MimeTypeAudioAMR  SupportedMimeType = "audio/amr"
	MimeTypeAudioOGG  SupportedMimeType = "audio/ogg"

	// Video MIME types
	MimeTypeVideoMP4  SupportedMimeType = "video/mp4"
	MimeTypeVideo3GPP SupportedMimeType = "video/3gpp"

	// Document MIME types
	MimeTypeDocumentPDF  SupportedMimeType = "application/pdf"
	MimeTypeDocumentDOCX SupportedMimeType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	MimeTypeDocumentPPTX SupportedMimeType = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	MimeTypeDocumentXLSX SupportedMimeType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"

	// Sticker MIME types
	MimeTypeStickerWebP SupportedMimeType = "image/webp"
)

// ValidateMediaSize checks if the media size is within the allowed limits for its type.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#supported-media-types
func ValidateMediaSize(mimeType string, size int64) error {
	switch {
	case size <= 0:
		return fmt.Errorf("invalid media size: %d", size)
	case mimeType == string(MimeTypeImageJPEG) || mimeType == string(MimeTypeImagePNG) || mimeType == string(MimeTypeImageWebP):
		if size > MaxImageSize {
			return fmt.Errorf("image size %d exceeds maximum allowed size %d", size, MaxImageSize)
		}
	case mimeType == string(MimeTypeAudioAAC) || mimeType == string(MimeTypeAudioMP4) ||
		mimeType == string(MimeTypeAudioMPEG) || mimeType == string(MimeTypeAudioAMR) ||
		mimeType == string(MimeTypeAudioOGG):
		if size > MaxAudioSize {
			return fmt.Errorf("audio size %d exceeds maximum allowed size %d", size, MaxAudioSize)
		}
	case mimeType == string(MimeTypeVideoMP4) || mimeType == string(MimeTypeVideo3GPP):
		if size > MaxVideoSize {
			return fmt.Errorf("video size %d exceeds maximum allowed size %d", size, MaxVideoSize)
		}
	case mimeType == string(MimeTypeDocumentPDF) || mimeType == string(MimeTypeDocumentDOCX) ||
		mimeType == string(MimeTypeDocumentPPTX) || mimeType == string(MimeTypeDocumentXLSX):
		if size > MaxDocumentSize {
			return fmt.Errorf("document size %d exceeds maximum allowed size %d", size, MaxDocumentSize)
		}
	case mimeType == string(MimeTypeStickerWebP):
		if size > MaxStickerSize {
			return fmt.Errorf("sticker size %d exceeds maximum allowed size %d", size, MaxStickerSize)
		}
	default:
		return fmt.Errorf("unsupported MIME type: %s", mimeType)
	}
	return nil
}

// ValidateMimeType checks if the provided MIME type is supported by WhatsApp.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#supported-media-types
func ValidateMimeType(mimeType string) error {
	supportedTypes := []SupportedMimeType{
		// Image types
		MimeTypeImageJPEG,
		MimeTypeImagePNG,
		MimeTypeImageWebP,
		// Audio types
		MimeTypeAudioAAC,
		MimeTypeAudioMP4,
		MimeTypeAudioMPEG,
		MimeTypeAudioAMR,
		MimeTypeAudioOGG,
		// Video types
		MimeTypeVideoMP4,
		MimeTypeVideo3GPP,
		// Document types
		MimeTypeDocumentPDF,
		MimeTypeDocumentDOCX,
		MimeTypeDocumentPPTX,
		MimeTypeDocumentXLSX,
		// Sticker types
		MimeTypeStickerWebP,
	}

	for _, supportedType := range supportedTypes {
		if mimeType == string(supportedType) {
			return nil
		}
	}

	return fmt.Errorf("unsupported MIME type: %s", mimeType)
}

// UploadMediaParams contains parameters for uploading media.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#upload-media
type UploadMediaParams struct {
	// File is the media file to upload (io.Reader)
	File io.Reader `json:"-"`
	// Filename is the name of the file being uploaded
	Filename string `json:"-"`
	// MimeType is the MIME type of the media file
	MimeType string `json:"-"`
	// MessagingProduct must be set to "whatsapp"
	MessagingProduct MessagingProduct `json:"messaging_product"`
}

// Validate validates the upload media parameters
func (ump *UploadMediaParams) Validate() error {
	if ump == nil {
		return fmt.Errorf("upload media parameters cannot be nil")
	}
	if ump.File == nil {
		return fmt.Errorf("file is required")
	}
	if ump.Filename == "" {
		return fmt.Errorf("filename is required")
	}
	if ump.MimeType == "" {
		return fmt.Errorf("mime type is required")
	}
	if ump.MessagingProduct != MessagingProductWhatsApp {
		return fmt.Errorf("messaging product must be 'whatsapp'")
	}
	return nil
}

// NewUploadMediaParams creates a new UploadMediaParams instance with validation.
// This is a convenience constructor that ensures all required fields are provided.
func NewUploadMediaParams(file io.Reader, filename, mimeType string) (*UploadMediaParams, error) {
	params := &UploadMediaParams{
		File:             file,
		Filename:         filename,
		MimeType:         mimeType,
		MessagingProduct: MessagingProductWhatsApp,
	}

	if err := params.Validate(); err != nil {
		return nil, err
	}

	return params, nil
}

// UploadMediaResponse represents the response from uploading media.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#upload-media
type UploadMediaResponse struct {
	// ID is the media object ID that can be used in messages
	ID string `json:"id"`
}

// DeleteMediaResponse represents the response from deleting media.
// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#delete-media
type DeleteMediaResponse struct {
	// Success indicates whether the media was successfully deleted
	Success bool `json:"success"`
}

// SendInteractiveParams contains parameters for sending an interactive message.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
type SendInteractiveParams struct {
	// Header is optional header for the interactive message
	Header *Header `json:"header,omitempty"`
	// Body is required body text for the interactive message
	Body *Body `json:"body"`
	// Footer is optional footer for the interactive message
	Footer *Footer `json:"footer,omitempty"`
	// Action contains the action parameters for the interactive message
	Action *Action `json:"action,omitempty"`
}

// Validate validates the interactive parameters
func (sip *SendInteractiveParams) Validate() error {
	if sip == nil {
		return fmt.Errorf("interactive parameters cannot be nil")
	}
	if sip.Body == nil {
		return fmt.Errorf("body is required")
	}
	if sip.Action == nil {
		return fmt.Errorf("action is required")
	}
	return nil
}

// NewSendInteractiveParams creates a new SendInteractiveParams instance with validation.
// This is a convenience constructor that ensures all required fields are provided.
func NewSendInteractiveParams(body *Body, action *Action) (*SendInteractiveParams, error) {
	params := &SendInteractiveParams{
		Body:   body,
		Action: action,
	}

	if err := params.Validate(); err != nil {
		return nil, err
	}

	return params, nil
}

// ValidateAction validates that an Action has the correct parameters for its name.
// This provides additional type safety by ensuring parameter types match action names.
func ValidateAction(action *Action) error {
	if action == nil {
		return fmt.Errorf("action cannot be nil")
	}

	if action.Parameters != nil {
		// Validate that the parameter type matches the action name
		expectedType := action.Parameters.ActionType()
		if action.Name != "" && action.Name != expectedType {
			return fmt.Errorf("action name '%s' does not match parameter type '%s'", action.Name, expectedType)
		}

		// Validate the parameters themselves
		if err := action.Parameters.Validate(); err != nil {
			return fmt.Errorf("parameter validation failed: %w", err)
		}
	}

	return nil
}

// ExtractMediaID extracts the media ID from a webhook message, if present.
// This is a convenience function to get the media ID from different message types.
// Returns empty string if the message doesn't contain media.
//
// Example usage:
//
//	mediaID := ExtractMediaID(webhookMessage)
//	if mediaID != "" {
//	    mediaInfo, reader, err := client.GetAndDownloadMedia(ctx, mediaID)
//	    if err != nil {
//	        log.Printf("Failed to download media: %v", err)
//	        return
//	    }
//	    defer reader.Close()
//	    // Process the media stream...
//	}
//
// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
func ExtractMediaID(message *WebhookMessage) string {
	if message == nil {
		return ""
	}

	switch message.Type {
	case MessageTypeImage:
		if message.Image != nil {
			return message.Image.ID
		}
	case MessageTypeAudio:
		if message.Audio != nil {
			return message.Audio.ID
		}
	case MessageTypeVideo:
		if message.Video != nil {
			return message.Video.ID
		}
	case MessageTypeDocument:
		if message.Document != nil {
			return message.Document.ID
		}
	case MessageTypeSticker:
		if message.Sticker != nil {
			return message.Sticker.ID
		}
	}

	return ""
}

// FlowActionPayload represents the payload for a flow action.
// https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
type FlowActionPayload struct {
	Screen string                 `json:"screen"`
	Data   map[string]interface{} `json:"data,omitempty"`
}
