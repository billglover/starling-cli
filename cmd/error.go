package cmd

import (
	"fmt"
	"os"

	"github.com/billglover/starling"
)

func check(err error, prefix string) {
	if _, ok := err.(starling.AuthError); ok {
		fmt.Println(prefix)
		fmt.Println("authentication failed: please check the validity of your access token")
		fmt.Println("https://developer.starlingbank.com/personal/list")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
