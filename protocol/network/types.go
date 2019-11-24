// Code generated by cdpgen. DO NOT EDIT.

package network

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/protocol/security"
)

// ResourceType Resource type as it was perceived by the rendering engine.
type ResourceType string

// ResourceType as enums.
const (
	ResourceTypeNotSet             ResourceType = ""
	ResourceTypeDocument           ResourceType = "Document"
	ResourceTypeStylesheet         ResourceType = "Stylesheet"
	ResourceTypeImage              ResourceType = "Image"
	ResourceTypeMedia              ResourceType = "Media"
	ResourceTypeFont               ResourceType = "Font"
	ResourceTypeScript             ResourceType = "Script"
	ResourceTypeTextTrack          ResourceType = "TextTrack"
	ResourceTypeXHR                ResourceType = "XHR"
	ResourceTypeFetch              ResourceType = "Fetch"
	ResourceTypeEventSource        ResourceType = "EventSource"
	ResourceTypeWebSocket          ResourceType = "WebSocket"
	ResourceTypeManifest           ResourceType = "Manifest"
	ResourceTypeSignedExchange     ResourceType = "SignedExchange"
	ResourceTypePing               ResourceType = "Ping"
	ResourceTypeCSPViolationReport ResourceType = "CSPViolationReport"
	ResourceTypeOther              ResourceType = "Other"
)

func (e ResourceType) Valid() bool {
	switch e {
	case "Document", "Stylesheet", "Image", "Media", "Font", "Script", "TextTrack", "XHR", "Fetch", "EventSource", "WebSocket", "Manifest", "SignedExchange", "Ping", "CSPViolationReport", "Other":
		return true
	default:
		return false
	}
}

func (e ResourceType) String() string {
	return string(e)
}

// LoaderID Unique loader identifier.
type LoaderID string

// RequestID Unique request identifier.
type RequestID string

// InterceptionID Unique intercepted request identifier.
type InterceptionID string

// ErrorReason Network level fetch failure reason.
type ErrorReason string

// ErrorReason as enums.
const (
	ErrorReasonNotSet               ErrorReason = ""
	ErrorReasonFailed               ErrorReason = "Failed"
	ErrorReasonAborted              ErrorReason = "Aborted"
	ErrorReasonTimedOut             ErrorReason = "TimedOut"
	ErrorReasonAccessDenied         ErrorReason = "AccessDenied"
	ErrorReasonConnectionClosed     ErrorReason = "ConnectionClosed"
	ErrorReasonConnectionReset      ErrorReason = "ConnectionReset"
	ErrorReasonConnectionRefused    ErrorReason = "ConnectionRefused"
	ErrorReasonConnectionAborted    ErrorReason = "ConnectionAborted"
	ErrorReasonConnectionFailed     ErrorReason = "ConnectionFailed"
	ErrorReasonNameNotResolved      ErrorReason = "NameNotResolved"
	ErrorReasonInternetDisconnected ErrorReason = "InternetDisconnected"
	ErrorReasonAddressUnreachable   ErrorReason = "AddressUnreachable"
	ErrorReasonBlockedByClient      ErrorReason = "BlockedByClient"
	ErrorReasonBlockedByResponse    ErrorReason = "BlockedByResponse"
)

func (e ErrorReason) Valid() bool {
	switch e {
	case "Failed", "Aborted", "TimedOut", "AccessDenied", "ConnectionClosed", "ConnectionReset", "ConnectionRefused", "ConnectionAborted", "ConnectionFailed", "NameNotResolved", "InternetDisconnected", "AddressUnreachable", "BlockedByClient", "BlockedByResponse":
		return true
	default:
		return false
	}
}

func (e ErrorReason) String() string {
	return string(e)
}

// TimeSinceEpoch UTC time in seconds, counted from January 1, 1970.
type TimeSinceEpoch = internal.NetworkTimeSinceEpoch

// MonotonicTime Monotonically increasing time in seconds since an arbitrary
// point in the past.
type MonotonicTime float64

// String calls (time.Time).String().
func (t MonotonicTime) String() string {
	return t.Time().String()
}

