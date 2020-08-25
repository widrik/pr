// +build unit

package rotator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUCB1(t *testing.T) {
	t.Run("New banner was showed", func(t *testing.T) {
		id := UCB1(initStateNew())

		assert.Equal(t, 0, id)
	})


	t.Run("Banner with greater reward was showed more times", func(t *testing.T) {
		state := initStateWithTries()

		count0 := 0.0
		count1 := 0.0

		for i := 1; i <= 1000; i++ {
			selectedID := UCB1(state)

			if (selectedID == 0) {
				state.Arms[selectedID].Reward += 0.5
				count0++
			} else {
				state.Arms[selectedID].Reward += 0.1
				count1++
			}

			state.Arms[selectedID].TriesCount++
		}

		assert.Greater(t, count0, count1)
	})

	t.Run("Banner was showed minimum one time", func(t *testing.T) {
		state := initStateWithDifferentArms()

		selectedArms := make(map[int]int)

		for i := 1; i <= 1000; i++ {
			selectedID := UCB1(state)
			selectedArms[selectedID]++

			state.Arms[selectedID].Reward += 0.5
			state.Arms[selectedID].TriesCount++
		}

		for _, arm := range state.Arms {
			assert.NotEqual(t, arm.TriesCount, 0.0)
		}

		for _, arm := range selectedArms {
			assert.NotEqual(t, arm, 0.0)
		}
	})

	t.Run("All banners are new", func(t *testing.T) {
		state := initStateWithEmptyArms()

		selectedArms := make(map[int]int)

		for i := 1; i <= 1000; i++ {
			selectedID := UCB1(state)
			selectedArms[selectedID]++

			state.Arms[selectedID].Reward += 0.5
			state.Arms[selectedID].TriesCount++
		}

		for _, arm := range state.Arms {
			assert.NotEqual(t, arm.TriesCount, 0.0)
		}
	})

	t.Run("Banner with same stats were showed minimum one time", func(t *testing.T) {
		state := initStateWithDifferentArms()

		selectedArms := make(map[int]int)

		for i := 1; i <= 1000; i++ {
			selectedID := UCB1(state)
			selectedArms[selectedID]++

			state.Arms[selectedID].Reward += 0.5
			state.Arms[selectedID].TriesCount++
		}

		for _, arm := range selectedArms {
			assert.NotEqual(t, arm, 0.0)
		}
	})
}

func initStateNew() State {
	arms := Arms{
		Arm{
			TriesCount:  0,
			Reward: 1,
		},
		Arm{
			TriesCount:  1,
			Reward: 1,
		},
	}

	return State{
		Arms: arms,
		TotalCount: 1,
	}
}

func initStateWithTries() State {
	arms := Arms{
		Arm{
			TriesCount:  1,
			Reward: 1,
		},
		Arm{
			TriesCount:  1,
			Reward: 1,
		},
	}

	return State{
		Arms: arms,
		TotalCount: 2,
	}
}

func initStateWithDifferentArms() State {
	arms := Arms{
		Arm{
			TriesCount:  10,
			Reward: 0.8,
		},
		Arm{
			TriesCount:  100,
			Reward: 0.5,
		},
		Arm{
			TriesCount:  20,
			Reward: 0.4,
		},
		Arm{
			TriesCount:  30,
			Reward: 0.1,
		},
		Arm{
			TriesCount:  40,
			Reward: 0.3,
		},
		Arm{
			TriesCount:  20,
			Reward: 0.2,
		},
		Arm{
			TriesCount:  0,
			Reward: 0,
		},
		Arm{
			TriesCount:  5,
			Reward: 0.1,
		},
	}

	return State{
		Arms: arms,
		TotalCount: 255,
	}
}

func initStateWithEmptyArms() State {
	arms := Arms{
		Arm{
			TriesCount:  0,
			Reward: 0,
		},
		Arm{
			TriesCount:  0,
			Reward: 0,
		},
		Arm{
			TriesCount:  0,
			Reward: 0,
		},
		Arm{
			TriesCount:  0,
			Reward: 0,
		},
		Arm{
			TriesCount:  0,
			Reward: 0,
		},
		Arm{
			TriesCount:  0,
			Reward: 0,
		},
		Arm{
			TriesCount:  0,
			Reward: 0,
		},
		Arm{
			TriesCount:  0,
			Reward: 0,
		},
	}

	return State{
		Arms: arms,
		TotalCount: 0,
	}
}

func initStateWithAmeArms() State {
	arms := Arms{
		Arm{
			TriesCount:  10,
			Reward: 10,
		},
		Arm{
			TriesCount:  10,
			Reward: 10,
		},
		Arm{
			TriesCount:  10,
			Reward: 10,
		},
		Arm{
			TriesCount:  10,
			Reward: 10,
		},
		Arm{
			TriesCount:  10,
			Reward: 10,
		},
		Arm{
			TriesCount:  10,
			Reward: 10,
		},
		Arm{
			TriesCount:  10,
			Reward: 10,
		},
		Arm{
			TriesCount:  10,
			Reward: 10,
		},
	}

	return State{
		Arms: arms,
		TotalCount: 0,
	}
}