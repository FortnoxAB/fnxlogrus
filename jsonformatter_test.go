package fnxlogrus

import (
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus"
)

func TestJSONFormatter(t *testing.T) {

	var tests = []struct {
		err         error
		contains    []string
		notContains []string
	}{
		{
			err: errors.New("test error"),
			contains: []string{
				"fnxlogrus.TestJSONFormatter",
				"stacktrace",
			},
		},
		{
			err:         nil,
			contains:    []string{"nil"},
			notContains: []string{`"error":null`},
		},
	}

	jf := &JSONFormatter{}
	for _, test := range tests {
		buf := &strings.Builder{}

		log := &logrus.Logger{
			Out:       buf,
			Formatter: jf,
			Level:     logrus.InfoLevel,
		}

		entry := logrus.NewEntry(log)
		entry.WithError(test.err).Error(test.err)

		t.Log(buf.String())
		for _, ex := range test.contains {
			assert.Contains(t, buf.String(), ex)
		}
		for _, ex := range test.notContains {
			assert.NotContains(t, buf.String(), ex)
		}
	}
}
