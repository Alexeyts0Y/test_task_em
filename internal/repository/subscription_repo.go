package repository

import (
	"context"
	"fmt"

	"github.com/Alexeyts0Y/test_task_em/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SubscriptionRepository interface {
	Create(ctx context.Context, sub models.Subscription) (int, error)
	GetByID(ctx context.Context, id int) (models.Subscription, error)
	Update(ctx context.Context, id int, sub models.Subscription) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, userID string, serviceName string) ([]models.Subscription, error)
}

type subRepo struct {
	db *pgxpool.Pool
}

func NewSubscriptionRepo(db *pgxpool.Pool) SubscriptionRepository {
	return &subRepo{db: db}
}

func (r *subRepo) Create(ctx context.Context, s models.Subscription) (int, error) {
	var id int
	query := `INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(ctx, query, s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate).Scan(&id)
	return id, err
}

func (r *subRepo) GetByID(ctx context.Context, id int) (models.Subscription, error) {
	var s models.Subscription
	query := `SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate)
	return s, err
}

func (r *subRepo) Update(ctx context.Context, id int, s models.Subscription) error {
	query := `UPDATE subscriptions SET service_name=$1, price=$2, user_id=$3, start_date=$4, end_date=$5 WHERE id=$6`
	_, err := r.db.Exec(ctx, query, s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate, id)
	return err
}

func (r *subRepo) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM subscriptions WHERE id=$1", id)
	return err
}

func (r *subRepo) List(ctx context.Context, userID string, serviceName string) ([]models.Subscription, error) {
	query := `SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE 1=1`
	args := []interface{}{}
	argID := 1

	if userID != "" {
		query += fmt.Sprintf(" AND user_id = $%d", argID)
		args = append(args, userID)
		argID++
	}
	if serviceName != "" {
		query += fmt.Sprintf(" AND service_name = $%d", argID)
		args = append(args, serviceName)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []models.Subscription
	for rows.Next() {
		var s models.Subscription
		if err := rows.Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}
	return subs, nil
}
