package main

import (
	"github.com/fortnoxab/fnxlogrus"
	"github.com/gin-gonic/gin"
	"github.com/jonaz/ginlogrus"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {

	fnxlogrus.Init(fnxlogrus.Config{Format: "json"})

	e := gin.Default()
	// fetches a contextual logger that have some fields from the request. (ginlogrus)
	e.Any("/a", func(c *gin.Context) {
		l := ginlogrus.GetLogger(c)
		l.Info("info")
		c.String(200, "hello")
	})
	// prints stack trace of the most "upper" error in the chain.
	e.GET("/error", handler)
	logrus.Error(e.Run(":8080"))

}
func handler(c *gin.Context) {
	err := a()
	logrus.WithError(err).Error(err)
}
func a() error {
	return errors.Wrap(b(), "error in a calling b")
}
func b() error {
	return errors.Wrap(c(), "b error")
}
func c() error {
	return errors.New("c error")
}
