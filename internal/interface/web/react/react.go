package react

import (
	"embed"

	"github.com/labstack/echo/v4"
)

var (
	dist          embed.FS
	indexHTML     embed.FS
	distDirFS     = echo.MustSubFS(dist, "internal/interface/web/react/my-app/dist")
	distIndexHtml = echo.MustSubFS(indexHTML, "internal/interface/web/react/my-app/dist")
)

func RegisterHandlersWebApp(e *echo.Echo) {
	e.FileFS("internal/interface/web/react/my-app", "index.html", distIndexHtml)
	e.StaticFS("/", distDirFS)
}
