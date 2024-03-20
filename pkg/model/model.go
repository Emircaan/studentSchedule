package model

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID               uint      `gorm:"primaryKey"`
	Ad               string    `gorm:"not null"`
	Soyad            string    `gorm:"not null"`
	Eposta           string    `gorm:"not null"`
	Sifre            string    `gorm:"not null"`
	OlusturmaTarihi  time.Time `gorm:"autoCreateTime"`
	GuncellemeTarihi time.Time `gorm:"autoUpdateTime"`
	Plans            []Plan    `gorm:"foreignKey:StudentID"`
}

type Plan struct {
	ID               uint   `gorm:"primaryKey"`
	StudentID        uint   `gorm:"not null"`
	Baslik           string `gorm:"not null"`
	Aciklama         string
	TarihVeSaat      time.Time
	Durum            string
	OlusturmaTarihi  time.Time `gorm:"autoCreateTime"`
	GuncellemeTarihi time.Time `gorm:"autoUpdateTime"`

	Student Student `gorm:"foreignKey:StudentID"`
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Student{}, &Plan{})
	return err
}
