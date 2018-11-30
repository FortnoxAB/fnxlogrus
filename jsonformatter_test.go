package fnxlogrus

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus"
)

func testWrap(err error, s string) error {
	return errors.Wrap(err, s)
}

func TestJSONFormatter(t *testing.T) {

	var tests = []struct {
		name        string
		err         error
		contains    []string
		notContains []string
	}{
		{
			name: "one error with stack",
			err:  errors.New("test error"),
			contains: []string{
				"fnxlogrus.TestJSONFormatter",
				"stacktrace",
			},
		},
		{
			name: "normal error wrapped",
			err:  errors.Wrap(fmt.Errorf("%s", "1"), "2"),
			contains: []string{
				"fnxlogrus.TestJSONFormatter",
				"stacktrace",
			},
		},
		{
			name: "normal error wrapped twice",
			err:  testWrap(errors.Wrap(fmt.Errorf("%s", "1"), "2"), "3"),
			contains: []string{
				"fnxlogrus.TestJSONFormatter",
				"stacktrace",
				"testWrap",
			},
		},
		{
			name:        "nil error",
			err:         nil,
			contains:    []string{"nil"},
			notContains: []string{`"error":null`, "stacktrace"},
		},
		{
			name:        "normal error without stack",
			err:         fmt.Errorf("%s", "1"),
			contains:    []string{`"msg":"1"`},
			notContains: []string{`"error":"1"`, "stacktrace"},
		},
	}

	jf := &JSONFormatter{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
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
		})
	}
}
