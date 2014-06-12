# migr

Utility for SQL schema migration

## Commands

- gen \[SQL command] \[table name] \[description (optional)]
- list \[branch name] \[directory path (optional)]
- concat \[branch name] \[directory path (optional)]

### gen

Generates schema file for migration.

#### Synopsis

    $ migr gen alter table_name "to index the created_on"
    # => it will generate sql file like `master~20130613000213-alter-table_name-to_index_the_created_on`

#### Command Format

gen \[SQL command] \[table name] \[description (optional)]

- SQL command: String
- table name: String
- description: String (optional)

If SQL command equals to any commands which are following to, generated schema file will contain skeleton.

- alter
- create
- drop

#### Output Format

Output file format is like so;

    [current git branch name]~[YYYYMMDDhhmmss]-[SQL command]-[table name]-[description (applied s/ /_/g)]

Such a file is placed in the current directory.

### list

Shows schema files list that is filtered by branch name.

### concat

Concats schema files that are filtered by branch name in chronological order.

## License

MIT
