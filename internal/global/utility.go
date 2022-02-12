package global

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ReadText(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Debugf("%#v", err)
		return nil, err
	}
	defer file.Close()
	var images []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		images = append(images, scanner.Text())
	}
	return images, scanner.Err()
}

func GetImages(path string) ([]string, error) {
	images, err := ReadText(path)
	if err != nil {
		log.Debugf("%#v", err)
		return nil, err
	}
	imageNum := len(images)
	if imageNum < 1 {
		err := fmt.Errorf("looks no container image names in %s", path)
		log.Debugf("%#v", err)
		return nil, err
	}

	log.Debug("split image names")
	for i, img := range images {
		images[i] = strings.Split(img, " ")[0]
	}

	return images, nil
}

func TrimExtention(s, ext string) string {
	if strings.HasSuffix(s, ext) {
		s = s[:len(s)-len(ext)]
	}
	return s
}

func GetLocalImages(path string) ([]string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Debugf("%#v", err.Error())
		err := fmt.Errorf("image directory %s not found", path)
		return nil, err
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Debugf("%#v", err.Error())
		err := fmt.Errorf("image directory %s not found", path)
		return nil, err
	}
	var imagePaths []string
	for _, f := range files {
		imagePaths = append(imagePaths, f.Name())
	}
	return imagePaths, nil
}

func CheckDockerCmdPath() error {
	log.Debug("check installed Docker cli")
	if _, err := exec.Command("sh", "-c", "docker version").Output(); err != nil {
		log.Debugf("%#v", err.Error())
		return err
	}
	return nil
}
