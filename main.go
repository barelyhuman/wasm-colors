package main

import (
	"fmt"
	"syscall/js"
)

func toHex(num int) string {
	return fmt.Sprintf("%02x", num)
}

func main() {
	js.Global().Set("rgbToHex", rgbWrapper())
	<-make(chan bool)
}

func rgbToHex(r int, g int, b int) string {
	return toHex(r) + toHex(g) + toHex(b)
}

func rgbWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 3 {
			return "Invalid no of arguments passed"
		}
		red := args[0].Int()
		green := args[1].Int()
		blue := args[2].Int()
		hex := rgbToHex(red, green, blue)
		return hex
	})
	return jsonFunc
}
