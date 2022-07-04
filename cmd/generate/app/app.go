package app

import (
	"fmt"
	"os"
	"path"

	v1 "github.com/ivanlsz/go-ssh/v1"
)

func Run() {
	homeRoot := path.Join(os.Getenv("HOME"), ".assh")
	os.MkdirAll(homeRoot, 0644)
	if _, err := os.Stat(path.Join(homeRoot, "private.pem")); err != nil {
		err = v1.RSAGenKey(2048, homeRoot)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}

	if _, err := os.Stat(path.Join(homeRoot, "public.pem")); err != nil {
		err = v1.RSAGenKey(2048, homeRoot)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}
}
