package rotator

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/widrik/pr/internal/entities"
	"github.com/widrik/pr/internal/repo"
)

var (
	ErrSlotNotFound        = errors.New("slot not found")
	ErrBannerNotFound      = errors.New("banner not found")
	ErrSocialGroupNotFound = errors.New("social group not found")
	ErrStatsNotFound       = errors.New("stats not found")
	ErrBannerDeleteError   = errors.New("banner delete error")
	ErrSlotIsEmpty         = errors.New("slot is empty")
)

type Rotator struct {
	repository *repo.Repository
}

func New(repo *repo.Repository) *Rotator {
	return &Rotator{
		repository: repo,
	}
}

func (rotator *Rotator) Add(bannerID, slotID uint) error {
	banner, err := rotator.findBanner(bannerID)
	if err != nil {
		return ErrBannerNotFound
	}

	slot, err := rotator.findSlot(slotID)
	if err != nil {
		return ErrSlotNotFound
	}

	rotator.repository.DB.Model(banner).Association("Slots").Append(slot)

	return nil
}

func (rotator *Rotator) Delete(bannerID, slotID uint) error {
	banner, err := rotator.findBanner(bannerID)
	if err != nil {
		return err
	}

	slot, err := rotator.findSlot(slotID)
	if err != nil {
		return err
	}

	if rotator.repository.DB.Model(banner).Association("Slots").Delete(slot).Error != nil {
		return ErrBannerDeleteError
	}

	return nil
}

func (rotator *Rotator) Hit(bannerID, slotID, socialGroupID uint) error {
	banner, err := rotator.findBanner(bannerID)
	if err != nil {
		return err
	}

	slot, err := rotator.findSlot(slotID)
	if err != nil {
		return err
	}

	stats, err := rotator.findOrCreateStats(banner.ID, slot.ID, socialGroupID)
	if err != nil {
		return err
	}

	expr := gorm.Expr("click_count + ?", 1)
	result := rotator.repository.DB.Model(stats).UpdateColumn("click_count", expr)

	if result.Error != nil {
		return err
	}

	return nil
}

func (rotator *Rotator) Get(slotID, socialGroupID uint) (*entities.Banner, error) {
	var err error

	slot, err := rotator.findSlot(slotID)
	if err != nil {
		return nil, err
	}

	rotator.repository.DB.Model(slot).Association("Banners").Find(&slot.Banners)
	if len(slot.Banners) == 0 {
		return nil, ErrSlotIsEmpty
	}

	socialGroup, err := rotator.findSocialGroup(socialGroupID)
	if err != nil {
		return nil, err
	}

	state, err := InitAlgoritm(rotator, slot, socialGroup)
	if err != nil {
		return nil, err
	}

	id := UCB1(state)
	selectedBanner := slot.Banners[id]

	stats, err := rotator.findOrCreateStats(selectedBanner.ID, slot.ID, socialGroupID)
	if err != nil {
		return nil, err
	}

	expr := gorm.Expr("show_count + ?", 1)
	result := rotator.repository.DB.Model(stats).UpdateColumn("show_count", expr)

	if result.Error != nil {
		return nil, ErrBannerNotFound
	}

	return slot.Banners[id], nil
}
