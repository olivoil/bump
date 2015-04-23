# BUMP

[![GoDoc](https://godoc.org/github.com/olivoil/bump?status.svg)](https://godoc.org/github.com/olivoil/bump)

A utility tool written in Go. Use it in your existing Golang code or run commands via the CLI.

```
GoCode   import github.com/olivoil/bump
CLI      go get -u github.com/olivoil/bump/bump
```

## Usage from terminal

```
bump patch file.json
bump minor file.txt
bump major file.txt file.json
```

## Usage in Go

Documentation available at: [http://godoc.org/github.com/olivoil/bump](http://godoc.org/github.com/olivoil/bump)

Example in a [gofer](https://github.com/chuckpreslar/gofer) task:

```
import "github.com/olivoil/bump"

var VersionPatch = gofer.Register(gofer.Task{
	Namespace:   "version",
	Label:       "patch",
	Description: "Increment version's patch number",
	Action: func(command, fileName string) error {

    return bump.File(command, fileName)

	},
})
```
