package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/bblfsh/sdk.v2/build"
)

func main() {
	flag.Parse()
	if err := runUpdate("."); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runUpdate(root string) error {
	return build.UpdateSDK(root, &build.UpdateOptions{
		Notice:  fmt.Printf,
		Warning: fmt.Printf,
	})
}
