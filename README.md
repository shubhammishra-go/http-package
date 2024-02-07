# net/http package in Go
It is a standard package for http that provides HTTP client and server implementations. 
HTTP is an acronym for hypertext transfer protocol, a communication protocol that ensures the transfer of data between a client and a server. A perfect instance of an HTTP client-server interaction is when you open your browser and type in a URL. Your browser acts as a client and fetches resources from a server which it then displays.

# to use net/http package

first and always to use any feature of golang you must have to initializes a new module in the current directory.
like ```github.com/shubhammishra-1```

just import ```"net/http"``` package and use it.


# About http.HandleFunc(pattern string, handler func(ResponseWriter, *Request))

HandleFunc() registers the handler function for the given pattern in DefaultServeMux.


# About http.ListenAndServe(addr string, handler Handler) error

ListenAndServe listens on the TCP network address ```addr``` and then calls ```Serve``` with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.

The handler is typically nil, in which case ```DefaultServeMux``` is used.

ListenAndServe always returns a non-nil error.


# What is DefaultServeMux

It is the default ```ServeMux``` used by ```Serve```. 

# What is ServeMux

ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL. 

# What is Serve

Serve accepts incoming HTTP connections on the listener l, creating a new service goroutine for each. The service goroutines read requests and then call handler to reply to them.

The handler is typically nil, in which case DefaultServeMux is used.

HTTP/2 support is only enabled if the Listener returns *tls.Conn connections and they were configured with "h2" in the TLS Config.NextProtos.

Serve always returns a non-nil error. 


# About http.ResponseWriter interface

A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
its instance generally parameterized in Haldeler Functions as ```w``` 

functions associated with ```ResponseWriter``` interface.

```w.Header()``` 
A Header represents the key-value pairs in an HTTP header. 
The keys should be in canonical form, as returned by CanonicalHeaderKey. 

```w.Header().Set("Trailer", "AtEnd1, AtEnd2")```

Set sets the header entries associated with key to the single element value. It replaces any existing values associated with key. The key is case insensitive; it is canonicalized by textproto.CanonicalMIMEHeaderKey. To use non-canonical keys, assign to the map directly. 

```w.Header().Add("Trailer", "AtEnd3")```

Add() adds the key, value pair to the header. It appends to any existing values associated with key. The key is case insensitive; it is canonicalized by CanonicalHeaderKey. 


```w.Write(byte[](data))``` Write() writes the byte data to the connection as part of an HTTP reply.

```WriteHeader(statusCode int)```  WriteHeader sends an HTTP response header with the provided status code.

## About http.Request struct

A Request represents an HTTP request received by a server or to be sent by a client. 
its pointer instance generally parameterized in Haldeler Functions as ```r``` 

variables associated with ```Request``` struct

```r.Method ``` Method is a variable which specifies the HTTP method (GET, POST, PUT, etc.).
For client requests, an empty string means GET.

```r.URL``` this variable returns wild pattern URL of client request

```r.Proto``` this variable returns protocol version for incoming server requests.

& many more please see documentation...


## About http.CanonicalHeaderKey(s string) string
CanonicalHeaderKey returns the canonical format of the header key s. The canonicalization converts the first letter and any letter following a hyphen to upper case; the rest are converted to lowercase. For example, the canonical key for "accept-encoding" is "Accept-Encoding". If s contains a space or invalid header field bytes, it is returned without modifications. 


# About Response

Response represents the response from an HTTP request.

The Client and Transport return Responses from servers once the response headers have been received. The response body is streamed on demand as the Body field is read. 

functions associated with response struct

```resp.Body```   Body represents the response body.
The response body is streamed on demand as the Body field is read. If the network connection fails or the server terminates the response, Body.Read calls return an error.

```resp.Body.Close()```  Close records whether the header directed that the connection be closed after reading Body. The value is advice for clients: neither ReadResponse nor Response.Write ever closes a connection.


## About http.Get(url string) (resp *Response, err error)

Get issues a GET to the specified URL. If the response is one of the following redirect codes, Get follows the redirect, up to a maximum of 10 redirects.

