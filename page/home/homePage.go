package home

import (
	"log"
	"time"

	"github.com/goinsta/finder"
	"github.com/tebeka/selenium"
)

type HomePage struct {
	Finder finder.Finder
}

func (homePage *HomePage) Search(text string) {
	log.Printf("%+v\n", "[ HomePage::Find ]")
	we := homePage.Finder.FindElementAndClick(`nav input[placeholder='Search'],nav [placeholder='Pesquisar']`)
	we.SendKeys(text)
	time.Sleep(4 * time.Second)
	we.SendKeys(selenium.ReturnKey)
	time.Sleep(1 * time.Second)
	we.SendKeys(selenium.ReturnKey)
	time.Sleep(3 * time.Second)
}
