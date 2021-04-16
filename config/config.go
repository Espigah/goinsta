package config

import (
	"fmt"
	"github.com/goinsta/page"
	"github.com/namsral/flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func Read(filename string) (*page.Conf, error) {

	max := getEnvAsInt("GOINSTA_MAX", 20)

	interval := getEnvAsInt("GOINSTA_INTERVAL", 1)

	c := &page.Conf{
		Max:      max,
		Login:    os.Getenv("GOINSTA_LOGIN"),
		Password: os.Getenv("GOINSTA_PASSWORD"),
		Search:   os.Getenv("GOINSTA_SEARCH"),
		Interval: interval,
	}

	file, err := loadFile(filename)
	err = yaml.Unmarshal(file, c)
	if err != nil {
		return nil, fmt.Errorf("file %q: %v", filename, err)
	}
	flag.IntVar(&c.Max, "max", c.Max, "MAX to follow")
	flag.StringVar(&c.Login, "login", c.Login, "login")
	flag.StringVar(&c.Password, "passwd", c.Password, "password")
	flag.StringVar(&c.Search, "search", c.Search, "category")
	flag.IntVar(&c.Interval, "interval", c.Interval, "Sleep duration between follow (seconds)")

	flag.Parse()

	fmt.Printf("%+v\n", c)

	return c, nil
}

func loadFile(filename string) ([]byte, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		home, _ := os.UserHomeDir()
		buf, err = ioutil.ReadFile(filepath.Join(home, ".goinsta", filename))
		if err != nil {
			return nil, fmt.Errorf("file %q: %v", filename, err)
		}
	}
	return buf, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
