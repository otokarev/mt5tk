package symbol

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/otokarev/mt5tk/pkg/client"
	"log"
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

func (s *Symbol) List() []string {
	payload, err := s.Client.Get("/api/symbol/list")
	if err != nil {
		log.Fatal(err)
	}
	resp := listResponse{}
	if nil != json.Unmarshal(payload, &resp) {
		log.Fatal(err)
	}

	return resp.Answer
}

func (s *Symbol) Get(symbol string) SymbolObject {
	req := getRequest{Symbol: symbol}
	q, err := query.Values(req)
	if err != nil {
		log.Fatal(err)
	}

	payload, err := s.Client.Get("/api/symbol/get?" + q.Encode())
	if err != nil {
		log.Fatal(err)
	}
	resp := getResponse{}
	if nil != json.Unmarshal(payload, &resp) {
		log.Fatal(err)
	}

	return resp.Answer
}
