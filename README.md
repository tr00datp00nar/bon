# :deciduous_tree: bon

This is `bon`. I have put it into a command branch for inclusion into my c Bonzai stateful command tree.

## Installation

If you just want to try it out, grab the release binary with curl and put into your PATH:

```
curl -L https://github.com/tr00datp00nar/bon/releases/latest/download/bon-linux-amd64 -o ~/.local/bin/bon
curl -L https://github.com/tr00datp00nar/bon/releases/latest/download/bon-darwin-amd64 -o ~/.local/bin/bon
curl -L https://github.com/tr00datp00nar/bon/releases/latest/download/bon-darwin-arm64 -o ~/.local/bin/bon
curl -L https://github.com/tr00datp00nar/bon/releases/latest/download/bon-windows-amd64 -o ~/.local/bin/bon
```

Or with `go`:

```shell
go install github.com/tr00datp00nar/bon/cmd/bon@latest
```

Composed

```go
package c

import (
	Z "github.com/rwxrob/bonzai/z"
    "github.com/tr00datp00nar/bon"
)

var Cmd = &Z.Cmd{
	Name:     'c',
    Commands: []*Z.Cmd{help.Cmd, bon.Cmd},
}
```

## Resources

To learn more about Bonzai stateful command trees: https://github.com/rwxrob/bonzai

To see my personal Bonzai stateful command tree: https://github.com/tr00datp00nar/c

To see the original Bonzai stateful command tree z: https://github.com/rwxrob/z
