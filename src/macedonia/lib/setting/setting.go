package setting

import (
	"github.com/jinzhu/configor"
)

type setting struct {
	Test       string `required:"true"`
	PrivateKey string `required:"true"`
	Password   string `required:"true"`
}

// Setting is settings for app
var Setting = setting{}

func init() {
	c := configor.New(&configor.Config{ENVPrefix: "APP"})
	if err := c.Load("setting/config.yaml", "service/macedonia/setting/config.yaml"); err != nil {
		panic(err)
	}
}
