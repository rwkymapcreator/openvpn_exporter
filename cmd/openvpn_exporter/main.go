package main

import (
	"os"

	"github.com/rwkymapcreator/openvpn_exporter/pkg/command"
)

func main() {
	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
