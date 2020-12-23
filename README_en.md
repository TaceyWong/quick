# Starter

[中文](./README.md)


+ Document:
+ License: MIT


## Features

**TODO:**

- [ ] Cross-platform: Windows, Mac, and Linux are officially supported.
- [ ] You don't have to know/write programming code to use Starter
- [ ] Project templates can be in any programming language or markup format: Python, JavaScript, Ruby, CoffeeScript, RST, Markdown, CSS, HTML, you name it. You can use multiple languages in the same project template.
- [ ] Simple command line usage:
- [ ] Use it at the command line with a local template:
- [ ] Or use it from Golang:
- [ ] Directory names and filenames can be templated. For example:
- [ ] Supports unlimited levels of directory nesting.(Only limit is os-system-limit)
- [ ] 100% of templating is done with pongo2. This includes file and directory names.
- [ ] Simply define your template variables in a `starter.json` file. For example:
- [ ] Unless you suppress it with `--no-input`, you are prompted for input:
	+ Prompts are the keys in `starter.json`.
	+ Default responses are the values in `starter.json`.
	+ Prompts are shown in order.
- [ ] Cross-platform support for `~/.starterrc` files:
- [ ] Starters (cloned Starter project templates) are put into ~/.starters/ by default, or starters_dir if specified.
- [ ] If you have already cloned a starter into ~/.starters/, you can reference it by directory name:
- [ ] You can use local starters, or remote starters directly from Git repos or from Mercurial repos on Bitbucket.
- [ ] Default context: specify key/value pairs that you want used as defaults whenever you generate a project.
- [ ] Inject extra context with command-line arguments:
- [ ] Direct access to the Starter API allows for injection of extra context.
- [ ] Pre- and post-generate hooks: Python or shell scripts to run before or after generating a project.
- [ ] Paths to local projects can be specified as absolute or relative.
- [ ] Projects generated to your current directory or to target directory if specified with -o option.


## Acknowledgements

+ [go-git](https://github.com/go-git/go-git) for git-clone
+ [chroma](https://github.com/alecthomas/chroma) for code highlight
+ [pongo2](https://github.com/flosch/pongo2) for template
+ [viper](github.com/spf13/viper) for loading config
+ [gjson](https://github.com/tidwall/gjson) for decoding JSON
+ [urfave/cli](https://github.com/urfave/cli) for CLI
+ [survey](https://github.com/AlecAivazis/survey) for interactive prompts
