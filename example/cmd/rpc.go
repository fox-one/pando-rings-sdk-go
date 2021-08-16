package cmd

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fox-one/pando-rings-sdk-go/example/rpc"
	"github.com/fox-one/pkg/uuid"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var rpcRequestPayCmd = cobra.Command{
	Use: "rpc-requestpay",
	Run: func(cmd *cobra.Command, args []string) {
		client := rpcProtoBufClient()

		req := rpc.PayReq{
			AssetId:    "4d8c508b-91c5-375b-92b0-ee702ed2dac5",
			Amount:     "0.01",
			TraceId:    uuid.New(),
			FollowId:   uuid.New(),
			MemoBase64: "AQIQONJAADfaSXCp+PK9A5Erkg==",
		}
		resp, err := client.PayRequest(cmd.Context(), &req)

		if err != nil {
			panic(err)
		}

		rbs, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			panic(err)
		}

		cmd.Println(string(rbs))
	},
}

var rpcTransactionsCmd = cobra.Command{
	Use: "rpc-transactions",
	Run: func(cmd *cobra.Command, args []string) {
		client := rpcProtoBufClient()

		t := timestamppb.New(time.Now().AddDate(0, 0, -400))
		req := rpc.TransactionReq{
			Offset: t,
			Limit:  50,
		}

		resp, err := client.Transactions(cmd.Context(), &req)
		if err != nil {
			panic(err)
		}

		rbs, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			panic(err)
		}

		cmd.Println(string(rbs))
	},
}

var rpcPriceRequestCmd = cobra.Command{
	Use: "rpc-price",
	Run: func(cmd *cobra.Command, args []string) {
		client := rpcProtoBufClient()
		resp, err := client.PriceRequest(cmd.Context(), &rpc.PriceReq{})
		if err != nil {
			panic(err)
		}

		rbs, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			panic(err)
		}

		cmd.Println(string(rbs))
	},
}

var rpcMarketCmd = cobra.Command{
	Use: "rpc-market",
	Run: func(cmd *cobra.Command, args []string) {
		client := rpcProtoBufClient()

		resp, err := client.AllMarkets(cmd.Context(), &rpc.MarketReq{})
		if err != nil {
			panic(err)
		}

		rbs, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			panic(err)
		}

		cmd.Println(string(rbs))
	},
}

func rpcProtoBufClient() rpc.Compound {
	return rpc.NewCompoundProtobufClient("https://compound-test-api.fox.one", &http.Client{})
}

func init() {
	Root.AddCommand(&rpcRequestPayCmd)
	Root.AddCommand(&rpcTransactionsCmd)
	Root.AddCommand(&rpcPriceRequestCmd)
	Root.AddCommand(&rpcMarketCmd)
}
