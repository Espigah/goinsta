package main

import (
	"github.com/goinsta/config"
	"github.com/goinsta/page"
	"github.com/goinsta/selenium"
	"log"
	"time"
)

const (
	pageTarget = `https://www.instagram.com/`
)

func main() {
	log.Printf("%+v\n", "run")

	config, err := config.Read("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	service := selenium.CreateService()
	//Delay service shutdown
	defer service.Stop()

	wd := selenium.CreateWebDriver()
	defer wd.Quit()

	if err := wd.Get(pageTarget); err != nil {
		log.Fatalf("Open url: %v", err)
	}

	page.Browse(config, wd)

	time.Sleep(40 * time.Second)
}
