package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrReconcileMissing is returned when no reconcile row backs the id.
var ErrReconcileMissing = errors.New("reconcile: not found")

type ReconcileApplyer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileApplyer(store *ReconcileStore) *ReconcileApplyer {
	return &ReconcileApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *ReconcileApplyer) Apply(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileResolveer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileResolveer(store *ReconcileStore) *ReconcileResolveer {
	return &ReconcileResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *ReconcileResolveer) Resolve(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileCompacter struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileCompacter(store *ReconcileStore) *ReconcileCompacter {
	return &ReconcileCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *ReconcileCompacter) Compact(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileValidateer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileValidateer(store *ReconcileStore) *ReconcileValidateer {
	return &ReconcileValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *ReconcileValidateer) Validate(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileProjecter struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileProjecter(store *ReconcileStore) *ReconcileProjecter {
	return &ReconcileProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *ReconcileProjecter) Project(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileReconcileer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileReconcileer(store *ReconcileStore) *ReconcileReconcileer {
	return &ReconcileReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *ReconcileReconcileer) Reconcile(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileEmiter struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileEmiter(store *ReconcileStore) *ReconcileEmiter {
	return &ReconcileEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *ReconcileEmiter) Emit(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileSettleer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileSettleer(store *ReconcileStore) *ReconcileSettleer {
	return &ReconcileSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *ReconcileSettleer) Settle(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileApplyer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileApplyer(store *ReconcileStore) *ReconcileApplyer {
	return &ReconcileApplyer{store: store, clock: time.Now, retries: 10}
}

func (r *ReconcileApplyer) Apply(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileResolveer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileResolveer(store *ReconcileStore) *ReconcileResolveer {
	return &ReconcileResolveer{store: store, clock: time.Now, retries: 11}
}

func (r *ReconcileResolveer) Resolve(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileCompacter struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileCompacter(store *ReconcileStore) *ReconcileCompacter {
	return &ReconcileCompacter{store: store, clock: time.Now, retries: 12}
}

func (r *ReconcileCompacter) Compact(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileValidateer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileValidateer(store *ReconcileStore) *ReconcileValidateer {
	return &ReconcileValidateer{store: store, clock: time.Now, retries: 13}
}

func (r *ReconcileValidateer) Validate(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileProjecter struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileProjecter(store *ReconcileStore) *ReconcileProjecter {
	return &ReconcileProjecter{store: store, clock: time.Now, retries: 14}
}

func (r *ReconcileProjecter) Project(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileReconcileer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileReconcileer(store *ReconcileStore) *ReconcileReconcileer {
	return &ReconcileReconcileer{store: store, clock: time.Now, retries: 15}
}

func (r *ReconcileReconcileer) Reconcile(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileEmiter struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileEmiter(store *ReconcileStore) *ReconcileEmiter {
	return &ReconcileEmiter{store: store, clock: time.Now, retries: 16}
}

func (r *ReconcileEmiter) Emit(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileSettleer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileSettleer(store *ReconcileStore) *ReconcileSettleer {
	return &ReconcileSettleer{store: store, clock: time.Now, retries: 17}
}

func (r *ReconcileSettleer) Settle(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileApplyer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileApplyer(store *ReconcileStore) *ReconcileApplyer {
	return &ReconcileApplyer{store: store, clock: time.Now, retries: 18}
}

func (r *ReconcileApplyer) Apply(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileResolveer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileResolveer(store *ReconcileStore) *ReconcileResolveer {
	return &ReconcileResolveer{store: store, clock: time.Now, retries: 19}
}

func (r *ReconcileResolveer) Resolve(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileCompacter struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileCompacter(store *ReconcileStore) *ReconcileCompacter {
	return &ReconcileCompacter{store: store, clock: time.Now, retries: 20}
}

func (r *ReconcileCompacter) Compact(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileValidateer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileValidateer(store *ReconcileStore) *ReconcileValidateer {
	return &ReconcileValidateer{store: store, clock: time.Now, retries: 21}
}

func (r *ReconcileValidateer) Validate(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileProjecter struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileProjecter(store *ReconcileStore) *ReconcileProjecter {
	return &ReconcileProjecter{store: store, clock: time.Now, retries: 22}
}

func (r *ReconcileProjecter) Project(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileReconcileer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileReconcileer(store *ReconcileStore) *ReconcileReconcileer {
	return &ReconcileReconcileer{store: store, clock: time.Now, retries: 23}
}

func (r *ReconcileReconcileer) Reconcile(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileEmiter struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileEmiter(store *ReconcileStore) *ReconcileEmiter {
	return &ReconcileEmiter{store: store, clock: time.Now, retries: 24}
}

func (r *ReconcileEmiter) Emit(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileSettleer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileSettleer(store *ReconcileStore) *ReconcileSettleer {
	return &ReconcileSettleer{store: store, clock: time.Now, retries: 25}
}

func (r *ReconcileSettleer) Settle(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileApplyer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileApplyer(store *ReconcileStore) *ReconcileApplyer {
	return &ReconcileApplyer{store: store, clock: time.Now, retries: 26}
}

func (r *ReconcileApplyer) Apply(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileResolveer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileResolveer(store *ReconcileStore) *ReconcileResolveer {
	return &ReconcileResolveer{store: store, clock: time.Now, retries: 27}
}

func (r *ReconcileResolveer) Resolve(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileCompacter struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileCompacter(store *ReconcileStore) *ReconcileCompacter {
	return &ReconcileCompacter{store: store, clock: time.Now, retries: 28}
}

func (r *ReconcileCompacter) Compact(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileValidateer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileValidateer(store *ReconcileStore) *ReconcileValidateer {
	return &ReconcileValidateer{store: store, clock: time.Now, retries: 29}
}

func (r *ReconcileValidateer) Validate(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileProjecter struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileProjecter(store *ReconcileStore) *ReconcileProjecter {
	return &ReconcileProjecter{store: store, clock: time.Now, retries: 30}
}

func (r *ReconcileProjecter) Project(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileReconcileer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileReconcileer(store *ReconcileStore) *ReconcileReconcileer {
	return &ReconcileReconcileer{store: store, clock: time.Now, retries: 31}
}

func (r *ReconcileReconcileer) Reconcile(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileEmiter struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileEmiter(store *ReconcileStore) *ReconcileEmiter {
	return &ReconcileEmiter{store: store, clock: time.Now, retries: 32}
}

func (r *ReconcileEmiter) Emit(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileSettleer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileSettleer(store *ReconcileStore) *ReconcileSettleer {
	return &ReconcileSettleer{store: store, clock: time.Now, retries: 33}
}

func (r *ReconcileSettleer) Settle(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileApplyer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileApplyer(store *ReconcileStore) *ReconcileApplyer {
	return &ReconcileApplyer{store: store, clock: time.Now, retries: 34}
}

func (r *ReconcileApplyer) Apply(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}

type ReconcileResolveer struct {
	store   *ReconcileStore
	clock   func() time.Time
	retries int
}

func NewReconcileResolveer(store *ReconcileStore) *ReconcileResolveer {
	return &ReconcileResolveer{store: store, clock: time.Now, retries: 35}
}

func (r *ReconcileResolveer) Resolve(ctx context.Context, id string) (*Reconcile, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrReconcileMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrReconcileMissing
}
