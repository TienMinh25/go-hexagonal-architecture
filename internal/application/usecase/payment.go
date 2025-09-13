package usecase

import (
	"context"

	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain"
	domainpayment "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/payment"
	portin "github.com/TienMinh25/go-hexagonal-architecture/internal/application/port/in"
	portout "github.com/TienMinh25/go-hexagonal-architecture/internal/application/port/out"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/util"
)

/**
 * paymentUsecase implements portin.PaymentService interface
 * and provides an access to the payment repository
 * and cache service
 */
type paymentUsecase struct {
	repo  portout.PaymentRepository
	cache portout.CacheRepository
}

// NewPaymentUsecase creates a new payment service instance
func NewPaymentUsecase(repo portout.PaymentRepository, cache portout.CacheRepository) portin.PaymentService {
	return &paymentUsecase{
		repo,
		cache,
	}
}

// CreatePayment creates a new payment
func (ps *paymentUsecase) CreatePayment(ctx context.Context, payment *domainpayment.Payment) (*domainpayment.Payment, error) {
	payment, err := ps.repo.CreatePayment(ctx, payment)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	cacheKey := util.GenerateCacheKey("payment", payment.ID)
	paymentSerialized, err := util.Serialize(payment)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = ps.cache.Set(ctx, cacheKey, paymentSerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = ps.cache.DeleteByPrefix(ctx, "payments:*")
	if err != nil {
		return nil, domain.ErrInternal
	}

	return payment, nil
}

// GetPayment retrieves a payment by id
func (ps *paymentUsecase) GetPayment(ctx context.Context, id uint64) (*domainpayment.Payment, error) {
	var payment *domainpayment.Payment

	cacheKey := util.GenerateCacheKey("payment", id)
	cachedPayment, err := ps.cache.Get(ctx, cacheKey)
	if err == nil {
		err := util.Deserialize(cachedPayment, &payment)
		if err != nil {
			return nil, domain.ErrInternal
		}

		return payment, nil
	}

	payment, err = ps.repo.GetPaymentByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	paymentSerialized, err := util.Serialize(payment)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = ps.cache.Set(ctx, cacheKey, paymentSerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return payment, nil
}

// ListPayments retrieves a list of payments
func (ps *paymentUsecase) ListPayments(ctx context.Context, skip, limit uint64) ([]domainpayment.Payment, error) {
	var payments []domainpayment.Payment

	params := util.GenerateCacheKeyParams(skip, limit)
	cacheKey := util.GenerateCacheKey("payments", params)

	cachedPayments, err := ps.cache.Get(ctx, cacheKey)
	if err == nil {
		err := util.Deserialize(cachedPayments, &payments)
		if err != nil {
			return nil, domain.ErrInternal
		}

		return payments, nil
	}

	payments, err = ps.repo.ListPayments(ctx, skip, limit)
	if err != nil {
		return nil, domain.ErrInternal
	}

	paymentsSerialized, err := util.Serialize(payments)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = ps.cache.Set(ctx, cacheKey, paymentsSerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return payments, nil

}

// UpdatePayment updates a payment
func (ps *paymentUsecase) UpdatePayment(ctx context.Context, payment *domainpayment.Payment) (*domainpayment.Payment, error) {
	existingPayment, err := ps.repo.GetPaymentByID(ctx, payment.ID)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	emptyData := payment.Name == "" && payment.Type == "" && payment.Logo == ""
	sameData := existingPayment.Name == payment.Name && existingPayment.Type == payment.Type && existingPayment.Logo == payment.Logo
	if emptyData || sameData {
		return nil, domain.ErrNoUpdatedData
	}

	_, err = ps.repo.UpdatePayment(ctx, payment)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	cacheKey := util.GenerateCacheKey("payment", payment.ID)

	err = ps.cache.Delete(ctx, cacheKey)
	if err != nil {
		return nil, domain.ErrInternal
	}

	paymentSerialized, err := util.Serialize(payment)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = ps.cache.Set(ctx, cacheKey, paymentSerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = ps.cache.DeleteByPrefix(ctx, "payments:*")
	if err != nil {
		return nil, domain.ErrInternal
	}

	return payment, nil
}

// DeletePayment deletes a payment
func (ps *paymentUsecase) DeletePayment(ctx context.Context, id uint64) error {
	_, err := ps.repo.GetPaymentByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return err
		}
		return domain.ErrInternal
	}

	cacheKey := util.GenerateCacheKey("payment", id)

	err = ps.cache.Delete(ctx, cacheKey)
	if err != nil {
		return domain.ErrInternal
	}

	err = ps.cache.DeleteByPrefix(ctx, "payments:*")
	if err != nil {
		return domain.ErrInternal
	}

	return ps.repo.DeletePayment(ctx, id)
}
