package merkletree

import (
	"crypto/sha256"
	"fmt"
)

type MerkleNode struct {
	data  string
	left  *MerkleNode
	right *MerkleNode
}

type Merkletree struct {
	Root *MerkleNode
}

func NewMerkleNode(data string) *MerkleNode {

	var newNode = new(MerkleNode)
	newNode.data = data
	newNode.left = nil
	newNode.right = nil

	return newNode

}

func hashData(data string) string {
	hash := sha256.Sum256([]byte(data))
	hashString := fmt.Sprintf("%x", hash)

	return hashString
}

func NewMerkleTree() Merkletree {
	newTree := Merkletree{
		Root: nil,
	}
	return newTree
}

func buildMerkleTree(merkletree *Merkletree, merkleNodes []*MerkleNode) {
	var outputNodes []*MerkleNode

	for i := 0; i < len(merkleNodes)-1; i += 2 {
		leftNode := merkleNodes[i]
		rightNode := merkleNodes[i+1]
		resHash := hashData(leftNode.data + rightNode.data)

		parent := new(MerkleNode)
		parent.left = leftNode
		parent.right = rightNode
		parent.data = hashData(resHash)

		outputNodes = append(outputNodes, parent)
	}

	if len(outputNodes) == 1 {
		merkletree.Root = outputNodes[len(outputNodes)-1]
		return
	}

	if len(outputNodes)%2 != 0 {
		outputNodes = append(outputNodes, outputNodes[len(outputNodes)-1])
	}

	buildMerkleTree(merkletree, outputNodes)

}

func (mt *Merkletree) Create(data []string) {

	var extendedData []string = data

	if len(data)%2 != 0 {
		extendedData = append(extendedData, data[len(data)-1])
	}

	var merkelNodes []*MerkleNode

	for i := 0; i < len(extendedData); i++ {

		node := new(MerkleNode)
		node.data = hashData(extendedData[i])
		node.left = nil
		node.right = nil
		merkelNodes = append(merkelNodes, node)
	}

	buildMerkleTree(mt, merkelNodes)

}

func ReduceHash(data string) string {
	return data[0:3] + "..." + data[len(data)-3:len(data)-1]
}

func PrintTree(node *MerkleNode, prefix string, isLeft bool) {
	if node != nil {
		fmt.Printf("%s", prefix)
		if isLeft {
			fmt.Printf("├──")
		} else {
			fmt.Printf("└──")
		}
		fmt.Printf("%s\n", ReduceHash(node.data))

		newPrefix := prefix
		if isLeft {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}

		PrintTree(node.left, newPrefix, true)
		PrintTree(node.right, newPrefix, false)
	}
}
