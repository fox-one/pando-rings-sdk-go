package cmd

import (
	"encoding/json"

	rings "github.com/fox-one/pando-rings-sdk-go"
	"github.com/fox-one/pkg/uuid"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var borrowCmd = cobra.Command{
	Use: "borrow",
	Run: func(cmd *cobra.Command, args []string) {
		fID, payInput, err := rings.RequestBorrow(cmd.Context(), uuid.New(), USDTAssetID, decimal.NewFromFloat(0.1))
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

var repayCmd = cobra.Command{
	Use: "repay",
	Run: func(cmd *cobra.Command, args []string) {
		fID, payInput, err := rings.RequestRepay(cmd.Context(), uuid.New(), USDTAssetID, decimal.NewFromFloat(0.1))
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

var quickBorrowCmd = cobra.Command{
	Use: "quickborrow",
	Run: func(cmd *cobra.Command, args []string) {
		fID, payInput, err := rings.RequestQuickBorrow(cmd.Context(), uuid.New(), ETHAssetID, decimal.NewFromFloat(0.001), USDTAssetID, decimal.NewFromFloat(0.01))
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
	Root.AddCommand(&borrowCmd)
	Root.AddCommand(&repayCmd)
	Root.AddCommand(&quickBorrowCmd)
}
