# TV Tracker

## App

API written in Go.

### Endpoints

#### Search

`/search/query string`

This will his the TVDB API's search endpoint. It will return the 1st 100 entries with no pagination.

NOTE: this may change over time!

## Web UI

Written using Vue.

## Database

It is recommended to call the database `tvtracer.db`.  This is set to be ignored by Git and it should never be committed to a public repository.

[goose](https://github.com/pressly/goose) has been used to create and update the DB schema.

### CliffsNotes

Warning: these steps were true at the time of writing, they may have changed as (hopefully) the DB schema becomes stable.

Install goose: `go install github.com/pressly/goose/v3/cmd/goose@latest`

Set the following environmental variables:

```bash
export GOOSE_DRIVER=sqlite3
export GOOSE_DBSTRING=tvtracker.db
export GOOSE_MIGRATION_DIR=database/db_definition
```

It is assumed that the `goose` command will be run at the root level of this repository.

To create you initial DB run `goose up`.