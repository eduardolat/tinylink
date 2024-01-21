package htmx

import "github.com/maragudk/gomponents"

// HxGet returns a gomponents node with the hx-get
// attribute set to the given path.
//
// https://htmx.org/attributes/hx-get/
func HxGet(path string) gomponents.Node {
	return gomponents.Attr("hx-get", path)
}

// HxPost returns a gomponents node with the hx-post
// attribute set to the given path.
//
// https://htmx.org/attributes/hx-post/
func HxPost(path string) gomponents.Node {
	return gomponents.Attr("hx-post", path)
}

// HxPut returns a gomponents node with the hx-put
// attribute set to the given path.
//
// https://htmx.org/attributes/hx-put/
func HxPut(path string) gomponents.Node {
	return gomponents.Attr("hx-put", path)
}

// HxPatch returns a gomponents node with the hx-patch
// attribute set to the given path.
//
// https://htmx.org/attributes/hx-patch/
func HxPatch(path string) gomponents.Node {
	return gomponents.Attr("hx-patch", path)
}

// HxDelete returns a gomponents node with the hx-delete
// attribute set to the given path.
//
// https://htmx.org/attributes/hx-delete/
func HxDelete(path string) gomponents.Node {
	return gomponents.Attr("hx-delete", path)
}

// HxTrigger returns a gomponents node with the hx-trigger
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-trigger/
func HxTrigger(value string) gomponents.Node {
	return gomponents.Attr("hx-trigger", value)
}

// HxTarget returns a gomponents node with the hx-target
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-target/
func HxTarget(value string) gomponents.Node {
	return gomponents.Attr("hx-target", value)
}

// HxSwap returns a gomponents node with the hx-swap
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-swap/
func HxSwap(value string) gomponents.Node {
	return gomponents.Attr("hx-swap", value)
}

// HxIndicator returns a gomponents node with the hx-indicator
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-indicator/
func HxIndicator(value string) gomponents.Node {
	return gomponents.Attr("hx-indicator", value)
}

// HxConfirm returns a gomponents node with the hx-confirm
// attribute set to the given value.
//
// https://htmx.org/attributes/hx-confirm/
func HxConfirm(value string) gomponents.Node {
	return gomponents.Attr("hx-confirm", value)
}
