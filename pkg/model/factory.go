package model

import (
	"github.com/otokarev/mt5tk/pkg/client"
	"github.com/otokarev/mt5tk/pkg/connection"
	"github.com/otokarev/mt5tk/pkg/model/account"
	"github.com/otokarev/mt5tk/pkg/model/group"
	"github.com/otokarev/mt5tk/pkg/model/symbol"
	"github.com/otokarev/mt5tk/pkg/model/user"
)

type Factory struct {
	connections []*connection.Connection
	clients     []*client.Client
}

func NewFactory(connects []*connection.Connection) *Factory {
	f := Factory{connections: connects}
	clients := make([]*client.Client, len(connects))
	for k, v := range connects {
		clients[k] = &client.Client{Connection: v}
	}
	f.clients = clients

	return &f
}

func (f *Factory) client() *client.Client {
	return f.clients[0]
}

func (f *Factory) clientPool() []*client.Client {
	return f.clients
}

func (f *Factory) Group() *group.Group {
	return &group.Group{Client: f.client(), ClientPool: f.clientPool()}
}

func (f *Factory) Symbol() *symbol.Symbol {
	return &symbol.Symbol{Client: f.client(), ClientPool: f.clientPool()}
}

func (f *Factory) User() *user.User {
	return &user.User{Client: f.client(), ClientPool: f.clientPool()}
}

func (f *Factory) Account() *account.Account {
	return &account.Account{Client: f.client(), ClientPool: f.clientPool()}
}
