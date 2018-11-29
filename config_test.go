package fnxlogrus

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestConfigCustomFormatter(t *testing.T) {
	f := &logrus.TextFormatter{}
	c := Config{Format: f}
	assert.Equal(t, f, c.formatter())
}
func TestConfigJSONFormatter(t *testing.T) {
	c := Config{Format: "json"}
	f := &JSONFormatter{}
	assert.Equal(t, f, c.formatter())
}
