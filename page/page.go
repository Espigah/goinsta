package page

import (
	"github.com/goinsta/page/category"
	"github.com/goinsta/page/home"
	"github.com/goinsta/page/login"
	"github.com/goinsta/web"
	"github.com/tebeka/selenium"
)

type Config struct {
	Max      int
	Login    string
	Password string
	Search   string
	Interval int
}

func Browse(c *Config, wd selenium.WebDriver) {
	finder := web.Finder{WebDriver: wd}

	loginPage := login.LoginPage{Finder: finder}
	loginPage.SetLogin(c.Login)
	loginPage.SetPassword(c.Password)
	loginPage.Submit()

	homePage := home.HomePage{Finder: finder}
	homePage.Search(c.Search)

	categoryPage := category.CategoryPage{Finder: finder}
	categoryPage.ShowFollowers()
	categoryPage.Follow(c.Max, c.Interval)
}
