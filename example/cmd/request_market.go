package cmd

import (
	"encoding/json"

	rings "github.com/fox-one/pando-rings-sdk-go"
	"github.com/spf13/cobra"
)

var requestMarketCmd = &cobra.Command{
	Use: "market",
	Run: func(cmd *cobra.Command, args []string) {
		markets, err := rings.RequestAllMarkets(cmd.Context())
		if err != nil {
			panic(err)
		}

		bytes, err := json.MarshalIndent(markets, "", "    ")
		if err != nil {
			panic(err)
		}

		cmd.Println(string(bytes))
	},
}

func init() {
	Root.AddCommand(requestMarketCmd)
}
