package model

import (
	"github.com/otokarev/mt5tk/pkg/client"
	"github.com/otokarev/mt5tk/pkg/model/group"
	"github.com/otokarev/mt5tk/pkg/model/symbol"
)

type Factory struct {
	Client *client.Client
}

func (f *Factory) Group() *group.Group {
	return &group.Group{Client: f.Client}
}

func (f *Factory) Symbol() *symbol.Symbol {
	return &symbol.Symbol{Client: f.Client}
}
