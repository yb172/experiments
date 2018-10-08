# Postgres experiments

Experiments to see how Postgre SQL will behave or perform.

## Setup

Currently Google Cloud SQL is used to run experiments (because you can get $300 credit). Please configure cloud sql instance and specify its params in `config.yaml` (or by providing env `PG_EXPS_GOOGLE_CLOUDSQL_HOST`, `PG_EXPS_GOOGLE_CLOUDSQL_NAME`, `PG_EXPS_GOOGLE_CLOUDSQL_USER`, `PG_EXPS_GOOGLE_CLOUDSQL_PASS` env vars)

## Usage

To see list of available experiments run:

```bash
./pg help
```

## Results

It takes ~10s to insert 100K records to a given table. Terminal output:

```text
$ go build pg.go && ./pg batch-insert --count=100000
2018/10/07 20:10:47 Connected to cloudsql "***"
2018/10/07 20:10:49 Current record count: 210006
2018/10/07 20:10:49 Start inserting records
2018/10/07 20:10:59 Insertion completed. Took 9.512064137s
2018/10/07 20:10:59 Record count after update: 310006
```
