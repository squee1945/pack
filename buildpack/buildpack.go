package buildpack

import (
	"strings"
)

type BuildpackTOML struct {
	Info   BuildpackInfo `toml:"buildpack"`
	Stacks []Stack
}

type BuildpackInfo struct {
	ID      string
	Latest  bool
	Version string
}

type Buildpack struct {
	ID      string
	Latest  bool
	Version string
	Stacks  []Stack
	Blob
}

type Stack struct {
	ID string
}

func (b *Buildpack) EscapedID() string {
	return strings.Replace(b.ID, "/", "_", -1)
}

func (b *Buildpack) SupportsStack(stackID string) bool {
	for _, stack := range b.Stacks {
		if stack.ID == stackID {
			return true
		}
	}
	return false
}
