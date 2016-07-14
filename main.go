package main

import (
    "flag"
    "fmt"
    "os"
)

const DEFAULT_BLOCK_SIZE = 2048

var (
    offset int64
    inFile string
)

func main() {
    flag.StringVar(&inFile, "inFile", "", "")
    flag.Int64Var(&offset, "offset", 0, "")
    flag.Parse()

    f, err := os.Open(inFile)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    buffer := make([]byte, DEFAULT_BLOCK_SIZE)
    f.ReadAt(buffer, offset)

    fmt.Println(buffer)
}
