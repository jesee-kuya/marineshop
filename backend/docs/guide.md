# Database Guide

## Prerequisites

- [PostgreSQL](https://www.postgresql.org/) installed and running
- [golang-migrate](https://github.com/golang-migrate/migrate) CLI installed
- `.env` file configured (see `.env.example`)

---

## Initial Setup

Run the init script as a PostgreSQL superuser to create the user and database:

```bash
psql -U postgres -f docs/init_database.sql
```

This will:
- Create the `marineshop_user` role
- Create the `marineshop` database
- Transfer ownership and grant privileges

---

## Running Migrations

```bash
# Apply all pending migrations
make up

# Roll back the last migration
make down
```

---

## Cleaning the Database

To completely wipe and recreate the database from scratch:

```bash
psql -U postgres -f docs/clean_database.sql
```

> **Warning:** This drops the database and user entirely. All data will be lost.

After cleaning, re-run the migrations to rebuild the schema:

```bash
make up
```
