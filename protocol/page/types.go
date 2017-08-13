// Code generated by cdpgen. DO NOT EDIT.

package page

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mafredri/cdp/protocol/network"
)

// FrameResourceTree Information about the Frame hierarchy along with their cached resources.
//
// Note: This type is experimental.
type FrameResourceTree struct {
	Frame       Frame               `json:"frame"`                 // Frame information for this tree item.
	ChildFrames []FrameResourceTree `json:"childFrames,omitempty"` // Child frames.
	Resources   []FrameResource     `json:"resources"`             // Information about frame resources.
}

// ScriptIdentifier Unique script identifier.
//
// Note: This type is experimental.
type ScriptIdentifier string

// TransitionType Transition type.
//
// Note: This type is experimental.
type TransitionType int

// TransitionType as enums.
const (
	TransitionTypeNotSet TransitionType = iota
	TransitionTypeLink
	TransitionTypeTyped
	TransitionTypeAutoBookmark
	TransitionTypeAutoSubframe
	TransitionTypeManualSubframe
	TransitionTypeGenerated
	TransitionTypeAutoToplevel
	TransitionTypeFormSubmit
	TransitionTypeReload
	TransitionTypeKeyword
	TransitionTypeKeywordGenerated
	TransitionTypeOther
)

// Valid returns true if enum is set.
func (e TransitionType) Valid() bool {
	return e >= 1 && e <= 12
}

