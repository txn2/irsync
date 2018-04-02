# Interval [rsync]

Command line utility and [Docker] container for running [rsync] on interval.



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