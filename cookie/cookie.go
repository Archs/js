package cookie

import (
	"github.com/gopherjs/gopherjs/js"
	"strings"
	"time"
)

var (
	doc = js.Global.Get("document")
)

// Get returns a given cookie by name. If the cookie is not set, ok will be
// set to false
func Get(name string) (value string, ok bool) {
	cookieStr := doc.Get("cookie").String()
	if cookieStr == "" {
		return "", false
	}
	cookiePairs := strings.Split(cookieStr, "; ")
	for _, c := range cookiePairs {
		cookie := strings.Split(c, "=")
		if cookie[0] == name {
			return cookie[1], true
		}
	}
	return "", false
}

// SetString sets a cookie given a correctly formatted cookie string
// i.e "username=John Smith; expires=Thu, 18 Dec 2013 12:00:00 UTC; path=/"
func SetString(cookie string) {
	doc.Set("cookie", cookie)
}

// Set adds a cookie to a user's browser with a name, value, expiry and path
// value, path and expires can be omitted
func Set(name string, value string, expires *time.Time, path string) {
	if name == "" {
		return
	}
	var expiry string
	if expires != nil {
		e := *expires
		e = e.UTC()
		t := e.Format("Mon, 02 Jan 2006 15:04:05 UTC")
		expiry = "expires=" + t + "; "
	}
	if path != "" {
		path = " path=" + path
	}
	c := name + "=" + value + "; " + expiry + path
	SetString(c)
}

// Delete removes a cookie specified by name
func Delete(name string) {
	c := name + "=; expires=Thu, 01 Jan 1970 00:00:01 UTC;"
	SetString(c)
}
