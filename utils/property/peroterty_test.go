package property

import (
	"github.com/gopherjs/gopherjs/js"
	"testing"
)

var (
	obj *js.Object
)

func init() {
	obj = js.Global.Get("Object").New()
	obj.Set("a", js.M{
		"b": js.M{
			"c": 100,
		},
	})
}

func TestGet(t *testing.T) {
	if Get(obj, "a.b.c").Int() != 100 {
		t.Fail()
	}
}

func TestSet(t *testing.T) {
	Set(obj, "a.b.c", 1000)
	if Get(obj, "a.b.c").Int() != 100 {
		t.Fail()
	}
}
