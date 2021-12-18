package symbol

import "encoding/json"

type listNamesResponse struct {
	Answer []string `json:"answer"`
}

func (s *Symbol) ListNames() ([]string, error) {
	payload, err := s.Client.Get("/api/symbol/list")
	if err != nil {
		return nil, err
	}
	resp := listNamesResponse{}
	if nil != json.Unmarshal(payload, &resp) {
		return nil, err
	}

	return resp.Answer, nil
}
