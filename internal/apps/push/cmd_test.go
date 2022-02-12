package push

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const (
	dummyTxtPath  = "../../../test/images.txt"
	dummyRegistry = "localhost:5000"
)

func TestRun(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	c := &client{
		filePath: dummyTxtPath,
		registry: dummyRegistry,
	}
	if err := c.run(); err != nil {
		t.Fatal(err)
	}
}
