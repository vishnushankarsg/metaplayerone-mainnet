package main

import (
	_ "embed"

	"github.com/vishnushankarsg/geth/command/root"
	"github.com/vishnushankarsg/geth/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
