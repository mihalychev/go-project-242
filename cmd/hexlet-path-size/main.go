package main

import (
	"fmt"
	"log"
	"os"
	"context"

  "github.com/urfave/cli/v3" // imports as package "cli"

  "code/internal/file"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "human", Aliases: []string{"H"}},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)
			size, err := file.GetSize(path)
			if err != nil {
				return err
			}

			fmt.Println(file.FormatSize(size, cmd.Bool("human")), path)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
