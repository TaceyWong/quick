# Starter

[English](./README_en.md)

+ 文档:
+ 协议: [MIT](./LICENSE)


## 功能特性

**TODO:**

- [ ] 跨平台: 官方支持Windows, Mac 和 Linux；
- [ ] 使用Starter无需懂得编写代码
- [ ] 项目模板可以用于任何编程语言或者标记格式: Python, JavaScript, Ruby, CoffeeScript, RST, Markdown, CSS, HTML 等等. 你也可以在同一项目下使用多种语言
- [ ] 简单命令行使用:
- [ ] 在终端命令行中使用:
- [ ] 或在Golang项目中作为一个package使用:
- [ ] 目录名和文件名同样可以模板。 例如:
- [ ] 支持目录的无限层级嵌套.(唯一的限制就是操作系统的限制)
- [ ] 模板化100%由pongo2完成 ，这包括文件名和目录名.
- [ ] 可以在 `starter.json` 文件中简单定义模板变量. 例如:
- [ ] 除非添加 `--no-input`选项, 你将会被交互式提示输入:
	+ 提示要输入的对象定义在`starter.json`.
	+ 缺省值为 `starter.json`中定义的.
	+ 输入提示是有序的.
- [ ] 跨平台支持`~/.starterrc` 文件:
- [ ] Starters (被克隆的Starter项目模板)缺省被放置在`~/.starters/` ,如果需要特殊指定，将会保存到指定的`starters_dir`。
- [ ] 如果你已经克隆了一个starter保存到`~/.starters/`,你可以通过目录名引用它:
- [ ] 你可以直接使用本地starters, 或远程Git仓库。
- [ ] 缺省上下文: 可以为生成项目指定键值对作为缺省值
- [ ] 使用命令行参数注入额外的上下文:
- [ ] 直接访问Starter的API允许注入额外上下文。
- [ ] 前置钩子、后置钩子: 在生成项目前后可以执行的一些额外动作（缺省支持starter预定义的一些命令，如果平台运行也可以是其他一些脚本，如Python、Shell、PowerShell
- [ ] 本地项目的路径可以指定为相对路径或绝对路径
- [ ] 缺省项目将会生成到当前目录，如果需要特别指定其他目录可以使用`-o`选项。


## Acknowledgements

+ [go-git](https://github.com/go-git/go-git) ：用于克隆git仓库
+ [chroma](https://github.com/alecthomas/chroma) ：用于代码高亮
+ [pongo2](https://github.com/flosch/pongo2) ：用于模板化
+ [viper](github.com/spf13/viper) ：用于加载配置
+ [gjson](https://github.com/tidwall/gjson) ：用于解码json
+ [urfave/cli](https://github.com/urfave/cli) ：用于命令行界面
+ [survey](https://github.com/AlecAivazis/survey) ：用于终端交互式输入
