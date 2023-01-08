# GoTorch Example

This project contains an example of using [gotorch][gotorch].

## Usage

To clone the repository (including the gotorch submodule:)

```shell
git clone https://github.com/Kautenja/gotorch-example --recursive
```

To download libtorch and compile the gotorch wrapper:

```shell
./pkg/gotorch/build.sh
```

To run the example script

```shell
go run cmd/example/main.go
```

[gotorch]: https://github.com/Kautenja/gotorch
