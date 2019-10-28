package main

import (
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell

func main() {
	// Where your local node is running on localhost:5001
	sh = shell.NewShell("localhost:5001")
	// sh.DagPut(`, "json", "cbor")

	cid, err := sh.DagPut(`{"x": "I","y": "<3", "z": "IPFS"}`, "json", "cbor")

	// cid, err := sh.Add(strings.NewReader("hello world!"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("WRITE: Added %s", string(cid+"\n"))

	// Fetch the details by reading the DAG for key "x"
	fmt.Println("READ: Value for key \"x\": ")
	res, err := GetDag(cid, "x")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	// Fetch the details by reading the DAG for key "y"
	fmt.Println("READ: Value for key \"y\": ")
	res, err = GetDag(cid, "y")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	// Fetch the details by reading the DAG for key "z"
	fmt.Println("READ: Value for key \"z\": ")
	res, err = GetDag(cid, "z")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func GetDag(ref, key string) (out interface{}, err error) {
	err = sh.DagGet(ref+"/"+key, &out)
	return
}
