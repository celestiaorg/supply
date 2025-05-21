# Supply

This repo contains:

1. Utilities for calculating the circulating and total supply of utia on a given date.
1. An API with endpoints that return the current circulating and total supply.

> [!NOTE]  
> Note on methodology: due to complexity, this supply API does not adjust the circulating supply for tokens that were retroactively locked after TGE, for example due to CIP-31. As a result, the reported circulating supply may include a minor margin of error (likely no more than around 1%), decreasing to zero as lockups elapse.

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
> Functions should ideally operate on utia values (of type `int64`) rather than TIA values (of type `float64`) to avoid loss of precsision. Utia values can be converted to TIA values (of type `float64`) prior to responding to API requests.

## Deployment

The [docker.yml](./.github/workflows/docker.yml) workflow builds a Docker image and pushes it to a Scaleway container registry. The supply server is deployed via a Scaleway serverless container. CI will not automatically deploy the latest Docker image to the serverless container so you must do that manually if you want to deploy a new version of the supply server.
