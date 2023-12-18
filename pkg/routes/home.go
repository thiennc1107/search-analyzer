package routes

import (
	"context"

	"github.com/mikestefanello/pagoda/enum"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/failure"

	"github.com/labstack/echo/v4"
)

type (
	home struct {
		controller.Controller
	}

	keyword struct {
		Keyword            string
		Status             string
		StatusTag          string
		AdsAmount          int
		LinksAmount        int
		SearchResultAmount int
		DownloadHtmlLink   string
	}
)

func (c *home) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "home"
	page.Metatags.Description = "Welcome to the homepage."
	page.Metatags.Keywords = []string{"Go", "MVC", "Web", "Software", "Search result", "Statistic"}
	page.Pager = controller.NewPager(ctx, 10)
	data, err := c.fetchKeywords(ctx.Request().Context(), &page.Pager)
	if err != nil {
		return failure.ErrWithTrace(err)
	}
	page.Data = data
	return c.RenderPage(ctx, page)
}

// fetchKeywords is an mock example of fetching posts to illustrate how paging works
func (c *home) fetchKeywords(ctx context.Context, pager *controller.Pager) ([]keyword, error) {
	keywordData, err := c.Container.ORM.Keywords.Query().All(ctx)
	if err != nil {
		return nil, failure.ErrWithTrace(err)
	}

	kws := make([]keyword, 0, len(keywordData))

	for _, val := range keywordData {
		kws = append(kws, keyword{
			Status:             val.Status.String(),
			StatusTag:          statusToTag(val.Status.String()),
			Keyword:            val.Keyword,
			AdsAmount:          val.AdsAmount,
			LinksAmount:        val.LinksAmount,
			SearchResultAmount: val.SearchResultAmount,
		})
	}

	start := 0
	end := len(kws)
	if pager.GetOffset() < len(kws) {
		start = pager.GetOffset()
	}

	if pager.GetOffset()+pager.ItemsPerPage < len(kws) {
		end = pager.GetOffset() + pager.ItemsPerPage
	}
	return kws[start:end], nil
}

func statusToTag(status string) string {
	switch status {
	case enum.StatusFailed:
		return "is-danger"
	case enum.StatusProcessing:
		return "is-warning"
	case enum.StatusFinished:
		return "is-success"
	default:
		return ""
	}
}
