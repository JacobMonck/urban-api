package scraper

import (
	"os"
	"regexp"

	"github.com/gocolly/colly/v2"
	"github.com/rs/zerolog"
)

const UrbanDictionaryDomain = "www.urbandictionary.com"
const TimeLayout = "January 2, 2006"

var logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

type Scraper struct {
	Collector *colly.Collector
}

func New() *Scraper {
	collector := colly.NewCollector(colly.AllowedDomains(UrbanDictionaryDomain))

	scraper := &Scraper{
		Collector: collector,
	}

	collector.OnRequest(scraper.logRequest)
	collector.OnError(scraper.logError)
}

func (scraper Scraper) Visit(word string) error {
	url := "https://" + UrbanDictionaryDomain + "/define.php?term=" + word

	if err := scraper.Collector.Visit(url); err != nil {
		return err
	}

	return nil
}

func (scraper Scraper) logRequest(request *colly.Request) {
	logger.Trace().Msgf("Request sent to: %s", request.URL)
}

func (scraper Scraper) logError(response *colly.Response, err error) {
	logger.
		Error().
		Err(err).
		Str("url", response.Request.URL.String()).
		Msgf("An error occurred while scraping.")
}
