package toggl

import (
	"encoding/json"
	"errors"
	"strconv"
)

type TClient struct {
	Name  string `json:"name"`
	Wid   int    `json:"wid"`
	Notes string `json:"wid"`
	At    string `json:"at"`
	ID    int    `json:"id"`
}

type TClients []TClient

func (repository TClients) FindByID(id int) (TClient, error) {
	for _, item := range repository {
		if item.ID == id {
			return item, nil
		}
	}
	return TClient{}, errors.New("Find Failed")
}

func (cl *Client) FetchWorkspaceClients(wid int) (TClients, error) {
	var clients TClients

	res, err := cl.do("GET", "/workspaces/"+strconv.Itoa(wid)+"/clients", nil)
	if err != nil {
		return TClients{}, err
	}

	enc := json.NewDecoder(res.Body)
	if err := enc.Decode(&clients); err != nil {
		return TClients{}, err
	}

	return clients, nil
}
