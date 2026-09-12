package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrUsageMissing is returned when no usage row backs the id.
var ErrUsageMissing = errors.New("usage: not found")

type UsageApplyer struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageApplyer(store *UsageStore) *UsageApplyer {
	return &UsageApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *UsageApplyer) Apply(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageResolveer struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageResolveer(store *UsageStore) *UsageResolveer {
	return &UsageResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *UsageResolveer) Resolve(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageCompacter struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageCompacter(store *UsageStore) *UsageCompacter {
	return &UsageCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *UsageCompacter) Compact(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageValidateer struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageValidateer(store *UsageStore) *UsageValidateer {
	return &UsageValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *UsageValidateer) Validate(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageProjecter struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageProjecter(store *UsageStore) *UsageProjecter {
	return &UsageProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *UsageProjecter) Project(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageReconcileer struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageReconcileer(store *UsageStore) *UsageReconcileer {
	return &UsageReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *UsageReconcileer) Reconcile(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageEmiter struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageEmiter(store *UsageStore) *UsageEmiter {
	return &UsageEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *UsageEmiter) Emit(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageSettleer struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageSettleer(store *UsageStore) *UsageSettleer {
	return &UsageSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *UsageSettleer) Settle(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageApplyer struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageApplyer(store *UsageStore) *UsageApplyer {
	return &UsageApplyer{store: store, clock: time.Now, retries: 10}
}

func (r *UsageApplyer) Apply(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageResolveer struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageResolveer(store *UsageStore) *UsageResolveer {
	return &UsageResolveer{store: store, clock: time.Now, retries: 11}
}

func (r *UsageResolveer) Resolve(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageCompacter struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageCompacter(store *UsageStore) *UsageCompacter {
	return &UsageCompacter{store: store, clock: time.Now, retries: 12}
}

func (r *UsageCompacter) Compact(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageValidateer struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageValidateer(store *UsageStore) *UsageValidateer {
	return &UsageValidateer{store: store, clock: time.Now, retries: 13}
}

func (r *UsageValidateer) Validate(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageProjecter struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageProjecter(store *UsageStore) *UsageProjecter {
	return &UsageProjecter{store: store, clock: time.Now, retries: 14}
}

func (r *UsageProjecter) Project(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageReconcileer struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageReconcileer(store *UsageStore) *UsageReconcileer {
	return &UsageReconcileer{store: store, clock: time.Now, retries: 15}
}

func (r *UsageReconcileer) Reconcile(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageEmiter struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageEmiter(store *UsageStore) *UsageEmiter {
	return &UsageEmiter{store: store, clock: time.Now, retries: 16}
}

func (r *UsageEmiter) Emit(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}

type UsageSettleer struct {
	store   *UsageStore
	clock   func() time.Time
	retries int
}

func NewUsageSettleer(store *UsageStore) *UsageSettleer {
	return &UsageSettleer{store: store, clock: time.Now, retries: 17}
}

func (r *UsageSettleer) Settle(ctx context.Context, id string) (*Usage, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrUsageMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrUsageMissing
}
