package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// formatter makes logrus use UTC and short filenames instead of full paths.
type formatter struct {
	lf logrus.Formatter
}

// Format satisfies the logrus.Formatter interface.
func (f *formatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	_, e.Caller.File = filepath.Split(e.Caller.File)
	return f.lf.Format(e)
}

// Init configures logrus according to our preferred standards.
func Init() {
	logrus.SetFormatter(&formatter{&logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.000000+00:00",
	}})
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
}
