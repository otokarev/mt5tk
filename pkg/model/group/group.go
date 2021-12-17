package group

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/otokarev/mt5tk/pkg/client"
	"github.com/otokarev/mt5tk/pkg/connection"
	"log"
	"strconv"
)

type Group struct {
	Client     *client.Client
	ClientPool []*client.Client
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

func (g *Group) GetTotal() (int, error) {
	payload, err := g.Client.Get("/api/group/total")
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

func (g *Group) Add(data GroupObject) (GroupObject, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return GroupObject{}, err
	}
	resultBytes, err := g.Client.Post("/api/group/add", body)
	if err != nil {
		return GroupObject{}, err
	}

	var result GroupObject
	err = json.Unmarshal(resultBytes, &result)
	if err != nil {
		return GroupObject{}, err
	}

	return result, err
}

func (g *Group) Exists(group string) (bool, error) {
	_, err := g.Get(group)
	if err == nil {
		return true, nil
	}

	if cerr, ok := err.(*connection.Error); ok == true && cerr.IsNotFound() == false {
		log.Fatalf("cannot verify group %s existance, error: %s", group, err.Error())
		return false, cerr
	}

	return false, nil
}

func (g *Group) Get(group string) (GroupObject, error) {
	req := getRequest{Group: group}
	q, err := query.Values(req)
	if err != nil {
		return GroupObject{}, err
	}

	payload, err := g.Client.Get("/api/group/get?" + q.Encode())
	if err != nil {
		return GroupObject{}, err
	}
	resp := getResponse{}
	if nil != json.Unmarshal(payload, &resp) {
		return GroupObject{}, err
	}

	return resp.Answer, nil
}
