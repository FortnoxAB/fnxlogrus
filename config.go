package fnxlogrus

import (
	"github.com/sirupsen/logrus"
)

// Config is the configuration
type Config struct {
	Level  string
	Format interface{}
	//Formatter field is used if you need a custom formatter
	Formatter logrus.Formatter
}

func (c Config) formatter() logrus.Formatter {
	if c.Formatter != nil {
		return c.Formatter
	}
	if c.Format == "json" {
		return &JSONFormatter{}
	}
	return &logrus.TextFormatter{FullTimestamp: true}
}
