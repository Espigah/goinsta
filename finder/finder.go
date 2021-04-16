package finder

import (
	"log"
	"time"

	"github.com/tebeka/selenium"
)

const waitingTimeBeforeFind = 500

type Finder struct {
	WebDriver selenium.WebDriver
}

func (f *Finder) FindElement(selector string) selenium.WebElement {
	log.Printf("%+v\n", selector)
	time.Sleep(waitingTimeBeforeFind * time.Millisecond)
	we, err := f.WebDriver.FindElement(selenium.ByCSSSelector, selector)

	if err != nil {
		log.Fatalf("FindElement: %v", err)
	}
	return we
}

func (f *Finder) FindElements(selector string) []selenium.WebElement {
	log.Printf("%+v\n", selector)
	time.Sleep(waitingTimeBeforeFind * time.Millisecond)
	elementList, err := f.WebDriver.FindElements(selenium.ByCSSSelector, selector)

	if err != nil {
		log.Fatalf("FindElements: %v", err)
	}
	return elementList
}

func (f *Finder) FindElementAndClick(selector string) selenium.WebElement {
	we := f.FindElement(selector)
	we.Click()
	return we
}
