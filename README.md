# crossplane-provider-tinkerbell

`crossplane-provider-tinkerbell` is an **experimental** [Crossplane](https://crossplane.io/) Provider for [Tinkerbell](https://tinkerbell.org). It intends to offer the following features (currently incomplete and experimental):

- A `ProviderConfig` type that only points to a credentials `Secret`.
  The value in the Secret key (identified by the `credentialsSecretRef`) should contain a json or yaml string that includes `grpc_authority` and `cert_url`.
- A `Hardware` resource type that serves as to managed Hardware resources.
- A `Template` resource type that serves as to manage Template resources.
- A `Workflow` resource type that serves as to manage Workflow resources.
- A managed resource controller that reconciles `Hardware` objects and simply
  prints their configuration in its `Observe` method.
- A managed resource controller that reconciles `Template` objects and simply
  prints their configuration in its `Observe` method.
- A managed resource controller that reconciles `Workflow` objects and simply
  prints their configuration in its `Observe` method.

## Developing

Run against a Kubernetes cluster:

```console
make run
```

Install `latest` into Kubernetes cluster where Crossplane is installed:

```console
make install
```

Install local build into [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
cluster where Crossplane is installed:

```console
make install-local
```

Build, push, and install:

```console
make all
```

Build image:

```console
make image
```

Push image:

```console
make push
```

Build binary:

```console
make build
```
