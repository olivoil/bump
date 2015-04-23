# BUMP

A utility tool written in Go. Use it in your existing Golang code or run commands via the CLI.

```
GoCode   import github.com/olivoil/bump
CLI      go get -u github.com/olivoil/bump/bump
```

## Usage from terminal

```
# bump version in files
bump patch file1.txt file2.json
```

## Usage in Go

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
