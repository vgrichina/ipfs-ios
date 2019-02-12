package ipfs

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ipsn/go-ipfs/core"
	"github.com/ipsn/go-ipfs/core/coreapi"
	iface "github.com/ipsn/go-ipfs/core/coreapi/interface"
	files "github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-ipfs-files"
)

func RunIPFS() {
	// Create a new IPFS network node
	node, err := core.NewNode(context.TODO(), &core.BuildCfg{Online: true})
	if err != nil {
		log.Fatalf("Failed to start IPFS node: %v", err)
	}
	path, _ := iface.ParsePath("QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB")

	// Resolve the IPFS welcome page
	nodeAPI, err := coreapi.NewCoreAPI(node)
	if err != nil {
		log.Fatalf("Failed to create API instance: %v", err)
	}

	reader, err := nodeAPI.Unixfs().Get(context.TODO(), path)
	if err != nil {
		log.Fatalf("Failed to look up IPFS welcome page: %v", err)
	}
	// Retrieve and print the welcome page
	blob, err := ioutil.ReadAll(files.ToFile(reader))
	if err != nil {
		log.Fatalf("Failed to retrieve IPFS welcome page: %v", err)
	}
	fmt.Println(string(blob))
}
