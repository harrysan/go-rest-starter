
# Golang Rest API Starter

> A rest API service based on golang.

## Quick Start
[](https://github.com/harrrysan/go-rest-starter#quick-start)
```shell
make start
```
## Build
[](https://github.com/harrrysan/go-rest-starter#build)
```shell
make build
```
## Generate wire inject files
[](https://github.com/harrrysan/go-rest-starter#generate-wire-inject-files)
```shell
make wire
```
## Generate swagger documents
[](https://github.com/harrrysan/go-rest-starter#generate-swagger-documents)
```shell
make swagger
```
## Seed Data
[](https://github.com/harrrysan/go-rest-starter#seed-data)
```shell
make seed
```

## Folder Structure

```
.
└── finance-tracker/
    ├── cmd/       # Entry point for starting the application.
    ├── docs/      # Swagger documentation generation file.
    ├── db/        # For migrating table.
    ├── internal/
    │   ├── data/    # Handles migrations (currently for PostgreSQL, future support dynamic).
    │   ├── mods/    # Core project layers.
    │   ├── routes/  # Routes api.
    │   ├── server/  # Creates and starts a new server instance.
    │   └── wirex/   # Generated dependency injection code.
    ├── pkg/       # Reusable global utilities and helpers.
    ├── seeder/    # Seed Data into database.
    ├── tests/     # Testing code for the project.
    └── Makefile   # Build, run, and utility commands via make.
```

### Description of Each Folder

1. **`cmd/`** : Contains the entry point of the application. The `main.go` file initializes the configuration, loads dependencies, and starts the server.
2. **`configs/`** : Holds the configuration files for the application, such as `config.toml`.
3. **`docs/`** : Stores the generated Swagger documentation files, including JSON, YAML, and the generator logic.
4. **`db/`** : For migrating tables into database.
5. **`internal/`** : Encapsulates the application's internal logic.
   `* data/` : Manages database migrations. Currently supports PostgreSQL, with future dynamic database support planned.
   `* mods/` : Contains the core layers: `dal` (Data Access), `biz` (Business Logic), `api` (API routes and handlers), and `schema` (data models and validation).
   `* server/` : Responsible for creating and starting the server.
   `* routes/` : List All Routes in API.
   `* wirex/` : Handles Google Wire-based dependency injection.
6. **`pkg/`** : A collection of reusable global utilities, such as error handling and utility functions.
7. **`seeder/`** : Seed Data for various purposes in database.
8. **`tests/`** : Contains all testing-related code, separated into unit and integration test directories.
9. **`Makefile`** : Defines build and run commands using `make`.
10. **`README.md`** : Provides a guide for developers working on the project.



Happy Coding 🔥🚀
