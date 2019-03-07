#!/usr/bin/env bash

cd $(dirname "${BASH_SOURCE[0]}")

docker build --build-arg base=packs/base:v3alpha2 -t packs/base:nolifecycle .
docker build --build-arg base=packs/build:v3alpha2 -t packs/build:nolifecycle .
docker build --build-arg base=packs/run:v3alpha2 -t packs/run:nolifecycle .
docker build --build-arg base=packs/samples:v3alpha2 -t packs/samples:nolifecycle .