// Time parses the Unix time.
func (t MonotonicTime) Time() time.Time {
	ts := float64(t) / 1
	secs := int64(ts)
	nsecs := int64((ts - float64(secs)) * 1000000000)
	return time.Unix(secs, nsecs)
}

// MarshalJSON implements json.Marshaler. Encodes to null if t is zero.
func (t MonotonicTime) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte("null"), nil
	}
	f := float64(t)
	return json.Marshal(&f)
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *MonotonicTime) UnmarshalJSON(data []byte) error {
	*t = 0
	if len(data) == 0 {
		return nil
	}
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return errors.New("network.MonotonicTime: " + err.Error())
	}
	*t = MonotonicTime(f)
	return nil
}

var _ json.Marshaler = (*MonotonicTime)(nil)
var _ json.Unmarshaler = (*MonotonicTime)(nil)

// Headers Request / response headers as keys / values of JSON object.
type Headers []byte

// MarshalJSON copies behavior of json.RawMessage.
func (h Headers) MarshalJSON() ([]byte, error) {
	if h == nil {
		return []byte("null"), nil
	}
	return h, nil
}

// UnmarshalJSON copies behavior of json.RawMessage.
func (h *Headers) UnmarshalJSON(data []byte) error {
	if h == nil {
		return errors.New("network.Headers: UnmarshalJSON on nil pointer")
	}
	*h = append((*h)[0:0], data...)
	return nil
}

var _ json.Marshaler = (*Headers)(nil)
var _ json.Unmarshaler = (*Headers)(nil)

// ConnectionType The underlying connection technology that the browser is
// supposedly using.
type ConnectionType string

// ConnectionType as enums.
const (
	ConnectionTypeNotSet     ConnectionType = ""
	ConnectionTypeNone       ConnectionType = "none"
	ConnectionTypeCellular2g ConnectionType = "cellular2g"
	ConnectionTypeCellular3g ConnectionType = "cellular3g"
	ConnectionTypeCellular4g ConnectionType = "cellular4g"
	ConnectionTypeBluetooth  ConnectionType = "bluetooth"
	ConnectionTypeEthernet   ConnectionType = "ethernet"
	ConnectionTypeWifi       ConnectionType = "wifi"
	ConnectionTypeWimax      ConnectionType = "wimax"
	ConnectionTypeOther      ConnectionType = "other"
)

func (e ConnectionType) Valid() bool {
	switch e {
	case "none", "cellular2g", "cellular3g", "cellular4g", "bluetooth", "ethernet", "wifi", "wimax", "other":
		return true
	default:
		return false
	}
}

func (e ConnectionType) String() string {
	return string(e)
}

// CookieSameSite Represents the cookie's 'SameSite' status:
// https://tools.ietf.org/html/draft-west-first-party-cookies
type CookieSameSite string

// CookieSameSite as enums.
const (
	CookieSameSiteNotSet CookieSameSite = ""
	CookieSameSiteStrict CookieSameSite = "Strict"
	CookieSameSiteLax    CookieSameSite = "Lax"
	CookieSameSiteNone   CookieSameSite = "None"
)

func (e CookieSameSite) Valid() bool {
	switch e {
	case "Strict", "Lax", "None":
		return true
	default:
		return false
	}
}

func (e CookieSameSite) String() string {
	return string(e)
}

