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

## Generate wire inject files (soon))

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

`project-root/
├── cmd/
│   └── main.go               # Entry point for starting the application.
│                             # Responsible for loading configuration, dependencies, and starting the server.
├── configs/
│   └── config.toml           # Application configuration file.
├── docs/
│   ├── docs.go               # Swagger documentation generation file.
│   ├── swagger.json          # Generated Swagger JSON file.
│   └── swagger.yaml          # Generated Swagger YAML file.
├── internal/
│   ├── data/                 # Database migration logic and management.
│   │   ├── migrate.go        # Handles migrations (currently for PostgreSQL, future support dynamic).
│   ├── mods/                 # Core project layers.
│   │   ├── dal/              # Data Access Layer.
│   │   ├── biz/              # Business Logic Layer.
│   │   ├── api/              # API layer with handlers and routes.
│   │   ├── schema/           # Data schemas (structs, validation).
│   │   └── mods.go           # Wire setup and route initialization.
│   ├── server/               # Server initialization logic.
│   │   └── server.go         # Creates and starts a new server instance.
│   ├── wirex/                # Google Wire setup for dependency injection.
│   │   └── wire_gen.go       # Generated dependency injection code.
├── pkg/                      # Reusable global utilities and helpers.
│   ├── errors/               # Custom error handling utilities.
│   ├── utils/                # General utility functions.
│   └── other_helpers.go      # Other global helper functions.
├── scripts/                  # Scripts for setting up and running the project in new environments.
│   └── setup_env.sh          # Environment setup script (future implementation).
├── tests/                    # Testing code for the project.
│   ├── unit/                 # Unit tests.
│   ├── integration/          # Integration tests.
│   └── test_helpers.go       # Helper functions for tests.
├── Makefile                  # Build, run, and utility commands via `


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
