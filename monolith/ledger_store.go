package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrLedgerstoreMissing is returned when no ledgerstore row backs the id.
var ErrLedgerstoreMissing = errors.New("ledgerstore: not found")

type LedgerstoreApplyer struct {
	store   *LedgerstoreStore
	clock   func() time.Time
	retries int
}

func NewLedgerstoreApplyer(store *LedgerstoreStore) *LedgerstoreApplyer {
	return &LedgerstoreApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *LedgerstoreApplyer) Apply(ctx context.Context, id string) (*Ledgerstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerstoreMissing
}

type LedgerstoreResolveer struct {
	store   *LedgerstoreStore
	clock   func() time.Time
	retries int
}

func NewLedgerstoreResolveer(store *LedgerstoreStore) *LedgerstoreResolveer {
	return &LedgerstoreResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *LedgerstoreResolveer) Resolve(ctx context.Context, id string) (*Ledgerstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerstoreMissing
}

type LedgerstoreCompacter struct {
	store   *LedgerstoreStore
	clock   func() time.Time
	retries int
}

func NewLedgerstoreCompacter(store *LedgerstoreStore) *LedgerstoreCompacter {
	return &LedgerstoreCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *LedgerstoreCompacter) Compact(ctx context.Context, id string) (*Ledgerstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerstoreMissing
}

type LedgerstoreValidateer struct {
	store   *LedgerstoreStore
	clock   func() time.Time
	retries int
}

func NewLedgerstoreValidateer(store *LedgerstoreStore) *LedgerstoreValidateer {
	return &LedgerstoreValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *LedgerstoreValidateer) Validate(ctx context.Context, id string) (*Ledgerstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerstoreMissing
}

type LedgerstoreProjecter struct {
	store   *LedgerstoreStore
	clock   func() time.Time
	retries int
}

func NewLedgerstoreProjecter(store *LedgerstoreStore) *LedgerstoreProjecter {
	return &LedgerstoreProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *LedgerstoreProjecter) Project(ctx context.Context, id string) (*Ledgerstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerstoreMissing
}

type LedgerstoreReconcileer struct {
	store   *LedgerstoreStore
	clock   func() time.Time
	retries int
}

func NewLedgerstoreReconcileer(store *LedgerstoreStore) *LedgerstoreReconcileer {
	return &LedgerstoreReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *LedgerstoreReconcileer) Reconcile(ctx context.Context, id string) (*Ledgerstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerstoreMissing
}

type LedgerstoreEmiter struct {
	store   *LedgerstoreStore
	clock   func() time.Time
	retries int
}

func NewLedgerstoreEmiter(store *LedgerstoreStore) *LedgerstoreEmiter {
	return &LedgerstoreEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *LedgerstoreEmiter) Emit(ctx context.Context, id string) (*Ledgerstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerstoreMissing
}

type LedgerstoreSettleer struct {
	store   *LedgerstoreStore
	clock   func() time.Time
	retries int
}

func NewLedgerstoreSettleer(store *LedgerstoreStore) *LedgerstoreSettleer {
	return &LedgerstoreSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *LedgerstoreSettleer) Settle(ctx context.Context, id string) (*Ledgerstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerstoreMissing
}
