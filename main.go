package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/coreos/go-semver/semver"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	De "github.com/visionmedia/go-debug"
)

var debug = De.Debug("http-healthcheck:main")

func main() {
	app := cli.NewApp()
	app.Name = "http-healthcheck"
	app.Version = version()
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "url, u",
			EnvVar: "HTTP_HEALTHCHECK_URL",
			Usage:  "Fully qualified http(s) url to healthcheck.",
		},
	}
	app.Run(os.Args)
}

func run(context *cli.Context) {
	url := getOpts(context)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if resp.StatusCode > 299 {
		log.Fatalln(fmt.Errorf("Received non 2xx status code: %v", resp.StatusCode))
	}

	os.Exit(0)
}

func getOpts(context *cli.Context) string {
	url := context.String("url")

	if url == "" {
		cli.ShowAppHelp(context)

		if url == "" {
			color.Red("  Missing required flag --url, -u or HTTP_HEALTHCHECK_URL")
		}
		os.Exit(1)
	}

	return url
}

func version() string {
	version, err := semver.NewVersion(VERSION)
	if err != nil {
		errorMessage := fmt.Sprintf("Error with version number: %v", VERSION)
		log.Panicln(errorMessage, err.Error())
	}
	return version.String()
}
