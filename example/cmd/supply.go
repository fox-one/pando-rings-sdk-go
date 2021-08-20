package cmd

import (
	"encoding/json"

	rings "github.com/fox-one/pando-rings-sdk-go"
	"github.com/fox-one/pkg/uuid"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var supplyCmd = cobra.Command{
	Use: "supply",
	Run: func(cmd *cobra.Command, args []string) {
		fID, url, err := rings.RequestSupply(cmd.Context(), uuid.New(), USDTAssetID, decimal.NewFromFloat(0.01))
		if err != nil {
			panic(err)
		}

		cmd.Println(fID, ",,", url)
	},
}

var pledgeCmd = cobra.Command{
	Use: "pledge",
	Run: func(cmd *cobra.Command, args []string) {
		fID, url, err := rings.RequestPledge(cmd.Context(), uuid.New(), cETHAssetID, decimal.NewFromFloat(0.0001))
		if err != nil {
			panic(err)
		}

		cmd.Println(fID, ",,", url)
	},
}

var unpledgeCmd = cobra.Command{
	Use: "unpledge",
	Run: func(cmd *cobra.Command, args []string) {
		fID, url, err := rings.RequestUnpledge(cmd.Context(), uuid.New(), cETHAssetID, decimal.NewFromFloat(0.0001))
		if err != nil {
			panic(err)
		}

		cmd.Println(fID, ",,", url)
	},
}

var redeemCmd = cobra.Command{
	Use: "redeem",
	Run: func(cmd *cobra.Command, args []string) {
		fID, payInput, err := rings.RequestRedeem(cmd.Context(), uuid.New(), cETHAssetID, decimal.NewFromFloat(0.0001))
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

var quickPledgeCmd = cobra.Command{
	Use: "quickpledge",
	Run: func(cmd *cobra.Command, args []string) {
		fID, payInput, err := rings.RequestQuickPledge(cmd.Context(), uuid.New(), ETHAssetID, decimal.NewFromFloat(0.0001))
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

var quickRedeemCmd = cobra.Command{
	Use: "quickredeem",
	Run: func(cmd *cobra.Command, args []string) {
		fID, payInput, err := rings.RequestQuickRedeem(cmd.Context(), uuid.New(), cETHAssetID, decimal.NewFromFloat(0.0001))
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
	Root.AddCommand(&supplyCmd)
	Root.AddCommand(&pledgeCmd)
	Root.AddCommand(&unpledgeCmd)
	Root.AddCommand(&redeemCmd)
	Root.AddCommand(&quickPledgeCmd)
	Root.AddCommand(&quickRedeemCmd)
}