// ResourceTiming Timing information for the request.
type ResourceTiming struct {
	RequestTime  float64 `json:"requestTime"`  // Timing's requestTime is a baseline in seconds, while the other numbers are ticks in milliseconds relatively to this requestTime.
	ProxyStart   float64 `json:"proxyStart"`   // Started resolving proxy.
	ProxyEnd     float64 `json:"proxyEnd"`     // Finished resolving proxy.
	DNSStart     float64 `json:"dnsStart"`     // Started DNS address resolve.
	DNSEnd       float64 `json:"dnsEnd"`       // Finished DNS address resolve.
	ConnectStart float64 `json:"connectStart"` // Started connecting to the remote host.
	ConnectEnd   float64 `json:"connectEnd"`   // Connected to the remote host.
	SSLStart     float64 `json:"sslStart"`     // Started SSL handshake.
	SSLEnd       float64 `json:"sslEnd"`       // Finished SSL handshake.
	// WorkerStart Started running ServiceWorker.
	//
	// Note: This property is experimental.
	WorkerStart float64 `json:"workerStart"`
	// WorkerReady Finished Starting ServiceWorker.
	//
	// Note: This property is experimental.
	WorkerReady float64 `json:"workerReady"`
	SendStart   float64 `json:"sendStart"` // Started sending request.
	SendEnd     float64 `json:"sendEnd"`   // Finished sending request.
	// PushStart Time the server started pushing request.
	//
	// Note: This property is experimental.
	PushStart float64 `json:"pushStart"`
	// PushEnd Time the server finished pushing request.
	//
	// Note: This property is experimental.
	PushEnd           float64 `json:"pushEnd"`
	ReceiveHeadersEnd float64 `json:"receiveHeadersEnd"` // Finished receiving response headers.
}

// ResourcePriority Loading priority of a resource request.
type ResourcePriority string

// ResourcePriority as enums.
const (
	ResourcePriorityNotSet   ResourcePriority = ""
	ResourcePriorityVeryLow  ResourcePriority = "VeryLow"
	ResourcePriorityLow      ResourcePriority = "Low"
	ResourcePriorityMedium   ResourcePriority = "Medium"
	ResourcePriorityHigh     ResourcePriority = "High"
	ResourcePriorityVeryHigh ResourcePriority = "VeryHigh"
)

func (e ResourcePriority) Valid() bool {
	switch e {
	case "VeryLow", "Low", "Medium", "High", "VeryHigh":
		return true
	default:
		return false
	}
}

func (e ResourcePriority) String() string {
	return string(e)
}

// Request HTTP request data.
type Request struct {
	URL              string                     `json:"url"`                        // Request URL (without fragment).
	URLFragment      *string                    `json:"urlFragment,omitempty"`      // Fragment of the requested URL starting with hash, if present.
	Method           string                     `json:"method"`                     // HTTP request method.
	Headers          Headers                    `json:"headers"`                    // HTTP request headers.
	PostData         *string                    `json:"postData,omitempty"`         // HTTP POST request data.
	HasPostData      *bool                      `json:"hasPostData,omitempty"`      // True when the request has POST data. Note that postData might still be omitted when this flag is true when the data is too long.
	MixedContentType *security.MixedContentType `json:"mixedContentType,omitempty"` // The mixed content type of the request.
	InitialPriority  ResourcePriority           `json:"initialPriority"`            // Priority of the resource request at the time request is sent.
	// ReferrerPolicy The referrer policy of the request, as defined in
	// https://www.w3.org/TR/referrer-policy/
	//
	// Values: "unsafe-url", "no-referrer-when-downgrade", "no-referrer", "origin", "origin-when-cross-origin", "same-origin", "strict-origin", "strict-origin-when-cross-origin".
	ReferrerPolicy string `json:"referrerPolicy"`
	IsLinkPreload  *bool  `json:"isLinkPreload,omitempty"` // Whether is loaded via link preload.
}

// SignedCertificateTimestamp Details of a signed certificate timestamp (SCT).
type SignedCertificateTimestamp struct {
	Status             string         `json:"status"`             // Validation status.
	Origin             string         `json:"origin"`             // Origin.
	LogDescription     string         `json:"logDescription"`     // Log name / description.
	LogID              string         `json:"logId"`              // Log ID.
	Timestamp          TimeSinceEpoch `json:"timestamp"`          // Issuance date.
	HashAlgorithm      string         `json:"hashAlgorithm"`      // Hash algorithm.
	SignatureAlgorithm string         `json:"signatureAlgorithm"` // Signature algorithm.
	SignatureData      string         `json:"signatureData"`      // Signature data.
}

