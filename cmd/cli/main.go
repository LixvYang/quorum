package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/huo-ju/quorum/cmd/cli/cache"
	"github.com/huo-ju/quorum/cmd/cli/config"
	"github.com/huo-ju/quorum/cmd/cli/ui"
	"github.com/huo-ju/quorum/internal/pkg/utils"
	logging "github.com/ipfs/go-log/v2"
)

var (
	ReleaseVersion string
	GitCommit      string
)

func main() {
	if ReleaseVersion == "" {
		ReleaseVersion = "0.0.0"
	}
	if GitCommit == "" {
		GitCommit = "devel"
	}
	version := flag.Bool("version", false, "Show the version")
	update := flag.Bool("update", false, "Update to the latest version")
	updateFrom := flag.String("from", "qingcloud", "Update from: github/qingcloud, default to qingcloud")

	flag.Parse()

	if *version {
		fmt.Printf("%s - %s\n", ReleaseVersion, GitCommit)
		return
	}

	if *update {
		mainLog := logging.Logger("main")
		lvl, _ := logging.LevelFromString("info")
		logging.SetAllLoggers(lvl)

		err := errors.New(fmt.Sprintf("invalid `-from`: %s", *updateFrom))
		if *updateFrom == "qingcloud" {
			err = utils.CheckUpdateQingCloud(ReleaseVersion, "rumcli")
		} else if *updateFrom == "github" {
			err = utils.CheckUpdate(ReleaseVersion, "rumcli")
		}
		if err != nil {
			mainLog.Fatalf("Failed to do self-update: %s\n", err.Error())
		}
		return
	}

	config.Init()
	cache.Init()
	ui.Init()

	if err := ui.App.Run(); err != nil {
		panic(err)
	}
}
