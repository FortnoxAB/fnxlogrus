package fnxlogrus

import (
	"log"

	"github.com/sirupsen/logrus"
)

// Init initializes logrus with level and formatter
func Init(config Config, l *logrus.Logger) {
	if config.Level == "" {
		config.Level = "info"
	}

	lvl, err := logrus.ParseLevel(config.Level)
	if err != nil {
		log.Fatalln(err)
	}

	l.SetLevel(lvl)
	if config.formatter() != nil {
		l.SetFormatter(config.formatter())
	}
}
