package cmd

import (
	"github.com/fox-one/compound-sdk-go"
	"github.com/fox-one/pkg/uuid"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var borrowCmd = cobra.Command{
	Use: "borrow",
	Run: func(cmd *cobra.Command, args []string) {
		fID, url, err := compound.RequestBorrow(cmd.Context(), uuid.New(), USDTAssetID, decimal.NewFromFloat(0.1))
		if err != nil {
			panic(err)
		}

		cmd.Println(fID, ",,", url)
	},
}

var repayCmd = cobra.Command{
	Use: "repay",
	Run: func(cmd *cobra.Command, args []string) {
		fID, url, err := compound.RequestRepay(cmd.Context(), uuid.New(), USDTAssetID, decimal.NewFromFloat(0.1))
		if err != nil {
			panic(err)
		}

		cmd.Println(fID, ",,", url)
	},
}

var quickBorrowCmd = cobra.Command{
	Use: "quickborrow",
	Run: func(cmd *cobra.Command, args []string) {
		fID, url, err := compound.RequestQuickBorrow(cmd.Context(), uuid.New(), ETHAssetID, decimal.NewFromFloat(0.001), USDTAssetID, decimal.NewFromFloat(0.01))
		if err != nil {
			panic(err)
		}

		cmd.Println(fID, ",,", url)
	},
}

func init() {
	Root.AddCommand(&borrowCmd)
	Root.AddCommand(&repayCmd)
	Root.AddCommand(&quickBorrowCmd)
}
