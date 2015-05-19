[![GoDoc](http://godoc.org/github.com/fabioberger/cookie?status.svg)](https://godoc.org/github.com/fabioberger/cookie)


Gopherjs Cookie Library
-----------------------

This Library implements convenience functions for manipulating cookies in a Gopherjs application.

# Installation

Install with go get:

```go get github.com/fabioberger/cookie```

Then include the package in your imports:

```import "github.com/fabioberger/cookie"```

# Example Usage

Set Cookie:

```go
expires := time.Now().Add(time.Hour) // Set expiry time to in one hour
cookie.Set("username", "John Doe", &expires, "/")

// Or

// set cookie with correctly formatted string
cookie.SetString("username=John Doe; expires=Thu, 18 Dec 2013 12:00:00 UTC; path=/")
```

Get Cookie:

```go
username, ok := cookie.Get("username")
if !ok {
	// Cookie was not found
}
```

Delete Cookie:

```go
cookie.Delete("username")
```