package main

import (
	"fmt"
	"go-merkle-tree/merkletree"
	"strings"
)

func main() {
	var data string
	fmt.Print("Enter data sep by comma: \n")
	fmt.Scanln(&data)

	arrayData := strings.Split(data, ",")
	mt := merkletree.NewMerkleTree()

	mt.Create(arrayData)

	merkletree.PrintTree(mt.Root, "", true)

}