// SecurityDetails Security details about a request.
type SecurityDetails struct {
	Protocol                          string                            `json:"protocol"`                          // Protocol name (e.g. "TLS 1.2" or "QUIC").
	KeyExchange                       string                            `json:"keyExchange"`                       // Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchangeGroup                  *string                           `json:"keyExchangeGroup,omitempty"`        // (EC)DH group used by the connection, if applicable.
	Cipher                            string                            `json:"cipher"`                            // Cipher name.
	MAC                               *string                           `json:"mac,omitempty"`                     // TLS MAC. Note that AEAD ciphers do not have separate MACs.
	CertificateID                     security.CertificateID            `json:"certificateId"`                     // Certificate ID value.
	SubjectName                       string                            `json:"subjectName"`                       // Certificate subject name.
	SanList                           []string                          `json:"sanList"`                           // Subject Alternative Name (SAN) DNS names and IP addresses.
	Issuer                            string                            `json:"issuer"`                            // Name of the issuing CA.
	ValidFrom                         TimeSinceEpoch                    `json:"validFrom"`                         // Certificate valid from date.
	ValidTo                           TimeSinceEpoch                    `json:"validTo"`                           // Certificate valid to (expiration) date
	SignedCertificateTimestampList    []SignedCertificateTimestamp      `json:"signedCertificateTimestampList"`    // List of signed certificate timestamps (SCTs).
	CertificateTransparencyCompliance CertificateTransparencyCompliance `json:"certificateTransparencyCompliance"` // Whether the request complied with Certificate Transparency policy
}

// CertificateTransparencyCompliance Whether the request complied with
// Certificate Transparency policy.
type CertificateTransparencyCompliance string

// CertificateTransparencyCompliance as enums.
const (
	CertificateTransparencyComplianceNotSet       CertificateTransparencyCompliance = ""
	CertificateTransparencyComplianceUnknown      CertificateTransparencyCompliance = "unknown"
	CertificateTransparencyComplianceNotCompliant CertificateTransparencyCompliance = "not-compliant"
	CertificateTransparencyComplianceCompliant    CertificateTransparencyCompliance = "compliant"
)

func (e CertificateTransparencyCompliance) Valid() bool {
	switch e {
	case "unknown", "not-compliant", "compliant":
		return true
	default:
		return false
	}
}

func (e CertificateTransparencyCompliance) String() string {
	return string(e)
}

// BlockedReason The reason why request was blocked.
type BlockedReason string

// BlockedReason as enums.
const (
	BlockedReasonNotSet            BlockedReason = ""
	BlockedReasonOther             BlockedReason = "other"
	BlockedReasonCsp               BlockedReason = "csp"
	BlockedReasonMixedContent      BlockedReason = "mixed-content"
	BlockedReasonOrigin            BlockedReason = "origin"
	BlockedReasonInspector         BlockedReason = "inspector"
	BlockedReasonSubresourceFilter BlockedReason = "subresource-filter"
	BlockedReasonContentType       BlockedReason = "content-type"
	BlockedReasonCollapsedByClient BlockedReason = "collapsed-by-client"
)

func (e BlockedReason) Valid() bool {
	switch e {
	case "other", "csp", "mixed-content", "origin", "inspector", "subresource-filter", "content-type", "collapsed-by-client":
		return true
	default:
		return false
	}
}

func (e BlockedReason) String() string {
	return string(e)
}

