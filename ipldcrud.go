package ipldcrud

import (
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

// InitShell is used to create a new IPFS Shell
func InitShell(ipfsURL string) *shell.Shell {
	// Where your local node is running on localhost:5001
	sh := shell.NewShell(ipfsURL)
	return sh
}

// func main() {

// 	sh := InitShell("https://ipfs.infura.io:5001")

// 	// fmt.Println("MMMMMMMMMMMMMMWX0kdlldk0XWMMMMMMMMMMMMMM\nMMMMMMMMMMWNKOxdoddooddodxOKNWMMMMMMMMMM\nMMMMMMMWX0kxdoddddddddddddodxk0XWMMMMMMM\nMMMWNKOxdoddddddddddddddddddddodxOKNWMMM\nMWKkdooddddddddddddddddddddddddddoodkKWM\nMNxcllloddddddddddddddddddddddddolllcxNM\nMNxooollllooddddddddddddddddoolllloooxNM\nMNxoddddoolllloodddddddddollllooddddoxNM\nMNxoddddddddolllllllllllllloddddddddoxNM\nMNxodddddddddddolcccccclodddddddddddoxNM\nMNxoddddddddddddolccccloddddddddddddoxNM\nMNxoddddddddddddddlcclddddddddddddddoxNM\nMNxoddddddddddddddoccoddddddddddddddoxNM\nMNxlodddddddddddddoccodddddddddddddolxNM\nMNklloddddddddddddoccoddddddddddddollkNM\nMMWX0kddddddddddddoccoddddddddddddk0XNMM\nMMMMMWN0OxdoddddddoccoddddddodxOKNWMMMMM\nMMMMMMMMWNXOkdodddoccodddodkOXNMMMMMMMMM\nMMMMMMMMMMMMWX0OxolccloxO0XWMMMMMMMMMMMM\nMMMMMMMMMMMMMMMWN0dlld0NWMMMMMMMMMMMMMMM\n")

// 	// fmt.Println("### ######  #######  #####  \n #  #     # #       #     # \n #  #     # #       #       \n #  ######  #####    #####  \n #  #       #             # \n #  #       #       #     # \n### #       #        #####  \n")

// 	// fmt.Println("###########################\n   Welcome to IPLD-CRUD!\n###########################\n")
// 	// fmt.Println("This client generates a dynamic key-value entry and stores it in IPFS!\n")

// 	// Map structure to record key-value information

// 	// fmt.Println("------\nOUTPUT\n------")
// 	// fmt.Printf("WRITE: Successfully added %sHere's the IPLD Explorer link: https://explore.ipld.io/#/explore/%s \n", string(cid+"\n"), string(cid+"\n"))

// }

// Set writes key-value data
func Set(sh *shell.Shell, data []byte) interface{} {
	// Dag PUT operation which will return the CID for futher access or pinning etc.
	cid, err := sh.DagPut(data, "json", "cbor")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	return cid
}

// Get handles READ operations of a DAG entry by CID, returning the corresponding value
func Get(sh *shell.Shell, ref, key string) (out interface{}, err error) {
	err = sh.DagGet(ref+"/"+key, &out)
	return
}
