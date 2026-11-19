package common

import (
	"time"

	"github.com/google/uuid"
)

type Audit struct {
	CreatedBy *uuid.UUID `gorm:"type:uuid"` // Maker
	UpdatedBy *uuid.UUID `gorm:"type:uuid"` // Updater
	CreatedAt time.Time  `gorm:"not null;default:now()"`
	UpdatedAt time.Time  `gorm:"not null;default:now()"`
}
