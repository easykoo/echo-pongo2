package main

import (
	"fmt"
	p "github.com/easykoo/echo-pongo2"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	fmt.Println("Prepare templates.")
	t := p.PrepareTemplates(p.Options{
		Directory:  "templates/",
		Extensions: []string{".html"},
	})
	e.SetRenderer(t)

	e.Get("/subdir1", func(c *echo.Context) error {
		return c.Render(http.StatusOK, "index1", map[string]interface{}{"title": "index1", "msg": "index1"})
	})
	e.Get("/subdir2", func(c *echo.Context) error {
		return c.Render(http.StatusOK, "index2", map[string]interface{}{"title": "index2", "msg": "index2"})
	})
	e.Get("/subdir3", func(c *echo.Context) error {
		return c.Render(http.StatusOK, "index3", map[string]interface{}{"title": "index3", "msg": "index3"})
	})

	fmt.Println("Start server.")
	e.Run(":8080")
}
