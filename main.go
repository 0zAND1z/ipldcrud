package main

import (
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	// Where your local node is running on localhost:5001
	sh := shell.NewShell("localhost:5001")
	// sh.DagPut(`, "json", "cbor")

	cid, err := sh.DagPut(`{"x": "omg","y":"gpk"}`, "json", "cbor")

	// cid, err := sh.Add(strings.NewReader("hello world!"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s", string(cid+"\n"))

	// sh.DagGet(cid,)
}
