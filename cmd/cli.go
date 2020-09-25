/*
 * @Description: 填写描述
 * @Author: WangXinyong/TaceyWong
 * @Date: 2020-09-25 10:27:14
 * @LastEditors: WangXinyong/TaceyWong
 * @LastEditTime: 2020-09-25 11:59:02
 * @FilePath: /quick/cmd/cli.go
 */
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.UseShortOptionHandling = true
	app.Version = "v0.1"
	app.Name = "Quick"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Tacey Wong",
			Email: "xinyong.wang@qq.com",
		}, &cli.Author{
			Name: "All Contributors",
		},
	}
	app.Description = `Create a project from a Cookiecutter project template (TEMPLATE).

	 Quick is free and open source software,managed by Tacey Wong ,developed by
	 Tacey Wong & volunteers. 
	 If you would like to help out or fund the project, please get in touch at 
	 https://github.com/TaceyWong/quick`

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "checkout",
			Aliases: []string{"c"},
			Usage:   "branch, tag or commit to checkout after git clone",
		}, &cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"verb"},
			Value:   false,
			Usage:   "Print debug information",
		}, &cli.BoolFlag{
			Name:    "replay",
			Aliases: []string{"r"},
			Usage:   "Do not prompt for parameters and only use information entered previously",
		}, &cli.StringFlag{
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
	action := func(c *cli.Context) error {
		fmt.Printf("Hello %q", c.Args().Get(0))
		return nil
	}
	app.Usage = "Project Template Tool Like Cookiecutter"
	app.Copyright = "(c) 2020 - Forever Tacey Wong & All Contributors"
	app.ArgsUsage = "TEMPLATE [EXTRA_CONTEXT]..."
	app.Action = action
	app.Flags = flags
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
