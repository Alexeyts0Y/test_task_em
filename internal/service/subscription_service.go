package service

import (
	"context"

	"github.com/Alexeyts0Y/test_task_em/internal/models"
	"github.com/Alexeyts0Y/test_task_em/internal/repository"
)

type SubscriptionService interface {
	Create(ctx context.Context, sub models.Subscription) (int, error)
	Get(ctx context.Context, id int) (models.Subscription, error)
	Update(ctx context.Context, id int, sub models.Subscription) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, userID, serviceName string) ([]models.Subscription, error)
	CalculateTotalCost(ctx context.Context, req models.CostRequest) (int, error)
}

type subService struct {
	repo repository.SubscriptionRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository) SubscriptionService {
	return &subService{repo: repo}
}

func (s *subService) Create(ctx context.Context, sub models.Subscription) (int, error) {
	return s.repo.Create(ctx, sub)
}

func (s *subService) Get(ctx context.Context, id int) (models.Subscription, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *subService) Update(ctx context.Context, id int, sub models.Subscription) error {
	return s.repo.Update(ctx, id, sub)
}

func (s *subService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *subService) List(ctx context.Context, u, sn string) ([]models.Subscription, error) {
	return s.repo.List(ctx, u, sn)
}

func (s *subService) CalculateTotalCost(ctx context.Context, req models.CostRequest) (int, error) {
	subs, err := s.repo.List(ctx, req.UserID, req.ServiceName)
	if err != nil {
		return 0, err
	}

	total := 0
	for _, sub := range subs {
		endDate := ""
		if sub.EndDate != nil {
			endDate = *sub.EndDate
		}
		months := CalculateOverlap(sub.StartDate, endDate, req.StartDate, req.EndDate)
		total += months * sub.Price
	}
	return total, nil
}