// Response HTTP response data.
type Response struct {
	URL                string           `json:"url"`                          // Response URL. This URL can be different from CachedResource.url in case of redirect.
	Status             int              `json:"status"`                       // HTTP response status code.
	StatusText         string           `json:"statusText"`                   // HTTP response status text.
	Headers            Headers          `json:"headers"`                      // HTTP response headers.
	HeadersText        *string          `json:"headersText,omitempty"`        // HTTP response headers text.
	MimeType           string           `json:"mimeType"`                     // Resource mimeType as determined by the browser.
	RequestHeaders     Headers          `json:"requestHeaders,omitempty"`     // Refined HTTP request headers that were actually transmitted over the network.
	RequestHeadersText *string          `json:"requestHeadersText,omitempty"` // HTTP request headers text.
	ConnectionReused   bool             `json:"connectionReused"`             // Specifies whether physical connection was actually reused for this request.
	ConnectionID       float64          `json:"connectionId"`                 // Physical connection id that was actually used for this request.
	RemoteIPAddress    *string          `json:"remoteIPAddress,omitempty"`    // Remote IP address.
	RemotePort         *int             `json:"remotePort,omitempty"`         // Remote port.
	FromDiskCache      *bool            `json:"fromDiskCache,omitempty"`      // Specifies that the request was served from the disk cache.
	FromServiceWorker  *bool            `json:"fromServiceWorker,omitempty"`  // Specifies that the request was served from the ServiceWorker.
	FromPrefetchCache  *bool            `json:"fromPrefetchCache,omitempty"`  // Specifies that the request was served from the prefetch cache.
	EncodedDataLength  float64          `json:"encodedDataLength"`            // Total number of bytes received for this request so far.
	Timing             *ResourceTiming  `json:"timing,omitempty"`             // Timing information for the given request.
	Protocol           *string          `json:"protocol,omitempty"`           // Protocol used to fetch this request.
	SecurityState      security.State   `json:"securityState"`                // Security state of the request resource.
	SecurityDetails    *SecurityDetails `json:"securityDetails,omitempty"`    // Security details for the request.
}

// WebSocketRequest WebSocket request data.
type WebSocketRequest struct {
	Headers Headers `json:"headers"` // HTTP request headers.
}

// WebSocketResponse WebSocket response data.
type WebSocketResponse struct {
	Status             int     `json:"status"`                       // HTTP response status code.
	StatusText         string  `json:"statusText"`                   // HTTP response status text.
	Headers            Headers `json:"headers"`                      // HTTP response headers.
	HeadersText        *string `json:"headersText,omitempty"`        // HTTP response headers text.
	RequestHeaders     Headers `json:"requestHeaders,omitempty"`     // HTTP request headers.
	RequestHeadersText *string `json:"requestHeadersText,omitempty"` // HTTP request headers text.
}

// WebSocketFrame WebSocket message data. This represents an entire WebSocket
// message, not just a fragmented frame as the name suggests.
type WebSocketFrame struct {
	Opcode      float64 `json:"opcode"`      // WebSocket message opcode.
	Mask        bool    `json:"mask"`        // WebSocket message mask.
	PayloadData string  `json:"payloadData"` // WebSocket message payload data. If the opcode is 1, this is a text message and payloadData is a UTF-8 string. If the opcode isn't 1, then payloadData is a base64 encoded string representing binary data.
}

// CachedResource Information about the cached resource.
type CachedResource struct {
	URL      string       `json:"url"`                // Resource URL. This is the url of the original network request.
	Type     ResourceType `json:"type"`               // Type of this resource.
	Response *Response    `json:"response,omitempty"` // Cached response data.
	BodySize float64      `json:"bodySize"`           // Cached response body size.
}

// Initiator Information about the request initiator.
type Initiator struct {
	// Type Type of this initiator.
	//
	// Values: "parser", "script", "preload", "SignedExchange", "other".
	Type       string              `json:"type"`
	Stack      *runtime.StackTrace `json:"stack,omitempty"`      // Initiator JavaScript stack trace, set for Script only.
	URL        *string             `json:"url,omitempty"`        // Initiator URL, set for Parser type or for Script type (when script is importing module) or for SignedExchange type.
	LineNumber *float64            `json:"lineNumber,omitempty"` // Initiator line number, set for Parser type or for Script type (when script is importing module) (0-based).
}

// Cookie Cookie object
type Cookie struct {
	Name     string         `json:"name"`               // Cookie name.
	Value    string         `json:"value"`              // Cookie value.
	Domain   string         `json:"domain"`             // Cookie domain.
	Path     string         `json:"path"`               // Cookie path.
	Expires  float64        `json:"expires"`            // Cookie expiration date as the number of seconds since the UNIX epoch.
	Size     int            `json:"size"`               // Cookie size.
	HTTPOnly bool           `json:"httpOnly"`           // True if cookie is http-only.
	Secure   bool           `json:"secure"`             // True if cookie is secure.
	Session  bool           `json:"session"`            // True in case of session cookie.
	SameSite CookieSameSite `json:"sameSite,omitempty"` // Cookie SameSite type.
}

