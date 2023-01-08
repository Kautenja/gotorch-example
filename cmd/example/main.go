package main

import (
    "fmt"
    "github.com/Kautenja/gotorch"
)

func main() {
    values := torch.RandN([]int64{10}, torch.NewTensorOptions())
    fmt.Println(values.String())
}
