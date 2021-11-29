package symbol

import (
	"encoding/json"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/otokarev/mt5tk/pkg/model"
	"log"
)

var modelFactory *model.Factory

type ListCommand struct {
}

var listCommand ListCommand

func (x *ListCommand) Execute(args []string) error {
	fmt.Println(modelFactory.Symbol().List())
	return nil
}

type GetCommand struct {
	File string `short:"s" long:"symbol" description:"Symbol" required:"true"`
}

var getCommand GetCommand

func (x *GetCommand) Execute(args []string) error {
	json, err := json.MarshalIndent(modelFactory.Symbol().Get("SPOT"), "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
	return nil
}

func Parse(factory *model.Factory, parser *flags.Parser, args []string) {
	modelFactory = factory
	parser.AddCommand("list",
		"List symbols",
		"List all symbols",
		&listCommand)
	parser.AddCommand("get",
		"Get symbol",
		"Get symbol description",
		&getCommand)
	parser.ParseArgs(args)
}
