package load

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const (
	dummyImageDir = "../../../test/images"
)

func TestRun(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	c := &client{
		dirPath: dummyImageDir,
	}
	if err := c.run(); err != nil {
		t.Fatal(err)
	}
}
