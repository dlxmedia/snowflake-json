package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "math/rand"
  "github.com/bwmarrin/snowflake"
)

func main() {
  snowflake.Epoch = 1545264000000

  node, err := snowflake.NewNode(int64(rand.Intn(100)))
  if err != nil {
    fmt.Println(err)
    return
  }

  num := flag.Int("n", 5, "# of iterations")
  flag.Parse()

  n := *num

  ids := make([]string, 0)
	for i := 0; i < n; i++ {
		ids = append(ids, node.Generate().String())
	}

  jsonResponse, err := json.Marshal(ids)
	if err != nil {
		panic(err)
	}

  fmt.Println(string(jsonResponse))
}
