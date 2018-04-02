# Interval rSync

rsync on interval, application and container.



## Environment Configuration

- `IRSYNC_INTERVAL=10` Start next interval 10 seconds after last completion.
- `IRSYNC_TIMEOUT=7200` Timout rsync and start next interval if time exceds 7200 seconds.
- `IRSYNC_FROM=./` rsync from location
- `IRSYNC_TO=./data` rync to location
- `IRSYNC_FLAGS=`-avzr` rsync flags (verbose is required)
- `IRSYNC_DELETE=false` resync --delete flag if set to true
