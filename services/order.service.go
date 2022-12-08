package services

import "WhatsappOrderServer/models"

type OrderService interface {
	SaveOrder(*models.Order) (string, error)
	GetAllOrder() ([]*models.DBResponse, error)
}
