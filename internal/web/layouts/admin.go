package layouts

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func Admin(title string, children []gomponents.Node) gomponents.Node {
	if title != "" {
		title = " - " + title
	}
	title = "TinyLink" + title

	return components.HTML5(components.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []gomponents.Node{
			html.Link(
				html.Rel("stylesheet"),
				html.Href("https://cdn.jsdelivr.net/npm/@picocss/pico@1/css/pico.min.css"),
			),
			html.Script(
				html.Src("https://cdn.jsdelivr.net/npm/@unocss/runtime@0.58.3/uno.global.min.js"),
			),
			html.Script(
				html.Src("https://cdn.jsdelivr.net/npm/htmx.org@1.9.10/dist/htmx.min.js"),
			),

			themeDetectorScript(),
		},
		Body: []gomponents.Node{
			html.Nav(
				html.Class("container"),
				html.Ul(
					html.Li(
						html.Img(
							html.Src("/static/images/mascot.png"),
							html.Width("40"),
						),
					),
					html.Li(
						html.Strong(
							gomponents.Text("TinyLink"),
						),
					),
				),
				html.Ul(
					html.Li(
						html.Select(
							gomponents.Attr(
								"onChange",
								"setAndStoreTheme(this.value)",
							),

							html.Option(
								html.Value(""),
								html.Selected(),
								html.Disabled(),
								gomponents.Text("Change theme"),
							),
							html.Option(
								html.Value("light"),
								gomponents.Text("Light"),
							),
							html.Option(
								html.Value("dark"),
								gomponents.Text("Dark"),
							),
						),
					),
				),
			),

			html.Main(
				html.Class("container"),
				gomponents.Group(children),
			),
		},
	})
}

func themeDetectorScript() gomponents.Node {
	return gomponents.Raw(`
		<script>
			let mode = "light"
			if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
				mode = "dark"
			}

			function isThemeValid(theme) {
				return theme === "light" || theme === "dark"
			}

			function setTheme(newMode) {
				if (!isThemeValid(newMode)) return;
				document.documentElement.setAttribute('data-theme', newMode)
			}

			function setAndStoreTheme(newMode) {
				if (!isThemeValid(newMode)) return;
				setTheme(newMode)
				localStorage.theme = newMode
			}

			setTheme(mode)
		</script>
	`)
}
