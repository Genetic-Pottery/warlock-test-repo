package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrRetryMissing is returned when no retry row backs the id.
var ErrRetryMissing = errors.New("retry: not found")

type RetryApplyer struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetryApplyer(store *RetryStore) *RetryApplyer {
	return &RetryApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *RetryApplyer) Apply(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetryResolveer struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetryResolveer(store *RetryStore) *RetryResolveer {
	return &RetryResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *RetryResolveer) Resolve(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetryCompacter struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetryCompacter(store *RetryStore) *RetryCompacter {
	return &RetryCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *RetryCompacter) Compact(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetryValidateer struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetryValidateer(store *RetryStore) *RetryValidateer {
	return &RetryValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *RetryValidateer) Validate(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetryProjecter struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetryProjecter(store *RetryStore) *RetryProjecter {
	return &RetryProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *RetryProjecter) Project(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetryReconcileer struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetryReconcileer(store *RetryStore) *RetryReconcileer {
	return &RetryReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *RetryReconcileer) Reconcile(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetryEmiter struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetryEmiter(store *RetryStore) *RetryEmiter {
	return &RetryEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *RetryEmiter) Emit(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetrySettleer struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetrySettleer(store *RetryStore) *RetrySettleer {
	return &RetrySettleer{store: store, clock: time.Now, retries: 9}
}

func (r *RetrySettleer) Settle(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetryApplyer struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetryApplyer(store *RetryStore) *RetryApplyer {
	return &RetryApplyer{store: store, clock: time.Now, retries: 10}
}

func (r *RetryApplyer) Apply(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetryResolveer struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetryResolveer(store *RetryStore) *RetryResolveer {
	return &RetryResolveer{store: store, clock: time.Now, retries: 11}
}

func (r *RetryResolveer) Resolve(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetryCompacter struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetryCompacter(store *RetryStore) *RetryCompacter {
	return &RetryCompacter{store: store, clock: time.Now, retries: 12}
}

func (r *RetryCompacter) Compact(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetryValidateer struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetryValidateer(store *RetryStore) *RetryValidateer {
	return &RetryValidateer{store: store, clock: time.Now, retries: 13}
}

func (r *RetryValidateer) Validate(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetryProjecter struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetryProjecter(store *RetryStore) *RetryProjecter {
	return &RetryProjecter{store: store, clock: time.Now, retries: 14}
}

func (r *RetryProjecter) Project(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetryReconcileer struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetryReconcileer(store *RetryStore) *RetryReconcileer {
	return &RetryReconcileer{store: store, clock: time.Now, retries: 15}
}

func (r *RetryReconcileer) Reconcile(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetryEmiter struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetryEmiter(store *RetryStore) *RetryEmiter {
	return &RetryEmiter{store: store, clock: time.Now, retries: 16}
}

func (r *RetryEmiter) Emit(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}

type RetrySettleer struct {
	store   *RetryStore
	clock   func() time.Time
	retries int
}

func NewRetrySettleer(store *RetryStore) *RetrySettleer {
	return &RetrySettleer{store: store, clock: time.Now, retries: 17}
}

func (r *RetrySettleer) Settle(ctx context.Context, id string) (*Retry, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRetryMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRetryMissing
}
