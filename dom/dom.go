// Package dom is a thin wrapper of some useful dom functions for gopherjs.
package dom

import (
	"github.com/gopherjs/gopherjs/js"
)

const (
	// Wheel Delta
	DeltaPixel = 0
	DeltaLine  = 1
	DeltaPage  = 2

	// Event Phase
	EvPhaseNone      = 0
	EvPhaseCapturing = 1
	EvPhaseAtTarget  = 2
	EvPhaseBubbling  = 3
)

// Events Types
const (
	// Window Event Attributes
	// Window Events triggered for a window object and apply in <body> tag
	// Attributes	Value	Description	In HTML5?
	EvtAfterprint       = "afterprint"       //	Script is run after the document is printed	NEW
	EvtBeforeprint      = "beforeprint"      //	Script is run before the document is printed	NEW
	EvtBeforeunload     = "beforeunload"     //	Script is run before the document is unloaded	NEW
	EvtError            = "error"            //	Script is run when any error occur	NEW
	EvtHaschange        = "haschange"        //	Script is run when document has changed	NEW
	EvtLoad             = "load"             //	Event fires after the page loading finished
	EvtDOMContentLoaded = "DOMContentLoaded" // Event fires after the page DOM is ready
	EvtMessage          = "message"          //	Script is run when document goes in offline	NEW
	EvtOffline          = "offline"          //	Script is run when document comes in Event = "line	NEW
	EvtPagehide         = "pagehide"         //	Script is run when document window is hidden	NEW
	EvtPageshow         = "pageshow"         //	Script is run when document window become visible	NEW
	EvtPopstate         = "popstate"         //	Script is run when document window history changes	NEW
	EvtRedo             = "redo"             //	Script is run when document perform redo	NEW
	EvtResize           = "resize"           //	Event fires when browser window is resized	NEW
	EvtStorage          = "storage"          //	Script is run when web storage area is updated	NEW
	EvtUndo             = "undo"             //	Script is run when document performs undo	NEW
	EvtUnload           = "unload"           //	Event fires when browser window has been closed

	// Form Events
	// Form Events triggered by perform some action inside HTML form elements.
	// Attributes	Value	Description	In HTML5?
	EvtBlur        = "blur"        //	Event fire when element loses focus
	EvtChange      = "change"      //	Event fire when element value is changed
	EvtContextmenu = "contextmenu" //	Event fire when context menu is triggered	NEW
	EvtFocus       = "focus"       //	Event fire when element gets focus
	EvtFormchange  = "formchange"  //	Event fire when form changes	NEW
	EvtForminput   = "forminput"   //	Event fire when form get input field
	EvtInput       = "input"       //	Event fire when element get input field	NEW
	EvtInvalid     = "invalid"     //	Event fire when element is invalid	NEW
	EvtReset       = "reset"       //	Event fire when clicked on form reset button	REMOVE
	EvtSelect      = "select"      //	Event fire after allow to select text in an element
	EvtSubmit      = "Submit"      //	Event fire when form is submitted

	// Keyboard Events
	// Attributes	Value	Description	In HTML5?
	EvtKeydown  = "keydown"  //	Event fire when pressing a key
	EvtKeypress = "keypress" //	Event fire when press a key
	EvtKeyup    = "keyup"    //	Event fire when releases a key

	// Mouse Events
	// Mouse Events triggered by mouse action.
	// Attributes	Value	Description	In HTML5?
	EvtClick      = "click"      //	Event fire when mouse click on element
	EvtDblclick   = "dblclick"   //	Event fire when mouse double click on element
	EvtDrag       = "drag"       //	Script is run when element is dragged	NEW
	EvtDragend    = "dragend"    //	Script is run at end of drag operation	NEW
	EvtDragenter  = "dragenter"  //	Script is run when element has dragged to a valid drop target	NEW
	EvtDragleave  = "dragleave"  //	Script is run when element leaves valid drop target	NEW
	EvtDragover   = "dragover"   //	Script is run when element is dragged over on valid drop target	NEW
	EvtDragstart  = "dragstart"  //	Script is run at start of drag operation	NEW
	EvtDrop       = "drop"       //	Script is run when dragged element is dropped	NEW
	EvtMousedown  = "mousedown"  //	Event fire when mouse button is pressed down on element
	EvtMousemove  = "mousemove"  //	Event fire when mouse pointer moves over an element
	EvtMouseout   = "mouseout"   //	Event fire when mouse pointer moves out an element
	EvtMouseover  = "mouseover"  //	Event fire when mouse pointer moves over on element
	EvtMouseup    = "mouseup"    //	Event fire when mouse button is released over an element
	EvtMousewheel = "mousewheel" //	Event fire when mouse wheel being rotated	NEW
	EvtScroll     = "scroll"     //	Event fire when element scrollbar being scrolled	NEW

	// Media Events
	// Media Events triggered by common media elements like <img>, <audio>, <embed>, <object>, and <video>.
	// Attributes	Value	Description	In HTML5?
	EvtAbort            = "abort"            //	Script is run when element is abort
	EvtCanplay          = "canplay"          //	Script is run when file is ready for start playing	NEW
	EvtCanplaythrough   = "canplaythrough"   //	Script is run when file is played all way without pausing for buffering	NEW
	EvtDurationchange   = "durationchange"   //	Script is run when media length changes	NEW
	EvtEmptied          = "emptied"          //	Script is run when something unavailable/disconnects	NEW
	EvtEnded            = "ended"            //	Script is run when media has reach to end position	NEW
	EvtLoadeddata       = "loadeddata"       //	Script is run when media is loaded	NEW
	EvtLoadedmetadata   = "loadedmetadata"   //	Script is run when meta data are loaded	NEW
	EvtLoadstart        = "loadstart"        //	Script is run when file being loaded	NEW
	EvtPause            = "pause"            //	Script is run when media is paused	NEW
	EvtPlay             = "play"             //	Script is run when media is ready to start playing	NEW
	EvtPlaying          = "playing"          //	Script is run when media is actually start for playing	NEW
	EvtProgress         = "progress"         //	Script is run when browser is process of getting media data	NEW
	EvtRatechange       = "ratechange"       //	Script is run when playback rate changes	NEW
	EvtReadystatechange = "readystatechange" //	Script is run when ready state changes for each time	NEW
	EvtSeeked           = "seeked"           //	Script is run when seeking attribute value set to false, that indicate seeking has ended	NEW
	EvtSeeking          = "seeking"          //	Script is run when seeking attribute value set to true, that indicate seeking has active	NEW
	EvtStalled          = "stalled"          //	Script is run when browser is unable to fetch media data for any reason	NEW
	EvtSuspend          = "suspend"          //	Script is run when fetching media data is stopped before it is completely loaded for any reason	NEW
	EvtTimeupdate       = "timeupdate"       //	Script is run when playing position has changed	NEW
	EvtVolumechange     = "volumechange"     //	Script is run each time volume is changed	NEW
	EvtWaiting          = "waiting"          //	Script is run when media has paused(for buffer more data)	NEW
	// onerror            //	Script is run when error occurs file loaded time	NEW

)

