package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrPayoutstoreMissing is returned when no payoutstore row backs the id.
var ErrPayoutstoreMissing = errors.New("payoutstore: not found")

type PayoutstoreApplyer struct {
	store   *PayoutstoreStore
	clock   func() time.Time
	retries int
}

func NewPayoutstoreApplyer(store *PayoutstoreStore) *PayoutstoreApplyer {
	return &PayoutstoreApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *PayoutstoreApplyer) Apply(ctx context.Context, id string) (*Payoutstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrPayoutstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrPayoutstoreMissing
}

type PayoutstoreResolveer struct {
	store   *PayoutstoreStore
	clock   func() time.Time
	retries int
}

func NewPayoutstoreResolveer(store *PayoutstoreStore) *PayoutstoreResolveer {
	return &PayoutstoreResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *PayoutstoreResolveer) Resolve(ctx context.Context, id string) (*Payoutstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrPayoutstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrPayoutstoreMissing
}

type PayoutstoreCompacter struct {
	store   *PayoutstoreStore
	clock   func() time.Time
	retries int
}

func NewPayoutstoreCompacter(store *PayoutstoreStore) *PayoutstoreCompacter {
	return &PayoutstoreCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *PayoutstoreCompacter) Compact(ctx context.Context, id string) (*Payoutstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrPayoutstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrPayoutstoreMissing
}

type PayoutstoreValidateer struct {
	store   *PayoutstoreStore
	clock   func() time.Time
	retries int
}

func NewPayoutstoreValidateer(store *PayoutstoreStore) *PayoutstoreValidateer {
	return &PayoutstoreValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *PayoutstoreValidateer) Validate(ctx context.Context, id string) (*Payoutstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrPayoutstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrPayoutstoreMissing
}

type PayoutstoreProjecter struct {
	store   *PayoutstoreStore
	clock   func() time.Time
	retries int
}

func NewPayoutstoreProjecter(store *PayoutstoreStore) *PayoutstoreProjecter {
	return &PayoutstoreProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *PayoutstoreProjecter) Project(ctx context.Context, id string) (*Payoutstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrPayoutstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrPayoutstoreMissing
}

type PayoutstoreReconcileer struct {
	store   *PayoutstoreStore
	clock   func() time.Time
	retries int
}

func NewPayoutstoreReconcileer(store *PayoutstoreStore) *PayoutstoreReconcileer {
	return &PayoutstoreReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *PayoutstoreReconcileer) Reconcile(ctx context.Context, id string) (*Payoutstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrPayoutstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrPayoutstoreMissing
}

type PayoutstoreEmiter struct {
	store   *PayoutstoreStore
	clock   func() time.Time
	retries int
}

func NewPayoutstoreEmiter(store *PayoutstoreStore) *PayoutstoreEmiter {
	return &PayoutstoreEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *PayoutstoreEmiter) Emit(ctx context.Context, id string) (*Payoutstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrPayoutstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrPayoutstoreMissing
}

type PayoutstoreSettleer struct {
	store   *PayoutstoreStore
	clock   func() time.Time
	retries int
}

func NewPayoutstoreSettleer(store *PayoutstoreStore) *PayoutstoreSettleer {
	return &PayoutstoreSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *PayoutstoreSettleer) Settle(ctx context.Context, id string) (*Payoutstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrPayoutstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrPayoutstoreMissing
}
