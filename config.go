package main

import (
	"os"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type appConfig struct {
	listenAddress string
	githubUserURL string
	hostPrefix    string
}

func parseFlags(args []string) appConfig {
	config := appConfig{hostPrefix: "docwhat.org/"} // TODO: Make hostPrefix a flag.
	app := kingpin.New("go-importd", "Redirector for go imports for custom domains.")
	app.Writer(os.Stdout)
	app.HelpFlag.Short('h')
	app.Author("Christian Höltje https://docwhat.org/")
	app.Version(appVersion)
	app.VersionFlag.Short('v')

	app.Flag("listen", "The address to server http on").
		Short('l').
		Default(":http").
		OverrideDefaultFromEnvar("GO_IMPORTD_LISTEN").
		StringVar(&config.listenAddress)

	app.Flag("github-user-url", "The base URL on github for your projects.").
		Short('g').
		Default("https://github.com/docwhat").
		OverrideDefaultFromEnvar("GO_IMPORTD_GITHUB_USER_URL").
		StringVar(&config.githubUserURL)

	if command, err := app.Parse(args); err != nil {
		app.Usage(nil)
	} else {
		if command != "" {
			app.Usage(nil)
		}
	}

	config.githubUserURL = strings.TrimSuffix(config.githubUserURL, "/") + "/"

	return config
}