type CSSStyleDeclaration struct {
	*js.Object
	// Textual representation of the declaration block. Setting this attribute changes the style.
	CssText string `js:"cssText"`
	// CSSStyleDeclaration.length

	// The number of properties. See the item method below.
	Length int `js:"length"`
	// CSSStyleDeclaration.parentRule

	// The containing CssRule.
	ParentRule *CSSStyleDeclaration `js:"parentRule"`

	// funcs
	RemoveProperty      func(name string)                            `js:"removeProperty"`
	GetPropertyValue    func(name string) string                     `js:"getPropertyValue"`
	GetPropertyPriority func(name string) string                     `js:"getPropertyPriority"`
	SetProperty         func(name, value string, priority ...string) `js:"setProperty"`
}

func (css *CSSStyleDeclaration) ToMap() map[string]string {
	m := make(map[string]string)
	N := css.Length
	for i := 0; i < N; i++ {
		name := css.Call("index", i).String()
		value := css.Call("getPropertyValue").String()
		m[name] = value
	}
	return m
}

func GetComputedStyle(e *Element) *CSSStyleDeclaration {
	return &CSSStyleDeclaration{
		Object: js.Global.Get("document").Get("defaultView").Call("getComputedStyle", e.Object),
	}
}

type Element struct {
	*js.Object
	// basic attr
	Id              string `js:"id"`
	InnerHTML       string `js:"innerHTML"`
	InnerText       string `js:"innerText"`
	TagName         string `js:"tagName"`
	ContentEditable bool   `js:"contentEditable"`
	// width & height, not all element supoort these attributes,
	// use Style to set width/height
	Width  int `js:"width"`
	Height int `js:"height"`
	// window size, for window object
	InnerWidth  int `js:"innerWidth"`
	InnerHeight int `js:"innerHeight"`
	// dom
	PreviousElementSibling *Element `js:"previousElementSibling"`
	NextElementSibling     *Element `js:"nextElementSibling"`
	FirstElementChild      *Element `js:"firstElementChild"`
	LastElementChild       *Element `js:"lastElementChild"`
	// img, script
	Src string `js:"src"`
	// style
	Style     *CSSStyleDeclaration `js:"style"`
	ClassName string               `js:"className"`
	ClassList []string             `js:"classList"`

	// funcs
	SetAttribute    func(attr string, val interface{}) `js:"setAttribute"`
	GetAttribute    func(attr string) *js.Object       `js:"getAttribute"`
	RemoveAttribute func(attr string)                  `js:"removeAttribute"`

	AppendChild func(child *Element) `js:"appendChild"`
	RemoveChild func(child *Element) `js:"removeChild"`
	Remove      func()               `js:"remove"`

	GetElementsByTagName func(tagName string) *HTMLCollection `js:"getElementsByTagName"`
	QuerySelector        func(sel string) *Element            `js:"querySelector"`
	QuerySelectorAll     func(sel string) *HTMLCollection     `js:"querySelectorAll"`

	// Event handling

	// Registers an event handler to a specific event type on the element.
	//   If true, useCapture indicates that the user wishes to initiate capture.
	//   After initiating capture, all events of the specified type will be dispatched to the registered listener before being dispatched to any EventTarget beneath it in the DOM tree.
	//   Events which are bubbling upward through the tree will not trigger a listener designated to use capture.
	AddEventListener    func(eventType string, listener func(*Event), useCapture ...bool) `js:"addEventListener"`
	RemoveEventListener func(eventType string, listener func(*Event), useCapture ...bool) `js:"removeEventListener"`
}

