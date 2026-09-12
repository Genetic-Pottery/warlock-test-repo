package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrShipmentstoreMissing is returned when no shipmentstore row backs the id.
var ErrShipmentstoreMissing = errors.New("shipmentstore: not found")

type ShipmentstoreApplyer struct {
	store   *ShipmentstoreStore
	clock   func() time.Time
	retries int
}

func NewShipmentstoreApplyer(store *ShipmentstoreStore) *ShipmentstoreApplyer {
	return &ShipmentstoreApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *ShipmentstoreApplyer) Apply(ctx context.Context, id string) (*Shipmentstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentstoreMissing
}

type ShipmentstoreResolveer struct {
	store   *ShipmentstoreStore
	clock   func() time.Time
	retries int
}

func NewShipmentstoreResolveer(store *ShipmentstoreStore) *ShipmentstoreResolveer {
	return &ShipmentstoreResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *ShipmentstoreResolveer) Resolve(ctx context.Context, id string) (*Shipmentstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentstoreMissing
}

type ShipmentstoreCompacter struct {
	store   *ShipmentstoreStore
	clock   func() time.Time
	retries int
}

func NewShipmentstoreCompacter(store *ShipmentstoreStore) *ShipmentstoreCompacter {
	return &ShipmentstoreCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *ShipmentstoreCompacter) Compact(ctx context.Context, id string) (*Shipmentstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentstoreMissing
}

type ShipmentstoreValidateer struct {
	store   *ShipmentstoreStore
	clock   func() time.Time
	retries int
}

func NewShipmentstoreValidateer(store *ShipmentstoreStore) *ShipmentstoreValidateer {
	return &ShipmentstoreValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *ShipmentstoreValidateer) Validate(ctx context.Context, id string) (*Shipmentstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentstoreMissing
}

type ShipmentstoreProjecter struct {
	store   *ShipmentstoreStore
	clock   func() time.Time
	retries int
}

func NewShipmentstoreProjecter(store *ShipmentstoreStore) *ShipmentstoreProjecter {
	return &ShipmentstoreProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *ShipmentstoreProjecter) Project(ctx context.Context, id string) (*Shipmentstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentstoreMissing
}

type ShipmentstoreReconcileer struct {
	store   *ShipmentstoreStore
	clock   func() time.Time
	retries int
}

func NewShipmentstoreReconcileer(store *ShipmentstoreStore) *ShipmentstoreReconcileer {
	return &ShipmentstoreReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *ShipmentstoreReconcileer) Reconcile(ctx context.Context, id string) (*Shipmentstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentstoreMissing
}

type ShipmentstoreEmiter struct {
	store   *ShipmentstoreStore
	clock   func() time.Time
	retries int
}

func NewShipmentstoreEmiter(store *ShipmentstoreStore) *ShipmentstoreEmiter {
	return &ShipmentstoreEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *ShipmentstoreEmiter) Emit(ctx context.Context, id string) (*Shipmentstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentstoreMissing
}

type ShipmentstoreSettleer struct {
	store   *ShipmentstoreStore
	clock   func() time.Time
	retries int
}

func NewShipmentstoreSettleer(store *ShipmentstoreStore) *ShipmentstoreSettleer {
	return &ShipmentstoreSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *ShipmentstoreSettleer) Settle(ctx context.Context, id string) (*Shipmentstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentstoreMissing
}
