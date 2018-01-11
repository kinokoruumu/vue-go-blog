package model

import (
	"testing"
)

func TestMigration(t *testing.T) {
	db.AutoMigrate(&Article{})
}
