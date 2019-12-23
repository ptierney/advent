package day07

import (
	"advent/common"

	"fmt"
)

func CreateComputers(input string) []*Computer {
	amps := make([]*Computer, 5)

	for i := 0; i < 5; i++ {
		c := NewComputer()

		c.LoadProgramString(input)

		amps[i] = c
	}

	return amps
}

func Part1() {
	input := common.GetInput("day07/input")

	ms, _ := FindMaxThrusterSignal(input[0])

	fmt.Printf("Max Thrust: %v\n", ms)
}

func Part2() {
	input := common.GetInput("day07/input")

	ms, _ := FindMaxFeedbackSignal(input[0])

	fmt.Printf("\n\nMax Feedback Thrust: %v\n", ms)
}

func FindMaxFeedbackSignal(input string) (maxSignal int64, phaseSetting []int64) {

	var ms int64
	maxSignalSet := false
	var maxPS []int64

	all_ps := getAllFeedbackPhaseSettings()

	for _, ps := range all_ps {
		amps := CreateComputers(input)

		signal := FeedbackSignalFromSetting(amps, ps)

		//fmt.Printf("%v sig from %v\n", signal, ps)

		if maxSignalSet == false {
			ms = signal
			maxSignalSet = true
			maxPS = ps

			continue
		}

		if ms > signal {
			continue
		}

		ms = signal
		maxPS = ps
	}

	return ms, maxPS
}

func FindMaxThrusterSignal(input string) (maxSignal int64, phaseSetting []int64) {

	var ms int64
	maxSignalSet := false
	var maxPS []int64

	all_ps := getAllPhaseSettings()

	for _, ps := range all_ps {
		amps := CreateComputers(input)

		signal := ThrusterSignalFromSetting(amps, ps)

		if maxSignalSet == false {
			ms = signal
			maxSignalSet = true
			maxPS = ps

			continue
		}

		if ms > signal {
			continue
		}

		ms = signal
		maxPS = ps
	}

	return ms, maxPS
}

func FeedbackSignalFromSetting(amps []*Computer, phaseSetting []int64) int64 {
	SetPhaseSettings(amps, phaseSetting)

	amps[0].AddInput(0)

	var output int64 = 0
	var err error = nil

	lastEOutputSet := false
	var lastEOutput int64

	lastValues := make([]int64, 5)

	i := 0

	for {
		output, err = amps[i].ExecuteUntilOutput()

		lastValues[i] = output

		if err != nil {
			// The engines are shutting down, this means just return the last E
			if lastEOutputSet == true {
				return lastEOutput
			} else {
				// This means the engines are shutting down early
				panic(err)
			}
		}

		amps[i].ClearOutput()

		if i == 4 {
			lastEOutput = output
			lastEOutputSet = true

			//fmt.Printf("%v sigs from %v\n", lastValues, phaseSetting)

			if amps[i].HaltFlag == true {
				break
			}

			i = 0
		} else {
			i++
		}

		if lastEOutputSet == true {
			amps[i].SetInput(output)
		} else {
			amps[i].AddInput(output)
		}
	}

	return output
}

func ThrusterSignalFromSetting(amps []*Computer, phaseSetting []int64) int64 {
	SetPhaseSettings(amps, phaseSetting)

	amps[0].AddInput(0)

	var output int64 = 0

	for i := 0; i < 5; i++ {
		output, _ = amps[i].ExecuteUntilOutput()

		if i == 4 {
			break
		}

		amps[i+1].AddInput(output)
	}

	return output
}

func SetPhaseSettings(amps []*Computer, phaseSetting []int64) {
	for i := 0; i < 5; i++ {
		amps[i].AddInput(phaseSetting[i])
	}
}

func getAllPhaseSettings() [][]int64 {
	all_ps := make([][]int64, 0)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == j {
				continue
			}

			for k := 0; k < 5; k++ {
				if k == i || k == j {
					continue
				}

				for l := 0; l < 5; l++ {
					if l == i || l == j || l == k {
						continue
					}

					for m := 0; m < 5; m++ {
						if m == i || m == j || m == k || m == l {
							continue
						}

						ps := make([]int64, 5)
						ps[0] = int64(i)
						ps[1] = int64(j)
						ps[2] = int64(k)
						ps[3] = int64(l)
						ps[4] = int64(m)

						all_ps = append(all_ps, ps)
					}
				}
			}
		}
	}

	return all_ps
}

func getAllFeedbackPhaseSettings() [][]int64 {
	all_ps := make([][]int64, 0)

	for i := 5; i < 10; i++ {
		for j := 5; j < 10; j++ {
			if i == j {
				continue
			}

			for k := 5; k < 10; k++ {
				if k == i || k == j {
					continue
				}

				for l := 5; l < 10; l++ {
					if l == i || l == j || l == k {
						continue
					}

					for m := 5; m < 10; m++ {
						if m == i || m == j || m == k || m == l {
							continue
						}

						ps := make([]int64, 5)
						ps[0] = int64(i)
						ps[1] = int64(j)
						ps[2] = int64(k)
						ps[3] = int64(l)
						ps[4] = int64(m)

						all_ps = append(all_ps, ps)
					}
				}
			}
		}
	}

	return all_ps
}
