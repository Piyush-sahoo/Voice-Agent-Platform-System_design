# Database Migrations

Database migration scripts using Flyway or Goose.

## Structure

```
migrations/
├── postgres/
│   ├── V001__initial_schema.sql
│   ├── V002__add_api_keys.sql
│   └── ...
└── mongodb/
    └── init-collections.js
```

## Running Migrations

```bash
# PostgreSQL (using Flyway)
flyway -url=jdbc:postgresql://localhost:5432/vap_platform migrate

# Or using Goose
goose -dir=./postgres postgres "host=localhost dbname=vap_platform sslmode=disable" up
```
