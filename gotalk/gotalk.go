// Package gotalk privides a gopherjs binding for github.com/rsms/gotalk
// to do browser side message processing
package gotalk

import (
	"github.com/gopherjs/gopherjs/js"
)

const (
	// Sock event types
	EvtOpen         = "open"
	EvtClose        = "close"
	EvtHeartBeat    = "heartbeat"
	EvtNotification = "notification"
)

var (
	gotalk = js.Global.Get("gotalk")
	// Default `Handlers` utilized by the module-level `handle*` functions
	defaultHandlers = gotalk.Get("defaultHandlers")
	// Default address to connect to. This is falsey if the JS library isn't served
	// by gotalk.
	DefaultResponderAddress = gotalk.Get("defaultResponderAddress")
)

type Handlers struct {
	*js.Object
	// Register a handler for an operation `op`. If `op` is the empty string the
	// handler will be registered as a "fallback" handler, meaning that if there are
	// no handlers registered for request "x", the fallback handler will be invoked.
	//
	//   interface ValueResult(any)
	//   interface ReqValueHandler(any, ValueResult, op string)
	HandleRequest       func(op string, ReqValueHandler interface{})  `js:"HandleRequest"`
	HandleBufferRequest func(op string, ReqBufferHandler interface{}) `js:"handleBufferRequest"`
	// Register a handler for notification `op`. Just as with request handlers,
	// registering a handler for the empty string means it's registered as the
	// fallback handler.
	//
	//   interface NotValueHandler(any, op string)
	//   handleNotification(op string, NotValueHandler)
	HandleNotification func(op string, NotValueHandler interface{}) `js:"handleNotification"`

	// handleBufferNotification(op string, NotBufferHandler)
	//   interface NotBufferHandler(Buf|string, op string)
	HandleBufferNotification func(op string, NotBufferHandler interface{}) `js:"handleBufferNotification"`
}

//     type Sock prototypeof EventEmitter {
type Sock struct {
	*js.Object
	Handlers *Handlers `js:"handlers"`
	//     type Sock prototypeof EventEmitter {
	//   handlers ➝ Handlers    // default: defaultHandlers
	//   protocol ⇄ ProtocolImp // default: protocol.binary

	// Open a connection to a gotalk responder.
	// If `addr` is not provided, `defaultResponderAddress` is used.
	//   open([addr string], [cb function(Error, Sock)]) ➝ Sock
	OpenAddr func(addr string, cb func(err *js.Object, s *Sock)) `js:"open"`
	Open     func(cb func(err *js.Object, s *Sock))              `js:"open"`

	// Start a persistent (keep-alive) connection to a gotalk responder.
	// If `addr` is not provided, `defaultResponderAddress` is used.
	// Because the "open" step is abstracted away, this function does not accept
	// any "open callback". You should listen to the "open" and "close" events
	// instead.
	// The Sock will stay connected, and reconnect as needed, until you call `end()`.
	//   openKeepAlive([addr string]) ➝ Sock
	OpenKeepAlive func(addr ...string) *Sock `js:"openKeepAlive"`

	// Send request for operation `op` with `msg` as the payload, using JSON
	// for encoding.
	//   request(op string, msg any, cb function(Error, result any))
	Request func(op string,
		msg interface{},
		cb func(err *js.Object, result *js.Object)) `js:"request"`

	// Send a request for operation `op` with raw-buffer `buf` as the payload,
	// if any. The type of result depends on the protocol used by the server
	// — a server sending a "text" frame means the result is a string, while a
	// server sending a "binary" frame causes the result to be a Buf.
	//   bufferRequest(op string,
	//                 buf Buf|string|null,
	//                 cb function(Error, result Buf|string))
	BufferRequest func(op string,
		buf interface{},
		cb func(err *js.Object, result *js.Object)) `js:"bufferRequest"`

	//   // Create a StreamRequest for operation `op` which is ready to be used.
	//   // Note that calling this method does not send any data — sending the request
	//   // and reading the response is performed by using the returned object.
	//   streamRequest(op string) ➝ StreamRequest
	//
	//   // Send notification `name` with raw-buffer `buf` as the payload, if any.
	//   bufferNotify(name string, buf Buf|string|null)
	BufferNotify func(op string, buf ...interface{}) `js:"bufferNotify"`

	// Send notification `name` with `msg`, using JSON for encoding.
	//   notify(name string, msg any)
	Notify func(op string, msg ...interface{}) `js:"notify"`

	// Send a heartbeat message with `load` which should be in the range [0-1]
	//   sendHeartbeat(load float)
	SendHeartbeat func(load float64) `js:"sendHeartbeat"`

	// Returns a string representing the address to which the socket is connected.
	//   address() ➝ string|null
	Address func() string `js:"address"`

	//   // Adopt a connection capable of being received from, written to and closed.
	//   // It should be in an "OPEN" ready-state.
	//   // You need to call `handshake` followed by `startReading` after adopting a previosuly
	//   // unadopted connection.
	//   // Throws an error if the provided connection type is not supported.
	//   // Currently only supports WebSocket.
	//   adopt(c Conn)

	//   // Perform protocol handshake.
	//   handshake()

	//   // Schedule reading from the underlying connection. Should only be called
	//   // once per connection.
	//   startReading()

	// Close the socket. If there are any outstanding responses from pending
	// requests, the socket will close when all pending requests has finished.
	// If you call this function a second time, the socket will close immediately,
	// even if there are outstanding responses.
	//   end()
	End func() `js:"end"`

	//   // Event emitted when the connection has opened.
	//   event "open" ()

	//   // Event emitted when the connection has closed. If it closed because of an
	//   // error, the error argument is non-falsey.
	//   event "close" (Error)

	//   event "heartbeat" ({time: Date, load: float})
	// }

	// event handling
	//     EvtOpen         func()
	//     EvtClose        func(Error)
	//     EvtHeartBeat    func({time: Date, load: float})
	//     EvtNotification func(name, value)
	AddListener        func(evtType string, fn interface{}) *Sock `js:"addListener"`
	On                 func(evtType string, fn interface{}) *Sock `js:"on"`
	Once               func(evtType string, fn interface{}) *Sock `js:"once"`
	RemoveListener     func(evtType string, fn interface{}) *Sock `js:"removeListener"`
	RemoveAllListeners func(evtType string)                       `js:"removeAllListeners"`

	// Emit               func(evtType string)                 `js:"emit"`
}

