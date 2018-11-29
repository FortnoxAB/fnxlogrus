package fnxlogrus

import (
	"log"

	"github.com/sirupsen/logrus"
)

//Init initializes logrus with level and formatter
func Init(config Config) {
	if config.Level == "" {
		config.Level = "info"
	}

	lvl, err := logrus.ParseLevel(config.Level)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Setting logrus logging level to %s\n", lvl)
	logrus.SetLevel(lvl)
	if config.Format == "json" {
		log.Println("Setting logrus logging formatter to json")
		logrus.SetFormatter(config.formatter())
	}
}
