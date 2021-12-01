package group

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/otokarev/mt5tk/pkg/client"
	"log"
	"strconv"
)

type Group struct {
	Client *client.Client
}

type getRequest struct {
	Group string `url:"group"`
}

type getResponse struct {
	Answer GroupObject `json:"answer"`
}

type totalResponse struct {
	Answer struct {
		Total string `json:"total"`
	} `json:"answer"`
}

func (s *Group) Get(group string) (GroupObject, error) {
	req := getRequest{Group: group}
	q, err := query.Values(req)
	if err != nil {
		return GroupObject{}, err
	}

	payload, err := s.Client.Get("/api/group/get?" + q.Encode())
	if err != nil {
		return GroupObject{}, err
	}
	resp := getResponse{}
	if nil != json.Unmarshal(payload, &resp) {
		return GroupObject{}, err
	}

	return resp.Answer, nil
}

func (s *Group) List() []GroupObject {
	payload, err := s.Client.Get("/api/group/total")
	if err != nil {
		log.Fatal(err)
	}
	resp := totalResponse{}
	if nil != json.Unmarshal(payload, &resp) {
		log.Fatalf("Cannot interpret `%s`", string(payload))
	}
	total, _ := strconv.Atoi(resp.Answer.Total)
	if nil != json.Unmarshal(payload, &resp) {
		log.Fatalf("Cannot interpret `%s`", string(payload))
	}

	results := []GroupObject{}
	for i := 0; i < total; i++ {
		payload, err := s.Client.Get(fmt.Sprintf("/api/group/next?index=%d", i))
		if err != nil {
			log.Fatal(err)
		}
		resp := getResponse{}
		if nil != json.Unmarshal(payload, &resp) {
			log.Fatal(err)
		}
		results = append(results, resp.Answer)
	}

	return results
}
