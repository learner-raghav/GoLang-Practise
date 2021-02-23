1. Http package has the http.Handler interface which has a function called ServeHTTP, the core 
    responsibility of the function is to write response headers and the body
   
2. In order to process an incoming request, all we need is a type that implements this interface. 
    The standard library provides some default types that implement the Handler interface.
   
3. ServeMux - It is a router, that holds the mapping of URL path and the HTTP method to the corresponding Handler.
    It simply routes the incoming request to the correct handler. 
   
4. ServeMuz is an HTTP request multiplexer. It matches the URL of each incoming request against a list 
   of registered patterns and calls the handler for the pattern that most closely matches the URL.
   
type ServeMux struct {
    mu sync.RWMutex
    m map[string]muxEntry
    es []muxEntry
    hosts bool
}