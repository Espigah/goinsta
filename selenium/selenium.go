package selenium

import (
	"github.com/tebeka/selenium"
	"log"
)

//const CapabilitiesKey = "goog:chromeOptions"

const (
	//Set constants separately chromedriver.exe Address and local call port of
	seleniumPath = `/usr/bin/chromedriver`
	port         = 9515
)

func CreateService() *selenium.Service {
	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	if err != nil {
		log.Printf("Error starting the ChromeDriver server: %v", err)
	}
	return service
}

func CreateWebDriver() selenium.WebDriver {
	//Set browser compatibility. We set the browser name to chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//Call browser urlPrefix: test reference: defaulturlprefix =“ http://127.0.0.1 :4444/wd/hub"
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	if err != nil {
		log.Fatalf("NewRemote: %v", err)
	}

	return wd
}
