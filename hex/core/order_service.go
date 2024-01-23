package core

import "errors"

// Primary port
type OrderService interface {
	CreateOrder(order Order) error
}

type orderServiceImpl struct {
	repo OrderRepository
}

func NewOrderServicer(repo OrderRepository) OrderService {
	return &orderServiceImpl{repo: repo}
}

// Business core
func (s *orderServiceImpl) CreateOrder(order Order) error {
	if order.Total <= 0 {
		return errors.New("Total must be positive")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}

	return nil
}
