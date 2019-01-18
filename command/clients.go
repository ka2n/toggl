package command

import (
	"strconv"

	"github.com/sachaos/toggl/cache"
	toggl "github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func (app *App) getClients(c *cli.Context) (clients toggl.TClients, err error) {
	clients = cache.GetContent().Clients
	if len(clients) == 0 || !c.GlobalBool("cache") {
		clients, err = app.client.FetchWorkspaceClients(viper.GetInt("wid"))
		cache.SetClients(clients)
		cache.Write()
	}
	return
}

func (app *App) CmdClients(c *cli.Context) error {
	clients, err := app.getClients(c)
	if err != nil {
		return err
	}

	writer := NewWriter(c)

	defer writer.Flush()

	for _, client := range clients {
		writer.Write([]string{strconv.Itoa(client.ID), client.Name})
	}

	return nil
}
