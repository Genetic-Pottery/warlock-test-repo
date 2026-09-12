package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrLedgerMissing is returned when no ledger row backs the id.
var ErrLedgerMissing = errors.New("ledger: not found")

type LedgerApplyer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerApplyer(store *LedgerStore) *LedgerApplyer {
	return &LedgerApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *LedgerApplyer) Apply(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerResolveer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerResolveer(store *LedgerStore) *LedgerResolveer {
	return &LedgerResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *LedgerResolveer) Resolve(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerCompacter struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerCompacter(store *LedgerStore) *LedgerCompacter {
	return &LedgerCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *LedgerCompacter) Compact(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerValidateer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerValidateer(store *LedgerStore) *LedgerValidateer {
	return &LedgerValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *LedgerValidateer) Validate(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerProjecter struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerProjecter(store *LedgerStore) *LedgerProjecter {
	return &LedgerProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *LedgerProjecter) Project(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerReconcileer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerReconcileer(store *LedgerStore) *LedgerReconcileer {
	return &LedgerReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *LedgerReconcileer) Reconcile(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerEmiter struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerEmiter(store *LedgerStore) *LedgerEmiter {
	return &LedgerEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *LedgerEmiter) Emit(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerSettleer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerSettleer(store *LedgerStore) *LedgerSettleer {
	return &LedgerSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *LedgerSettleer) Settle(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerApplyer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerApplyer(store *LedgerStore) *LedgerApplyer {
	return &LedgerApplyer{store: store, clock: time.Now, retries: 10}
}

func (r *LedgerApplyer) Apply(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerResolveer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerResolveer(store *LedgerStore) *LedgerResolveer {
	return &LedgerResolveer{store: store, clock: time.Now, retries: 11}
}

func (r *LedgerResolveer) Resolve(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerCompacter struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerCompacter(store *LedgerStore) *LedgerCompacter {
	return &LedgerCompacter{store: store, clock: time.Now, retries: 12}
}

func (r *LedgerCompacter) Compact(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerValidateer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerValidateer(store *LedgerStore) *LedgerValidateer {
	return &LedgerValidateer{store: store, clock: time.Now, retries: 13}
}

func (r *LedgerValidateer) Validate(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerProjecter struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerProjecter(store *LedgerStore) *LedgerProjecter {
	return &LedgerProjecter{store: store, clock: time.Now, retries: 14}
}

func (r *LedgerProjecter) Project(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerReconcileer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerReconcileer(store *LedgerStore) *LedgerReconcileer {
	return &LedgerReconcileer{store: store, clock: time.Now, retries: 15}
}

func (r *LedgerReconcileer) Reconcile(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerEmiter struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerEmiter(store *LedgerStore) *LedgerEmiter {
	return &LedgerEmiter{store: store, clock: time.Now, retries: 16}
}

func (r *LedgerEmiter) Emit(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerSettleer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerSettleer(store *LedgerStore) *LedgerSettleer {
	return &LedgerSettleer{store: store, clock: time.Now, retries: 17}
}

func (r *LedgerSettleer) Settle(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerApplyer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerApplyer(store *LedgerStore) *LedgerApplyer {
	return &LedgerApplyer{store: store, clock: time.Now, retries: 18}
}

func (r *LedgerApplyer) Apply(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerResolveer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerResolveer(store *LedgerStore) *LedgerResolveer {
	return &LedgerResolveer{store: store, clock: time.Now, retries: 19}
}

func (r *LedgerResolveer) Resolve(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerCompacter struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerCompacter(store *LedgerStore) *LedgerCompacter {
	return &LedgerCompacter{store: store, clock: time.Now, retries: 20}
}

func (r *LedgerCompacter) Compact(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerValidateer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerValidateer(store *LedgerStore) *LedgerValidateer {
	return &LedgerValidateer{store: store, clock: time.Now, retries: 21}
}

func (r *LedgerValidateer) Validate(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerProjecter struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerProjecter(store *LedgerStore) *LedgerProjecter {
	return &LedgerProjecter{store: store, clock: time.Now, retries: 22}
}

func (r *LedgerProjecter) Project(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerReconcileer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerReconcileer(store *LedgerStore) *LedgerReconcileer {
	return &LedgerReconcileer{store: store, clock: time.Now, retries: 23}
}

func (r *LedgerReconcileer) Reconcile(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerEmiter struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerEmiter(store *LedgerStore) *LedgerEmiter {
	return &LedgerEmiter{store: store, clock: time.Now, retries: 24}
}

func (r *LedgerEmiter) Emit(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerSettleer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerSettleer(store *LedgerStore) *LedgerSettleer {
	return &LedgerSettleer{store: store, clock: time.Now, retries: 25}
}

func (r *LedgerSettleer) Settle(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerApplyer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerApplyer(store *LedgerStore) *LedgerApplyer {
	return &LedgerApplyer{store: store, clock: time.Now, retries: 26}
}

func (r *LedgerApplyer) Apply(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerResolveer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerResolveer(store *LedgerStore) *LedgerResolveer {
	return &LedgerResolveer{store: store, clock: time.Now, retries: 27}
}

func (r *LedgerResolveer) Resolve(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerCompacter struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerCompacter(store *LedgerStore) *LedgerCompacter {
	return &LedgerCompacter{store: store, clock: time.Now, retries: 28}
}

func (r *LedgerCompacter) Compact(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerValidateer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerValidateer(store *LedgerStore) *LedgerValidateer {
	return &LedgerValidateer{store: store, clock: time.Now, retries: 29}
}

func (r *LedgerValidateer) Validate(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerProjecter struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerProjecter(store *LedgerStore) *LedgerProjecter {
	return &LedgerProjecter{store: store, clock: time.Now, retries: 30}
}

func (r *LedgerProjecter) Project(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerReconcileer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerReconcileer(store *LedgerStore) *LedgerReconcileer {
	return &LedgerReconcileer{store: store, clock: time.Now, retries: 31}
}

func (r *LedgerReconcileer) Reconcile(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerEmiter struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerEmiter(store *LedgerStore) *LedgerEmiter {
	return &LedgerEmiter{store: store, clock: time.Now, retries: 32}
}

func (r *LedgerEmiter) Emit(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerSettleer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerSettleer(store *LedgerStore) *LedgerSettleer {
	return &LedgerSettleer{store: store, clock: time.Now, retries: 33}
}

func (r *LedgerSettleer) Settle(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerApplyer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerApplyer(store *LedgerStore) *LedgerApplyer {
	return &LedgerApplyer{store: store, clock: time.Now, retries: 34}
}

func (r *LedgerApplyer) Apply(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}

type LedgerResolveer struct {
	store   *LedgerStore
	clock   func() time.Time
	retries int
}

func NewLedgerResolveer(store *LedgerStore) *LedgerResolveer {
	return &LedgerResolveer{store: store, clock: time.Now, retries: 35}
}

func (r *LedgerResolveer) Resolve(ctx context.Context, id string) (*Ledger, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrLedgerMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrLedgerMissing
}
