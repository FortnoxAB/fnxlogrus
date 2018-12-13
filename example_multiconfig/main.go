package main

import (
	"github.com/fortnoxab/fnxlogrus"
	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"
)

type config struct {
	Log fnxlogrus.Config
}

func main() {
	c := &config{}
	multiconfig.MustLoad(c)
	fnxlogrus.Init(c.Log, logrus.StandardLogger())

}

// Example run: CONFIG_LOG_LEVEL=error go run main.go
// go run main.go --help
