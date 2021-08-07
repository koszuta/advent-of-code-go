package rtg

import "fmt"

const (
	// NElements = 2
	NElements = 7
	NFloors   = 4
)

type State struct {
	Elevator   int
	Generators [NElements]int
	Microchips [NElements]int
}

type Hash struct {
	Elevator                      int
	NPairs, NLoneGens, NLoneChips [NFloors]int
}

var (
	FinalState     State
	FinalStateHash Hash
)

func init() {
	var allTopFloor [NElements]int
	for i := 0; i < NElements; i++ {
		allTopFloor[i] = NFloors - 1
	}
	FinalState = NewState(NFloors-1, allTopFloor, allTopFloor)
	FinalStateHash = FinalState.Hash()
}

func NewState(elevator int, generators, microchips [NElements]int) State {
	return State{Elevator: elevator, Generators: generators, Microchips: microchips}
}

func (state *State) Copy() State {
	return NewState(state.Elevator, state.Generators, state.Microchips)
}

func (state *State) Hash() (hash Hash) {
	hash.Elevator = state.Elevator
	for i := 0; i < NElements; i++ {
		genFloor, chipFloor := state.Generators[i], state.Microchips[i]
		if genFloor == chipFloor {
			hash.NPairs[genFloor]++
		} else {
			hash.NLoneGens[genFloor]++
			hash.NLoneChips[chipFloor]++
		}
	}
	return hash
}

func (state *State) ToString() string {
	s := ""
	for floor := NFloors - 1; floor >= 0; floor-- {
		row := fmt.Sprintf("F%d", floor+1)
		if state.Elevator == floor {
			row += " E "
		} else {
			row += " . "
		}
		for element := 0; element < NElements; element++ {
			if state.Generators[element] == floor {
				row += fmt.Sprintf(" G%d", element)
			} else {
				row += " . "
			}
			if state.Microchips[element] == floor {
				row += fmt.Sprintf(" M%d", element)
			} else {
				row += " . "
			}
		}
		s += row
		if floor > 0 {
			s += "\n"
		}
	}
	return s
}

func (state *State) MoveGenUp(element int) *State {
	genFloor := state.Generators[element]
	if state.Elevator != genFloor || state.Elevator == NFloors-1 {
		return nil
	}
	aboveIsEmpty := true
	for e := 0; e < NElements; e++ {
		if state.Generators[e] > state.Elevator || state.Microchips[e] > state.Elevator {
			aboveIsEmpty = false
			break
		}
	}
	if aboveIsEmpty {
		return nil
	}
	newState := state.Copy()
	newState.Elevator++
	newState.Generators[element] = newState.Elevator
	if !newState.Valid() {
		return nil
	}
	return &newState
}

func (state *State) MoveGenDown(element int) *State {
	genFloor := state.Generators[element]
	if state.Elevator != genFloor || state.Elevator == 0 {
		return nil
	}
	belowIsEmpty := true
	for e := 0; e < NElements; e++ {
		if state.Generators[e] < state.Elevator || state.Microchips[e] < state.Elevator {
			belowIsEmpty = false
			break
		}
	}
	if belowIsEmpty {
		return nil
	}
	newState := state.Copy()
	newState.Elevator--
	newState.Generators[element] = newState.Elevator
	if !newState.Valid() {
		return nil
	}
	return &newState
}

func (state *State) MoveChipUp(element int) *State {
	chipFloor := state.Microchips[element]
	if state.Elevator != chipFloor || state.Elevator == NFloors-1 {
		return nil
	}
	aboveIsEmpty := true
	for e := 0; e < NElements; e++ {
		if state.Generators[e] > state.Elevator || state.Microchips[e] > state.Elevator {
			aboveIsEmpty = false
			break
		}
	}
	if aboveIsEmpty {
		return nil
	}
	newState := state.Copy()
	newState.Elevator++
	newState.Microchips[element] = newState.Elevator
	if !newState.Valid() {
		return nil
	}
	return &newState
}

func (state *State) MoveChipDown(element int) *State {
	chipFloor := state.Microchips[element]
	if state.Elevator != chipFloor || state.Elevator == 0 {
		return nil
	}
	belowIsEmpty := true
	for e := 0; e < NElements; e++ {
		if state.Generators[e] < state.Elevator || state.Microchips[e] < state.Elevator {
			belowIsEmpty = false
			break
		}
	}
	if belowIsEmpty {
		return nil
	}
	newState := state.Copy()
	newState.Elevator--
	newState.Microchips[element] = newState.Elevator
	if !newState.Valid() {
		return nil
	}
	return &newState
}

