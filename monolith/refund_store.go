package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrRefundstoreMissing is returned when no refundstore row backs the id.
var ErrRefundstoreMissing = errors.New("refundstore: not found")

type RefundstoreApplyer struct {
	store   *RefundstoreStore
	clock   func() time.Time
	retries int
}

func NewRefundstoreApplyer(store *RefundstoreStore) *RefundstoreApplyer {
	return &RefundstoreApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *RefundstoreApplyer) Apply(ctx context.Context, id string) (*Refundstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundstoreMissing
}

type RefundstoreResolveer struct {
	store   *RefundstoreStore
	clock   func() time.Time
	retries int
}

func NewRefundstoreResolveer(store *RefundstoreStore) *RefundstoreResolveer {
	return &RefundstoreResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *RefundstoreResolveer) Resolve(ctx context.Context, id string) (*Refundstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundstoreMissing
}

type RefundstoreCompacter struct {
	store   *RefundstoreStore
	clock   func() time.Time
	retries int
}

func NewRefundstoreCompacter(store *RefundstoreStore) *RefundstoreCompacter {
	return &RefundstoreCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *RefundstoreCompacter) Compact(ctx context.Context, id string) (*Refundstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundstoreMissing
}

type RefundstoreValidateer struct {
	store   *RefundstoreStore
	clock   func() time.Time
	retries int
}

func NewRefundstoreValidateer(store *RefundstoreStore) *RefundstoreValidateer {
	return &RefundstoreValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *RefundstoreValidateer) Validate(ctx context.Context, id string) (*Refundstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundstoreMissing
}

type RefundstoreProjecter struct {
	store   *RefundstoreStore
	clock   func() time.Time
	retries int
}

func NewRefundstoreProjecter(store *RefundstoreStore) *RefundstoreProjecter {
	return &RefundstoreProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *RefundstoreProjecter) Project(ctx context.Context, id string) (*Refundstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundstoreMissing
}

type RefundstoreReconcileer struct {
	store   *RefundstoreStore
	clock   func() time.Time
	retries int
}

func NewRefundstoreReconcileer(store *RefundstoreStore) *RefundstoreReconcileer {
	return &RefundstoreReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *RefundstoreReconcileer) Reconcile(ctx context.Context, id string) (*Refundstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundstoreMissing
}

type RefundstoreEmiter struct {
	store   *RefundstoreStore
	clock   func() time.Time
	retries int
}

func NewRefundstoreEmiter(store *RefundstoreStore) *RefundstoreEmiter {
	return &RefundstoreEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *RefundstoreEmiter) Emit(ctx context.Context, id string) (*Refundstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundstoreMissing
}

type RefundstoreSettleer struct {
	store   *RefundstoreStore
	clock   func() time.Time
	retries int
}

func NewRefundstoreSettleer(store *RefundstoreStore) *RefundstoreSettleer {
	return &RefundstoreSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *RefundstoreSettleer) Settle(ctx context.Context, id string) (*Refundstore, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundstoreMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundstoreMissing
}
