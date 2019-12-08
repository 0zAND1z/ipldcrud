package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

// Global variable to handle all the IPFS API client calls
var sh *shell.Shell

// ipfsURL is a constant string pointing to the IPFS API endpoint.
// Use "https://ipfs.infura.io:5001" if you are not running ipfs daemon locally.
const ipfsURL = "http://localhost:5001"

// SampleStruct defines the benchmark payload
type SampleStruct struct {
	ID    string
	Name  string
	Value string
}

func writeEntry(entryData []byte) string {
	cid, err := sh.DagPut(entryData, "json", "cbor")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	return cid
}

func readSuperKey(superKeyCID, _nodeID string) {
	// Fetch the details by ID
	// start := time.Now()
	fmt.Printf("READ: Value for key \"%s\" is: ", _nodeID)
	res, err := GetDag(superKeyCID, _nodeID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	entryCID := fmt.Sprintf("%v", res)
	queryDAG(entryCID, "Value")
	// elapsed := time.Since(start)
	// log.Printf("***IPLD GET OPERATION TOOK %s", elapsed)
}

func createComplexMapping() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter value for ID: ")
	scanner.Scan()
	inputID := scanner.Text()

	fmt.Println("Enter value for Name: ")
	scanner.Scan()
	inputName := scanner.Text()

	fmt.Println("Enter value for Value: ")
	scanner.Scan()
	inputValue := scanner.Text()

	structObject := SampleStruct{inputID, inputName, inputValue}
	// Converting into JSON object
	entryJSON, err := json.Marshal(structObject)
	if err != nil {
		fmt.Println(err)
	}

	// Display the marshaled JSON object before sending it to IPFS
	jsonStr := string(entryJSON)
	fmt.Println("The JSON object of your key-value entry is:")
	fmt.Println(jsonStr)

	// Dag PUT operation which will return the CID for futher access or pinning etc.
	fmt.Println("------\nPUT OUTPUT\n------")
	start := time.Now()
	cid := writeEntry(entryJSON)
	masterCID := updateMapping(inputID, cid)
	elapsed := time.Since(start)
	fmt.Println("masterCID: ", masterCID)
	log.Printf("***IPLD PUT OPERATION TOOK %s", elapsed)

	fmt.Println("------\nGET OUTPUT\n------")
	start = time.Now()
	readSuperKey(masterCID, inputID)
	elapsed = time.Since(start)
	log.Printf("***IPLD GET OPERATION TOOK %s", elapsed)
	// fmt.Printf("WRITE: Successfully added %sHere's the IPLD Explorer link: https://explore.ipld.io/#/explore/%s \n", string(cid+"\n"), string(cid+"\n"))

}

func updateMapping(_structID, _CID string) string {
	// Map structure to record key-value information
	m := make(map[string]interface{})
	m[_structID] = _CID
	// Converting into JSON object
	mappingJSON, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	// Display the marshaled JSON object before sending it to IPFS
	jsonStr := string(mappingJSON)
	fmt.Println("The JSON object of the mapping is:")
	fmt.Println(jsonStr)

	// Dag PUT operation which will return the CID for futher access or pinning etc.
	masterCID, err := sh.DagPut(mappingJSON, "json", "cbor")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	// fmt.Println("------\nOUTPUT\n------")
	// fmt.Printf("WRITE: Successfully added the mapping as well. Here's the IPLD Explorer link: https://explore.ipld.io/#/explore/%s \n", string(masterCID+"\n"))
	return masterCID
}

func queryDAG(_CID, _queryString string) {
	// Fetch the details of the data from the calculated hash
	res, err := GetDag(_CID, _queryString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("QUERY RESULT: ")
	fmt.Println(res)
}

func main() {

	// Where your local node is running on localhost:5001
	sh = shell.NewShell(ipfsURL)

	fmt.Println("MMMMMMMMMMMMMMWX0kdlldk0XWMMMMMMMMMMMMMM\nMMMMMMMMMMWNKOxdoddooddodxOKNWMMMMMMMMMM\nMMMMMMMWX0kxdoddddddddddddodxk0XWMMMMMMM\nMMMWNKOxdoddddddddddddddddddddodxOKNWMMM\nMWKkdooddddddddddddddddddddddddddoodkKWM\nMNxcllloddddddddddddddddddddddddolllcxNM\nMNxooollllooddddddddddddddddoolllloooxNM\nMNxoddddoolllloodddddddddollllooddddoxNM\nMNxoddddddddolllllllllllllloddddddddoxNM\nMNxodddddddddddolcccccclodddddddddddoxNM\nMNxoddddddddddddolccccloddddddddddddoxNM\nMNxoddddddddddddddlcclddddddddddddddoxNM\nMNxoddddddddddddddoccoddddddddddddddoxNM\nMNxlodddddddddddddoccodddddddddddddolxNM\nMNklloddddddddddddoccoddddddddddddollkNM\nMMWX0kddddddddddddoccoddddddddddddk0XNMM\nMMMMMWN0OxdoddddddoccoddddddodxOKNWMMMMM\nMMMMMMMMWNXOkdodddoccodddodkOXNMMMMMMMMM\nMMMMMMMMMMMMWX0OxolccloxO0XWMMMMMMMMMMMM\nMMMMMMMMMMMMMMMWN0dlld0NWMMMMMMMMMMMMMMM\n")
	fmt.Println("### ######  #######  #####  \n #  #     # #       #     # \n #  #     # #       #       \n #  ######  #####    #####  \n #  #       #             # \n #  #       #       #     # \n### #       #        #####  \n")
	fmt.Println("###########################\n   Welcome to IPLD-CRUD!\n###########################\n")
	fmt.Println("This client generates a dynamic key-value entry and stores it in IPFS!\n")
	createComplexMapping()
}

// GetDag handles READ operations of a DAG entry by CID, returning the corresponding value
func GetDag(ref, key string) (out interface{}, err error) {
	err = sh.DagGet(ref+"/"+key, &out)
	return
}
