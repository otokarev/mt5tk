package position

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/otokarev/mt5tk/pkg/client"
	"github.com/otokarev/mt5tk/pkg/connection"
	"github.com/otokarev/mt5tk/pkg/model/entities"
	"log"
	"strconv"
)

type Position struct {
	Client     *client.Client
	ClientPool []*client.Client
}

type getRequest struct {
	User   string `url:"login"`
	Symbol string `url:"symbol"`
}

type getResponse struct {
	Answer entities.Position `json:"answer"`
}

type totalResponse struct {
	Answer struct {
		Total string `json:"total"`
	} `json:"answer"`
}

func (g *Position) GetTotal() (int, error) {
	payload, err := g.Client.Get("/api/position/get_total")
	if err != nil {
		return 0, err
	}
	resp := totalResponse{}
	if nil != json.Unmarshal(payload, &resp) {
		return 0, err
	}
	total, err := strconv.Atoi(resp.Answer.Total)
	if nil != err {
		return 0, err
	}

	return total, nil
}

func (g *Position) Exists(user string, symbol string) (bool, error) {
	_, err := g.Get(user, symbol)
	if err == nil {
		return true, nil
	}

	if cerr, ok := err.(*connection.Error); ok == true && cerr.IsNotFound() == false {
		log.Fatalf("cannot verify position (user: %s, symbol: %s) existance, error: %s", user, symbol, err.Error())
		return false, cerr
	}

	return false, nil
}

func (g *Position) Get(user string, symbol string) (entities.Position, error) {
	req := getRequest{User: user, Symbol: symbol}
	q, err := query.Values(req)
	if err != nil {
		return entities.Position{}, err
	}

	payload, err := g.Client.Get("/api/position/get?" + q.Encode())
	if err != nil {
		return entities.Position{}, err
	}
	resp := getResponse{}
	if err := json.Unmarshal(payload, &resp); err != nil {
		return entities.Position{}, err
	}

	return resp.Answer, nil
}
