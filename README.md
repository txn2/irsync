[![irsync: interval rsync](https://raw.githubusercontent.com/cjimti/irsync/master/irsync.png)](https://github.com/cjimti/irsync)
# Interval [rsync]

Source: https://github.com/cjimti/irsync

[![Maintainability](https://api.codeclimate.com/v1/badges/a99a88d28ad37a79dbf6/maintainability)](https://codeclimate.com/github/codeclimate/codeclimate/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/cjimti/irsync)](https://goreportcard.com/report/github.com/cjimti/irsync)
[![GoDoc](https://godoc.org/github.com/cjimti/irsync/irsync?status.svg)](https://godoc.org/github.com/cjimti/iotwifi/irsync)

[![Docker Container Image Size](https://shields.beevelop.com/docker/image/image-size/cjimti/irsync/1.0.0.svg)](https://hub.docker.com/r/cjimti/irsync/)
[![Docker Container Layers](https://shields.beevelop.com/docker/image/layers/cjimti/irsync/1.0.0.svg)](https://hub.docker.com/r/cjimti/irsync/)
[![Docker Container Pulls](https://img.shields.io/docker/pulls/cjimti/irsync.svg)](https://hub.docker.com/r/cjimti/irsync/)

Command line utility and [Docker] container for running [rsync] on an interval.

## Demo

Setup a quick demo using [Docker]s `docker-compose` command. Included with this project is a `docker-compose.yaml` with a simple client/server setup. In this composer configuration `irsync` is set to check the server every 30 seconds (after sync is complete. The server mounts the `./data/source` directory and the client mounts the `./data/dest` directory. Drop files in `./data/source` and see them appear in `./data/dest`.

**Setup and run demo (requires [Docker]):**

```bash
# create a source and dest directories (mounted from the docker-compose)
mkdir -p ./data/source
mkdir -p ./data/dest

# make a couple of sample files
touch ./data/source/test1.txt
touch ./data/source/test2.txt

# get the docker-compose.yml
curl https://raw.githubusercontent.com/cjimti/irsync/master/docker-compose.yml >docker-compose.yml

# run docker-compose in the background (-d flag)
docker-compose up -d

# view logs
docker-compose logs -f

# drop some more files in the ./data/source directory
# irsync is configured to check every 30 seconds in this demo.

#### Cleanup

# stop containers
docker-compose stop

# remove containers
docker-compose rm

```

## Run Container

#### Example #1

Sync files from a remote server on 30 second interval

```bash
docker run --rm \
    -v "$(pwd)"/data:/data \
    -e RSYNC_PASSWORD=password \
    -e IRSYNC_INTERVAL=30 \
    -e IRSYNC_FROM=rsync://user@example.com:873/data/"\
    -e IRSYNC_TO=./data \
    -e IRSYNC_DELETE=true \
    cjimti/irsync
```

## Environment Configuration

- `IRSYNC_INTERVAL=10` Start next interval 10 seconds after the last completion.
- `IRSYNC_TIMEOUT=7200` Timeout rsync and start next interval if the time exceeds 7200 seconds.
- `IRSYNC_FROM=./` rsync from location
- `IRSYNC_TO=./data` rsync to location
- `IRSYNC_FLAGS=`-avzr` rsync flags (verbose is required)
- `IRSYNC_DELETE=false` resync --delete flag added if set to true.

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
- Run: `GITHUB_TOKEN=$GITHUB_TOKEN goreleaser --rm-dist`

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