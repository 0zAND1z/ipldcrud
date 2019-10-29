package main

import (
	"bufio"
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
	Value     int
}

func main() {

	// Where your local node is running on localhost:5001
	sh = shell.NewShell("https://ipfs.infura.io:5001")

	fmt.Println("MMMMMMMMMMMMMMWX0kdlldk0XWMMMMMMMMMMMMMM\nMMMMMMMMMMWNKOxdoddooddodxOKNWMMMMMMMMMM\nMMMMMMMWX0kxdoddddddddddddodxk0XWMMMMMMM\nMMMWNKOxdoddddddddddddddddddddodxOKNWMMM\nMWKkdooddddddddddddddddddddddddddoodkKWM\nMNxcllloddddddddddddddddddddddddolllcxNM\nMNxooollllooddddddddddddddddoolllloooxNM\nMNxoddddoolllloodddddddddollllooddddoxNM\nMNxoddddddddolllllllllllllloddddddddoxNM\nMNxodddddddddddolcccccclodddddddddddoxNM\nMNxoddddddddddddolccccloddddddddddddoxNM\nMNxoddddddddddddddlcclddddddddddddddoxNM\nMNxoddddddddddddddoccoddddddddddddddoxNM\nMNxlodddddddddddddoccodddddddddddddolxNM\nMNklloddddddddddddoccoddddddddddddollkNM\nMMWX0kddddddddddddoccoddddddddddddk0XNMM\nMMMMMWN0OxdoddddddoccoddddddodxOKNWMMMMM\nMMMMMMMMWNXOkdodddoccodddodkOXNMMMMMMMMM\nMMMMMMMMMMMMWX0OxolccloxO0XWMMMMMMMMMMMM\nMMMMMMMMMMMMMMMWN0dlld0NWMMMMMMMMMMMMMMM\n")

	fmt.Println("### ######  #######  #####  \n #  #     # #       #     # \n #  #     # #       #       \n #  ######  #####    #####  \n #  #       #             # \n #  #       #       #     # \n### #       #        #####  \n")

	fmt.Println("###########################\n   Welcome to IPLD-CRUD!\n###########################\n")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the value for Subject field: ")
	subject, _ := reader.ReadString('\n')

	fmt.Println("Enter the value for Predicate field: ")
	predicate, _ := reader.ReadString('\n')

	fmt.Println("Enter the value for Value field: ")
	var value int
	_, err := fmt.Scan(&value)

	// Creating an entry as per the definition of struct. New struct, new schema!
	// entry := Schema{"IPFS", "is awesome!", 007}
	entry := Schema{subject, predicate, value}

	// Converting into JSON object
	entryJSON, err := json.Marshal(entry)

	// Dag PUT operation which will return the CID for futher access or pinning etc.
	cid, err := sh.DagPut(entryJSON, "json", "cbor")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Println("------\nOUTPUT\n------")
	fmt.Printf("WRITE: Successfully added %sHere's the IPLD Explorer link: https://explore.ipld.io/#/explore/%s \n", string(cid+"\n"), string(cid+"\n"))

	// Fetch the details by reading the DAG for key "Subject"
	fmt.Println("READ: Value for key \"Subject\": ")
	res, err := GetDag(cid, "Subject")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	// Fetch the details by reading the DAG for key "Predicate"
	fmt.Println("READ: Value for key \"Predicate\": ")
	res, err = GetDag(cid, "Predicate")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	// Fetch the details by reading the DAG for key "Value"
	fmt.Println("READ: Value for key \"Value\": ")
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
