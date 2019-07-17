package builder

import (
	"github.com/buildpack/pack/lifecycle"
)

const MetadataLabel = "io.buildpacks.builder.metadata"

type Metadata struct {
	Description string              `json:"description"`
	Buildpacks  []BuildpackMetadata `json:"buildpacks"`
	Groups      []GroupMetadata     `json:"groups"`
	Stack       StackMetadata       `json:"stack"`
	Lifecycle   lifecycle.Metadata  `json:"lifecycle"`
}

type BuildpackMetadata struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	Latest  bool   `json:"latest"`
}

type GroupMetadata struct {
	Buildpacks []GroupBuildpack `json:"buildpacks" toml:"buildpacks"`
}

type OrderMetadata2 struct {
	// TODO: Group
	Buildpacks []GroupBuildpack `json:"group" toml:"group"`
}

type OrderTOML struct {
	Order []OrderMetadata2 `toml:"order"`
}

type GroupBuildpack struct {
	ID       string `json:"id" toml:"id"`
	Version  string `json:"version" toml:"version"`
	Optional bool   `json:"optional,omitempty" toml:"optional,omitempty"`
}

type StackMetadata struct {
	RunImage RunImageMetadata `toml:"run-image" json:"runImage"`
}

type RunImageMetadata struct {
	Image   string   `toml:"image" json:"image"`
	Mirrors []string `toml:"mirrors" json:"mirrors"`
}

func OrderToGroups(orders []OrderMetadata2) []GroupMetadata {
	var groups []GroupMetadata
	for _, order := range orders {
		groups = append(groups, GroupMetadata{
			Buildpacks: order.Buildpacks,
		})
	}

	return groups
}
func GroupsToOrder(groups []GroupMetadata) []OrderMetadata2 {
	var orders []OrderMetadata2
	for _, group := range groups {
		orders = append(orders, OrderMetadata2{
			Buildpacks: group.Buildpacks,
		})
	}
	return orders
}
