/*
 * @Description: 填写描述
 * @Author: WangXinyong/TaceyWong
 * @Date: 2020-09-25 10:27:14
 * @LastEditors: WangXinyong/TaceyWong
 * @LastEditTime: 2020-09-25 13:14:46
 * @FilePath: /quick/cmd/cli.go
 */
package main

import (
	"log"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.UseShortOptionHandling = true
	app.Version = "v0.1"
	app.Name = "Starter"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Tacey Wong",
			Email: "xinyong.wang@qq.com",
		}, &cli.Author{
			Name: "All Contributors",
		},
	}
	app.Description = `Create a project from a project template (TEMPLATE).

	 Starter is free and open source software,managed by Tacey Wong ,developed by
	 Tacey Wong & volunteers. 
	 If you would like to help out or fund the project, please get in touch at 
	 https://github.com/TaceyWong/starter`

	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "no-input",
			Aliases: []string{"n"},
			Value:   false,
			Usage:   "Do not prompt for parameters and only use cookiecutter.json file content",
		}, &cli.StringFlag{
			Name:    "checkout",
			Aliases: []string{"c"},
			Usage:   "branch, tag or commit to checkout after git clone",
		}, &cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"verb"},
			Value:   false,
			Usage:   "Print debug information",
		}, &cli.BoolFlag{
			Name:    "overwrite-if-exists",
			Aliases: []string{"f"},
			Usage:   "Overwrite the contents of the output directory if it already exists",
		}, &cli.BoolFlag{
			Name:    "skip-if-file-exists",
			Aliases: []string{"s"},
			Value:   false,
			Usage:   "Skip the files in the corresponding directories if they already exist",
		}, &cli.PathFlag{
			Name:    "output-dir",
			Aliases: []string{"o"},
			Value:   ".",
			Usage:   "Where to output the generated project dir into",
		}, &cli.PathFlag{
			Name:    "config-file",
			Aliases: []string{"conf"},
			Value:   "",
			Usage:   "User configuration file",
		}, &cli.BoolFlag{
			Name:    "default-config",
			Aliases: []string{"d"},
			Value:   true,
			Usage:   "Do not load a config file. Use the defaults instead",
		}, &cli.PathFlag{
			Name:    "debug-file",
			Aliases: []string{"debug"},
			Value:   "",
			Usage:   "File to be used as a stream for DEBUG logging",
		},
	}
	action := action
	app.Usage = "Project Template Tool Like Cookiecutter"
	app.Copyright = "(c) 2020 - Forever Tacey Wong & All Contributors"
	app.ArgsUsage = "TEMPLATE [EXTRA_CONTEXT]..."
	app.Commands = []*cli.Command{
		&cli.Command{
			Name:    "list",
			Aliases: []string{"list-installed"},
			Usage:   "List installed templates",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		&cli.Command{
			Name:    "replay",
			Aliases: []string{"replay-file"},
			Usage:   "Replay entered previously",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}
	app.Action = action
	app.Flags = flags
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	sets := []string{"🌍", "🌎", "🌏"}
	s := spinner.New(sets, 100*time.Millisecond) // Build our new spinner
	s.Suffix = "正在准备中……"
	s.Color("red", "bold")
	s.FinalMSG = "失败"
	s.Start() // Start the spinner
	s.FinalMSG = "成功"
	time.Sleep(4 * time.Second) // Run for some time to simulate work
	s.Stop()
	return nil
}
