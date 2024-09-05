package main

import (
	"fmt"
	"github.com/m41denx/alligator/options"
	"os"
	"strings"

	gator "github.com/m41denx/alligator"
)

func main() {
	url := os.Getenv("CROC_URL")
	app, _ := gator.NewApp(url, os.Getenv("CROC_KEY"))

	node, err := app.CreateNode(gator.CreateNodeDescriptor{
		Name:               "croc-node-1",
		LocationID:         1,
		Public:             true,
		FQDN:               fmt.Sprintf("test.nodes.%s", strings.Split(url, "//")[1]),
		Scheme:             "https",
		BehindProxy:        false,
		Memory:             16000,
		MemoryOverallocate: 0,
		Disk:               1024,
		DiskOverallocate:   0,
		DaemonBase:         "/var/lib/pterodactyl/volumes",
		DaemonSftp:         2022,
		DaemonListen:       8080,
		UploadSize:         100,
	})
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}

	fmt.Printf("ID: %d - Name: %s - Public: %v\n", node.ID, node.Name, node.Public)

	data := node.UpdateDescriptor()
	data.Public = false
	node, err = app.UpdateNode(node.ID, *data)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}

	fmt.Printf("ID: %d - Name: %s - Public: %v\n", node.ID, node.Name, node.Public)

	nodes, err := app.ListNodes(options.ListNodesOptions{
		Include: options.IncludeNodes{
			Location: true,
			Servers:  true,
		}})
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}

	for _, n := range nodes {
		fmt.Printf("%d: %s\n", n.ID, n.Name)
	}

	if err = app.DeleteNode(node.ID); err != nil {
		handleError(err)
	}
}
