# Makefile

## Usage Notes for Makefile

- currently a work in progress
- goal is using go templating to facilitate code generation

## Targets

### `all`

- **Description**: Runs tests and builds the project.
- **Usage**: `make all`

### `init`

- **Description**: Initializes a new Go module and creates the standard project directory structure. Will not run if the module is already initialized. If `URL_PATH` is set, it uses `URL_PATH` as part of the module path.
- **Usage**: `make init`

### `build`

- **Description**: Builds the project and places the binary in the specified directory.
- **Usage**: `make build`

### `test`

- **Description**: Runs tests for all packages except those in the `/test/` directory.
- **Usage**: `make test`

### `test-with-cover`

- **Description**: Runs tests with coverage for all packages except those in the `/test/` directory. Generates a coverage report.
- **Usage**: `make test-with-cover`

### `generate-mocks`

- **Description**: Generates mocks using `mockery` with specific options.
- **Usage**: `make generate-mocks`

### `clean`

- **Description**: Cleans the build and removes the binary and vendor directories.
- **Usage**: `make clean`

### `run`

- **Description**: Builds the project and runs the main Go application.
- **Usage**: `make run`

### `deps`

- **Description**: Fetches all dependencies.
- **Usage**: `make deps`

### `mod`

- **Description**: Downloads and tidies up modules and creates a vendor directory.
- **Usage**: `make mod`

### `prod`

- **Description**: Builds the project for production using vendored dependencies and without debug symbols.
- **Usage**: `make prod`

### `asm`

- **Description**: Generates assembly output for debugging and optimization.
- **Usage**: `make asm`

### `lint`

- **Description**: Runs linting on the code with all enabled checks.
- **Usage**: `make lint`

---