func (e TransitionType) String() string {
	switch e {
	case 0:
		return "TransitionTypeNotSet"
	case 1:
		return "link"
	case 2:
		return "typed"
	case 3:
		return "auto_bookmark"
	case 4:
		return "auto_subframe"
	case 5:
		return "manual_subframe"
	case 6:
		return "generated"
	case 7:
		return "auto_toplevel"
	case 8:
		return "form_submit"
	case 9:
		return "reload"
	case 10:
		return "keyword"
	case 11:
		return "keyword_generated"
	case 12:
		return "other"
	}
	return fmt.Sprintf("TransitionType(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e TransitionType) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("page.TransitionType: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *TransitionType) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"link\"":
		*e = 1
	case "\"typed\"":
		*e = 2
	case "\"auto_bookmark\"":
		*e = 3
	case "\"auto_subframe\"":
		*e = 4
	case "\"manual_subframe\"":
		*e = 5
	case "\"generated\"":
		*e = 6
	case "\"auto_toplevel\"":
		*e = 7
	case "\"form_submit\"":
		*e = 8
	case "\"reload\"":
		*e = 9
	case "\"keyword\"":
		*e = 10
	case "\"keyword_generated\"":
		*e = 11
	case "\"other\"":
		*e = 12
	default:
		return fmt.Errorf("page.TransitionType: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}

// NavigationEntry Navigation history entry.
//
// Note: This type is experimental.
type NavigationEntry struct {
	ID             int            `json:"id"`             // Unique id of the navigation history entry.
	URL            string         `json:"url"`            // URL of the navigation history entry.
	UserTypedURL   string         `json:"userTypedURL"`   // URL that the user typed in the url bar.
	Title          string         `json:"title"`          // Title of the navigation history entry.
	TransitionType TransitionType `json:"transitionType"` // Transition type.
}

// ScreencastFrameMetadata Screencast frame metadata.
//
// Note: This type is experimental.
type ScreencastFrameMetadata struct {
	// OffsetTop Top offset in DIP.
	//
	// Note: This property is experimental.
	OffsetTop float64 `json:"offsetTop"`
	// PageScaleFactor Page scale factor.
	//
	// Note: This property is experimental.
	PageScaleFactor float64 `json:"pageScaleFactor"`
	// DeviceWidth Device screen width in DIP.
	//
	// Note: This property is experimental.
	DeviceWidth float64 `json:"deviceWidth"`
	// DeviceHeight Device screen height in DIP.
	//
	// Note: This property is experimental.
	DeviceHeight float64 `json:"deviceHeight"`
	// ScrollOffsetX Position of horizontal scroll in CSS pixels.
	//
	// Note: This property is experimental.
	ScrollOffsetX float64 `json:"scrollOffsetX"`
	// ScrollOffsetY Position of vertical scroll in CSS pixels.
	//
	// Note: This property is experimental.
	ScrollOffsetY float64 `json:"scrollOffsetY"`
	// Timestamp Frame swap timestamp.
	//
	// Note: This property is experimental.
	Timestamp network.TimeSinceEpoch `json:"timestamp,omitempty"`
}

// DialogType Javascript dialog type.
//
// Note: This type is experimental.
type DialogType int

// DialogType as enums.
const (
	DialogTypeNotSet DialogType = iota
	DialogTypeAlert
	DialogTypeConfirm
	DialogTypePrompt
	DialogTypeBeforeunload
)

// Valid returns true if enum is set.
func (e DialogType) Valid() bool {
	return e >= 1 && e <= 4
}

func (e DialogType) String() string {
	switch e {
	case 0:
		return "DialogTypeNotSet"
	case 1:
		return "alert"
	case 2:
		return "confirm"
	case 3:
		return "prompt"
	case 4:
		return "beforeunload"
	}
	return fmt.Sprintf("DialogType(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e DialogType) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("page.DialogType: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *DialogType) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"alert\"":
		*e = 1
	case "\"confirm\"":
		*e = 2
	case "\"prompt\"":
		*e = 3
	case "\"beforeunload\"":
		*e = 4
	default:
		return fmt.Errorf("page.DialogType: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}

// AppManifestError Error while paring app manifest.
//
// Note: This type is experimental.
type AppManifestError struct {
	Message  string `json:"message"`  // Error message.
	Critical int    `json:"critical"` // If criticial, this is a non-recoverable parse error.
	Line     int    `json:"line"`     // Error line.
	Column   int    `json:"column"`   // Error column.
}

// NavigationResponse Proceed: allow the navigation; Cancel: cancel the navigation; CancelAndIgnore: cancels the navigation and makes the requester of the navigation acts like the request was never made.
//
// Note: This type is experimental.
type NavigationResponse int

// NavigationResponse as enums.
const (
	NavigationResponseNotSet NavigationResponse = iota
	NavigationResponseProceed
	NavigationResponseCancel
	NavigationResponseCancelAndIgnore
)

// Valid returns true if enum is set.
func (e NavigationResponse) Valid() bool {
	return e >= 1 && e <= 3
}

func (e NavigationResponse) String() string {
	switch e {
	case 0:
		return "NavigationResponseNotSet"
	case 1:
		return "Proceed"
	case 2:
		return "Cancel"
	case 3:
		return "CancelAndIgnore"
	}
	return fmt.Sprintf("NavigationResponse(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e NavigationResponse) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("page.NavigationResponse: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *NavigationResponse) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"Proceed\"":
		*e = 1
	case "\"Cancel\"":
		*e = 2
	case "\"CancelAndIgnore\"":
		*e = 3
	default:
		return fmt.Errorf("page.NavigationResponse: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}

// LayoutViewport Layout viewport position and dimensions.
//
// Note: This type is experimental.
type LayoutViewport struct {
	PageX        int `json:"pageX"`        // Horizontal offset relative to the document (CSS pixels).
	PageY        int `json:"pageY"`        // Vertical offset relative to the document (CSS pixels).
	ClientWidth  int `json:"clientWidth"`  // Width (CSS pixels), excludes scrollbar if present.
	ClientHeight int `json:"clientHeight"` // Height (CSS pixels), excludes scrollbar if present.
}

// VisualViewport Visual viewport position, dimensions, and scale.
//
// Note: This type is experimental.
type VisualViewport struct {
	OffsetX      float64 `json:"offsetX"`      // Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetY      float64 `json:"offsetY"`      // Vertical offset relative to the layout viewport (CSS pixels).
	PageX        float64 `json:"pageX"`        // Horizontal offset relative to the document (CSS pixels).
	PageY        float64 `json:"pageY"`        // Vertical offset relative to the document (CSS pixels).
	ClientWidth  float64 `json:"clientWidth"`  // Width (CSS pixels), excludes scrollbar if present.
	ClientHeight float64 `json:"clientHeight"` // Height (CSS pixels), excludes scrollbar if present.
	Scale        float64 `json:"scale"`        // Scale relative to the ideal viewport (size at width=device-width).
}

// Viewport Viewport for capturing screenshot.
//
// Note: This type is experimental.
type Viewport struct {
	X      float64 `json:"x"`      // X offset in CSS pixels.
	Y      float64 `json:"y"`      // Y offset in CSS pixels
	Width  float64 `json:"width"`  // Rectangle width in CSS pixels
	Height float64 `json:"height"` // Rectangle height in CSS pixels
	Scale  float64 `json:"scale"`  // Page scale factor.
}