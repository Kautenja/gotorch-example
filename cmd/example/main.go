package main

import (
    "fmt"
    "github.com/Kautenja/gotorch"
)

func main() {
    fmt.Println("Torch version", torch.TorchVersion())
    fmt.Println("GoTorch version", torch.Version())
    values := torch.RandN([]int64{10}, torch.NewTensorOptions())
    fmt.Println(values.String())
}
