# Editors

There is an [editor support page](https://golang.org/doc/editors.html), the wiki contains a more comprehensive list:

* [IDEs and Plugins for Go](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins)

For a good cross-platform experience, the [Visual Studio
Code](https://code.visualstudio.com/) and
[vscode-go](https://github.com/Microsoft/vscode-go) (11,926,535 downloads
currently).

Depending on your setup, I can recommend [vim-go](https://github.com/fatih/vim-go).

A huge project currently in development is
[gopls](https://github.com/golang/tools/blob/master/gopls/README.md), the

> official Go language server developed by the Go team

![vim-go](static/vim-go.png)

For exploring a larger codebase:

* Jump back and forth to definitions (e.g. `CTRL-g d`, "Go to definition", ...)
* A call graph visualizer

## Call graphs

Install [go-callvis](https://github.com/ofabry/go-callvis), then clone a repo you want to analyze:

```
$ go-callvis -file callgraph -format png -nostd \
    -focus github.com/miku/microblob \
    github.com/miku/microblob/cmd/microblob
```

![](static/callgraph.png)
