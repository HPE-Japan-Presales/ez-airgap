package apps

import (
	"github.com/HPE-Japan-Presales/ez-airgap/internal/apps/load"
	"github.com/HPE-Japan-Presales/ez-airgap/internal/apps/pull"
	"github.com/HPE-Japan-Presales/ez-airgap/internal/apps/push"
	"github.com/urfave/cli/v2"
)

var (
	Cmds = []*cli.Command{
		pull.Cmd,
		load.Cmd,
		push.Cmd,
	}
)
