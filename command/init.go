package command

import "github.com/df-mc/dragonfly/server/cmd"

func InitCommands() {
	cmd.Register(cmd.New("/fill", "Fills a region with blocks.", []string{"/set"}, Fill{}))
	cmd.Register(cmd.New("/wand", "Gives an wand to player.", nil, Wand{}))
	cmd.Register(cmd.New("/undo", "Undoes the last action.", nil, Undo{}))
	cmd.Register(cmd.New("/copy", "Copies the selected region.", nil, Copy{}))
	cmd.Register(cmd.New("/paste", "Pastes the copied region.", nil, Paste{}))
}
