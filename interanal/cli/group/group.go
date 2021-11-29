package group

import (
	"github.com/jessevdk/go-flags"
	"github.com/otokarev/mt5tk/pkg/model"
)

func Parse(factory *model.Factory, parser *flags.Parser, args []string) {
	parser.ParseArgs(args)
}