// func (e *Element) Style() *CSSStyleDeclaration {
// 	return &CSSStyleDeclaration{
// 		Object: e.Get("style"),
// 	}
// }

type HTMLCollection struct {
	*js.Object
	// length Read only
	// Returns the number of items in the collection.
	Length int `js:"length"`
	// HTMLCollection.item(index number)
	// Returns the specific node at the given zero-based index into the list. Returns null if the index is out of range.
	Item func(index int) *Element `js:"item"`
}

func Wrap(el *js.Object) *Element {
	if el == js.Undefined || el == nil {
		return nil
	}
	return &Element{Object: el}
}

func Window() *Element {
	return Wrap(js.Global)
}

func Document() *Element {
	return Wrap(js.Global.Get("document"))
}

func Body() *Element {
	return Wrap(Document().Get("body"))
}

// Create an element instance
func CreateElement(tagName string) *Element {
	obj := Document().Call("createElement", tagName)
	return Wrap(obj)
}

func GetElementById(id string) *Element {
	return Wrap(Document().Call("getElementById", id))
}

func Alert(msg string) {
	js.Global.Call("alert", msg)
}

// Type Event implements the Event interface and is embedded by
// concrete event types.
type Event struct {
	*js.Object
	Type string `js:"type"`
	// close event
	Code     int    `js:"code"`
	Reason   string `js:"reason"`
	WasClean bool   `js:"wasClean"`
	// wheel event
	DeltaX    float64 `js:"deltaX"`
	DeltaY    float64 `js:"deltaY"`
	DeltaZ    float64 `js:"deltaZ"`
	DeltaMode int     `js:"deltaMode"`
	// keyboard event
	AltKey        bool   `js:"altKey"`
	CharCode      int    `js:"charCode"`
	CtrlKey       bool   `js:"ctrlKey"`
	Key           string `js:"key"`
	KeyIdentifier string `js:"keyIdentifier"`
	KeyCode       int    `js:"keyCode"`
	Locale        string `js:"locale"`
	Location      int    `js:"location"`
	KeyLocation   int    `js:"keyLocation"`
	MetaKey       bool   `js:"metaKey"`
	Repeat        bool   `js:"repeat"`
	ShiftKey      bool   `js:"shiftKey"`
	// mouse event
	Button int `js:"button"`
	// mouse position
	ClientX   int `js:"clientX"`
	ClientY   int `js:"clientY"`
	MovementX int `js:"movementX"`
	MovementY int `js:"movementY"`
	ScreenX   int `js:"screenX"`
	ScreenY   int `js:"screenY"`
	// 	UIEvent.layerX  Read only
	LayerX int `js:"layerX"`
	// Returns the horizontal coordinate of the event relative to the current layer(element).
	// UIEvent.layerY  Read only
	LayerY int `js:"layerY"`
	// Returns the vertical coordinate of the event relative to the current layer.
	// message event
	Data *js.Object `js:"data"`

	// Event control

	// A boolean indicating whether the event bubbles up through the DOM or not.
	Bubbles bool `js:"bubbles"`
	// A boolean indicating whether the event is cancelable.
	Cancelable bool `js:"cancelable"`
	// A reference to the currently registered target for the event.
	CurrentTarget    *Element `js:"currentTarget"`
	DefaultPrevented bool     `js:"defaultPrevented"`
	// Indicates which phase of the event flow is being processed.
	EventPhase int `js:"eventPhase"`
	// A reference to the target to which the event was originally dispatched.
	Target *Element `js:"target"`
	// The time that the event was created. timestamp in ms
	Timestamp                int    `js:"timeStamp"`
	PreventDefault           func() `js:"preventDefault"`
	StopImmediatePropagation func() `js:"stopImmediatePropagation"`
	StopPropagation          func() `js:"stopPropagation"`
	// The KeyboardEvent.getModifierState() method returns the current state of the specified modifier key,
	// true if the modifier is active (that is the modifier key is pressed or locked), otherwise, false.
	//
	// keyArg
	// 	A modifier key value. The value must be one of the KeyboardEvent.key values which represent modifier keys or "Accel". This is case-sensitive.
	GetModifierState func(keyArg string) bool `js:"getModifierState"`
}

