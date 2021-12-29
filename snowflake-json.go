package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"hash/adler32"
	"os"
	"strconv"
)

func main() {
	snowflake.Epoch = 1545264000000
	identifier := 1

	hostname, hostnamePresent := os.LookupEnv("HOSTNAME")
	if hostnamePresent {
        // 	fmt.Println("Checksum Int: ", adler32.Checksum([]byte(hostname)))
		checksum := strconv.FormatUint(uint64(adler32.Checksum([]byte(hostname))), 10)
		// 	fmt.Println("Checksum String: ", checksum)
		checkId, err := strconv.Atoi(checksum[len(checksum)-3:])

		if err != nil {
			panic(err)
		}

		identifier = checkId
	}

	nodeId := flag.Int64("i", int64(identifier), "identifier")
	num := flag.Int("n", 5, "# of iterations")
	flag.Parse()

	i := *nodeId
	n := *num

// 	fmt.Println("Identifier: ", i)

	node, err := snowflake.NewNode(i)
	if err != nil {
		fmt.Println(err)
		return
	}

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
