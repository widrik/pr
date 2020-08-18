package rotator

import (
	"math"

	"github.com/widrik/pr/internal/entities"
)

type Arm struct {
	TriesCount float64
	Reward     float64
}

type Arms []Arm

type State struct {
	Arms       Arms
	TotalCount int64
}

func InitAlgoritm(r *Rotator, slot *entities.Slot, socialGroup *entities.SocialGroup) (State, error) {
	var totalCount int64

	bannersCount := len(slot.Banners)
	arms := make(Arms, bannersCount)
	totalCount = 0

	bannerStats := make(map[uint]*entities.Stats)

	state := State{
		Arms:       arms,
		TotalCount: totalCount,
	}

	for _, banner := range slot.Banners {
		stats, err := r.findOrCreateStats(banner.ID, slot.ID, socialGroup.ID)
		if err != nil {
			return state, err
		}

		bannerStats[banner.ID] = stats
		newArm := Arm{
			TriesCount: float64(stats.ShowCount),
			Reward:     float64(stats.ClickCount),
		}

		arms[banner.ID] = newArm

		totalCount += stats.ShowCount
	}

	state.Arms = arms
	state.TotalCount = totalCount

	return state, nil
}

func UCB1(state State) int {
	newIndex := searchIndexOfNew(state)

	if newIndex != -1 {
		return newIndex
	}

	totalCount := float64(state.TotalCount)
	var resultIndex int
	var maxValue float64

	for id, arm := range state.Arms {
		avg := arm.Reward / arm.TriesCount
		armValue := avg + math.Sqrt((2*math.Log(totalCount))/arm.TriesCount)

		if armValue > maxValue {
			maxValue = armValue
			resultIndex = id
		}
	}

	return resultIndex
}

func searchIndexOfNew(state State) int {
	for id, arm := range state.Arms {
		if arm.TriesCount == 0 {
			return id
		}
	}

	return -1
}
