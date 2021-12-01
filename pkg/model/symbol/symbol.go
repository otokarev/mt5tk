package symbol

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/otokarev/mt5tk/pkg/client"
)

type Symbol struct {
	Client *client.Client
}

type listResponse struct {
	Answer []string `json:"answer"`
}

type getRequest struct {
	Symbol string `url:"symbol"`
}

type getResponse struct {
	Answer SymbolObject `json:"answer"`
}

func (s *Symbol) List() ([]string, error) {
	payload, err := s.Client.Get("/api/symbol/list")
	if err != nil {
		return nil, err
	}
	resp := listResponse{}
	if nil != json.Unmarshal(payload, &resp) {
		return nil, err
	}

	return resp.Answer, nil
}

func (s *Symbol) Get(symbol string) (SymbolObject, error) {
	req := getRequest{Symbol: symbol}
	q, err := query.Values(req)
	if err != nil {
		return SymbolObject{}, err
	}

	payload, err := s.Client.Get("/api/symbol/get?" + q.Encode())
	if err != nil {
		return SymbolObject{}, err
	}
	resp := getResponse{}
	if nil != json.Unmarshal(payload, &resp) {
		return SymbolObject{}, err
	}

	return resp.Answer, nil
}
