package pack

import (
	"context"

	"github.com/Masterminds/semver"
	"github.com/buildpack/imgutil"

	"github.com/buildpack/pack/blob"
)

//go:generate mockgen -package mocks -destination mocks/image_fetcher.go github.com/buildpack/pack ImageFetcher

type ImageFetcher interface {
	Fetch(ctx context.Context, name string, daemon, pull bool) (imgutil.Image, error)
}

//go:generate mockgen -package mocks -destination mocks/blob_fetcher.go github.com/buildpack/pack BlobFetcher

type BlobFetcher interface {
	FetchBuildpack(uri string) (blob.Buildpack, error)
	FetchLifecycle(version *semver.Version, uri string) (blob.Lifecycle, error)
}
