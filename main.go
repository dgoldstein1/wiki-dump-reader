package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"strconv"
)

// checks environment for required env vars
var logFatalf = log.Fatalf
var logMsg = log.Infof
var logErr = log.Errorf
var logWarn = log.Warnf

func parseEnv() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	requiredEnvs := []string{
		"GRAPH_DB_ENDPOINT",
		"TWO_WAY_KV_ENDPOINT",
		"METRICS_PORT",
	}
	for _, v := range requiredEnvs {
		if os.Getenv(v) == "" {
			logFatalf("'%s' was not set", v)
		} else {
			// print out config
			logMsg("%s=%s", v, os.Getenv(v))
		}
	}
	numberVars := []string{"PARALLELISM"}
	for _, e := range numberVars {
		i, err := strconv.Atoi(os.Getenv(e))
		if err != nil {
			logFatalf("Could not parse %s for env variable %s. Reccieve: %v", e, os.Getenv(e), err.Error())
		}
		if i < 1 && i != -1 {
			logFatalf("%s must be greater than 1 but was '%i'", e, i)
		}

	}
}

// runs parser with given functions
func runParser(file string) {
	// assert environment
	parseEnv()
	// parse with passed args
	ServeMetrics()
	Run(file)
}

func main() {
	app := cli.NewApp()
	app.Name = "wiki dump parser"
	app.Usage = " wiki-dump-parser [action] [file.xml]"
	app.Description = "parses dumps from wikipedia and indexes nodes to github.com/dgoldstein1/graphApi"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "parse",
			Aliases: []string{"p"},
			Usage:   "wiki-dump-parser [parse] [file.xml]",
			Action: func(c *cli.Context) error {
				runParser(c.Args().Get(0))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
