package main

import (
	"fmt"
	"os"

	rings "github.com/fox-one/pando-rings-sdk-go"
	"github.com/fox-one/pando-rings-sdk-go/example/cmd"
	"github.com/spf13/cobra"
)

func init() {
	fmt.Println("init main")
	cobra.OnInitialize(func() {
		rings.Endpoint = "https://compound-test-api.fox.one"
	})
}

func main() {
	fmt.Println("hello compound sdk")

	if err := cmd.Root.Execute(); err != nil {
		os.Exit(1)
	}
}
