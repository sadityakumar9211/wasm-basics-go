package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	ch := make(chan struct{}, 0)
	htmlString := "<h3>I am an HTML snippet from Go</h3>"
	fmt.Println("Welcome to Web assembly from Go")
	js.Global().Set("getHtml", GetHTML(htmlString))
	<-ch
}

func GetHTML(htmlString string) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return htmlString
	})
}
