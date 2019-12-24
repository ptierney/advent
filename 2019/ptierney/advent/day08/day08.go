package day08

import (
	"advent/common"

	"fmt"
	"strconv"
)

type Layer struct {
	Pixels [][]int
}

type Picture struct {
	Layers []*Layer
}

func NewLayer(w, h int) *Layer {
	l := new(Layer)

	l.Pixels = make([][]int, w)

	for i := 0; i < w; i++ {
		l.Pixels[i] = make([]int, h)
	}

	return l
}

func NewPicture() *Picture {
	p := new(Picture)
	return p
}

func (l *Layer) NumZeros() int {
	return l.NumValues(0)
}

func (l *Layer) NumValues(value int) int {
	var sum int = 0

	for i := 0; i < PictureWidth; i++ {
		for j := 0; j < PictureHeight; j++ {
			if l.Pixels[i][j] == value {
				sum++
			}
		}
	}

	return sum
}

var PictureWidth int = 25
var PictureHeight int = 6

func Part1() {
	input := common.GetInput("day08/input")

	c := GetChecksum(input)

	fmt.Printf("Checksum: %v\n", c)
}

func GetChecksum(input []string) int {
	p := PictureFromString(input[0])

	var lowestLayer *Layer = nil
	var lowest int

	for _, l := range p.Layers {
		zeros := l.NumZeros()

		if lowestLayer == nil {
			lowest = zeros
			lowestLayer = l
		}

		if zeros < lowest {
			lowest = zeros
			lowestLayer = l
		}
	}

	ones := lowestLayer.NumValues(1)
	twos := lowestLayer.NumValues(2)

	return ones * twos
}

func PictureFromString(input string) *Picture {

	layerStride := PictureWidth * PictureHeight

	pixelStr := common.StringArrayFromString(input)

	numLayers := len(pixelStr) / layerStride

	picture := NewPicture()
	picture.Layers = make([]*Layer, numLayers)

	for li := 0; li < numLayers; li++ {
		layer := NewLayer(PictureWidth, PictureHeight)

		for h := 0; h < PictureHeight; h++ {
			for w := 0; w < PictureWidth; w++ {
				index := layerStride*li + h*PictureWidth + w

				p, err := strconv.Atoi(pixelStr[index])

				if err != nil {
					panic(err)
				}

				layer.Pixels[w][h] = p
			}
		}

		picture.Layers[li] = layer
	}

	return picture
}
