package load

import (
	"fmt"
	"os/exec"

	"github.com/HPE-Japan-Presales/ez-airgap/internal"
	"github.com/HPE-Japan-Presales/ez-airgap/internal/global"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	Cmd = &cli.Command{
		Name:      "load",
		Usage:     "load container images on local",
		UsageText: fmt.Sprintf("%s load <YOUR_IMAGE_DIR_PATH>", internal.AppName),
		Action: func(c *cli.Context) error {
			path := c.Args().Get(0)
			if path == "" {
				err := fmt.Errorf("not set images directory path")
				log.Debugf("%#v", err.Error())
				return err
			}
			client := &client{
				dirPath: path,
			}
			//load images
			if err := client.run(); err != nil {
				log.Debugf("%#v", err.Error())
				return err
			}
			return nil
		},
	}
)

type client struct {
	dirPath string
}

func (c *client) getDirPath() string {
	if c != nil {
		return c.dirPath
	}
	return ""
}

func (c *client) run() error {
	log.Debug("execute pull client")
	if err := global.CheckDockerCmdPath(); err != nil {
		log.Debugf("%#v", err.Error())
		return err
	}
	dirPath := c.getDirPath()
	dirPath = global.TrimExtention(dirPath, "/")
	imgNames, err := global.GetLocalImages(dirPath)
	if err != nil {
		log.Debugf("%#v", err.Error())
		return err
	}

	for _, img := range imgNames {
		imgPath := fmt.Sprintf("%s/%s", dirPath, img)
		log.Infof("trying to load %s", imgPath)
		if _, err := exec.Command("sh", "-c", fmt.Sprintf("docker load -i %s", imgPath)).Output(); err != nil {
			log.Debugf("%#v", err.Error())
			log.Warnf("failed to load images from %s", imgPath)
		}
	}

	return nil
}
