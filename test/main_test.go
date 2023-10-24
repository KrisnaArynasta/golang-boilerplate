package test

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	log.Print("Before Unit Test")

	m.Run()

	log.Print("After Unit Test")
}
