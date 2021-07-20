package cmd

import "github.com/spf13/cobra"

const (
	USDTAssetID  = "4d8c508b-91c5-375b-92b0-ee702ed2dac5"
	ETHAssetID   = "43d61dcd-e413-450d-80b8-101d5e903357"
	cUSDTAssetID = "f8abf8be-2ead-3638-afa4-8a0b08872998"
	cETHAssetID  = "186d9355-a5e9-31f9-bed9-73ea2db016f8"
)

var Root = cobra.Command{
	Use: "compoundsdk",
}
