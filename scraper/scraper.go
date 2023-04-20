package scraper

import (
	"fmt"
	Db "github.com/Juandiaz5673/go-scraper-test-v2/database_wrk"
	"github.com/gocolly/colly"
	"strings"
)

func Scraper() []string {
	Verses := make([]string, 32)

	var Url = "https://www.biblegateway.com/passage/?search=Genesis%201&version=KJV"
	c := colly.NewCollector()

	Dbc, err := Db.DbConnection()
	if err != nil {
		panic(err.Error())
	}

	// Step 1 Url | Line 13
	// Step 2 Update the rows | Line 25
	// Step 3 Update CSS Identifier | 27

	for i := 1; i <= 32; i++ {
		c := colly.NewCollector()
		cssVerse := fmt.Sprintf("#en-KJV-%d.text.Gen-1-%d", i, i)

		c.OnHTML(cssVerse, func(h *colly.HTMLElement) {
			verseText := h.Text
			trim := fmt.Sprintf("%d", i)
			verseText1 := strings.TrimLeft(verseText, trim)
			verseText2 := strings.TrimSpace(verseText1)
			Verses[i-1] = verseText2

			Bible := "New King James Version"
			Testament := 0

			y := fmt.Sprintf("INSERT INTO newkingjamesversion (ID, BookNumber, BookName, ChapterName, ChapterNumber, VerseNumber, VerseText, Testament, BookPreference) VALUES('%d', '1', 'Genesis', 'The History of Creation', '1', '%d', '%s', '%d', '%s')", i, i, Verses[i-1], Testament, Bible)
			insert, err := Dbc.Query(y)

			if err != nil {
				panic(err.Error())
			}
			insert.Close()
		})

		c.Visit(Url)
	}

	c.Visit(Url)

	return nil
}
