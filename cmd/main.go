package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/HPE-Japan-Presales/ez-airgap/internal"
	"github.com/HPE-Japan-Presales/ez-airgap/internal/apps"
	"github.com/HPE-Japan-Presales/ez-airgap/internal/global"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	version     = "Dev"
	appUsageTxt = fmt.Sprintf("%s [global-options] [command] [options] </image/text/file/path>", internal.AppName)
	globalFlags = global.Flags
	cmds        = apps.Cmds
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "version",
	}

	app := &cli.App{
		Name:      internal.AppName,
		Usage:     internal.AppUsage,
		Flags:     globalFlags,
		UsageText: appUsageTxt,
		Commands:  cmds,
		Before: func(c *cli.Context) error {
			global.EnableDebug(c.Bool("debug"))
			return nil
		},
		Version: version,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
