package rotator

import (
	"github.com/widrik/pr/internal/entities"
)

func (rotator *Rotator) findBanner(bannerID uint) (*entities.Banner, error) {
	banner := &entities.Banner{}
	result := rotator.repository.DB.First(banner, bannerID)

	if result.RecordNotFound() || result.Error != nil {
		return nil, ErrBannerNotFound
	}

	return banner, nil
}

func (rotator *Rotator) findSlot(slotID uint) (*entities.Slot, error) {
	slot := &entities.Slot{}
	result := rotator.repository.DB.First(slot, slotID)

	if result.RecordNotFound() || result.Error != nil {
		return nil, ErrSlotNotFound
	}

	return slot, nil
}

func (rotator *Rotator) findSocialGroup(socialGroupID uint) (*entities.SocialGroup, error) {
	socialGroup := &entities.SocialGroup{}
	result := rotator.repository.DB.First(socialGroup, socialGroupID)

	if result.RecordNotFound() || result.Error != nil {
		return nil, ErrSocialGroupNotFound
	}

	return socialGroup, nil
}

func (rotator *Rotator) findOrCreateStats(bannerID, slotID, socialGroupID uint) (*entities.Stats, error) {
	stats := &entities.Stats{}
	searchStat := &entities.Stats{
		BannerID:      bannerID,
		SlotID:        slotID,
		SocialGroupID: socialGroupID,
	}
	result := rotator.repository.DB.FirstOrCreate(stats, searchStat)

	if result.Error != nil {
		return nil, ErrStatsNotFound
	}

	return stats, nil
}
