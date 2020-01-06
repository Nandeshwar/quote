package env

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetStringWithDefault(t *testing.T) {

	Convey("GetStringWithDefault", t, func() {
		logrus.SetOutput(ioutil.Discard)
		varName := "URL"
		defer os.Unsetenv(varName)

		Convey("Read environment variable", func() {
			os.Setenv(varName, "localhost")
			So(GetStringWithDefault(varName, ""), ShouldEqual, "localhost")
		})

		Convey("Read unset environment variable and set default value", func() {
			So(GetStringWithDefault(varName, "localhost"), ShouldEqual, "localhost")
		})

		Convey("Read in an empty string", func() {
			os.Setenv(varName, "")
			So(GetStringWithDefault(varName, "localhost"), ShouldEqual, "")
		})
	})
}

func TestGetString(t *testing.T) {

	Convey("GetString", t, func() {
		logrus.SetOutput(ioutil.Discard)
		varName := "URL"
		defer os.Unsetenv(varName)

		Convey("Read environment variable", func() {
			os.Setenv(varName, "localhost")
			So(GetString(varName), ShouldEqual, "localhost")
		})
	})
}

func TestGetIntWithDefault(t *testing.T) {

	Convey("GetIntWithDefault", t, func() {
		logrus.SetOutput(ioutil.Discard)
		varName := "NUMBER"
		defer os.Unsetenv(varName)

		Convey("Read environment variable", func() {
			os.Setenv(varName, "123")
			So(GetIntWithDefault(varName, 0), ShouldEqual, 123)
		})

		Convey("Read unset environment variable and set default value", func() {
			So(GetIntWithDefault(varName, 123), ShouldEqual, 123)
		})
	})
}

func TestGetLogLevelWithDefault(t *testing.T) {
	Convey("GetLogLevelWithDefault", t, func() {
		logrus.SetOutput(ioutil.Discard)
		varName := "LOG_LEVEL"
		defer os.Unsetenv(varName)

		Convey("Read environment variable", func() {
			os.Setenv(varName, "debug")
			So(GetLogLevelWithDefault(varName, logrus.DebugLevel), ShouldEqual, logrus.DebugLevel)
		})

		Convey("Read unset environment variable and set default value", func() {
			So(GetLogLevelWithDefault(varName, logrus.DebugLevel), ShouldEqual, logrus.DebugLevel)
		})
	})
}

func TestGetBoolWithDefault(t *testing.T) {
	Convey("GetBoolWithDefault", t, func() {
		varName := "MY_BOOL"
		defer os.Unsetenv(varName)

		Convey("Read environment variable", func() {
			os.Setenv(varName, "false")
			So(GetBoolWithDefault(varName, true), ShouldEqual, false)
		})

		Convey("Read unset environment variable and set default value", func() {
			So(GetBoolWithDefault(varName, true), ShouldEqual, true)
		})
	})
}
