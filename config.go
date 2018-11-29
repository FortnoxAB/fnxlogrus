package fnxlogrus

import (
	"log"

	"github.com/sirupsen/logrus"
)

// Config is the configuration
type Config struct {
	Level  string
	Format interface{}
}

func (c Config) formatter() logrus.Formatter {
	switch v := c.Format.(type) {
	case logrus.Formatter:
		return v
	case string:
		if v == "json" {
			return &JSONFormatter{}
		}
	default:
		log.Fatalln("fnxlog: unknown logging format: ", v)
	}
	return nil
}
