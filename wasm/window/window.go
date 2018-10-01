package window

import "syscall/js"

func Window() (window js.Value) {
	return js.Global().Get("window")
}
func Document() (document js.Value) {
	return js.Global().Get("document")
}
func Head() (head js.Value) {
	return Document().Get("head")
}
func Body() (body js.Value) {
	return Document().Get("body")
}

func GetElementById(id string) (element js.Value) {
	return Document().Call("getElementById", id)
}

func AddEventHandler(element js.Value, onEvent string, callback func(values []js.Value)) {
	element.Set(onEvent, js.NewCallback(callback))
}

func AddOnClick(element js.Value, callback func(values []js.Value)) {
	AddEventHandler(element, "onclick", callback)
}

func AddOnMouseDown(element js.Value, callback func(values []js.Value)) {
	AddEventHandler(element, "onmousedown", callback)
}

func AddOnMouseUp(element js.Value, callback func(values []js.Value)) {
	AddEventHandler(element, "onmouseup", callback)
}

func AddOnMouseMove(element js.Value, callback func(values []js.Value)) {
	AddEventHandler(element, "onmousemove", callback)
}
