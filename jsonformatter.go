package fnxlogrus

import (
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// JSONFormatter is a logrus.JSONFormatter but with ability to decode stacktraces from github.com/pkg/errors
type JSONFormatter struct {
	logrus.JSONFormatter
}

//Format makes sure we implement logrus.Formatter
func (jf *JSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	if _, ok := entry.Data[logrus.ErrorKey]; !ok {
		return jf.JSONFormatter.Format(entry)
	}

	if entry.Data[logrus.ErrorKey] == nil {
		delete(entry.Data, logrus.ErrorKey)
		return jf.JSONFormatter.Format(entry)
	}

	if err, ok := entry.Data[logrus.ErrorKey].(error); ok {
		b := &strings.Builder{}
		writeStack(b, errors.Cause(err))
		if b.String() != "" {
			entry.Data["stacktrace"] = b.String()
			delete(entry.Data, logrus.ErrorKey)
		}
	}

	return jf.JSONFormatter.Format(entry)
}

func writeStack(w io.Writer, err error) {
	type stackTracer interface {
		StackTrace() errors.StackTrace
		Error() string
	}
	if err, ok := err.(stackTracer); ok {
		fmt.Fprintf(w, "%s\n", err.Error())
		for _, f := range err.StackTrace() {
			fmt.Fprintf(w, "%+v\n", f)
		}
	}

}
