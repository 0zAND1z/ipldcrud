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

// Set writes key-value data
func Set(sh *shell.Shell, data []byte) string {
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
