package options

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/HPE-Japan-Presales/ez-airgap/internal/global"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

//Usage description
var (
	SaveImagesFlag = &cli.StringFlag{
		Name: "save",
		Aliases: []string{
			"s",
		},
		Usage:    "`PATH` to save images as tar file to use other servers.",
		Required: false,
	}
)

func SaveImages(path string, images []string) error {
	log.Debugf("save images into %s", path)

	path = global.TrimExtention(path, "/")
	log.Infof("image will be saved in %s", path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := fmt.Errorf("%s not exist", path)
		log.Debugf("%#v", err.Error())
		return err
	}
	for _, img := range images {
		replacer := strings.NewReplacer("/", "_", ":", "_")
		compPath := fmt.Sprintf("%s/%s", path, replacer.Replace(img))
		if _, err := os.Stat(compPath); os.IsNotExist(err) {
			log.Infof("trying to save image: %s", compPath)
			if _, err := exec.Command("sh", "-c", fmt.Sprintf("docker save -o %s %s", compPath, img)).Output(); err != nil {
				log.Warnf("failed to save: %s", compPath)
			}
		} else {
			log.Infof("%s exists...skip...", compPath)
		}

	}
	return nil
}
