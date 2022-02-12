package global

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

const (
	dummyTxtPath  = "../../test/ezmeral_5.3_k8s_images.txt"
	dummyImageDir = "../../test/images"
)

func TestGetImages(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	images, err := GetImages(dummyTxtPath)
	if err != nil {
		t.Fatal(err)
	}
	num := len(images)
	if num < 1 {
		t.Fatal("could not get image names")
	}
	t.Logf("Total images: %v", num)
	t.Logf("%v", images)
}

func TestGetLocalImages(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	imagePaths, err := GetLocalImages(dummyImageDir)
	if err != nil {
		t.Fatal(err)
	}
	num := len(imagePaths)
	if num < 1 {
		t.Fatal("could not parse image paths")
	}
	t.Logf("Total images: %v", num)
	t.Logf("%v", imagePaths)
}
