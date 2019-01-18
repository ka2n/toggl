package command

import (
	"strconv"

	"github.com/sachaos/toggl/cache"
	toggl "github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func (app *App) getProjects(c *cli.Context) (projects toggl.Projects, err error) {
	projects = cache.GetContent().Projects
	if len(projects) == 0 || !c.GlobalBool("cache") {
		projects, err = app.client.FetchWorkspaceProjects(viper.GetInt("wid"))
		cache.SetProjects(projects)
		cache.Write()
	}
	return
}

func (app *App) CmdProjects(c *cli.Context) error {
	projects, err := app.getProjects(c)
	if err != nil {
		return err
	}

	var clients toggl.TClients
	if c.IsSet("show-client") {
		clients, err = app.getClients(c)
		if err != nil {
			return err
		}
	}

	writer := NewWriter(c)

	defer writer.Flush()

	for _, project := range projects {
		if clients != nil {
			var clientName string
			if project.Cid != 0 {
				client, _ := clients.FindByID(project.Cid)
				clientName = client.Name
			}
			writer.Write([]string{strconv.Itoa(project.ID), clientName, project.Name})
		} else {
			writer.Write([]string{strconv.Itoa(project.ID), project.Name})
		}
	}

	return nil
}
