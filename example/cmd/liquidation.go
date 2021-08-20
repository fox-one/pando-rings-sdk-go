package cmd

import (
	"encoding/json"

	rings "github.com/fox-one/pando-rings-sdk-go"
	"github.com/fox-one/pkg/uuid"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var liquidateCmd = cobra.Command{
	Use: "liquidate",
	Run: func(cmd *cobra.Command, args []string) {
		fID, payInput, err := rings.RequestLiquidate(cmd.Context(), uuid.New(), "8be122b4-596f-4e4f-a307-978bed0ffb75", ETHAssetID, USDTAssetID, decimal.NewFromFloat(0.1))
		if err != nil {
			panic(err)
		}

		cmd.Println("followID:", fID)
		pbs, err := json.MarshalIndent(payInput, "", "    ")
		if err != nil {
			panic(err)
		}

		cmd.Println(string(pbs))
	},
}

func init() {
	Root.AddCommand(&liquidateCmd)
}
