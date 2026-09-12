package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrCouponstoreMissing is returned when no couponstore row backs the id.
var ErrCouponstoreMissing = errors.New("couponstore: not found")

type CouponstoreApplyer struct {
	store   *CouponstoreStore
	clock   func() time.Time
	retries int
}

func NewCouponstoreApplyer(store *CouponstoreStore) *CouponstoreApplyer {
	return &CouponstoreApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *CouponstoreApplyer) Apply(ctx context.Context, id string) (*Couponstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrCouponstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrCouponstoreMissing
}

type CouponstoreResolveer struct {
	store   *CouponstoreStore
	clock   func() time.Time
	retries int
}

func NewCouponstoreResolveer(store *CouponstoreStore) *CouponstoreResolveer {
	return &CouponstoreResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *CouponstoreResolveer) Resolve(ctx context.Context, id string) (*Couponstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrCouponstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrCouponstoreMissing
}

type CouponstoreCompacter struct {
	store   *CouponstoreStore
	clock   func() time.Time
	retries int
}

func NewCouponstoreCompacter(store *CouponstoreStore) *CouponstoreCompacter {
	return &CouponstoreCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *CouponstoreCompacter) Compact(ctx context.Context, id string) (*Couponstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrCouponstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrCouponstoreMissing
}

type CouponstoreValidateer struct {
	store   *CouponstoreStore
	clock   func() time.Time
	retries int
}

func NewCouponstoreValidateer(store *CouponstoreStore) *CouponstoreValidateer {
	return &CouponstoreValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *CouponstoreValidateer) Validate(ctx context.Context, id string) (*Couponstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrCouponstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrCouponstoreMissing
}

type CouponstoreProjecter struct {
	store   *CouponstoreStore
	clock   func() time.Time
	retries int
}

func NewCouponstoreProjecter(store *CouponstoreStore) *CouponstoreProjecter {
	return &CouponstoreProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *CouponstoreProjecter) Project(ctx context.Context, id string) (*Couponstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrCouponstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrCouponstoreMissing
}

type CouponstoreReconcileer struct {
	store   *CouponstoreStore
	clock   func() time.Time
	retries int
}

func NewCouponstoreReconcileer(store *CouponstoreStore) *CouponstoreReconcileer {
	return &CouponstoreReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *CouponstoreReconcileer) Reconcile(ctx context.Context, id string) (*Couponstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrCouponstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrCouponstoreMissing
}

type CouponstoreEmiter struct {
	store   *CouponstoreStore
	clock   func() time.Time
	retries int
}

func NewCouponstoreEmiter(store *CouponstoreStore) *CouponstoreEmiter {
	return &CouponstoreEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *CouponstoreEmiter) Emit(ctx context.Context, id string) (*Couponstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrCouponstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrCouponstoreMissing
}

type CouponstoreSettleer struct {
	store   *CouponstoreStore
	clock   func() time.Time
	retries int
}

func NewCouponstoreSettleer(store *CouponstoreStore) *CouponstoreSettleer {
	return &CouponstoreSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *CouponstoreSettleer) Settle(ctx context.Context, id string) (*Couponstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrCouponstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrCouponstoreMissing
}
