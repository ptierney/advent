package day06

import (
	"advent/common"

	"fmt"
	"strings"
)

type Planet struct {
	Name string

	OrbitParent *Planet
}

func NewPlanet(n string) *Planet {
	p := new(Planet)
	p.Name = n
	p.OrbitParent = nil
	return p
}

var PlanetMap map[string]*Planet

func Part1() {
	input := common.GetInput("day06/input")

	orbits := GetDirectAndIndirectInputs(input)

	fmt.Printf("Total Orbits: %v", orbits)
}

func CreateAndAddPlanet(name string) {
	if PlanetMap[name] != nil {
		return
	}

	p := NewPlanet(name)

	PlanetMap[name] = p
}

func GetDirectAndIndirectInputs(input []string) int {
	CreatePlanetsAndMap(input)

	var sum int = 0

	for _, v := range PlanetMap {
		orbits := CountAllOrbits(v)

		sum += orbits
	}

	return sum
}

func CreatePlanetsAndMap(input []string) {
	PlanetMap = make(map[string]*Planet)

	// create the planets
	for _, row := range input {
		elems := strings.Split(row, ")")

		p1 := elems[0]
		p2 := elems[1]

		CreateAndAddPlanet(p1)
		CreateAndAddPlanet(p2)
	}

	// Link Planets
	for _, row := range input {
		elems := strings.Split(row, ")")

		p1Name := elems[0]
		p2Name := elems[1]

		p1 := PlanetMap[p1Name]
		p2 := PlanetMap[p2Name]

		p2.OrbitParent = p1
	}
}

func CountAllOrbits(planet *Planet) int {
	var sum int = 0

	p := planet

	for {
		p = p.OrbitParent

		if p == nil {
			break
		}

		sum++
	}

	return sum
}
