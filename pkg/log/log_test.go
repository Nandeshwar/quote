package log

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLog(t *testing.T) {

	Convey("Log a message", t, func() {
		Init()
		b := bytes.Buffer{}
		logrus.SetOutput(&b)
		msg := "this is a test"
		logrus.Info(msg)
		So(b.String(), ShouldContainSubstring, msg)
	})

}
