package state

import (
	"math/rand"
)

// State is a mock robot state to test the hashing methods.
type State struct {
	batteryState   batteryState
	stepsCharging  int64
	lastDeepCharge int64
	step           int64
}

// batteryState is a part of state.
type batteryState struct {
	totalCapacity   float64
	currentCapacity float64
	temperature     float64
}

// GenStates generates a number of random states.
func GenStates(size int) []*State {
	states := make([]*State, size)
	for i := 0; i < size; i++ {
		states[i] = genState()
	}
	return states
}

// genState generates a random state.
func genState() *State {
	batteryState := genBatteryState()
	state := &State{
		batteryState:   *batteryState,
		stepsCharging:  int64(rand.Intn(24 * 7 * 4)),
		lastDeepCharge: int64(rand.Intn(24 * 7 * 4)),
		step:           int64(rand.Intn(24 * 7 * 4)),
	}
	return state
}

// genBatteryState generates a random battery state.
func genBatteryState() *batteryState {
	state := &batteryState{
		totalCapacity:   50.0 * rand.Float64(),
		currentCapacity: 50.0 * rand.Float64(),
		temperature:     50.0 * rand.Float64(),
	}
	return state
}
