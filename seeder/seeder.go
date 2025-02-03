package seeder

import (
	"log"

	"gorm.io/gorm"
)

type Seed struct {
	DB *gorm.DB
}

func (a *Seed) SeederData() error {
	if err := SeedUsers(a.DB); err != nil {
		return err
	}
	log.Println("Success seed Users.")

	return nil
}
