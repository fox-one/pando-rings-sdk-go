package cmd

import (
	"github.com/fox-one/compound-sdk-go"
	"github.com/fox-one/pkg/uuid"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var supplyCmd = cobra.Command{
	Use: "supply",
	Run: func(cmd *cobra.Command, args []string) {
		fID, url, err := compound.RequestSupply(cmd.Context(), uuid.New(), USDTAssetID, decimal.NewFromFloat(0.01))
		if err != nil {
			panic(err)
		}

		cmd.Println(fID, ",,", url)
	},
}

var pledgeCmd = cobra.Command{
	Use: "pledge",
	Run: func(cmd *cobra.Command, args []string) {
		fID, url, err := compound.RequestPledge(cmd.Context(), uuid.New(), cETHAssetID, decimal.NewFromFloat(0.0001))
		if err != nil {
			panic(err)
		}

		cmd.Println(fID, ",,", url)
	},
}

var unpledgeCmd = cobra.Command{
	Use: "unpledge",
	Run: func(cmd *cobra.Command, args []string) {
		fID, url, err := compound.RequestUnpledge(cmd.Context(), uuid.New(), cETHAssetID, decimal.NewFromFloat(0.0001))
		if err != nil {
			panic(err)
		}

		cmd.Println(fID, ",,", url)
	},
}

var redeemCmd = cobra.Command{
	Use: "redeem",
	Run: func(cmd *cobra.Command, args []string) {
		fID, url, err := compound.RequestRedeem(cmd.Context(), uuid.New(), cETHAssetID, decimal.NewFromFloat(0.0001))
		if err != nil {
			panic(err)
		}

		cmd.Println(fID, ",,", url)
	},
}

var quickPledgeCmd = cobra.Command{
	Use: "quickpledge",
	Run: func(cmd *cobra.Command, args []string) {
		fID, url, err := compound.RequestQuickPledge(cmd.Context(), uuid.New(), ETHAssetID, decimal.NewFromFloat(0.0001))
		if err != nil {
			panic(err)
		}

		cmd.Println(fID, ",,", url)
	},
}

var quickRedeemCmd = cobra.Command{
	Use: "quickredeem",
	Run: func(cmd *cobra.Command, args []string) {
		fID, url, err := compound.RequestQuickRedeem(cmd.Context(), uuid.New(), cETHAssetID, decimal.NewFromFloat(0.0001))
		if err != nil {
			panic(err)
		}

		cmd.Println(fID, ",,", url)
	},
}

func init() {
	Root.AddCommand(&supplyCmd)
	Root.AddCommand(&pledgeCmd)
	Root.AddCommand(&unpledgeCmd)
	Root.AddCommand(&redeemCmd)
	Root.AddCommand(&quickPledgeCmd)
	Root.AddCommand(&quickRedeemCmd)
}
