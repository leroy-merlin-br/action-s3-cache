package main

import (
	"fmt"
	"log"
	"os"

	"strings"

	"github.com/urfave/cli"
)

var (
	version = "1.0.0"
	build   = "1"
)

func main() {
	app := cli.NewApp()
	app.Name = "s3 Cache artifacts"
	app.Description = "Send artifacts to s3 to improve build time"
	app.Action = run
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "action",
			Usage:    "Action to perform. Options are: put, get, delete",
			EnvVar:   "ACTION",
			Required: true,
		},
		cli.StringFlag{
			Name:     "aws-access-key-id",
			Usage:    "AWS access key id to access your bucket",
			EnvVar:   "AWS_ACCESS_KEY_ID",
			Required: true,
		},
		cli.StringFlag{
			Name:     "aws-secret-access-key",
			Usage:    "AWS secret access key to access your bucket",
			EnvVar:   "AWS_SECRET_ACCESS_KEY",
			Required: true,
		},
		cli.StringFlag{
			Name:     "bucket",
			Usage:    "AWS s3 bucket to store the artifacts",
			EnvVar:   "BUCKET",
			Required: true,
		},
		cli.StringFlag{
			Name:     "key",
			Usage:    "An explicit key for restoring and saving the cache",
			EnvVar:   "KEY",
			Required: true,
		},
		cli.StringFlag{
			Name:   "artifacts",
			Usage:  "A list of files, directories and wildcard patterns to cache and restore",
			EnvVar: "ARTIFACTS",
			// Only required if action arg is == "put"
			Required: GetOsArgValue(os.Args, os.Getenv("ACTION"), "action") == PutAction,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	action := Action{
		Action:             c.String("action"),
		AwsSecretAccessKey: c.String("aws-secret-access-key"),
		AwsAccessKeyId:     c.String("aws-access-key-id"),
		Bucket:             c.String("bucket"),
		Key:                c.String("key"),
		Artifacts:          strings.Split(strings.TrimSpace(c.String("artifacts")), "\n"),
	}

	return action.Exec()
}
