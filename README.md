# worldedit
worldedit is a library for [Dragonfly](https://github.com/df-mc/dragonfly) that provides a set of tools for manipulating blocks in the world.

## Usage
To use worldedit, install the library first:
```bash
go get github.com/alvin0319/worldedit
```

Then, most of code are same, but with following:

```go
package main

import (
	"github.com/alvin0319/worldedit"
	"github.com/alvin0319/worldedit/command"
	"github.com/alvin0319/worldedit/session"
	"github.com/alvin0319/worldedit/util"
	"github.com/df-mc/dragonfly/server"
)

func main() {
	util.InitCustomItem()
	command.InitCommands()

	// rest are same
	var srv *server.Server

	srv.Listen()

	// create a session, and assign handler to player
	for p := range srv.Accept() {
		s := session.CreateSession(p)
		p.Handle(worldedit.NewHandler(s))
	}
}
```

For detailed example, example server provided [here](./example/main.go).
