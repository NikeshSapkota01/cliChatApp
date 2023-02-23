### Install these

- https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

## CREATING A MIGRATION FOLDER:

$ migrate create -ext sql -dir db/migrations add_users_table

- To run the migration
  $ make migrateup

To view changes
$ make postgres
$ \c chatApi
$ \d users
