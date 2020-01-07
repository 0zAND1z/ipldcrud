# ipld-crud

A simple Golang based hack to experiment the uses of IPLD through the DAG operations using the [go-ipfs-api](https://github.com/ipfs/go-ipfs-api) package.

I have built a tiny client that receives key-value entries from user and stores it on IPFS DAG, returning an explorable URL to play with.

## Instructions

Clone the repo and just run `go get && go run main.go`

Once you run the main.go, a nice ASCII version of the IPFS logo pops up, asking you to enter a key and a value.

After entering two string values, it will create a dag entry based on the input data and return back the hash along with the query results.

### Tutorial

For a more detailed understanding of IPLD and how the code works, check out this [article on SimpleAsWater.com](https://simpleaswater.com/hands-on-ipld-tutorial-in-golang/)!
