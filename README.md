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

    # Query the circulating supply on a particular date
    curl http://0.0.0.0:8080/v0/circulating-supply?date=2025-05-31

    # Query the total supply on a particular date
    curl http://0.0.0.0:8080/v0/total-supply?date=2025-05-31
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

- The [docker.yml](./.github/workflows/docker.yml) workflow builds a Docker image on every release tag and pushes it to GitHub Container Registry (GHCR).
- The production deployment runs on a **Scaleway Kapsule** (managed Kubernetes) cluster. All manifests live in [celestiaorg/infrastructure](https://github.com/celestiaorg/infrastructure) under `k8s/websites/supply/`.
- The deployment runs 2 replicas in the `websites` namespace, pinned to a specific image tag (e.g. `ghcr.io/celestiaorg/supply:v1.0.0`). It does **not** auto-update.
- Two hostnames are served:
  - **Production**: https://supply.celestia.org
  - **Staging**: https://supply.celestia.dev

### Updating the production deployment

Every Monday at 09:00 UTC, a [workflow in the infrastructure repo](https://github.com/celestiaorg/infrastructure/blob/main/.github/workflows/supply-version-check.yml) checks for new supply releases and automatically opens a PR bumping the image tag if a newer version is found.

To manually update or apply a change:

1. Open a PR in [celestiaorg/infrastructure](https://github.com/celestiaorg/infrastructure) updating the image tag in `k8s/websites/supply/deployment.yaml`.
2. After the PR is merged, apply the change:

    ```shell
    kubectl apply -f k8s/websites/supply/deployment.yaml
    ```

3. Verify the rollout:

    ```shell
    curl https://supply.celestia.org/version
    ```
