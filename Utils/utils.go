package utils

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

//GenerateJSONFiles Scrap docker to find the labels!
func GenerateJSONFiles() {

	fmt.Println("Generating the files!")

	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {

		panic(err)
	}

	swarmNodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}

	//List all nodes - works only in Swarm Mode
	fmt.Print("\n\n\n")
	fmt.Println("LIST SWARM NODES\n-----------------------")
	fmt.Println("Name | Role | Leader | Status")
	for _, swarmNode := range swarmNodes {
		fmt.Printf("%s | %s | isLeader = %t | %s\n", swarmNode.Description.Hostname, swarmNode.Spec.Role, swarmNode.ManagerStatus.Leader, swarmNode.Status.State)
	}

}