func WrapEvent(event *js.Object) *Event {
	return &Event{
		Object: event,
	}
}

// func (e *Element) AddEventListener(typ string, listener func(*Event), useCapture ...bool) func(*js.Object) {
// 	capture := false
// 	if len(useCapture) >= 1 {
// 		capture = useCapture[0]
// 	}
// 	wrapper := func(o *js.Object) {
// 		ev := &Event{Object: o}
// 		listener(ev)
// 	}
// 	e.Call("addEventListener", typ, wrapper, capture)
// 	return wrapper
// }

// func (e *Element) RemoveEventListener(typ string, listener func(*js.Object), useCapture ...bool) {
// 	capture := false
// 	if len(useCapture) >= 1 {
// 		capture = useCapture[0]
// 	}
// 	e.Call("removeEventListener", typ, listener, capture)
// }

func OnLoad(callback func()) {
	Window().AddEventListener(EvtLoad, func(*Event) {
		callback()
	}, false)
}

// The event "DOMContentLoaded" will be fired when the document has been parsed completely,
// that is without stylesheets* and additional images.
// If you need to wait for images and stylesheets, use "load" instead.
func OnDOMContentLoaded(callback func()) {
	Window().AddEventListener(EvtDOMContentLoaded, func(*Event) {
		callback()
	}, false)
}
