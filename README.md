# ipld-crud

A simple Golang based hack to experiment the uses of IPLD through the DAG operations using the [go-ipfs-api](https://github.com/ipfs/go-ipfs-api) package.

I have built a tiny client that receives key-value entries from user and stores it on IPFS DAG, returning an explorable URL to play with.

## Usage

Here is a sample main.go file for your quick reference:

```go
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	ipldcrud "github.com/0zAND1z/ipldcrud"
)

func main() {
	sh := ipldcrud.InitShell("https://ipfs.infura.io:5001")

	keyValueMap := make(map[string]interface{})

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter value for the key field: ")
	scanner.Scan()
	inputKey := scanner.Text()

	fmt.Println("Enter value for value field: ")
	scanner.Scan()
	inputValue := scanner.Text()

	keyValueMap[inputKey] = inputValue

	// Converting into JSON object
	entryJSON, err := json.Marshal(keyValueMap)
	if err != nil {
		fmt.Println(err)
	}

	// Display the marshaled JSON object before sending it to IPFS
	jsonStr := string(entryJSON)
	fmt.Println("The JSON object of your key-value entry is:")
	fmt.Println(jsonStr)
	cid := ipldcrud.Set(sh, entryJSON)
	fmt.Println("CID: ", cid)

	// Fetch the details by reading the DAG for key "inputKey"
	fmt.Printf("READ: Value for key \"%s\" is: ", inputKey)
	res, err := ipldcrud.Get(sh, cid, inputKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

```

1. Once you run the main.go, a simple console pops up, asking you to enter a key and a value.

2. After entering two string values, it will create a dag entry based on the input data. It will subsequently query the same data and return back the result along with the query results.

### Tutorial

For a more detailed understanding of IPLD and how the code works, check out this [article on SimpleAsWater.com](https://simpleaswater.com/hands-on-ipld-tutorial-in-golang/)!
