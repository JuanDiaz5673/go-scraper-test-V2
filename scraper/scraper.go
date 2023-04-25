package scraper

import (
	"fmt"
	Db "github.com/Juandiaz5673/go-scraper-test-v2/database_wrk"
	"github.com/gocolly/colly"
	"strings"
)

func Scraper() []string {
	// Needs to be updated to total number of verses
	Verses := make([]string, 1000)

	var Url = "https://www.biblegateway.com/passage/?search=Genesis%202&version=NKJV"
	c := colly.NewCollector()

	Dbc, err := Db.DbConnection()
	if err != nil {
		panic(err.Error())
	}
	for i := 32; i <= 57; i++ {
		c := colly.NewCollector()
		cssVerse := fmt.Sprintf("#en-NKJV-%d.text.Gen-2-2", i)

		c.OnHTML(cssVerse, func(h *colly.HTMLElement) {
			verseText := h.Text
			fmt.Println(verseText)
			trim := fmt.Sprintf("%d", i)
			verseText1 := strings.TrimLeft(verseText, trim)
			verseText2 := strings.TrimSpace(verseText1)
			Verses[i-1] = verseText2

			Bible := "New King James Version"
			Testament := 0

			y := fmt.Sprintf("INSERT INTO newkingjamesversion (ID, BookNumber, BiblePreference, Testament, BookNumber, BookName, ChapterNumber, ChapterName, VerseNumber, VerseText) VALUES('%d', '2', 'Genesis2', 'Book2chaptermustupdate', '1', '%d', '%s', '%d', '%s')", i, i, Verses[i-1], Testament, Bible)
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
