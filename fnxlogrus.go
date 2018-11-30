package fnxlogrus

import (
	"log"

	"github.com/sirupsen/logrus"
)

//Init initializes logrus with level and formatter
func Init(config Config, l *logrus.Logger) {
	if config.Level == "" {
		config.Level = "info"
	}

	lvl, err := logrus.ParseLevel(config.Level)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("fnxlogrus: setting logrus logging level to %s\n", lvl)
	l.SetLevel(lvl)
	if config.formatter() != nil {
		log.Printf("fnxlogrus: setting logrus logging formatter to %T\n", config.formatter())
		l.SetFormatter(config.formatter())
	}
}
