package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Banner struct {
	gorm.Model
	Description string
	Slots       []*Slot `gorm:"many2many:banner_slots;"`
}

type Slot struct {
	gorm.Model
	ID          uint
	Description string
	Banners     []*Banner `gorm:"many2many:banner_slots;"`
}

type SocialGroup struct {
	gorm.Model
	Description string
}

type Stats struct {
	ShowCount     int64 `gorm:"DEFAULT:0"`
	ClickCount    int64 `gorm:"DEFAULT:0"`
	BannerID      uint  `gorm:"primary_key;auto_increment:false"`
	Banner        Banner
	SlotID        uint `gorm:"primary_key;auto_increment:false"`
	Slot          Slot
	SocialGroupID uint `gorm:"primary_key;auto_increment:false"`
	SocialGroup   SocialGroup
	CreatedAt     time.Time
}
