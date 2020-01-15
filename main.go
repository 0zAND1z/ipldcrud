package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

// Global variable to handle all the IPFS API client calls
var sh *shell.Shell

// SampleStruct defines the benchmark payload
type SampleStruct struct {
	ID     string `json:"ID"`
	Name   string `json:"Name"`
	Salary string `json:"Salary"`
}

func main() {

	sh = shell.NewShell("https://ipfs.infura.io:5001") // Replace URL with "localhost:5001" if you'd like to use local IPFS daemon

	fmt.Println("MMMMMMMMMMMMMMWX0kdlldk0XWMMMMMMMMMMMMMM\nMMMMMMMMMMWNKOxdoddooddodxOKNWMMMMMMMMMM\nMMMMMMMWX0kxdoddddddddddddodxk0XWMMMMMMM\nMMMWNKOxdoddddddddddddddddddddodxOKNWMMM\nMWKkdooddddddddddddddddddddddddddoodkKWM\nMNxcllloddddddddddddddddddddddddolllcxNM\nMNxooollllooddddddddddddddddoolllloooxNM\nMNxoddddoolllloodddddddddollllooddddoxNM\nMNxoddddddddolllllllllllllloddddddddoxNM\nMNxodddddddddddolcccccclodddddddddddoxNM\nMNxoddddddddddddolccccloddddddddddddoxNM\nMNxoddddddddddddddlcclddddddddddddddoxNM\nMNxoddddddddddddddoccoddddddddddddddoxNM\nMNxlodddddddddddddoccodddddddddddddolxNM\nMNklloddddddddddddoccoddddddddddddollkNM\nMMWX0kddddddddddddoccoddddddddddddk0XNMM\nMMMMMWN0OxdoddddddoccoddddddodxOKNWMMMMM\nMMMMMMMMWNXOkdodddoccodddodkOXNMMMMMMMMM\nMMMMMMMMMMMMWX0OxolccloxO0XWMMMMMMMMMMMM\nMMMMMMMMMMMMMMMWN0dlld0NWMMMMMMMMMMMMMMM\n")

	fmt.Println("### ######  #######  #####  \n #  #     # #       #     # \n #  #     # #       #       \n #  ######  #####    #####  \n #  #       #             # \n #  #       #       #     # \n### #       #        #####  \n")

	fmt.Println("###########################\n   Welcome to IPLD-CRUD!\n###########################\n")
	fmt.Println("This client generates a dynamic key-value entry and stores it in IPFS!\n")

	// Map structure to record key-value information
	DocStoreMap := make(map[string]SampleStruct)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the ID of the employee: ")
	scanner.Scan()
	inputID := scanner.Text()

	fmt.Println("Enter the name of the employee: ")
	scanner.Scan()
	inputName := scanner.Text()

	fmt.Println("Enter the salary of the employee: ")
	scanner.Scan()
	inputSalary := scanner.Text()

	// Create a struct instance and assign the input values to the corresponding fields
	employeeObject := SampleStruct{ID: inputID, Name: inputName, Salary: inputSalary}

	// Map the struct instance to the mapping
	DocStoreMap[inputID] = employeeObject

	// Converting the map into JSON object
	entryJSON, err := json.Marshal(DocStoreMap)
	if err != nil {
		fmt.Println(err)
	}

	// Display the marshaled JSON object before sending it to IPFS
	jsonStr := string(entryJSON)
	fmt.Println("The JSON object of your document entry is:")
	fmt.Println(jsonStr)

	start := time.Now()
	// Dag PUT operation which will return the CID for futher access or pinning etc.
	cid, err := sh.DagPut(entryJSON, "json", "cbor")
	elapsed := time.Since(start)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	fmt.Println("------\nOUTPUT\n------")
	fmt.Printf("WRITE: Successfully added %sHere's the IPLD Explorer link: https://explore.ipld.io/#/explore/%s", string(cid+"\n"), string(cid+"\n"))
	fmt.Println("WRITE: IPLD PUT call took ", elapsed)

	// Fetch the details by reading the DAG for key "inputKey"
	fmt.Printf("READ: Reading the document details of employee by ID: \"%s\"\n", inputID)
	start = time.Now()
	document, err := GetDocument(cid, inputID)
	elapsed = time.Since(start)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("READ: Salary of employee ID %s is %s\n", string(inputID), string(document.Salary))
	fmt.Println("READ: IPLD GET call took ", elapsed)
}

// GetDocument handles READ operations of a DAG entry by CID, returning the corresponding document
func GetDocument(ref, key string) (out SampleStruct, err error) {
	err = sh.DagGet(ref+"/"+key, &out)
	return
}
