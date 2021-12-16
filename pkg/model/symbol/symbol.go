package symbol

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/otokarev/mt5tk/pkg/client"
	"strconv"
)

type Symbol struct {
	Client     *client.Client
	ClientPool []*client.Client
}

type listResponse struct {
	Answer []string `json:"answer"`
}

type getRequest struct {
	Symbol string `url:"symbol"`
}

type getForGroupRequest struct {
	Symbol string `url:"symbol"`
	Group  string `url:"group"`
}

type totalResponse struct {
	Answer struct {
		Total string `json:"total"`
	} `json:"answer"`
}

type getResponse struct {
	Answer SymbolObject `json:"answer"`
}

func (s *Symbol) ListNames() ([]string, error) {
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

func (s *Symbol) GetTotal() (int, error) {
	payload, err := s.Client.Get("/api/symbol/total")
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

func (s *Symbol) GetForGroup(symbol string, group string) (SymbolObject, error) {
	req := getForGroupRequest{Symbol: symbol, Group: group}
	q, err := query.Values(req)
	if err != nil {
		return SymbolObject{}, err
	}

	payload, err := s.Client.Get("/api/symbol/get_group?" + q.Encode())
	if err != nil {
		return SymbolObject{}, err
	}
	resp := getResponse{}
	if nil != json.Unmarshal(payload, &resp) {
		return SymbolObject{}, err
	}

	return resp.Answer, nil
}
