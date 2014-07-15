package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/modouwifi/md/api"
)

var cmdSystem = cli.Command{
	Name:  "system",
	Usage: "",
	//Action: runSystem,
	Subcommands: []cli.Command{
		{
			Name:   "current",
			Usage:  "Display currently version",
			Action: runSystemGetVersionInfo,
		},
	},
}

func runSystemGetVersionInfo(c *cli.Context) {
	req, err := client.NewRequest("GET", "/system/get_version_info")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.SystemVersion
	err = req.ToJSON(&result)
	fmt.Println(result.Code)
	fmt.Println(result.Track)
	fmt.Println(result.Version1)
	fmt.Println(result.Version2)
}