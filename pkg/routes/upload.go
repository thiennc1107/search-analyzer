package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/controller"
)

type (
	upload struct {
		controller.Controller
	}
)

func (c *upload) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Cache.Enabled = false
	page.Layout = "main"
	page.Name = "upload"
	page.Metatags.Description = "Upload keyword page."
	page.Metatags.Keywords = []string{"Go", "MVC", "Web", "Software", "Search result", "Statistic"}
	return c.RenderPage(ctx, page)
}

func (c *upload) Post(ctx echo.Context) error {
	err := c.handleFileUpload(ctx)
	if err != nil {
		return err
	}
	page := controller.NewPage(ctx)
	page.Cache.Enabled = false
	page.Layout = "main"
	page.Name = "upload-success"
	page.Metatags.Description = "Upload keyword success page."
	page.Metatags.Keywords = []string{"Go", "MVC", "Web", "Software", "Search result", "Statistic"}
	return c.RenderPage(ctx, page)
}

func (c *upload) handleFileUpload(ctx echo.Context) error {
	time.Sleep(5 * time.Second)
	return nil
}
