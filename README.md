# Supply

Supply contains utilities for calculating the total available and circulating supply of utia on a given date.

## Usage

1. Run the API server

    ```shell
    go run main.go
    ```

2. Verify the API server is running

    ```shell
    # Query the current supply
    curl http://localhost:8080/supply

    # Query the supply at a given date
    curl http://localhost:8080/supply/2024-10-31
    ```

## Contributing

### Helpful commands

```shell
# Run all tests
make test
```
