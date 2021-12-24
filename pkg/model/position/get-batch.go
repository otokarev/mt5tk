package position

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/otokarev/mt5tk/pkg/model/entities"
	"strings"
)

type getBatchRequest struct {
	User   string `url:"login,omitempty"`
	Symbol string `url:"symbol,omitempty"`
	Group  string `url:"group,omitempty"`
	Ticket string `url:"ticket,omitempty"`
}

type getBatchResponse struct {
	Answer []entities.Position `json:"answer"`
}

func (g *Position) GetBatch(tickets []string, users []string, groups []string, symbols []string) ([]entities.Position, error) {
	req := getBatchRequest{
		User:   strings.Join(users, ", "),
		Symbol: strings.Join(symbols, ","),
		Group:  strings.Join(groups, ","),
		Ticket: strings.Join(tickets, ","),
	}
	q, err := query.Values(req)
	if err != nil {
		return []entities.Position{}, err
	}

	payload, err := g.Client.Get("/api/position/get_batch?" + q.Encode())
	if err != nil {
		return []entities.Position{}, err
	}
	resp := getBatchResponse{}
	if err := json.Unmarshal(payload, &resp); err != nil {
		return []entities.Position{}, err
	}

	return resp.Answer, nil
}
