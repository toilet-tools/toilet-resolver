package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/toilet-tools/toilet-resolver/core"
	log "github.com/toilet-tools/toilet-resolver/logger"
	"github.com/toilet-tools/toilet-resolver/utils"
	"github.com/urfave/cli/v2"
)

var (
	// flags
	domain          string
	customUserAgent string
	verbose         bool

	// app defaults
	titleBox = box.New(box.Config{Px: 2, Py: 1, Type: "Round", TitlePos: "Bottom", Color: "Blue", TitleColor: "Cyan", ContentColor: "White", AllowWrapping: true})
	logger   = log.New("<TIME>")
)

func main() {
	app := &cli.App{
		Name:     "Toilet-Resolver",
		Usage:    "Golang CloudFlare Resolver",
		Version:  "1.0.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name: "github.com/kobley",
			},
			&cli.Author{
				Name: "github.com/resetd3v",
			},
		},
		CommandNotFound: func(cCtx *cli.Context, command string) {
			fmt.Fprintf(cCtx.App.Writer, "Command %q not found.\n", command)
		},
		OnUsageError: func(cCtx *cli.Context, err error, isSubcommand bool) error {
			if isSubcommand {
				return err
			}
			fmt.Fprintf(cCtx.App.Writer, "Missing Flag: %#v\n", err.Error())
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "domain",
				Aliases:     []string{"d"},
				Usage:       "target domain to resolve",
				Destination: &domain,
			},
			&cli.StringFlag{
				Name:        "useragent",
				Aliases:     []string{"ua"},
				Value:       "random",
				Usage:       "custom useragent",
				Destination: &customUserAgent,
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Value:       false,
				Usage:       "verbosity",
				Destination: &verbose,
			},
		},
		Action: func(ctx *cli.Context) error {
			if domain == "" {
				logger.Error("No domain specified!")
				os.Exit(1)
			}

			utils.Title("Toilet-Resolver")

			core.ProcessUserAgent(customUserAgent)

			ua := ""
			if customUserAgent == "random" {
				ua = core.UserAgent
			} else {
				ua = customUserAgent
			}

			titleBox.Println(fmt.Sprintf("version: %s", utils.LGreen(ctx.App.Version)), fmt.Sprintf("%s \nDomain: %s\nUser-Agent: %s",
				utils.Cyan("╔╦╗╔═╗╦╦  ╔═╗╔╦╗  ╦═╗╔═╗╔═╗╔═╗╦ ╦  ╦╔═╗╦═╗\n")+
					utils.Blue(" ║ ║ ║║║  ║╣  ║───╠╦╝║╣ ╚═╗║ ║║ ╚╗╔╝║╣ ╠╦╝\n")+
					utils.DBlue(" ╩ ╚═╝╩╩═╝╚═╝ ╩   ╩╚═╚═╝╚═╝╚═╝╩═╝╚╝ ╚═╝╩╚═"), utils.Cyan(domain), utils.Cyan(ua)))

			core.Resolve(domain, verbose)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
