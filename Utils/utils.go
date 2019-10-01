package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/jarzamendia/regionmaker/sigma"
)

//GenerateFiles Generate the JSON files
func GenerateFiles() {

	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}

	nodes, _ := cli.NodeList(ctx, types.NodeListOptions{})

	var i = 0
	var edge = []sigma.Edges{}

	var nodeslicePrd = []sigma.Nodes{}
	var nodesliceHom = []sigma.Nodes{}
	var nodesliceDes = []sigma.Nodes{}
	var nodesliceCMS = []sigma.Nodes{}
	var nodesliceSup = []sigma.Nodes{}
	var nodesliceStream = []sigma.Nodes{}

	var payloadPrd = sigma.Payload{Edges: edge}
	var payloadHom = sigma.Payload{Edges: edge}
	var payloadDes = sigma.Payload{Edges: edge}
	var payloadCMS = sigma.Payload{Edges: edge}
	var payloadSup = sigma.Payload{Edges: edge}
	var payloadStream = sigma.Payload{Edges: edge}

	for _, element := range nodes {

		id := "n" + strconv.Itoa(i)
		color := ""
		x := 0
		y := 0

		if x%4 == 0 && x != 0 {

			x = x + 1

		}

		if y == 5 {

			y = 0

		}

		nodeAddr := element.Status.Addr
		nodeHostname := element.Description.Hostname
		nodeAvailability := element.Spec.Availability
		nodeState := element.Status.State
		nodeLabel := element.Spec.Labels["region"]

		if nodeAvailability == "active" {

			color = "#008000"

		} else {

			color = "#DC143C"
		}

		node := sigma.Nodes{
			Color:   color,
			ID:      id,
			Label:   nodeHostname,
			Size:    1,
			X:       x,
			Y:       y,
			Custom1: nodeAddr,
			Custom2: string(nodeAvailability),
			Custom3: string(nodeState),
			Custom4: nodeLabel,
		}

		fmt.Println(node)

		if nodeLabel == "prd" {

			nodeslicePrd = append(nodeslicePrd, node)

		}

		if nodeLabel == "hom" {

			nodesliceHom = append(nodesliceHom, node)

		}

		if nodeLabel == "des" {

			nodesliceDes = append(nodesliceDes, node)

		}

		if nodeLabel == "sup" {

			nodesliceSup = append(nodesliceSup, node)

		}

		if nodeLabel == "cms" {

			nodesliceCMS = append(nodesliceCMS, node)

		}

		if nodeLabel == "stream" {

			nodesliceStream = append(nodesliceStream, node)

		}

		i++

	}

	payloadPrd.Nodes = nodeslicePrd
	payloadHom.Nodes = nodesliceHom
	payloadDes.Nodes = nodesliceDes
	payloadSup.Nodes = nodesliceSup
	payloadCMS.Nodes = nodesliceCMS
	payloadStream.Nodes = nodesliceStream

	filePrd, _ := json.MarshalIndent(payloadPrd, "", " ")
	_ = ioutil.WriteFile("/data/prod_data.json", filePrd, 0644)

	fileHom, _ := json.MarshalIndent(payloadHom, "", " ")
	_ = ioutil.WriteFile("/data/hom_data.json", fileHom, 0644)

	fileDes, _ := json.MarshalIndent(payloadDes, "", " ")
	_ = ioutil.WriteFile("/data/des_data.json", fileDes, 0644)

	fileSup, _ := json.MarshalIndent(payloadSup, "", " ")
	_ = ioutil.WriteFile("/data/sup_data.json", fileSup, 0644)

	fileCMS, _ := json.MarshalIndent(payloadCMS, "", " ")
	_ = ioutil.WriteFile("/data/cms_data.json", fileCMS, 0644)

	fileStream, _ := json.MarshalIndent(payloadStream, "", " ")
	_ = ioutil.WriteFile("/data/stream_data.json", fileStream, 0644)

}
