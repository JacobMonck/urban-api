package v1

import (
	"regexp"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/jacobmonck/urban-api/internal/scraper"
)

type Definition struct {
	Word string `json:"word"`
}

var contributorRegex = regexp.MustCompile(`by\s(.+?)\s(\w+\s\d{1,2},\s\d{4})`)

func GetDefinition(ctx *fiber.Ctx) {
	var definitionBody Definition
	if err := ctx.BodyParser(&Definition); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error":   true,
			"message": "Invalid request body.",
		})
	}

	scraper := scraper.New()

	scraper.Collector.OnHTML("div.definition", func(html *colly.HTMLElement) {
		definition := html.DOM.Find("div.break-words.meaning").Text()
		example := html.DOM.Find("div.break-words.meaning").Text()
		contributorAndDate := html.DOM.Find("div.contributor").Text()

		match := contributorRegex.FindStringSubmatch(contributorAndDate)

		contributor := match[1]
		dateString := match[2]
		createdAt, err := time.Parse(TimeLayout, dateString)
	})

	scraper.Visit(defin)

}