func NewSock(handlers ...*Handlers) *Sock {
	var o *js.Object
	if len(handlers) != 0 {
		o = gotalk.Call("Sock", handlers[0])
	} else {
		o = gotalk.Call("Sock")
	}
	return wrapSock(o)
}

func wrapSock(o *js.Object) *Sock {
	return &Sock{
		Object: o,
	}
}

// Open a connection to a gotalk responder.
//
// open(addr string[, onConnect(Error, Sock)]) -> Sock
//   Connect to gotalk responder at `addr`
//
// open([onConnect(Error, Sock)]) -> Sock
//   Connect to default gotalk responder.
//   Throws an error if `gotalk.defaultResponderAddress` isn't defined.
func Open(onConnect func(err *js.Object, s *Sock), addr ...string) *Sock {
	if len(addr) > 0 {
		return wrapSock(gotalk.Call("open", addr[0], onConnect))
	}
	return wrapSock(gotalk.Call("open", onConnect))
}

// Start a persistent (keep-alive) connection to a gotalk responder.
// If `addr` is not provided, `defaultResponderAddress` is used.
// Equivalent to `Sock(defaultHandlers).openKeepAlive(addr)`
// connection([addr string]) ➝ Sock
func Connection(addr ...string) *Sock {
	if len(addr) > 0 {
		return wrapSock(gotalk.Call("connection", addr[0]))
	}
	return wrapSock(gotalk.Call("connection"))
}

// Convenience "shortcuts" to `defaultHandlers`
// handle(op string, Handlers.ReqValueHandler)

func Handle(op string, ReqValueHandler interface{}) {
	gotalk.Call("handle", op, ReqValueHandler)
}

// handleNotification(name string, Handlers.NotValueHandler)
func HandleNotification(op string, NotValueHandler interface{}) {
	gotalk.Call("handleNotification", op, NotValueHandler)
}

// handleBufferRequest(op string, Handlers.ReqBufferHandler)
func HandleBufferRequest(op string, ReqBufferHandler interface{}) {
	gotalk.Call("handleBufferRequest", op, ReqBufferHandler)
}

// handleBufferNotification(name string, Handlers.NotBufferHandler)
func HandleBufferNotification(op string, NotBufferHandler interface{}) {
	gotalk.Call("handleBufferNotification", op, NotBufferHandler)
}
