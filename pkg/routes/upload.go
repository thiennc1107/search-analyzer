package routes

import (
	"encoding/csv"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/keywords"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/failure"
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
	data, err := c.bindCsv(ctx, "csvFile")
	if err != nil {
		return failure.ErrWithTrace(err)
	}

	dtos := make([]*ent.KeywordsCreate, 0, len(data))
	c.Container.Web.Logger.Debugf("Processing data %v", data)

	for i, row := range data {
		if len(row) < 1 {
			return failure.ErrWithTrace(fmt.Errorf("invalid amount of column in row %d", i+1))
		}

		keyword := row[0]
		dto := c.Container.ORM.Keywords.Create().
			SetKeyword(keyword).
			SetStatus(keywords.StatusProcessing).
			SetAdsAmount(0).
			SetLinksAmount(0).
			SetSearchResultAmount(0).
			SetHTMLCode("")
		dtos = append(dtos, dto)
	}

	c.Container.Web.Logger.Debug("Saving data")

	_, err = c.Container.ORM.Keywords.CreateBulk(dtos...).Save(ctx.Request().Context())
	if err != nil {
		return failure.ErrWithTrace(err)
	}

	return nil
}

func (c *upload) bindCsv(ctx echo.Context, fileName string) ([][]string, error) {
	csvHeader, err := ctx.FormFile(fileName)
	if err != nil {
		return nil, failure.ErrWithTrace(err)
	}

	file, err := csvHeader.Open()
	if err != nil {
		return nil, failure.ErrWithTrace(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			c.Container.
				Web.
				Logger.
				Errorf("Failed to close file, err=%v", err)
		}
	}()

	if err != nil {
		return nil, failure.ErrWithTrace(err)
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, failure.ErrWithTrace(err)
	}
	return records, nil
}
