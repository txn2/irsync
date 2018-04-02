# Interval [rsync]

Source: https://github.com/cjimti/irsync

[![Go Report Card](https://goreportcard.com/badge/github.com/cjimti/irsync)](https://goreportcard.com/report/github.com/cjimti/irsync)
[![GoDoc](https://godoc.org/github.com/cjimti/iotwifi/irsync?status.svg)](https://godoc.org/github.com/cjimti/iotwifi/irsync)
[![Docker Container Image Size](https://shields.beevelop.com/docker/image/image-size/cjimti/irsync/1.0.0.svg)](https://hub.docker.com/r/cjimti/irsync/)
[![Docker Container Layers](https://shields.beevelop.com/docker/image/layers/cjimti/irsync/1.0.0.svg)](https://hub.docker.com/r/cjimti/irsync/)
[![Docker Container Pulls](https://img.shields.io/docker/pulls/cjimti/irsync.svg)](https://hub.docker.com/r/cjimti/irsync/)

Command line utility and [Docker] container for running [rsync] on interval.

## Run Container

#### Example #1

Sync files from a remote server on 30 second interval

```bash
docker run --rm \
-v "$(pwd)"/data:/data \
-e RSYNC_PASSWORD=password \
-e IRSYNC_INTERVAL=30 \
-e IRSYNC_FROM="rsync://user@example.com:873/data/" \
-e IRSYNC_TO="./data" \
-e IRSYNC_DELETE=true \
cjimti/irsync:1.0.0
```

## Environment Configuration

- `IRSYNC_INTERVAL=10` Start next interval 10 seconds after last completion.
- `IRSYNC_TIMEOUT=7200` Timout rsync and start next interval if time exceds 7200 seconds.
- `IRSYNC_FROM=./` rsync from location
- `IRSYNC_TO=./data` rync to location
- `IRSYNC_FLAGS=`-avzr` rsync flags (verbose is required)
- `IRSYNC_DELETE=false` resync --delete flag if set to true

## Development

### Building and Releasing

Interval [rsync] uses [GORELEASER] to build binaries and [Docker] containers.

#### Test Release Steps

Install [GORELEASER] with [brew] (MacOS):
```bash
brew install goreleaser/tap/goreleaser
```

Build without releasing:
```bash
goreleaser --skip-publish --rm-dist --skip-validate
```

#### Release Steps

- Commit latest changes
- [Tag] a version `git tag -a v1.4 -m "my version 1.4"`
- Push tag `git push origin v1.5`
- Run: `goreleaser --rm-dist`

## Resources

- [GORELEASER]
- [Docker]
- [rsync]
- [homebrew]

[homebrew]: https://brew.sh/
[brew]: https://brew.sh/
[GORELEASER]: https://goreleaser.com/
[Docker]: https://www.docker.com/
[rsync]: https://en.wikipedia.org/wiki/Rsync
[Tag]: https://git-scm.com/book/en/v2/Git-Basics-Tagging