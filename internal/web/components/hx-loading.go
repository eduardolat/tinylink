package components

import (
	"github.com/maragudk/gomponents"
	gcomponents "github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

// LoadingProps are the props for the loading indicator.
type LoadingProps struct {
	// Center indicates whether the loading indicator should be centered.
	Center bool
	// Size is the size of the loading indicator. Can be "sm" "md" (the default) or "lg".
	Size string
}

// Loading returns a loading indicator.
func Loading(props LoadingProps) gomponents.Node {
	if props.Size == "" {
		props.Size = "md"
	}

	return html.Div(
		gcomponents.Classes{
			"htmx-indicator": true,
			"w-full h-full flex justify-center items-center": props.Center,
		},
		html.Div(
			gcomponents.Classes{
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
