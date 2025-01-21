# Finance Tracker API

> A finance tracker API service based on golang.

## Quick Start

[](https://github.com/harrrysan/finance-tracker#quick-start)

```shell
make start
```

## Build

[](https://github.com/harrrysan/finance-tracker#build)

```shell
make build
```

## Generate wire inject files (soon)

[](https://github.com/harrrysan/finance-tracker#generate-wire-inject-files)

```shell
make wire
```

## Generate swagger documents

[](https://github.com/harrrysan/finance-tracker#generate-swagger-documents)

```shell
make swagger
```

## Folder Structure

`.
└── finance-tracker/
    ├── cmd/       # Entry point for starting the application.
    ├── docs/      # Swagger documentation generation file.
    ├── internal/
    │   ├── data/    # Handles migrations (currently for PostgreSQL, future support dynamic).
    │   ├── mods/    # Core project layers.
    │   ├── server/  # Creates and starts a new server instance.
    │   └── wirex/   # Generated dependency injection code.
    ├── pkg/       # Reusable global utilities and helpers.
    ├── scripts/   # Environment setup script (future implementation).
    ├── tests/     # Testing code for the project.
    └── Makefile   # Build, run, and utility commands via `


### Description of Each Folder/File

1. **`cmd/`** : Contains the entry point of the application. The `main.go` file initializes the configuration, loads dependencies, and starts the server.
2. **`configs/`** : Holds the configuration files for the application, such as `config.toml`.
3. **`docs/`** : Stores the generated Swagger documentation files, including JSON, YAML, and the generator logic.
4. **`internal/`** : Encapsulates the application's internal logic.

* **`data/`** : Manages database migrations. Currently supports PostgreSQL, with future dynamic database support planned.
* **`mods/`** : Contains the core layers: `dal` (Data Access), `biz` (Business Logic), `api` (API routes and handlers), and `schema` (data models and validation).
* **`server/`** : Responsible for creating and starting the server.
* **`wirex/`** : Handles Google Wire-based dependency injection.

1. **`pkg/`** : A collection of reusable global utilities, such as error handling and utility functions.
2. **`scripts/`** : Placeholder for scripts to set up the environment or other automation tasks.
3. **`tests/`** : Contains all testing-related code, separated into unit and integration test directories.
4. **`Makefile`** : Defines build and run commands using `make`.
5. **`README.md`** : Provides a guide for developers working on the project.


Happy Coding :D

.
