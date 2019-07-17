package buildpack

import (
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"

	"github.com/buildpack/pack/internal/archive"
)

//go:generate mockgen -package mocks -destination mocks/downloader.go github.com/buildpack/pack/lifecycle Downloader

type Downloader interface {
	Download(uri string) (string, error)
}

type buildpackTOML struct {
	Buildpacks []struct {
		ID      string  `toml:"id"`
		Version string  `toml:"version"`
		Stacks  []Stack `toml:"stacks"`
	} `toml:"buildpacks"`
}

type Fetcher struct {
	downloader Downloader
}

func NewFetcher(downloader Downloader) *Fetcher {
	return &Fetcher{downloader: downloader}
}

func (f *Fetcher) FetchBuildpack(uri string) (Buildpack, error) {
	downloadedPath, err := f.downloader.Download(uri)
	if err != nil {
		return Buildpack{}, errors.Wrap(err, "fetching buildpack")
	}

	data, err := readTOML(downloadedPath)
	if err != nil {
		return Buildpack{}, err
	}

	if len(data.Buildpacks) <= 0 {
		return Buildpack{}, errors.New("no buildpacks defined in buildpack TOML")
	}

	return Buildpack{ // TODO: return multiple buildpacks
		Path:    downloadedPath,
		ID:      data.Buildpacks[0].ID, // TODO: check length, filter based on
		Version: data.Buildpacks[0].Version,
		Stacks:  data.Buildpacks[0].Stacks,
	}, err
}

func readTOML(path string) (buildpackTOML, error) {
	var (
		buf []byte
		err error
	)
	if filepath.Ext(path) == ".tgz" {
		_, buf, err = archive.ReadTarEntry(path, "./buildpack.toml", "buildpack.toml", "/buildpack.toml")
	} else {
		buf, err = ioutil.ReadFile(filepath.Join(path, "buildpack.toml"))
	}

	if err != nil {
		return buildpackTOML{}, err
	}

	bpTOML := buildpackTOML{}
	_, err = toml.Decode(string(buf), &bpTOML)
	if err != nil {
		return buildpackTOML{}, errors.Wrapf(err, "reading buildpack.toml from path %s", path)
	}
	return bpTOML, nil
}
