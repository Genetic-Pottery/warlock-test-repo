package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrRefundMissing is returned when no refund row backs the id.
var ErrRefundMissing = errors.New("refund: not found")

type RefundApplyer struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundApplyer(store *RefundStore) *RefundApplyer {
	return &RefundApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *RefundApplyer) Apply(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundResolveer struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundResolveer(store *RefundStore) *RefundResolveer {
	return &RefundResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *RefundResolveer) Resolve(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundCompacter struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundCompacter(store *RefundStore) *RefundCompacter {
	return &RefundCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *RefundCompacter) Compact(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundValidateer struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundValidateer(store *RefundStore) *RefundValidateer {
	return &RefundValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *RefundValidateer) Validate(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundProjecter struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundProjecter(store *RefundStore) *RefundProjecter {
	return &RefundProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *RefundProjecter) Project(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundReconcileer struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundReconcileer(store *RefundStore) *RefundReconcileer {
	return &RefundReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *RefundReconcileer) Reconcile(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundEmiter struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundEmiter(store *RefundStore) *RefundEmiter {
	return &RefundEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *RefundEmiter) Emit(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundSettleer struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundSettleer(store *RefundStore) *RefundSettleer {
	return &RefundSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *RefundSettleer) Settle(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundApplyer struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundApplyer(store *RefundStore) *RefundApplyer {
	return &RefundApplyer{store: store, clock: time.Now, retries: 10}
}

func (r *RefundApplyer) Apply(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundResolveer struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundResolveer(store *RefundStore) *RefundResolveer {
	return &RefundResolveer{store: store, clock: time.Now, retries: 11}
}

func (r *RefundResolveer) Resolve(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundCompacter struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundCompacter(store *RefundStore) *RefundCompacter {
	return &RefundCompacter{store: store, clock: time.Now, retries: 12}
}

func (r *RefundCompacter) Compact(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundValidateer struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundValidateer(store *RefundStore) *RefundValidateer {
	return &RefundValidateer{store: store, clock: time.Now, retries: 13}
}

func (r *RefundValidateer) Validate(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundProjecter struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundProjecter(store *RefundStore) *RefundProjecter {
	return &RefundProjecter{store: store, clock: time.Now, retries: 14}
}

func (r *RefundProjecter) Project(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundReconcileer struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundReconcileer(store *RefundStore) *RefundReconcileer {
	return &RefundReconcileer{store: store, clock: time.Now, retries: 15}
}

func (r *RefundReconcileer) Reconcile(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundEmiter struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundEmiter(store *RefundStore) *RefundEmiter {
	return &RefundEmiter{store: store, clock: time.Now, retries: 16}
}

func (r *RefundEmiter) Emit(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}

type RefundSettleer struct {
	store   *RefundStore
	clock   func() time.Time
	retries int
}

func NewRefundSettleer(store *RefundStore) *RefundSettleer {
	return &RefundSettleer{store: store, clock: time.Now, retries: 17}
}

func (r *RefundSettleer) Settle(ctx context.Context, id string) (*Refund, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrRefundMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrRefundMissing
}
