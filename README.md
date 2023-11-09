# Supply

Supply contains utilities for calculating the total available and circulating supply of utia on a given date.

## Usage

1. Run the API server

    ```shell
    go run main.go
    ```

2. Verify the API server is running

    ```shell

    curl http://0.0.0.0:8080

    # Query the circulating supply
    curl http://0.0.0.0:8080/v0/circulating-supply

    # Query the total supply
    curl http://0.0.0.0:8080/v0/total-supply
    ```

## Contributing

### Helpful commands

```shell
# Get more info on available make commands.
make help
```
