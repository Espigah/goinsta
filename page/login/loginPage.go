package login

import (
	"log"
	"time"

	"github.com/goinsta/web"
)

type LoginPage struct {
	Finder web.Finder
}

func (loginPage *LoginPage) SetLogin(login string) {
	log.Printf("%+v\n", "[ LoginPage::SetLogin ]")
	we := loginPage.Finder.FindElementAndClick(`#loginForm input[name='username']`)
	we.SendKeys(login)
}

func (loginPage *LoginPage) SetPassword(password string) {
	log.Printf("%+v\n", "[ LoginPage::SetPassword ]")
	we := loginPage.Finder.FindElementAndClick(`#loginForm input[name='password']`)
	we.SendKeys(password)
}

func (loginPage *LoginPage) Submit() {
	log.Printf("%+v\n", "[ LoginPage::Submit ]")
	loginPage.Finder.FindElementAndClick(`#loginForm > div > div:nth-child(3) > button`)
	time.Sleep(5 * time.Second)
}
