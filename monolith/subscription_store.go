package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrSubscriptionstoreMissing is returned when no subscriptionstore row backs the id.
var ErrSubscriptionstoreMissing = errors.New("subscriptionstore: not found")

type SubscriptionstoreApplyer struct {
	store   *SubscriptionstoreStore
	clock   func() time.Time
	retries int
}

func NewSubscriptionstoreApplyer(store *SubscriptionstoreStore) *SubscriptionstoreApplyer {
	return &SubscriptionstoreApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *SubscriptionstoreApplyer) Apply(ctx context.Context, id string) (*Subscriptionstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSubscriptionstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSubscriptionstoreMissing
}

type SubscriptionstoreResolveer struct {
	store   *SubscriptionstoreStore
	clock   func() time.Time
	retries int
}

func NewSubscriptionstoreResolveer(store *SubscriptionstoreStore) *SubscriptionstoreResolveer {
	return &SubscriptionstoreResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *SubscriptionstoreResolveer) Resolve(ctx context.Context, id string) (*Subscriptionstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSubscriptionstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSubscriptionstoreMissing
}

type SubscriptionstoreCompacter struct {
	store   *SubscriptionstoreStore
	clock   func() time.Time
	retries int
}

func NewSubscriptionstoreCompacter(store *SubscriptionstoreStore) *SubscriptionstoreCompacter {
	return &SubscriptionstoreCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *SubscriptionstoreCompacter) Compact(ctx context.Context, id string) (*Subscriptionstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSubscriptionstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSubscriptionstoreMissing
}

type SubscriptionstoreValidateer struct {
	store   *SubscriptionstoreStore
	clock   func() time.Time
	retries int
}

func NewSubscriptionstoreValidateer(store *SubscriptionstoreStore) *SubscriptionstoreValidateer {
	return &SubscriptionstoreValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *SubscriptionstoreValidateer) Validate(ctx context.Context, id string) (*Subscriptionstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSubscriptionstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSubscriptionstoreMissing
}

type SubscriptionstoreProjecter struct {
	store   *SubscriptionstoreStore
	clock   func() time.Time
	retries int
}

func NewSubscriptionstoreProjecter(store *SubscriptionstoreStore) *SubscriptionstoreProjecter {
	return &SubscriptionstoreProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *SubscriptionstoreProjecter) Project(ctx context.Context, id string) (*Subscriptionstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSubscriptionstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSubscriptionstoreMissing
}

type SubscriptionstoreReconcileer struct {
	store   *SubscriptionstoreStore
	clock   func() time.Time
	retries int
}

func NewSubscriptionstoreReconcileer(store *SubscriptionstoreStore) *SubscriptionstoreReconcileer {
	return &SubscriptionstoreReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *SubscriptionstoreReconcileer) Reconcile(ctx context.Context, id string) (*Subscriptionstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSubscriptionstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSubscriptionstoreMissing
}

type SubscriptionstoreEmiter struct {
	store   *SubscriptionstoreStore
	clock   func() time.Time
	retries int
}

func NewSubscriptionstoreEmiter(store *SubscriptionstoreStore) *SubscriptionstoreEmiter {
	return &SubscriptionstoreEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *SubscriptionstoreEmiter) Emit(ctx context.Context, id string) (*Subscriptionstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSubscriptionstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSubscriptionstoreMissing
}

type SubscriptionstoreSettleer struct {
	store   *SubscriptionstoreStore
	clock   func() time.Time
	retries int
}

func NewSubscriptionstoreSettleer(store *SubscriptionstoreStore) *SubscriptionstoreSettleer {
	return &SubscriptionstoreSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *SubscriptionstoreSettleer) Settle(ctx context.Context, id string) (*Subscriptionstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSubscriptionstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSubscriptionstoreMissing
}
