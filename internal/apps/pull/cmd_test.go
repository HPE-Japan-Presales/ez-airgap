package pull

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const (
	dummyTxtPath = "../../../test/images.txt"
)

func TestCheckDockerCmdPath(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	c := &client{}
	if err := c.checkDockerCmdPath(); err != nil {
		t.Fatal(err)
	}
}

func TestReadImages(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	c := &client{
		filePath: dummyTxtPath,
	}
	if err := c.readImages(); err != nil {
		t.Fatal(err)
	}
	if len(c.images) < 1 {
		t.Fatal("could not get images list")
	}
	t.Logf("%v", c.images)
}

func TestRun(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	c := &client{
		filePath: dummyTxtPath,
	}
	if err := c.run(); err != nil {
		t.Fatal(err)
	}
}
