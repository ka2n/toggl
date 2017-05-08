package command

import (
	"fmt"

	"github.com/sachaos/toggl/cache"
	"github.com/urfave/cli"
)

func (app *App) CmdRestart(c *cli.Context) error {
	entries, err := app.client.GetTimeEntries()

	if err != nil {
		return err
	}

	if len(entries) == 0 {
		fmt.Println("No time entry")
		return nil
	}

	response, err := app.client.PostStartTimeEntry(entries[len(entries)-1])
	if err != nil {
		return err
	}

	cache.SetCurrentTimeEntry(response.Data)
	cache.Write()

	return nil
}
