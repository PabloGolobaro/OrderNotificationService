package services

import "WhatsappOrderServer/models"

type OrderService interface {
	SaveOrder(*models.Order) (int, error)
	GetAllOrder() ([]*models.DBResponse, error)
}
