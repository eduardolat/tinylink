package components

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// HxResWrapper is a wrapper for htmx responses.
// It is hidden by default and only shown when there is content inside
// of it. This avoids problems with margins and padding when the
// response is empty.
//
// https://developer.mozilla.org/en-US/docs/Web/CSS/:empty
func HxResWrapper(id string) gomponents.Node {
	return html.Div(
		html.Class("empty:hidden"),
		html.ID(id),
	)
}
