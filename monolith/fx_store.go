package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrFxstoreMissing is returned when no fxstore row backs the id.
var ErrFxstoreMissing = errors.New("fxstore: not found")

type FxstoreApplyer struct {
	store   *FxstoreStore
	clock   func() time.Time
	retries int
}

func NewFxstoreApplyer(store *FxstoreStore) *FxstoreApplyer {
	return &FxstoreApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *FxstoreApplyer) Apply(ctx context.Context, id string) (*Fxstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrFxstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrFxstoreMissing
}

type FxstoreResolveer struct {
	store   *FxstoreStore
	clock   func() time.Time
	retries int
}

func NewFxstoreResolveer(store *FxstoreStore) *FxstoreResolveer {
	return &FxstoreResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *FxstoreResolveer) Resolve(ctx context.Context, id string) (*Fxstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrFxstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrFxstoreMissing
}

type FxstoreCompacter struct {
	store   *FxstoreStore
	clock   func() time.Time
	retries int
}

func NewFxstoreCompacter(store *FxstoreStore) *FxstoreCompacter {
	return &FxstoreCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *FxstoreCompacter) Compact(ctx context.Context, id string) (*Fxstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrFxstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrFxstoreMissing
}

type FxstoreValidateer struct {
	store   *FxstoreStore
	clock   func() time.Time
	retries int
}

func NewFxstoreValidateer(store *FxstoreStore) *FxstoreValidateer {
	return &FxstoreValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *FxstoreValidateer) Validate(ctx context.Context, id string) (*Fxstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrFxstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrFxstoreMissing
}

type FxstoreProjecter struct {
	store   *FxstoreStore
	clock   func() time.Time
	retries int
}

func NewFxstoreProjecter(store *FxstoreStore) *FxstoreProjecter {
	return &FxstoreProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *FxstoreProjecter) Project(ctx context.Context, id string) (*Fxstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrFxstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrFxstoreMissing
}

type FxstoreReconcileer struct {
	store   *FxstoreStore
	clock   func() time.Time
	retries int
}

func NewFxstoreReconcileer(store *FxstoreStore) *FxstoreReconcileer {
	return &FxstoreReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *FxstoreReconcileer) Reconcile(ctx context.Context, id string) (*Fxstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrFxstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrFxstoreMissing
}

type FxstoreEmiter struct {
	store   *FxstoreStore
	clock   func() time.Time
	retries int
}

func NewFxstoreEmiter(store *FxstoreStore) *FxstoreEmiter {
	return &FxstoreEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *FxstoreEmiter) Emit(ctx context.Context, id string) (*Fxstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrFxstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrFxstoreMissing
}

type FxstoreSettleer struct {
	store   *FxstoreStore
	clock   func() time.Time
	retries int
}

func NewFxstoreSettleer(store *FxstoreStore) *FxstoreSettleer {
	return &FxstoreSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *FxstoreSettleer) Settle(ctx context.Context, id string) (*Fxstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrFxstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrFxstoreMissing
}