func (state *State) MovePairUp(element int) *State {
	genFloor, chipFloor := state.Generators[element], state.Microchips[element]
	if state.Elevator != genFloor || state.Elevator != chipFloor || state.Elevator == NFloors-1 {
		return nil
	}
	newState := state.Copy()
	newState.Elevator++
	newState.Generators[element] = newState.Elevator
	newState.Microchips[element] = newState.Elevator
	if !newState.Valid() {
		return nil
	}
	return &newState
}

func (state *State) MovePairDown(element int) *State {
	genFloor, chipFloor := state.Generators[element], state.Microchips[element]
	if state.Elevator != genFloor || state.Elevator != chipFloor || state.Elevator == 0 {
		return nil
	}
	belowIsEmpty := true
	for e := 0; e < NElements; e++ {
		if state.Generators[e] < state.Elevator || state.Microchips[e] < state.Elevator {
			belowIsEmpty = false
			break
		}
	}
	if belowIsEmpty {
		return nil
	}
	newState := state.Copy()
	newState.Elevator--
	newState.Generators[element], newState.Microchips[element] = newState.Elevator, newState.Elevator
	if !newState.Valid() {
		return nil
	}
	return &newState
}

func (state *State) MoveGensUp(element1, element2 int) *State {
	genFloor1, genFloor2 := state.Generators[element1], state.Generators[element2]
	if state.Elevator != genFloor1 || state.Elevator != genFloor2 || state.Elevator == NFloors-1 {
		return nil
	}
	newState := state.Copy()
	newState.Elevator++
	newState.Generators[element1] = newState.Elevator
	newState.Generators[element2] = newState.Elevator
	if !newState.Valid() {
		return nil
	}
	return &newState
}

func (state *State) MoveGensDown(element1, element2 int) *State {
	genFloor1, genFloor2 := state.Generators[element1], state.Generators[element2]
	if state.Elevator != genFloor1 || state.Elevator != genFloor2 || state.Elevator == 0 {
		return nil
	}
	belowIsEmpty := true
	for e := 0; e < NElements; e++ {
		if state.Generators[e] < state.Elevator || state.Microchips[e] < state.Elevator {
			belowIsEmpty = false
			break
		}
	}
	if belowIsEmpty {
		return nil
	}
	newState := state.Copy()
	newState.Elevator--
	newState.Generators[element1] = newState.Elevator
	newState.Generators[element2] = newState.Elevator
	if !newState.Valid() {
		return nil
	}
	return &newState
}

func (state *State) MoveChipsUp(element1, element2 int) *State {
	chipFloor1, chipFloor2 := state.Microchips[element1], state.Microchips[element2]
	if state.Elevator != chipFloor1 || state.Elevator != chipFloor2 || state.Elevator == NFloors-1 {
		return nil
	}
	newState := state.Copy()
	newState.Elevator++
	newState.Microchips[element1] = newState.Elevator
	newState.Microchips[element2] = newState.Elevator
	if !newState.Valid() {
		return nil
	}
	return &newState
}

func (state *State) MoveChipsDown(element1, element2 int) *State {
	chipFloor1, chipFloor2 := state.Microchips[element1], state.Microchips[element2]
	if state.Elevator != chipFloor1 || state.Elevator != chipFloor2 || state.Elevator == 0 {
		return nil
	}
	belowIsEmpty := true
	for e := 0; e < NElements; e++ {
		if state.Generators[e] < state.Elevator || state.Microchips[e] < state.Elevator {
			belowIsEmpty = false
			break
		}
	}
	if belowIsEmpty {
		return nil
	}
	newState := state.Copy()
	newState.Elevator--
	newState.Microchips[element1] = newState.Elevator
	newState.Microchips[element2] = newState.Elevator
	if !newState.Valid() {
		return nil
	}
	return &newState
}

func (state *State) Valid() bool {
	for element := 0; element < NElements; element++ {
		genFloor, chipFloor := state.Generators[element], state.Microchips[element]
		if genFloor < 0 || chipFloor < 0 || genFloor >= NFloors || chipFloor >= NFloors {
			return false
		}
		if genFloor != chipFloor {
			hash := state.Hash()
			if hash.NPairs[chipFloor] != 0 || hash.NLoneGens[chipFloor] != 0 {
				return false
			}
		}
	}
	return true
}
