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
	Buildpacks []BuildpackRefMetadata `json:"buildpacks"`
}

type BuildpackRefMetadata struct {
	ID       string `json:"id"`
	Version  string `json:"version"`
	Optional bool   `json:"optional,omitempty"`
}

type StackMetadata struct {
	RunImage RunImageMetadata `json:"runImage"`
}

type RunImageMetadata struct {
	Image   string   `json:"image"`
	Mirrors []string `json:"mirrors"`
}

// TODO: Move these?
func OrderConfigToGroupMetadata(order OrderConfig) []GroupMetadata {
	var groups []GroupMetadata
	for _, gp := range order {
		var buildpacks []BuildpackRefMetadata
		for _, bp := range gp.Group {
			buildpacks = append(buildpacks, BuildpackRefMetadata{
				ID:       bp.ID,
				Version:  bp.Version,
				Optional: bp.Optional,
			})
		}

		groups = append(groups, GroupMetadata{
			Buildpacks: buildpacks,
		})
	}

	return groups
}

func GroupMetadataToOrderConfig(groups []GroupMetadata) OrderConfig {
	var order []GroupConfig

	for _, group := range groups {
		var buildpacks []GroupBuildpackConfig
		for _, bp := range group.Buildpacks {
			buildpacks = append(buildpacks, GroupBuildpackConfig{
				ID:       bp.ID,
				Version:  bp.Version,
				Optional: bp.Optional,
			})
		}

		order = append(order, GroupConfig{
			Group: buildpacks,
		})
	}

	return order
}
