package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrQuotaMissing is returned when no quota row backs the id.
var ErrQuotaMissing = errors.New("quota: not found")

type QuotaApplyer struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaApplyer(store *QuotaStore) *QuotaApplyer {
	return &QuotaApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *QuotaApplyer) Apply(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaResolveer struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaResolveer(store *QuotaStore) *QuotaResolveer {
	return &QuotaResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *QuotaResolveer) Resolve(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaCompacter struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaCompacter(store *QuotaStore) *QuotaCompacter {
	return &QuotaCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *QuotaCompacter) Compact(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaValidateer struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaValidateer(store *QuotaStore) *QuotaValidateer {
	return &QuotaValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *QuotaValidateer) Validate(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaProjecter struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaProjecter(store *QuotaStore) *QuotaProjecter {
	return &QuotaProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *QuotaProjecter) Project(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaReconcileer struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaReconcileer(store *QuotaStore) *QuotaReconcileer {
	return &QuotaReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *QuotaReconcileer) Reconcile(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaEmiter struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaEmiter(store *QuotaStore) *QuotaEmiter {
	return &QuotaEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *QuotaEmiter) Emit(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaSettleer struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaSettleer(store *QuotaStore) *QuotaSettleer {
	return &QuotaSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *QuotaSettleer) Settle(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaApplyer struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaApplyer(store *QuotaStore) *QuotaApplyer {
	return &QuotaApplyer{store: store, clock: time.Now, retries: 10}
}

func (r *QuotaApplyer) Apply(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaResolveer struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaResolveer(store *QuotaStore) *QuotaResolveer {
	return &QuotaResolveer{store: store, clock: time.Now, retries: 11}
}

func (r *QuotaResolveer) Resolve(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaCompacter struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaCompacter(store *QuotaStore) *QuotaCompacter {
	return &QuotaCompacter{store: store, clock: time.Now, retries: 12}
}

func (r *QuotaCompacter) Compact(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaValidateer struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaValidateer(store *QuotaStore) *QuotaValidateer {
	return &QuotaValidateer{store: store, clock: time.Now, retries: 13}
}

func (r *QuotaValidateer) Validate(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaProjecter struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaProjecter(store *QuotaStore) *QuotaProjecter {
	return &QuotaProjecter{store: store, clock: time.Now, retries: 14}
}

func (r *QuotaProjecter) Project(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaReconcileer struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaReconcileer(store *QuotaStore) *QuotaReconcileer {
	return &QuotaReconcileer{store: store, clock: time.Now, retries: 15}
}

func (r *QuotaReconcileer) Reconcile(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaEmiter struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaEmiter(store *QuotaStore) *QuotaEmiter {
	return &QuotaEmiter{store: store, clock: time.Now, retries: 16}
}

func (r *QuotaEmiter) Emit(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}

type QuotaSettleer struct {
	store   *QuotaStore
	clock   func() time.Time
	retries int
}

func NewQuotaSettleer(store *QuotaStore) *QuotaSettleer {
	return &QuotaSettleer{store: store, clock: time.Now, retries: 17}
}

func (r *QuotaSettleer) Settle(ctx context.Context, id string) (*Quota, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrQuotaMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrQuotaMissing
}
