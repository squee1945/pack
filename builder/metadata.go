package builder

import (
	"github.com/buildpack/pack/lifecycle"
)

const MetadataLabel = "io.buildpacks.builder.metadata"

type Metadata struct {
	Description string              `json:"description"`
	Buildpacks  []BuildpackMetadata `json:"buildpacks"`
	Groups      OrderMetadata       `json:"groups"`
	Stack       StackMetadata       `json:"stack"`
	Lifecycle   lifecycle.Metadata  `json:"lifecycle"`
}

type BuildpackMetadata struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	Latest  bool   `json:"latest"`
}

type OrderMetadata []GroupMetadata

func (o OrderMetadata) ToConfig() OrderConfig {
	var order OrderConfig

	for _, group := range o {
		var buildpacks []BuildpackRefConfig
		for _, bp := range group.Buildpacks {
			buildpacks = append(buildpacks, BuildpackRefConfig{
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