An error is returned if there were too many redirects or if there was an HTTP protocol error. A non-2xx response doesn't cause an error. Any returned error will be of type *url.Error. The url.Error value's Timeout method will report true if the request timed out.

When err is nil, resp always contains a non-nil resp.Body. Caller should close resp.Body when done reading from it.

Get is a wrapper around DefaultClient.Get.

To make a request with custom headers, use NewRequest and DefaultClient.Do.

To make a request with a specified context.Context, use NewRequestWithContext and DefaultClient.Do. 

```go 
    res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
```


## About http.Post(url, contentType string, body io.Reader) (resp *Response, err error)

Post issues a POST to the specified URL.

Caller should close resp.Body when done reading from it.

If the provided body is an io.Closer, it is closed after the request.

Post is a wrapper around DefaultClient.Post.

To set custom headers, use NewRequest and DefaultClient.Do.

See the Client.Do method documentation for details on how redirects are handled.

To make a request with a specified context.Context, use NewRequestWithContext and DefaultClient.Do. 



# Predefined HTTP status code in Go

you can use these predefined ```constant``` http status code like ```code:=http.StatusContinue```

```go
const (
	StatusContinue           = 100 // RFC 9110, 15.2.1
	StatusSwitchingProtocols = 101 // RFC 9110, 15.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusOK                   = 200 // RFC 9110, 15.3.1
	StatusCreated              = 201 // RFC 9110, 15.3.2
	StatusAccepted             = 202 // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 9110, 15.3.4
	StatusNoContent            = 204 // RFC 9110, 15.3.5
	StatusResetContent         = 205 // RFC 9110, 15.3.6
	StatusPartialContent       = 206 // RFC 9110, 15.3.7
	StatusMultiStatus          = 207 // RFC 4918, 11.1
	StatusAlreadyReported      = 208 // RFC 5842, 7.1
	StatusIMUsed               = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices  = 300 // RFC 9110, 15.4.1
	StatusMovedPermanently = 301 // RFC 9110, 15.4.2
	StatusFound            = 302 // RFC 9110, 15.4.3
	StatusSeeOther         = 303 // RFC 9110, 15.4.4
	StatusNotModified      = 304 // RFC 9110, 15.4.5
	StatusUseProxy         = 305 // RFC 9110, 15.4.6

	StatusTemporaryRedirect = 307 // RFC 9110, 15.4.8
	StatusPermanentRedirect = 308 // RFC 9110, 15.4.9

	StatusBadRequest                   = 400 // RFC 9110, 15.5.1
	StatusUnauthorized                 = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired              = 402 // RFC 9110, 15.5.3
	StatusForbidden                    = 403 // RFC 9110, 15.5.4
	StatusNotFound                     = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed             = 405 // RFC 9110, 15.5.6
	StatusNotAcceptable                = 406 // RFC 9110, 15.5.7
	StatusProxyAuthRequired            = 407 // RFC 9110, 15.5.8
	StatusRequestTimeout               = 408 // RFC 9110, 15.5.9
	StatusConflict                     = 409 // RFC 9110, 15.5.10
	StatusGone                         = 410 // RFC 9110, 15.5.11
	StatusLengthRequired               = 411 // RFC 9110, 15.5.12
	StatusPreconditionFailed           = 412 // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge        = 413 // RFC 9110, 15.5.14
	StatusRequestURITooLong            = 414 // RFC 9110, 15.5.15
	StatusUnsupportedMediaType         = 415 // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable = 416 // RFC 9110, 15.5.17
	StatusExpectationFailed            = 417 // RFC 9110, 15.5.18
	StatusTeapot                       = 418 // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest           = 421 // RFC 9110, 15.5.20
	StatusUnprocessableEntity          = 422 // RFC 9110, 15.5.21
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 9110, 15.5.22
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 9110, 15.6.1
	StatusNotImplemented                = 501 // RFC 9110, 15.6.2
	StatusBadGateway                    = 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable            = 503 // RFC 9110, 15.6.4
	StatusGatewayTimeout                = 504 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)
```

# Reference

```https://pkg.go.dev/net/http```