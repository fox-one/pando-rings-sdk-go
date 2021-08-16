package cmd

import (
	rings "github.com/fox-one/pando-rings-sdk-go"
	"github.com/fox-one/pkg/uuid"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var liquidateCmd = cobra.Command{
	Use: "liquidate",
	Run: func(cmd *cobra.Command, args []string) {
		fID, url, err := rings.RequestLiquidate(cmd.Context(), uuid.New(), "8be122b4-596f-4e4f-a307-978bed0ffb75", ETHAssetID, USDTAssetID, decimal.NewFromFloat(0.1))
		if err != nil {
			panic(err)
		}

		cmd.Println(fID, ",,", url)
	},
}

func init() {
	Root.AddCommand(&liquidateCmd)
}
