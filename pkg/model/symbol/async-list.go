package symbol

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	client2 "github.com/otokarev/mt5tk/pkg/client"
	"github.com/otokarev/mt5tk/pkg/connection"
	"github.com/otokarev/mt5tk/pkg/cracker"
)

// List loads all available symbols by their indexes from 0 to total (see GetTotal)
func (s *Symbol) List() ([]SymbolObject, error) {
	total, err := s.GetTotal()
	if nil != err {
		return nil, err
	}

	cmds := prepareListCommands(total)

	return s.processListCommands(cmds)
}

// ListByNames loads all available symbols by their names retrieved by ListNames
func (s *Symbol) ListByNames() ([]SymbolObject, error) {
	names, err := s.ListNames()
	if nil != err {
		return nil, err
	}

	cmds := prepareListByNamesCommands(names)

	return s.processListCommands(cmds)
}

func (s *Symbol) processListCommands(cmds []cracker.Command) ([]SymbolObject, error) {
	rawResults := cracker.ProcessBatch(cmds, convertClientPoolToResources(s.ClientPool))

	var results []SymbolObject
	for i := 0; i < len(cmds); i++ {
		err := rawResults[i].Error()
		if err != nil {
			return nil, err
		}
		if rawResults[i].Result() == nil {
			continue
		}

		results = append(results, rawResults[i].Result().(SymbolObject))
	}

	return results, nil
}

func prepareListCommands(total int) []cracker.Command {
	var proc cracker.Proc
	proc = func(resource cracker.Resource, arg cracker.Argument) (cracker.Result, error) {
		client := resource.(*client2.Client)
		i := arg.(int)
		payload, err := client.Get(fmt.Sprintf("/api/symbol/next?index=%d", i))
		// skip not found symbols
		if cerr, ok := err.(*connection.Error); ok == true && cerr.IsNotFound() == true {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

		resp := getResponse{}
		if nil != json.Unmarshal(payload, &resp) {
			return nil, err
		}
		return cracker.Result(resp.Answer), nil
	}
	cmds := make([]cracker.Command, total)

	for i := 0; i < total; i++ {
		cmds[i] = cracker.NewCommand(proc, cracker.Argument(i), 1)
	}

	return cmds
}

func prepareListByNamesCommands(names []string) []cracker.Command {
	var proc cracker.Proc
	proc = func(resource cracker.Resource, arg cracker.Argument) (cracker.Result, error) {
		client := resource.(*client2.Client)
		name := arg.(string)
		req := getRequest{Symbol: name}
		q, err := query.Values(req)
		if err != nil {
			return SymbolObject{}, err
		}

		payload, err := client.Get("/api/symbol/get?" + q.Encode())
		if err != nil {
			return SymbolObject{}, err
		}
		resp := getResponse{}
		if nil != json.Unmarshal(payload, &resp) {
			return SymbolObject{}, err
		}
		return cracker.Result(resp.Answer), nil
	}
	cmds := make([]cracker.Command, len(names))

	for i, name := range names {
		cmds[i] = cracker.NewCommand(proc, cracker.Argument(name), 1)
	}

	return cmds
}

func convertClientPoolToResources(resources []*client2.Client) []cracker.Resource {
	r := make([]cracker.Resource, len(resources))
	for k, v := range resources {
		r[k] = cracker.Resource(v)
	}
	return r
}
