package database

import (
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-classroom/internal/subscription/domain"
	merrors "github.com/sumelms/microservice-classroom/pkg/errors"
)

const (
	whereSubscriptionID = "UUID = ?"
)

type Repository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepository(db *gorm.DB, logger log.Logger) *Repository {
	db.AutoMigrate(&Subscription{})

	return &Repository{db: db, logger: logger}
}

func (r *Repository) Create(subscription *domain.Subscription) (domain.Subscription, error) {
	entity := toDBModel(subscription)

	if err := r.db.Create(&entity).Error; err != nil {
		return domain.Subscription{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't create subscription")
	}
	return toDomainModel(&entity), nil
}

func (r *Repository) Find(id string) (domain.Subscription, error) {
	var subscription Subscription

	query := r.db.Where(whereSubscriptionID, id).First(&subscription)
	if query.RecordNotFound() {
		return domain.Subscription{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "subscription not found")
	}
	if err := query.Error; err != nil {
		return domain.Subscription{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "find subscription")
	}

	return toDomainModel(&subscription), nil
}

func (r *Repository) Update(s *domain.Subscription) (domain.Subscription, error) {
	var subscription Subscription

	query := r.db.Where(whereSubscriptionID, s.UUID).First(&subscription)

	if query.RecordNotFound() {
		return domain.Subscription{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "subscription not found")
	}

	query = r.db.Model(&subscription).Updates(&s)

	if err := query.Error; err != nil {
		return domain.Subscription{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't update subscription")
	}

	return *s, nil
}

func (r *Repository) Delete(id string) error {
	query := r.db.Where(whereSubscriptionID, id).Delete(&Subscription{})

	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return merrors.WrapErrorf(err, merrors.ErrCodeNotFound, "subscription not found")
		}
	}

	return nil
}

func (r *Repository) List(filters map[string]interface{}) ([]domain.Subscription, error) {
	var subscriptions []Subscription

	query := r.db.Find(&subscriptions, filters)
	if query.RecordNotFound() {
		return []domain.Subscription{}, nil
	}

	var list []domain.Subscription
	for i := range subscriptions {
		s := subscriptions[i]
		list = append(list, toDomainModel(&s))
	}

	return list, nil
}