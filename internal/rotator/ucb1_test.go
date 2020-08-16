package rotator

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUCB1WithNew(t *testing.T) {
	id := UCB1(initStateNew())

	require.Equal(t, 0, id)
}

func TestUCB1(t *testing.T) {
	state := initStateWithTries()

	rand.Seed(time.Now().Unix())

	count0 := 0.0
	count1 := 0.0

	for i := 1; i <= 1000; i++ {
		selectedID := UCB1(state)
		selectedArm := state.Arms[selectedID]

		if (selectedID == 0) {
			selectedArm.Reward += 10
			count0++
		} else {
			selectedArm.Reward += 1
			count1++
		}

		selectedArm.TriesCount++
	}

	assert.Greater(t, count0, count1)
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