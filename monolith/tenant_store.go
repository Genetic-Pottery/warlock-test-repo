package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrTenantstoreMissing is returned when no tenantstore row backs the id.
var ErrTenantstoreMissing = errors.New("tenantstore: not found")

type TenantstoreApplyer struct {
	store   *TenantstoreStore
	clock   func() time.Time
	retries int
}

func NewTenantstoreApplyer(store *TenantstoreStore) *TenantstoreApplyer {
	return &TenantstoreApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *TenantstoreApplyer) Apply(ctx context.Context, id string) (*Tenantstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTenantstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTenantstoreMissing
}

type TenantstoreResolveer struct {
	store   *TenantstoreStore
	clock   func() time.Time
	retries int
}

func NewTenantstoreResolveer(store *TenantstoreStore) *TenantstoreResolveer {
	return &TenantstoreResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *TenantstoreResolveer) Resolve(ctx context.Context, id string) (*Tenantstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTenantstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTenantstoreMissing
}

type TenantstoreCompacter struct {
	store   *TenantstoreStore
	clock   func() time.Time
	retries int
}

func NewTenantstoreCompacter(store *TenantstoreStore) *TenantstoreCompacter {
	return &TenantstoreCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *TenantstoreCompacter) Compact(ctx context.Context, id string) (*Tenantstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTenantstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTenantstoreMissing
}

type TenantstoreValidateer struct {
	store   *TenantstoreStore
	clock   func() time.Time
	retries int
}

func NewTenantstoreValidateer(store *TenantstoreStore) *TenantstoreValidateer {
	return &TenantstoreValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *TenantstoreValidateer) Validate(ctx context.Context, id string) (*Tenantstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTenantstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTenantstoreMissing
}

type TenantstoreProjecter struct {
	store   *TenantstoreStore
	clock   func() time.Time
	retries int
}

func NewTenantstoreProjecter(store *TenantstoreStore) *TenantstoreProjecter {
	return &TenantstoreProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *TenantstoreProjecter) Project(ctx context.Context, id string) (*Tenantstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTenantstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTenantstoreMissing
}

type TenantstoreReconcileer struct {
	store   *TenantstoreStore
	clock   func() time.Time
	retries int
}

func NewTenantstoreReconcileer(store *TenantstoreStore) *TenantstoreReconcileer {
	return &TenantstoreReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *TenantstoreReconcileer) Reconcile(ctx context.Context, id string) (*Tenantstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTenantstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTenantstoreMissing
}

type TenantstoreEmiter struct {
	store   *TenantstoreStore
	clock   func() time.Time
	retries int
}

func NewTenantstoreEmiter(store *TenantstoreStore) *TenantstoreEmiter {
	return &TenantstoreEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *TenantstoreEmiter) Emit(ctx context.Context, id string) (*Tenantstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTenantstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTenantstoreMissing
}

type TenantstoreSettleer struct {
	store   *TenantstoreStore
	clock   func() time.Time
	retries int
}

func NewTenantstoreSettleer(store *TenantstoreStore) *TenantstoreSettleer {
	return &TenantstoreSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *TenantstoreSettleer) Settle(ctx context.Context, id string) (*Tenantstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTenantstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTenantstoreMissing
}
