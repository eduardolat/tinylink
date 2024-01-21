package components

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// HxLoadingProps are the props for the loading indicator.
type HxLoadingProps struct {
	// ID is the HTML ID of the loading indicator.
	ID string
	// Center indicates whether the loading indicator should be centered.
	Center bool
	// Size is the size of the loading indicator. Can be "sm" "md" (the default) or "lg".
	Size string
}

// HxLoading returns a loading indicator.
func HxLoading(props HxLoadingProps) gomponents.Node {
	if props.Size == "" {
		props.Size = "md"
	}

	return html.Div(
		gomponents.If(
			props.ID != "",
			html.ID(props.ID),
		),
		Classes{
			"htmx-indicator": true,
			"w-full h-full flex justify-center items-center": props.Center,
		},
		html.Div(
			Classes{
				"border-solid border-gray-500 border-t-transparent animate-spin rounded-full": true,

				"w-[15px] h-[15px]": props.Size == "sm",
				"w-[25px] h-[25px]": props.Size == "md",
				"w-[40px] h-[40px]": props.Size == "lg",

				"border-2": props.Size == "sm",
				"border-3": props.Size == "md",
				"border-5": props.Size == "lg",
			},
		),
	)
}
