package components

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents-heroicons/v2/solid"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

// FlagCircleProps is the props type for FlagCircle.
type FlagCircleProps struct {
	// Flag is the boolean flag to render.
	Flag bool
	// TrueTooltip is the tooltip to show when the flag is true.
	TrueTooltip string
	// FalseTooltip is the tooltip to show when the flag is false.
	FalseTooltip string
}

// FlagCircle renders a mark for a boolean flag.
// If the flag is true, a green check circle is rendered.
// If the flag is false, a red x circle is rendered.
func FlagCircle(props FlagCircleProps) gomponents.Node {
	if props.Flag {
		return html.Div(
			html.Class("inline border-b-0!"),
			gomponents.If(
				props.TrueTooltip != "",
				html.DataAttr("tooltip", props.TrueTooltip),
			),
			solid.CheckCircle(
				components.Classes{
					"w-4 h-4":        true,
					"text-green-500": true,
				},
			),
		)
	}

	return html.Div(
		html.Class("inline border-b-0!"),
		gomponents.If(
			props.FalseTooltip != "",
			html.DataAttr("tooltip", props.FalseTooltip),
		),
		solid.XCircle(
			components.Classes{
				"w-4 h-4":      true,
				"text-red-500": true,
			},
		),
	)
}
