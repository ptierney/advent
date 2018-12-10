package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() []string {
	file, err := os.Open("input")
	//file, err := os.Open("test_input")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputList []string

	for scanner.Scan() {
		inputList = append(inputList, scanner.Text())
	}

	return inputList
}

var inputNums []int

type Node struct {
	NumChildren int
	NumMetadata int
	Children    []*Node
	Metadata    []int

	Value int
}

func NewNode(numChildren int, numMetadata int) *Node {
	n := new(Node)

	n.NumChildren = numChildren
	n.NumMetadata = numMetadata

	n.Children = make([]*Node, 0)
	n.Metadata = make([]int, numMetadata)

	return n
}

func (node *Node) AddChild(c *Node) {
	node.Children = append(node.Children, c)
}

// data should start at the start of the first children
// returns the number of elements parsed
func ParseChildren(parentNode *Node, data []int) int {
	if parentNode.NumChildren == 0 {
		return 0
	}

	offset := 0

	for i := 0; i < parentNode.NumChildren; i++ {
		nc := data[offset]
		nm := data[offset+1]

		newNode := NewNode(nc, nm)
		parentNode.AddChild(newNode)

		parsed := ParseChildren(newNode, data[(offset+2):])

		CopyMetadata(newNode, data[(offset+2+parsed):])

		offset += 2 + parsed + nm
	}

	return offset
}

func CopyMetadata(node *Node, data []int) {
	for i := 0; i < node.NumMetadata; i++ {
		node.Metadata[i] = data[i]
	}
}

func GetMetadataSum(node *Node) int {
	sum := 0

	for _, val := range node.Metadata {
		sum += val
	}

	for _, c := range node.Children {
		sum += GetMetadataSum(c)
	}

	return sum
}

func ComputeValue(node *Node) {
	if len(node.Children) == 0 {
		node.Value = GetMetadataSum(node)
		return
	}

	for _, c := range node.Children {
		ComputeValue(c)
	}

	// now that all children's values are calculated, calculate this

	var valueSum int = 0

	for _, m := range node.Metadata {
		mIndex := m - 1

		if mIndex >= len(node.Children) {
			continue
		}

		valueSum += node.Children[mIndex].Value
	}

	node.Value = valueSum
}

func main() {
	inputLines := getInput()

	input := strings.Split(inputLines[0], " ")

	inputNums = make([]int, len(input))

	for i, numStr := range input {
		val, err := strconv.Atoi(numStr)

		if err != nil {
			log.Fatal(err)
		}

		inputNums[i] = val
	}

	rootNode := NewNode(inputNums[0], inputNums[1])

	parsed := ParseChildren(rootNode, inputNums[2:])

	CopyMetadata(rootNode, inputNums[(2+parsed):])

	sum := GetMetadataSum(rootNode)

	fmt.Printf("Metadata Sum: %v\n", sum)

	ComputeValue(rootNode)

	fmt.Printf("Root Node Value: %v\n", rootNode.Value)
}
