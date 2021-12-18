package symbol

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/otokarev/mt5tk/pkg/model/entities"
)

type getForGroupRequest struct {
	Symbol string `url:"symbol"`
	Group  string `url:"group"`
}

func (s *Symbol) GetForGroup(symbol string, group string) (entities.Symbol, error) {
	req := getForGroupRequest{Symbol: symbol, Group: group}
	q, err := query.Values(req)
	if err != nil {
		return entities.Symbol{}, err
	}

	payload, err := s.Client.Get("/api/symbol/get_group?" + q.Encode())
	if err != nil {
		return entities.Symbol{}, err
	}
	resp := getResponse{}
	if nil != json.Unmarshal(payload, &resp) {
		return entities.Symbol{}, err
	}

	return resp.Answer, nil
}