// SetCookieBlockedReason Types of reasons why a cookie may not be stored from
// a response.
//
// Note: This type is experimental.
type SetCookieBlockedReason string

// SetCookieBlockedReason as enums.
const (
	SetCookieBlockedReasonNotSet                          SetCookieBlockedReason = ""
	SetCookieBlockedReasonSecureOnly                      SetCookieBlockedReason = "SecureOnly"
	SetCookieBlockedReasonSameSiteStrict                  SetCookieBlockedReason = "SameSiteStrict"
	SetCookieBlockedReasonSameSiteLax                     SetCookieBlockedReason = "SameSiteLax"
	SetCookieBlockedReasonSameSiteUnspecifiedTreatedAsLax SetCookieBlockedReason = "SameSiteUnspecifiedTreatedAsLax"
	SetCookieBlockedReasonSameSiteNoneInsecure            SetCookieBlockedReason = "SameSiteNoneInsecure"
	SetCookieBlockedReasonUserPreferences                 SetCookieBlockedReason = "UserPreferences"
	SetCookieBlockedReasonSyntaxError                     SetCookieBlockedReason = "SyntaxError"
	SetCookieBlockedReasonSchemeNotSupported              SetCookieBlockedReason = "SchemeNotSupported"
	SetCookieBlockedReasonOverwriteSecure                 SetCookieBlockedReason = "OverwriteSecure"
	SetCookieBlockedReasonInvalidDomain                   SetCookieBlockedReason = "InvalidDomain"
	SetCookieBlockedReasonInvalidPrefix                   SetCookieBlockedReason = "InvalidPrefix"
	SetCookieBlockedReasonUnknownError                    SetCookieBlockedReason = "UnknownError"
)

func (e SetCookieBlockedReason) Valid() bool {
	switch e {
	case "SecureOnly", "SameSiteStrict", "SameSiteLax", "SameSiteUnspecifiedTreatedAsLax", "SameSiteNoneInsecure", "UserPreferences", "SyntaxError", "SchemeNotSupported", "OverwriteSecure", "InvalidDomain", "InvalidPrefix", "UnknownError":
		return true
	default:
		return false
	}
}

func (e SetCookieBlockedReason) String() string {
	return string(e)
}

// CookieBlockedReason Types of reasons why a cookie may not be sent with a
// request.
//
// Note: This type is experimental.
type CookieBlockedReason string

// CookieBlockedReason as enums.
const (
	CookieBlockedReasonNotSet                          CookieBlockedReason = ""
	CookieBlockedReasonSecureOnly                      CookieBlockedReason = "SecureOnly"
	CookieBlockedReasonNotOnPath                       CookieBlockedReason = "NotOnPath"
	CookieBlockedReasonDomainMismatch                  CookieBlockedReason = "DomainMismatch"
	CookieBlockedReasonSameSiteStrict                  CookieBlockedReason = "SameSiteStrict"
	CookieBlockedReasonSameSiteLax                     CookieBlockedReason = "SameSiteLax"
	CookieBlockedReasonSameSiteUnspecifiedTreatedAsLax CookieBlockedReason = "SameSiteUnspecifiedTreatedAsLax"
	CookieBlockedReasonSameSiteNoneInsecure            CookieBlockedReason = "SameSiteNoneInsecure"
	CookieBlockedReasonUserPreferences                 CookieBlockedReason = "UserPreferences"
	CookieBlockedReasonUnknownError                    CookieBlockedReason = "UnknownError"
)

func (e CookieBlockedReason) Valid() bool {
	switch e {
	case "SecureOnly", "NotOnPath", "DomainMismatch", "SameSiteStrict", "SameSiteLax", "SameSiteUnspecifiedTreatedAsLax", "SameSiteNoneInsecure", "UserPreferences", "UnknownError":
		return true
	default:
		return false
	}
}

func (e CookieBlockedReason) String() string {
	return string(e)
}

