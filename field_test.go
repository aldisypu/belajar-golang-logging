package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "putra").Info("Hello World")

	logger.WithField("username", "aldi").
		WithField("name", "Aldi Syah").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "aldi",
		"name":     "Aldi Syah",
	}).Info("Hello World")
}
