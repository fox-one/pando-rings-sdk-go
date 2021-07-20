package cmd

import (
	"encoding/json"
	"time"

	"github.com/fox-one/compound-sdk-go"
	"github.com/spf13/cobra"
)

var requestTransactionCmd = cobra.Command{
	Use: "transaction",
	Run: func(cmd *cobra.Command, args []string) {
		transactions, err := compound.RequestTransactions(cmd.Context(), 50, time.Now().AddDate(0, 0, -10))
		if err != nil {
			panic(err)
		}

		bytes, err := json.MarshalIndent(transactions, "", "    ")
		if err != nil {
			panic(err)
		}

		cmd.Println(string(bytes))
	},
}

func init() {
	Root.AddCommand(&requestTransactionCmd)
}
