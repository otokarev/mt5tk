package symbol

// !!!WARNING!!! It is a autogenerated file!

import (
	"encoding/json"
	"fmt"
	client2 "github.com/otokarev/mt5tk/pkg/client"
	"github.com/otokarev/mt5tk/pkg/connection"
	"github.com/otokarev/mt5tk/pkg/cracker"
	"github.com/otokarev/mt5tk/pkg/model/entities"
)

// List loads all available objects by their indexes from 0 to total (see GetTotal)
func (s *Symbol) List() ([]entities.Symbol, error) {
	total, err := s.GetTotal()
	if nil != err {
		return nil, err
	}

	cmds := prepareListCommands(total)

	return s.processListCommands(cmds)
}

func (s *Symbol) processListCommands(cmds []cracker.Command) ([]entities.Symbol, error) {
	rawResults := cracker.ProcessBatch(cmds, convertClientPoolToResources(s.ClientPool))

	var results []entities.Symbol
	for i := 0; i < len(cmds); i++ {
		err := rawResults[i].Error()
		if err != nil {
			return nil, err
		}
		if rawResults[i].Result() == nil {
			continue
		}

		results = append(results, rawResults[i].Result().(entities.Symbol))
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

func convertClientPoolToResources(resources []*client2.Client) []cracker.Resource {
	r := make([]cracker.Resource, len(resources))
	for k, v := range resources {
		r[k] = cracker.Resource(v)
	}
	return r
}
