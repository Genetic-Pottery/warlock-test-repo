package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrAuditMissing is returned when no audit row backs the id.
var ErrAuditMissing = errors.New("audit: not found")

type AuditApplyer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditApplyer(store *AuditStore) *AuditApplyer {
	return &AuditApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *AuditApplyer) Apply(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditResolveer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditResolveer(store *AuditStore) *AuditResolveer {
	return &AuditResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *AuditResolveer) Resolve(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditCompacter struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditCompacter(store *AuditStore) *AuditCompacter {
	return &AuditCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *AuditCompacter) Compact(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditValidateer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditValidateer(store *AuditStore) *AuditValidateer {
	return &AuditValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *AuditValidateer) Validate(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditProjecter struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditProjecter(store *AuditStore) *AuditProjecter {
	return &AuditProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *AuditProjecter) Project(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditReconcileer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditReconcileer(store *AuditStore) *AuditReconcileer {
	return &AuditReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *AuditReconcileer) Reconcile(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditEmiter struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditEmiter(store *AuditStore) *AuditEmiter {
	return &AuditEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *AuditEmiter) Emit(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditSettleer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditSettleer(store *AuditStore) *AuditSettleer {
	return &AuditSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *AuditSettleer) Settle(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditApplyer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditApplyer(store *AuditStore) *AuditApplyer {
	return &AuditApplyer{store: store, clock: time.Now, retries: 10}
}

func (r *AuditApplyer) Apply(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditResolveer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditResolveer(store *AuditStore) *AuditResolveer {
	return &AuditResolveer{store: store, clock: time.Now, retries: 11}
}

func (r *AuditResolveer) Resolve(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditCompacter struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditCompacter(store *AuditStore) *AuditCompacter {
	return &AuditCompacter{store: store, clock: time.Now, retries: 12}
}

func (r *AuditCompacter) Compact(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditValidateer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditValidateer(store *AuditStore) *AuditValidateer {
	return &AuditValidateer{store: store, clock: time.Now, retries: 13}
}

func (r *AuditValidateer) Validate(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditProjecter struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditProjecter(store *AuditStore) *AuditProjecter {
	return &AuditProjecter{store: store, clock: time.Now, retries: 14}
}

func (r *AuditProjecter) Project(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditReconcileer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditReconcileer(store *AuditStore) *AuditReconcileer {
	return &AuditReconcileer{store: store, clock: time.Now, retries: 15}
}

func (r *AuditReconcileer) Reconcile(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditEmiter struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditEmiter(store *AuditStore) *AuditEmiter {
	return &AuditEmiter{store: store, clock: time.Now, retries: 16}
}

func (r *AuditEmiter) Emit(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditSettleer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditSettleer(store *AuditStore) *AuditSettleer {
	return &AuditSettleer{store: store, clock: time.Now, retries: 17}
}

func (r *AuditSettleer) Settle(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditApplyer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditApplyer(store *AuditStore) *AuditApplyer {
	return &AuditApplyer{store: store, clock: time.Now, retries: 18}
}

func (r *AuditApplyer) Apply(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditResolveer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditResolveer(store *AuditStore) *AuditResolveer {
	return &AuditResolveer{store: store, clock: time.Now, retries: 19}
}

func (r *AuditResolveer) Resolve(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditCompacter struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditCompacter(store *AuditStore) *AuditCompacter {
	return &AuditCompacter{store: store, clock: time.Now, retries: 20}
}

func (r *AuditCompacter) Compact(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditValidateer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditValidateer(store *AuditStore) *AuditValidateer {
	return &AuditValidateer{store: store, clock: time.Now, retries: 21}
}

func (r *AuditValidateer) Validate(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditProjecter struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditProjecter(store *AuditStore) *AuditProjecter {
	return &AuditProjecter{store: store, clock: time.Now, retries: 22}
}

func (r *AuditProjecter) Project(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditReconcileer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditReconcileer(store *AuditStore) *AuditReconcileer {
	return &AuditReconcileer{store: store, clock: time.Now, retries: 23}
}

func (r *AuditReconcileer) Reconcile(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditEmiter struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditEmiter(store *AuditStore) *AuditEmiter {
	return &AuditEmiter{store: store, clock: time.Now, retries: 24}
}

func (r *AuditEmiter) Emit(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditSettleer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditSettleer(store *AuditStore) *AuditSettleer {
	return &AuditSettleer{store: store, clock: time.Now, retries: 25}
}

func (r *AuditSettleer) Settle(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditApplyer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditApplyer(store *AuditStore) *AuditApplyer {
	return &AuditApplyer{store: store, clock: time.Now, retries: 26}
}

func (r *AuditApplyer) Apply(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditResolveer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditResolveer(store *AuditStore) *AuditResolveer {
	return &AuditResolveer{store: store, clock: time.Now, retries: 27}
}

func (r *AuditResolveer) Resolve(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditCompacter struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditCompacter(store *AuditStore) *AuditCompacter {
	return &AuditCompacter{store: store, clock: time.Now, retries: 28}
}

func (r *AuditCompacter) Compact(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditValidateer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditValidateer(store *AuditStore) *AuditValidateer {
	return &AuditValidateer{store: store, clock: time.Now, retries: 29}
}

func (r *AuditValidateer) Validate(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditProjecter struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditProjecter(store *AuditStore) *AuditProjecter {
	return &AuditProjecter{store: store, clock: time.Now, retries: 30}
}

func (r *AuditProjecter) Project(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditReconcileer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditReconcileer(store *AuditStore) *AuditReconcileer {
	return &AuditReconcileer{store: store, clock: time.Now, retries: 31}
}

func (r *AuditReconcileer) Reconcile(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditEmiter struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditEmiter(store *AuditStore) *AuditEmiter {
	return &AuditEmiter{store: store, clock: time.Now, retries: 32}
}

func (r *AuditEmiter) Emit(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditSettleer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditSettleer(store *AuditStore) *AuditSettleer {
	return &AuditSettleer{store: store, clock: time.Now, retries: 33}
}

func (r *AuditSettleer) Settle(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditApplyer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditApplyer(store *AuditStore) *AuditApplyer {
	return &AuditApplyer{store: store, clock: time.Now, retries: 34}
}

func (r *AuditApplyer) Apply(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}

type AuditResolveer struct {
	store   *AuditStore
	clock   func() time.Time
	retries int
}

func NewAuditResolveer(store *AuditStore) *AuditResolveer {
	return &AuditResolveer{store: store, clock: time.Now, retries: 35}
}

func (r *AuditResolveer) Resolve(ctx context.Context, id string) (*Audit, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrAuditMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrAuditMissing
}
