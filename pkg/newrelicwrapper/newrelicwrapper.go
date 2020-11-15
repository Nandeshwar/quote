package newrelicwrapper

import (
	"os"

	"github.com/newrelic/go-agent"
	"github.com/sirupsen/logrus"

	"quote/pkg/env"
)

type transaction struct {
	name        string
	transaction newrelic.Transaction
}

type segment struct {
	segment newrelic.Segment
}

var newRelicKey = env.GetStringWithDefault("NEWRELIC_KEY", "96a09863fbfaa78fc259467b19aa298a8a1dNRAL")
var newRelicAppName = env.GetStringWithDefault("NEWRELIC_APP", "Quote")
var newRelicDebugLogs = env.GetBoolWithDefault("NEWRELIC_DEBUG_LOGS", false)
var NewRelicApplication newrelic.Application

func init() {
	newRelicEnabled := (newRelicKey != "") || (newRelicAppName != "")

	if !newRelicEnabled {
		logrus.Info("New Relic agent is disabled\n")
		return
	}

	logrus.Infof("Starting New Relic agent for NEWRELIC_APP='%s' using NEWRELIC_KEY='%s'\n",
		newRelicAppName, newRelicKey[0:5])

	config := newrelic.NewConfig(newRelicAppName, newRelicKey)
	if newRelicDebugLogs == true {
		config.Logger = newrelic.NewDebugLogger(os.Stdout)
		logrus.Info("Enabling New Relic Debug Level Logs")
	} else {
		config.Logger = newrelic.NewLogger(os.Stdout)
		logrus.Info("Enabling New Relic Info Level Logs")
	}

	var err error
	NewRelicApplication, err = newrelic.NewApplication(config)
	if err == nil {
		logrus.Info("Started New Relic agent\n")
	} else {
		logrus.Errorf("Failed to start New Relic agent: %s\n", err)
	}
}

func StartTransaction(name string) transaction {
	if NewRelicApplication == nil {
		return transaction{
			name:        name,
			transaction: nil,
		}
	}

	return transaction{
		name:        name,
		transaction: NewRelicApplication.StartTransaction(name, nil, nil),
	}
}

func (t transaction) End() {
	if NewRelicApplication == nil {
		return
	}

	t.transaction.End()
}

func (t transaction) AddAttribute(key string, value interface{}) error {
	if NewRelicApplication == nil {
		return nil
	}

	return t.transaction.AddAttribute(key, value)
}

func (t transaction) NoticeError(err error) error {
	if NewRelicApplication == nil {
		return nil
	}

	return t.transaction.NoticeError(err)
}

func (t transaction) StartSegment(name string) segment {
	if NewRelicApplication == nil {
		return segment{
			segment: newrelic.Segment{Name: name},
		}
	}

	return segment{
		segment: *newrelic.StartSegment(t.transaction, name),
	}
}

func (s segment) End() {
	if NewRelicApplication == nil {
		return
	}

	s.segment.End()
}
