package gotalk

import (
	"github.com/gopherjs/gopherjs/js"
)

// type Handlers {
//   // Callable for producing the results value of an operation.
//   interface BufferResult(Buf|string) {
//     // Callable for producing an error result
//     error(Error)
//   }

//   // Like BufferResult, but accepts any value which will be encoded as JSON.
//   interface ValueResult(any) {
//     error(Error)
//   }

//   // Signature for request handlers dealing with raw data
//   interface ReqBufferHandler(Buf|string, BufferResult, op string)

//   // Signature for request handlers dealing with JSON-decoded value
//   interface ReqValueHandler(any, ValueResult, op string)

//   // Signature for notification handlers dealing with raw data
//   interface NotBufferHandler(Buf|string, name string)

//   // Signature for request handlers dealing with JSON-decoded value
//   interface NotValueHandler(any, name string)

//   // Register a handler for an operation `op`. If `op` is the empty string the
//   // handler will be registered as a "fallback" handler, meaning that if there are
//   // no handlers registered for request "x", the fallback handler will be invoked.
//   handleRequest(op string, ReqValueHandler)
//   handleBufferRequest(op string, ReqBufferHandler)

//   // Register a handler for notification `name`. Just as with request handlers,
//   // registering a handler for the empty string means it's registered as the
//   // fallback handler.
//   handleNotification(name string, NotValueHandler)
//   handleBufferNotification(name string, NotBufferHandler)

//   // Find request and notification handlers
//   findRequestHandler(op string) ➝ ReqBufferHandler|null
//   findNotificationHandler(name string) ➝ NotBufferHandler|null
// }

type Handlers struct {
	*js.Object
	// Register a handler for an operation `op`. If `op` is the empty string the
	// handler will be registered as a "fallback" handler, meaning that if there are
	// no handlers registered for request "x", the fallback handler will be invoked.
	//
	//   interface ValueResult(any)
	//   interface ReqValueHandler(any, ValueResult, op string)
	HandleRequest func(op string, handler ...interface{}) `js:"HandleRequest"`
	// Register a handler for notification `name`. Just as with request handlers,
	// registering a handler for the empty string means it's registered as the
	// fallback handler.
	//
	//   interface NotValueHandler(any, name string)
	//   handleNotification(name string, NotValueHandler)
	HandleNotification func(op string, handler ...interface{}) `js:"HandleNotification"`
}

//     type Sock prototypeof EventEmitter {
type Sock struct {
	*js.Object
	Handlers *Handlers `js:"Handlers"`
	//     type Sock prototypeof EventEmitter {
	//   handlers ➝ Handlers    // default: defaultHandlers
	//   protocol ⇄ ProtocolImp // default: protocol.binary

	// Open a connection to a gotalk responder.
	// If `addr` is not provided, `defaultResponderAddress` is used.
	//   open([addr string], [cb function(Error, Sock)]) ➝ Sock
	Open        func(addr string, cb func(err *js.Object, s *Sock)) `js:"open"`
	OpenDefault func(cb func(err *js.Object, s *Sock))              `js:"open"`

	// Start a persistent (keep-alive) connection to a gotalk responder.
	// If `addr` is not provided, `defaultResponderAddress` is used.
	// Because the "open" step is abstracted away, this function does not accept
	// any "open callback". You should listen to the "open" and "close" events
	// instead.
	// The Sock will stay connected, and reconnect as needed, until you call `end()`.
	//   openKeepAlive([addr string]) ➝ Sock

	// Send request for operation `op` with `value` as the payload, using JSON
	// for encoding.
	//   request(op string, value any, cb function(Error, result any))
	Request func(op string, value interface{}, cb func(err *js.Object, result *js.Object)) `js:"request"`

	//   // Send a request for operation `op` with raw-buffer `buf` as the payload,
	//   // if any. The type of result depends on the protocol used by the server
	//   // — a server sending a "text" frame means the result is a string, while a
	//   // server sending a "binary" frame causes the result to be a Buf.
	//   bufferRequest(op string,
	//                 buf Buf|string|null,
	//                 cb function(Error, result Buf|string))

	//   // Create a StreamRequest for operation `op` which is ready to be used.
	//   // Note that calling this method does not send any data — sending the request
	//   // and reading the response is performed by using the returned object.
	//   streamRequest(op string) ➝ StreamRequest

	//   // Send notification `name` with raw-buffer `buf` as the payload, if any.
	//   bufferNotify(name string, buf Buf|string|null)

	// Send notification `name` with `value`, using JSON for encoding.
	//   notify(name string, value any)
	Notify func(op string, value interface{}) `js:"value"`

	//   // Send a heartbeat message with `load` which should be in the range [0-1]
	//   sendHeartbeat(load float)

	//   // Returns a string representing the address to which the socket is connected.
	//   address() ➝ string|null

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

}
