package blob

import (
	"strings"
)

type Buildpack struct {
	Info   BuildpackInfo `toml:"buildpack"`
	Stacks []Stack       `toml:"stacks"`
	Blob   `toml:"-"`
}

type BuildpackInfo struct {
	ID      string `toml:"id"`
	Latest  bool   `toml:"latest"`
	Version string `toml:"version"`
}

type Stack struct {
	ID string
}

func (b *Buildpack) EscapedID() string {
	return strings.Replace(b.Info.ID, "/", "_", -1)
}

func (b *Buildpack) SupportsStack(stackID string) bool {
	for _, stack := range b.Stacks {
		if stack.ID == stackID {
			return true
		}
	}
	return false
}
