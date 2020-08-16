package rotator

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/widrik/pr/internal/entities"
	"github.com/widrik/pr/internal/repo"
)

var (
	SlotNotFound      	= errors.New("slot not found")
	BannerNotFound      = errors.New("banner not found")
	SocialGroupNotFound = errors.New("social group not found")
	StatsNotFound 		= errors.New("stats not found")
	BannerDeleteError   = errors.New("banner delete error")
	SlotIsEmpty  	 	= errors.New("slot is empty")
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
		return BannerNotFound
	}

	slot, err := rotator.findSlot(slotID)
	if err != nil {
		return SlotNotFound
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
		return BannerDeleteError
	}

	return nil
}

func (rotator *Rotator) Hit(bannerID, slotID, socialGroupId uint) error {
	banner, err := rotator.findBanner(bannerID)
	if err != nil {
		return err
	}

	slot, err := rotator.findSlot(slotID)
	if err != nil {
		return err
	}

	stats, err := rotator.findOrCreateStats(banner.ID, slot.ID, socialGroupId)
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

func (rotator *Rotator) Get(slotID, socialGroupId uint) (*entities.Banner, error) {
	slot, err := rotator.findSlot(slotID)
	if err != nil {
		return nil, err
	}

	rotator.repository.DB.Model(slot).Association("Banners").Find(&slot.Banners)
	if len(slot.Banners) == 0 {
		return nil, SlotIsEmpty
	}

	socialGroup, err := rotator.findSocialGroup(socialGroupId)
	if err != nil {
		return nil, err
	}

	state, err := Init(rotator, slot, socialGroup)
	if err != nil {
		return nil, err
	}

	id := UCB1(state)
	selectedBanner := slot.Banners[id]

	stats, err := rotator.findOrCreateStats(selectedBanner.ID, slot.ID, socialGroupId)
	expr := gorm.Expr("show_count + ?", 1)
	result := rotator.repository.DB.Model(stats).UpdateColumn("show_count", expr)

	if result.Error != nil {
		return nil, BannerNotFound
	}

	return slot.Banners[id], nil
}