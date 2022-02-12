package pull

import (
	"fmt"
	"os/exec"

	"github.com/HPE-Japan-Presales/ez-airgap/internal"
	"github.com/HPE-Japan-Presales/ez-airgap/internal/apps/pull/options"
	"github.com/HPE-Japan-Presales/ez-airgap/internal/global"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	Cmd = &cli.Command{
		Name:      "pull",
		Usage:     "pull container images",
		UsageText: fmt.Sprintf("%s pull <YOUR_IMAGE_LIST_PATH>", internal.AppName),
		Flags: []cli.Flag{
			options.SaveImagesFlag,
		},
		Action: func(c *cli.Context) error {
			path := c.Args().Get(0)
			if path == "" {
				err := fmt.Errorf("not set images text")
				log.Debugf("%#v", err.Error())
				return err
			}
			client := &client{
				filePath: path,
			}

			//Pull images
			if err := client.run(); err != nil {
				log.Debugf("%#v", err.Error())
				return err
			}

			//Save images
			savePath := c.String("save")
			if savePath != "" {
				if err := options.SaveImages(savePath, client.getImages()); err != nil {
					log.Debugf("%#v", err.Error())
					return err
				}
			} else {
				log.Info("skip saving images")
			}
			return nil
		},
	}
)

type client struct {
	filePath     string
	images       []string
	failedImages []string
}

func (c *client) getFilePath() string {
	if c != nil {
		return c.filePath
	}
	return ""
}

func (c *client) getImages() []string {
	if c != nil {
		return c.images
	}
	return nil
}

func (c *client) readImages() error {
	log.Debug("read text file")
	images, err := global.GetImages(c.getFilePath())
	if err != nil {
		log.Debugf("%#v", err.Error())
		return err
	}
	c.images = images

	return nil
}

func (c *client) run() error {
	log.Debug("execute pull client")
	if err := global.CheckDockerCmdPath(); err != nil {
		log.Debugf("%#v", err.Error())
		return err
	}

	if err := c.readImages(); err != nil {
		log.Debugf("%#v", err.Error())
		return err
	}

	for _, img := range c.images {
		log.Infof("trying to pull image: %s", img)
		if _, err := exec.Command("sh", "-c", fmt.Sprintf("docker pull %s", img)).Output(); err != nil {
			log.Warnf("failed to pull: %s", img)
			c.failedImages = append(c.failedImages, img)
		}
	}

	failed_num := len(c.failedImages)
	if failed_num > 0 {
		log.Warnf("Failed to pull images : %v total", failed_num)
		for _, img := range c.failedImages {
			log.Warnf("%s", img)
		}
	}
	return nil
}
