# provider-upjet-warpgate

`provider-upjet-warpgate` is a [Crossplane](https://crossplane.io/) provider
generated with [Upjet](https://github.com/crossplane/upjet) on top of the
[Warpgate](https://github.com/warp-tech/warpgate) Terraform provider
([`warp-tech/warpgate`](https://registry.terraform.io/providers/warp-tech/warpgate/latest)).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/fulsiram/provider-upjet-warpgate/issues).
