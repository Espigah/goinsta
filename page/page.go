package page

import (
	"github.com/goinsta/finder"
	"github.com/goinsta/page/category"
	"github.com/goinsta/page/home"
	"github.com/goinsta/page/login"
	"github.com/tebeka/selenium"
)

type Conf struct {
	Max      int
	Login    string
	Password string
	Search   string
	Interval int
}

func Browse(config *Conf, wd selenium.WebDriver) {
	finder := finder.Finder{WebDriver: wd}

	loginPage := login.LoginPage{Finder: finder}
	loginPage.SetLogin(config.Login)
	loginPage.SetPassword(config.Password)
	loginPage.Submit()

	homePage := home.HomePage{Finder: finder}
	homePage.Search(config.Search)

	categoryPage := target.CategoryPage{Finder: finder}
	categoryPage.ShowFollowers()
	categoryPage.Follow(config.Max, config.Interval)
}
