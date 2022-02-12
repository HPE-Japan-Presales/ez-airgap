package push

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
		Name:      "push",
		Usage:     "push container images",
		UsageText: fmt.Sprintf("%s push <YOUR_IMAGE_LIST_PATH> <YOUR_PRIVATE_REGISTRY_ADDRESS>", internal.AppName),
		Action: func(c *cli.Context) error {
			path := c.Args().Get(0)
			if path == "" {
				err := fmt.Errorf("not set images text")
				log.Debugf("%#v", err.Error())
				return err
			}
			registry := c.Args().Get(1)
			if path == "" {
				err := fmt.Errorf("not set private registry")
				log.Debugf("%#v", err.Error())
				return err
			}
			client := &client{
				filePath: path,
				registry: registry,
			}
			//rename and push images
			if err := client.run(); err != nil {
				log.Debugf("%#v", err.Error())
				return err
			}

			return nil
		},
	}
)

type client struct {
	filePath     string
	registry     string
	images       []string
	failedImages []string
}

func (c *client) getFilePath() string {
	if c != nil {
		return c.filePath
	}
	return ""
}

func (c *client) getRegistry() string {
	if c != nil {
		return c.registry
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
		log.Debug("rename images for private registry")
		renamedImg := fmt.Sprintf("%s/%s", c.getRegistry(), img)
		if _, err := exec.Command("sh", "-c", fmt.Sprintf("docker tag %s %s", img, renamedImg)).Output(); err != nil {
			log.Warnf("failed to tag %s to %s", img, renamedImg)
			c.failedImages = append(c.failedImages, img)
		}else{
			log.Infof("trying to push image: %s", renamedImg)
			if _, err := exec.Command("sh", "-c", fmt.Sprintf("docker push %s", renamedImg)).Output(); err != nil {
				log.Warnf("failed to push: %s", img)
				c.failedImages = append(c.failedImages, img)
			}
		log.Debug("remove renamed image")
		if _, err := exec.Command("sh", "-c", fmt.Sprintf("docker rmi %s",renamedImg)).Output(); err != nil {
			log.Warnf("failed to remove renamed image: %s", renamedImg)
		}
		}

	}

	failed_num := len(c.failedImages)
	if failed_num > 0 {
		log.Warnf("Failed to push images : %v total", failed_num)
		for _, img := range c.failedImages {
			log.Warnf("%s", img)
		}
	}
	return nil
}
