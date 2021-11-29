package symbol

import (
	"fmt"
	"github.com/spf13/cobra"
)

func buildList() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List symbols",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(modelFactory.Symbol().List())
		},
	}

	return listCmd
}
