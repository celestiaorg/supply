# Supply

This repo contains:

1. Utilities for calculating the circulating and total supply of utia on a given date.
1. An API with endpoints that return the current circulating and total supply.

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

> [!NOTE]
> Functions should ideally operate on utia values (of type `int64`) rather than TIA values (of type `float`) to avoid loss of precsision. Utia values can be converted to TIA prior to responding to API requests.
