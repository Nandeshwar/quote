package env

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

// GetStringWithDefault reads a string from an environment variable. If the
// environment variable is not set, a default value is returned.
func GetStringWithDefault(name string, defaultValue string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		logrus.WithFields(logrus.Fields{
			name: defaultValue,
		}).Warn("Did not find environment variable, using default value.")
		return defaultValue
	}
	logrus.WithFields(logrus.Fields{
		name:         value,
		"value_type": fmt.Sprintf("%T", value),
	}).Info("Found environment variable")
	return value
}

// GetString reads a string from an environment variable.
func GetString(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		logrus.WithField("env_name", name).Fatal("Cannot proceed did not find string parameter.")
	}
	logrus.WithFields(logrus.Fields{
		name:         value,
		"value_type": fmt.Sprintf("%T", value),
	}).Info("Found environment variable")
	return value
}

// GetIntWithDefault reads an integer from an environment variable. If the
// environment variable is not set, a default value is returned.
func GetIntWithDefault(name string, defaultValue int) int {
	value, ok := os.LookupEnv(name)
	if !ok {
		logrus.WithFields(logrus.Fields{
			name: defaultValue,
		}).Warn("Did not find environment variable, using default value.")
		return defaultValue
	}
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		logrus.WithError(err).WithField(name, value).Fatal("Cannot proceed, failed to parse int")
	}
	logrus.WithFields(logrus.Fields{
		name:         value,
		"value_type": fmt.Sprintf("%T", value),
	}).Info("Found environment variable")
	return valueInt
}

// GetLogLevelWithDefault reads a log level from an environment variable. If
// the environment variable is not set, the default info level is returned.
func GetLogLevelWithDefault(name string, defaultValue logrus.Level) logrus.Level {
	value, ok := os.LookupEnv(name)
	if !ok {
		logrus.WithFields(logrus.Fields{
			name: defaultValue,
		}).Warn("Did not find environment variable, using default value")
		return defaultValue
	}
	level, err := logrus.ParseLevel(value)
	if err != nil {
		logrus.WithError(err).WithField(name, value).Fatal("Cannot proceed, failed to parse log level")
	}
	return level
}

// Get a boolean from an environment variable
func GetBoolWithDefault(name string, defaultValue bool) bool {
	var value, err = strconv.ParseBool(os.Getenv(name))

	if err != nil {
		log.Printf("Did not find boolean, using default value %s: %t (%T)\n", name, defaultValue, defaultValue)
		return defaultValue
	}

	log.Printf("Found parameter %s: %t (%T)\n", name, value, value)
	return value
}
