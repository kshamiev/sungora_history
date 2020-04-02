package model

import (
	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pb/typsun"
	"github.com/kshamiev/sungora/pkg/models"
)

// бизнес модель
type Order struct {
	cm    *config.Component
	Type  *typsun.User
	Model *models.User
}

// NewOrder создания безнес модели
func NewOrder(cm *config.Component) *Order { return &Order{cm: cm} }
