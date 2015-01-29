// Package dom is a thin wrapper of some useful dom functions.

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

// Events
const (
	// Window Event Attributes
	// Window Events triggered for a window object and apply in <body> tag
	// Attributes	Value	Description	In HTML5?
	EvtAfterprint   = "afterprint"   //	Script is run after the document is printed	NEW
	EvtBeforeprint  = "beforeprint"  //	Script is run before the document is printed	NEW
	EvtBeforeunload = "beforeunload" //	Script is run before the document is unloaded	NEW
	EvtError        = "error"        //	Script is run when any error occur	NEW
	EvtHaschange    = "haschange"    //	Script is run when document has changed	NEW
	EvtLoad         = "load"         //	Event fires after the page loading finished
	EvtMessage      = "message"      //	Script is run when document goes in offline	NEW
	EvtOffline      = "offline"      //	Script is run when document comes in Event = "line	NEW
	EvtPagehide     = "pagehide"     //	Script is run when document window is hidden	NEW
	EvtPageshow     = "pageshow"     //	Script is run when document window become visible	NEW
	EvtPopstate     = "popstate"     //	Script is run when document window history changes	NEW
	EvtRedo         = "redo"         //	Script is run when document perform redo	NEW
	EvtResize       = "resize"       //	Event fires when browser window is resized	NEW
	EvtStorage      = "storage"      //	Script is run when web storage area is updated	NEW
	EvtUndo         = "undo"         //	Script is run when document performs undo	NEW
	EvtUnload       = "unload"       //	Event fires when browser window has been closed

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

var (
	document = js.Global.Get("document")
	Body     = &Element{Object: document.Get("body")}
)

type Element struct {
	js.Object
	Id                     string   `js:"id"`
	InnerHTML              string   `js:"innerHTML"`
	InnerText              string   `js:"innerText"`
	TagName                string   `js:"tagName"`
	ContentEditable        bool     `js:"contentEditable"`
	PreviousElementSibling *Element `js:"previousElementSibling"`
	NextElementSibling     *Element `js:"nextElementSibling"`
	FirstElementChild      *Element `js:"firstElementChild"`
	LastElementChild       *Element `js:"lastElementChild"`
}

func CreateElement(tagName string) *Element {
	obj := document.Call("createElement", tagName)
	return &Element{Object: obj}
}

func GetElementById(id string) *Element {
	obj := document.Call("getElementById", id)
	return &Element{Object: obj}
}

func Alert(msg string) {
	js.Global.Call("alert", msg)
}

func (e *Element) AppendChild(child *Element) {
	e.Call("appendChild", child.Object)
}

func (e *Element) Remove() {
	e.Call("remove")
}

func (e *Element) RemoveChild(child *Element) {
	e.Call("removeChild", child.Object)
}

func (e *Element) SetAttribute(attr string, val interface{}) {
	e.Call("setAttribute", attr, val)
}

func (e *Element) GetAttribute(attr string) js.Object {
	return e.Call("getAttribute", attr)
}

func (e *Element) RemoveAttribute(attr string) {
	e.Call("removeAttribute", attr)
}

func (e *Element) QuerySelector(sel string) *Element {
	obj := e.Call("querySelector", sel)
	return &Element{Object: obj}
}

func (e *Element) QuerySelectorAll(sel string) []*Element {
	var out []*Element
	objs := e.Call("querySelectorAll", sel)
	for i := 0; i < objs.Length(); i++ {
		out = append(out, &Element{Object: objs.Index(i)})
	}
	return out
}

// Type Event implements the Event interface and is embedded by
// concrete event types.
type Event struct {
	js.Object
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
	Button    int `js:"button"`
	ClientX   int `js:"clientX"`
	ClientY   int `js:"clientY"`
	MovementX int `js:"movementX"`
	MovementY int `js:"movementY"`
	ScreenX   int `js:"screenX"`
	ScreenY   int `js:"screenY"`
	// message event
	Data js.Object `js:"data"`
}

func (ev *Event) Bubbles() bool {
	return ev.Get("bubbles").Bool()
}

func (ev *Event) Cancelable() bool {
	return ev.Get("cancelable").Bool()
}

func (ev *Event) CurrentTarget() *Element {
	return &Element{Object: ev.Get("currentTarget")}
}

func (ev *Event) DefaultPrevented() bool {
	return ev.Get("defaultPrevented").Bool()
}

func (ev *Event) EventPhase() int {
	return ev.Get("eventPhase").Int()
}

func (ev *Event) Target() *Element {
	return &Element{Object: ev.Get("target")}
}

// timestamp in ms
func (ev *Event) Timestamp() int {
	return ev.Get("timeStamp").Int()
}

func (ev *Event) Type() string {
	return ev.Get("type").String()
}

func (ev *Event) PreventDefault() {
	ev.Call("preventDefault")
}

func (ev *Event) StopImmediatePropagation() {
	ev.Call("stopImmediatePropagation")
}

func (ev *Event) StopPropagation() {
	ev.Call("stopPropagation")
}

func (ev *Event) ModifierState(mod string) bool {
	return ev.Call("getModifierState", mod).Bool()
}

func (e *Element) AddEventListener(typ string, useCapture bool, listener func(*Event)) func(js.Object) {
	wrapper := func(o js.Object) {
		ev := &Event{Object: o}
		listener(ev)
	}
	e.Call("addEventListener", typ, wrapper, useCapture)
	return wrapper
}

func (e *Element) RemoveEventListener(typ string, useCapture bool, listener func(js.Object)) {
	e.Call("removeEventListener", typ, listener, useCapture)
}
