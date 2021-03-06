package deal

// !!!WARNING!!! It is a autogenerated file!

import (
	"github.com/otokarev/mt5tk/internal/cmd/util/output"
	"github.com/spf13/cobra"
	"log"
)

func buildGet() *cobra.Command {
	var id string
	var outputFormat string

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get deal's details",
		Run: func(cmd *cobra.Command, args []string) {
			results, err := modelFactory.Deal().Get(id)
			if err != nil {
				log.Fatal(err)
			}

			if err := output.Print(outputFormat, results); err != nil {
				log.Fatal(err)
			}
		},
	}
	getCmd.Flags().StringVarP(&outputFormat, "output", "o", "json", "json|jsonpath=<pattern>")
	getCmd.Flags().StringVarP(&id, "deal", "k", "", "deal identifier")
	getCmd.MarkFlagRequired("deal")

	return getCmd
}
