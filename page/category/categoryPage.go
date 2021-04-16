package target

import (
	"log"
	"time"

	"github.com/goinsta/finder"
)

type CategoryPage struct {
	Finder finder.Finder
}

func (categoryPage *CategoryPage) ShowFollowers() {
	log.Printf("%+v\n", "[ CategoryPage::ShowFollowers ]")
	categoryPage.Finder.FindElementAndClick(`header section li:nth-child(2)`)
}

func (categoryPage *CategoryPage) ScrollDown() {
	log.Printf("%+v\n", "[ CategoryPage::ScrollDown ]")
	categoryPage.Finder.WebDriver.ExecuteScript(`document.querySelector("div[role='dialog'] div:first-child div:nth-child(2)").scrollBy(0,400)`, nil)
}

func (categoryPage *CategoryPage) Follow(max int, interval int) {
	log.Printf("%+v\n", "[ CategoryPage::Follow ]")

	clicks := 0

	var findAndFollow func()
	findAndFollow = func() {
		elements := categoryPage.Finder.FindElements(`li button`)
		for _, element := range elements {
			if clicks >= max {
				break
			}

			elementTextResponse, err := element.Text()
			if err != nil {
				log.Fatalf("Get text: %v", err)
			}
			if elementTextResponse != "Follow" && elementTextResponse != "Seguir" {
				continue
			}
			element.Click()

			clicks++
			log.Printf("%+v\n", clicks)
			time.Sleep(time.Duration(interval) * time.Second)
		}

		if clicks < max {
			categoryPage.ScrollDown()
			findAndFollow()
		}

	}

	findAndFollow()
	log.Printf("%+v\n", "FIM")

}
