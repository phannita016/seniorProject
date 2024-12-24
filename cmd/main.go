package main

import (
	"github.com/phannita016/seniorProject/apps"
	"github.com/spf13/cobra"
)

func main() {
	root := apps.NewAppsRoot()
	cobra.CheckErr(root.Execute())
}
