package main

import (
	"encoding/json"
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell

// Schema is a sample definition of how you could create schemas/models around DAG entries
type Schema struct {
	Subject   string
	Predicate string
	Value     int64
}

func main() {

	// Where your local node is running on localhost:5001
	sh = shell.NewShell("localhost:5001")

	// Creating an entry as per the definition of struct. New struct, new schema!
	entry := Schema{"IPFS", "is awesome!", 007}

	// Converting into JSON object
	entryJSON, err := json.Marshal(entry)

	// Dag PUT operation which will return the CID for futher access or pinning etc.
	cid, err := sh.DagPut(entryJSON, "json", "cbor")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("\nWRITE: Added %sIPLD EXPLORER LINK: https://explore.ipld.io/#/explore/%s", string(cid+"\n"), string(cid+"\n"))

	// Fetch the details by reading the DAG for key "Subject"
	fmt.Println("\nREAD: Value for key \"Subject\": ")
	res, err := GetDag(cid, "Subject")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	// Fetch the details by reading the DAG for key "Predicate"
	fmt.Println("\nREAD: Value for key \"Predicate\": ")
	res, err = GetDag(cid, "Predicate")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	// Fetch the details by reading the DAG for key "Value"
	fmt.Println("\nREAD: Value for key \"Value\": ")
	res, err = GetDag(cid, "Value")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

// GetDag handles READ operations of a DAG entry by CID, returning the corresponding value
func GetDag(ref, key string) (out interface{}, err error) {
	err = sh.DagGet(ref+"/"+key, &out)
	return
}
