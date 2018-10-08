# Postgres experiments

Experiments to see how Postgre SQL will behave or perform.

## Setup

Currently Google Cloud SQL is used to run experiments (because you can get $300 credit). Please configure cloud sql instance and specify its params in `config.yaml` (or by providing env `PG_EXPS_GOOGLE_CLOUDSQL_HOST`, `PG_EXPS_GOOGLE_CLOUDSQL_NAME`, `PG_EXPS_GOOGLE_CLOUDSQL_USER`, `PG_EXPS_GOOGLE_CLOUDSQL_PASS` env vars)

## Usage

To see list of available experiments run:

```bash
./pg help
```
