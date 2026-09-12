package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrSettlementstoreMissing is returned when no settlementstore row backs the id.
var ErrSettlementstoreMissing = errors.New("settlementstore: not found")

type SettlementstoreApplyer struct {
	store   *SettlementstoreStore
	clock   func() time.Time
	retries int
}

func NewSettlementstoreApplyer(store *SettlementstoreStore) *SettlementstoreApplyer {
	return &SettlementstoreApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *SettlementstoreApplyer) Apply(ctx context.Context, id string) (*Settlementstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSettlementstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSettlementstoreMissing
}

type SettlementstoreResolveer struct {
	store   *SettlementstoreStore
	clock   func() time.Time
	retries int
}

func NewSettlementstoreResolveer(store *SettlementstoreStore) *SettlementstoreResolveer {
	return &SettlementstoreResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *SettlementstoreResolveer) Resolve(ctx context.Context, id string) (*Settlementstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSettlementstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSettlementstoreMissing
}

type SettlementstoreCompacter struct {
	store   *SettlementstoreStore
	clock   func() time.Time
	retries int
}

func NewSettlementstoreCompacter(store *SettlementstoreStore) *SettlementstoreCompacter {
	return &SettlementstoreCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *SettlementstoreCompacter) Compact(ctx context.Context, id string) (*Settlementstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSettlementstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSettlementstoreMissing
}

type SettlementstoreValidateer struct {
	store   *SettlementstoreStore
	clock   func() time.Time
	retries int
}

func NewSettlementstoreValidateer(store *SettlementstoreStore) *SettlementstoreValidateer {
	return &SettlementstoreValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *SettlementstoreValidateer) Validate(ctx context.Context, id string) (*Settlementstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSettlementstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSettlementstoreMissing
}

type SettlementstoreProjecter struct {
	store   *SettlementstoreStore
	clock   func() time.Time
	retries int
}

func NewSettlementstoreProjecter(store *SettlementstoreStore) *SettlementstoreProjecter {
	return &SettlementstoreProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *SettlementstoreProjecter) Project(ctx context.Context, id string) (*Settlementstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSettlementstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSettlementstoreMissing
}

type SettlementstoreReconcileer struct {
	store   *SettlementstoreStore
	clock   func() time.Time
	retries int
}

func NewSettlementstoreReconcileer(store *SettlementstoreStore) *SettlementstoreReconcileer {
	return &SettlementstoreReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *SettlementstoreReconcileer) Reconcile(ctx context.Context, id string) (*Settlementstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSettlementstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSettlementstoreMissing
}

type SettlementstoreEmiter struct {
	store   *SettlementstoreStore
	clock   func() time.Time
	retries int
}

func NewSettlementstoreEmiter(store *SettlementstoreStore) *SettlementstoreEmiter {
	return &SettlementstoreEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *SettlementstoreEmiter) Emit(ctx context.Context, id string) (*Settlementstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSettlementstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSettlementstoreMissing
}

type SettlementstoreSettleer struct {
	store   *SettlementstoreStore
	clock   func() time.Time
	retries int
}

func NewSettlementstoreSettleer(store *SettlementstoreStore) *SettlementstoreSettleer {
	return &SettlementstoreSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *SettlementstoreSettleer) Settle(ctx context.Context, id string) (*Settlementstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrSettlementstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrSettlementstoreMissing
}