// BlockedSetCookieWithReason A cookie which was not stored from a response
// with the corresponding reason.
//
// Note: This type is experimental.
type BlockedSetCookieWithReason struct {
	BlockedReasons []SetCookieBlockedReason `json:"blockedReasons"`   // The reason(s) this cookie was blocked.
	CookieLine     string                   `json:"cookieLine"`       // The string representing this individual cookie as it would appear in the header. This is not the entire "cookie" or "set-cookie" header which could have multiple cookies.
	Cookie         *Cookie                  `json:"cookie,omitempty"` // The cookie object which represents the cookie which was not stored. It is optional because sometimes complete cookie information is not available, such as in the case of parsing errors.
}

// BlockedCookieWithReason A cookie with was not sent with a request with the
// corresponding reason.
//
// Note: This type is experimental.
type BlockedCookieWithReason struct {
	BlockedReasons []CookieBlockedReason `json:"blockedReasons"` // The reason(s) the cookie was blocked.
	Cookie         Cookie                `json:"cookie"`         // The cookie object representing the cookie which was not sent.
}

// CookieParam Cookie parameter object
type CookieParam struct {
	Name     string         `json:"name"`               // Cookie name.
	Value    string         `json:"value"`              // Cookie value.
	URL      *string        `json:"url,omitempty"`      // The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
	Domain   *string        `json:"domain,omitempty"`   // Cookie domain.
	Path     *string        `json:"path,omitempty"`     // Cookie path.
	Secure   *bool          `json:"secure,omitempty"`   // True if cookie is secure.
	HTTPOnly *bool          `json:"httpOnly,omitempty"` // True if cookie is http-only.
	SameSite CookieSameSite `json:"sameSite,omitempty"` // Cookie SameSite type.
	Expires  TimeSinceEpoch `json:"expires,omitempty"`  // Cookie expiration date, session cookie if not set
}

// AuthChallenge Authorization challenge for HTTP status code 401 or 407.
//
// Note: This type is experimental.
type AuthChallenge struct {
	// Source Source of the authentication challenge.
	//
	// Values: "Server", "Proxy".
	Source *string `json:"source,omitempty"`
	Origin string  `json:"origin"` // Origin of the challenger.
	Scheme string  `json:"scheme"` // The authentication scheme used, such as basic or digest
	Realm  string  `json:"realm"`  // The realm of the challenge. May be empty.
}

// AuthChallengeResponse Response to an AuthChallenge.
//
// Note: This type is experimental.
type AuthChallengeResponse struct {
	// Response The decision on what to do in response to the
	// authorization challenge. Default means deferring to the default
	// behavior of the net stack, which will likely either the Cancel
	// authentication or display a popup dialog box.
	//
	// Values: "Default", "CancelAuth", "ProvideCredentials".
	Response string  `json:"response"`
	Username *string `json:"username,omitempty"` // The username to provide, possibly empty. Should only be set if response is ProvideCredentials.
	Password *string `json:"password,omitempty"` // The password to provide, possibly empty. Should only be set if response is ProvideCredentials.
}

// InterceptionStage Stages of the interception to begin intercepting. Request
// will intercept before the request is sent. Response will intercept after the
// response is received.
//
// Note: This type is experimental.
type InterceptionStage string

// InterceptionStage as enums.
const (
	InterceptionStageNotSet          InterceptionStage = ""
	InterceptionStageRequest         InterceptionStage = "Request"
	InterceptionStageHeadersReceived InterceptionStage = "HeadersReceived"
)

func (e InterceptionStage) Valid() bool {
	switch e {
	case "Request", "HeadersReceived":
		return true
	default:
		return false
	}
}

func (e InterceptionStage) String() string {
	return string(e)
}

// RequestPattern Request pattern for interception.
//
// Note: This type is experimental.
type RequestPattern struct {
	URLPattern        *string           `json:"urlPattern,omitempty"`        // Wildcards ('*' -> zero or more, '?' -> exactly one) are allowed. Escape character is backslash. Omitting is equivalent to "*".
	ResourceType      ResourceType      `json:"resourceType,omitempty"`      // If set, only requests for matching resource types will be intercepted.
	InterceptionStage InterceptionStage `json:"interceptionStage,omitempty"` // Stage at which to begin intercepting requests. Default is Request.
}

