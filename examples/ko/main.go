package main

import (
	"github.com/Archs/js/gopherjs-ko"
	_ "github.com/Archs/js/gopherjs-ko/components/clock"
	_ "github.com/Archs/js/gopherjs-ko/components/vCard"
)

func main() {
	ko.ApplyBindings(100)
}