// SignedExchangeSignature Information about a signed exchange signature.
// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#rfc.section.3.1
//
// Note: This type is experimental.
type SignedExchangeSignature struct {
	Label        string   `json:"label"`                  // Signed exchange signature label.
	Signature    string   `json:"signature"`              // The hex string of signed exchange signature.
	Integrity    string   `json:"integrity"`              // Signed exchange signature integrity.
	CertURL      *string  `json:"certUrl,omitempty"`      // Signed exchange signature cert Url.
	CertSha256   *string  `json:"certSha256,omitempty"`   // The hex string of signed exchange signature cert sha256.
	ValidityURL  string   `json:"validityUrl"`            // Signed exchange signature validity Url.
	Date         int      `json:"date"`                   // Signed exchange signature date.
	Expires      int      `json:"expires"`                // Signed exchange signature expires.
	Certificates []string `json:"certificates,omitempty"` // The encoded certificates.
}

// SignedExchangeHeader Information about a signed exchange header.
// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cbor-representation
//
// Note: This type is experimental.
type SignedExchangeHeader struct {
	RequestURL      string                    `json:"requestUrl"`      // Signed exchange request URL.
	ResponseCode    int                       `json:"responseCode"`    // Signed exchange response code.
	ResponseHeaders Headers                   `json:"responseHeaders"` // Signed exchange response headers.
	Signatures      []SignedExchangeSignature `json:"signatures"`      // Signed exchange response signature.
	HeaderIntegrity string                    `json:"headerIntegrity"` // Signed exchange header integrity hash in the form of "sha256-<base64-hash-value>".
}

// SignedExchangeErrorField Field type for a signed exchange related error.
//
// Note: This type is experimental.
type SignedExchangeErrorField string

// SignedExchangeErrorField as enums.
const (
	SignedExchangeErrorFieldNotSet               SignedExchangeErrorField = ""
	SignedExchangeErrorFieldSignatureSig         SignedExchangeErrorField = "signatureSig"
	SignedExchangeErrorFieldSignatureIntegrity   SignedExchangeErrorField = "signatureIntegrity"
	SignedExchangeErrorFieldSignatureCertURL     SignedExchangeErrorField = "signatureCertUrl"
	SignedExchangeErrorFieldSignatureCertSha256  SignedExchangeErrorField = "signatureCertSha256"
	SignedExchangeErrorFieldSignatureValidityURL SignedExchangeErrorField = "signatureValidityUrl"
	SignedExchangeErrorFieldSignatureTimestamps  SignedExchangeErrorField = "signatureTimestamps"
)

func (e SignedExchangeErrorField) Valid() bool {
	switch e {
	case "signatureSig", "signatureIntegrity", "signatureCertUrl", "signatureCertSha256", "signatureValidityUrl", "signatureTimestamps":
		return true
	default:
		return false
	}
}

func (e SignedExchangeErrorField) String() string {
	return string(e)
}

// SignedExchangeError Information about a signed exchange response.
//
// Note: This type is experimental.
type SignedExchangeError struct {
	Message        string                   `json:"message"`                  // Error message.
	SignatureIndex *int                     `json:"signatureIndex,omitempty"` // The index of the signature which caused the error.
	ErrorField     SignedExchangeErrorField `json:"errorField,omitempty"`     // The field which caused the error.
}

// SignedExchangeInfo Information about a signed exchange response.
//
// Note: This type is experimental.
type SignedExchangeInfo struct {
	OuterResponse   Response              `json:"outerResponse"`             // The outer response of signed HTTP exchange which was received from network.
	Header          *SignedExchangeHeader `json:"header,omitempty"`          // Information about the signed exchange header.
	SecurityDetails *SecurityDetails      `json:"securityDetails,omitempty"` // Security details for the signed exchange header.
	Errors          []SignedExchangeError `json:"errors,omitempty"`          // Errors occurred while handling the signed exchagne.
}
