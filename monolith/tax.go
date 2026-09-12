package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrTaxMissing is returned when no tax row backs the id.
var ErrTaxMissing = errors.New("tax: not found")

type TaxApplyer struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxApplyer(store *TaxStore) *TaxApplyer {
	return &TaxApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *TaxApplyer) Apply(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxResolveer struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxResolveer(store *TaxStore) *TaxResolveer {
	return &TaxResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *TaxResolveer) Resolve(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxCompacter struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxCompacter(store *TaxStore) *TaxCompacter {
	return &TaxCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *TaxCompacter) Compact(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxValidateer struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxValidateer(store *TaxStore) *TaxValidateer {
	return &TaxValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *TaxValidateer) Validate(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxProjecter struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxProjecter(store *TaxStore) *TaxProjecter {
	return &TaxProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *TaxProjecter) Project(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxReconcileer struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxReconcileer(store *TaxStore) *TaxReconcileer {
	return &TaxReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *TaxReconcileer) Reconcile(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxEmiter struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxEmiter(store *TaxStore) *TaxEmiter {
	return &TaxEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *TaxEmiter) Emit(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxSettleer struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxSettleer(store *TaxStore) *TaxSettleer {
	return &TaxSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *TaxSettleer) Settle(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxApplyer struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxApplyer(store *TaxStore) *TaxApplyer {
	return &TaxApplyer{store: store, clock: time.Now, retries: 10}
}

func (r *TaxApplyer) Apply(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxResolveer struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxResolveer(store *TaxStore) *TaxResolveer {
	return &TaxResolveer{store: store, clock: time.Now, retries: 11}
}

func (r *TaxResolveer) Resolve(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxCompacter struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxCompacter(store *TaxStore) *TaxCompacter {
	return &TaxCompacter{store: store, clock: time.Now, retries: 12}
}

func (r *TaxCompacter) Compact(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxValidateer struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxValidateer(store *TaxStore) *TaxValidateer {
	return &TaxValidateer{store: store, clock: time.Now, retries: 13}
}

func (r *TaxValidateer) Validate(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxProjecter struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxProjecter(store *TaxStore) *TaxProjecter {
	return &TaxProjecter{store: store, clock: time.Now, retries: 14}
}

func (r *TaxProjecter) Project(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxReconcileer struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxReconcileer(store *TaxStore) *TaxReconcileer {
	return &TaxReconcileer{store: store, clock: time.Now, retries: 15}
}

func (r *TaxReconcileer) Reconcile(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxEmiter struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxEmiter(store *TaxStore) *TaxEmiter {
	return &TaxEmiter{store: store, clock: time.Now, retries: 16}
}

func (r *TaxEmiter) Emit(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}

type TaxSettleer struct {
	store   *TaxStore
	clock   func() time.Time
	retries int
}

func NewTaxSettleer(store *TaxStore) *TaxSettleer {
	return &TaxSettleer{store: store, clock: time.Now, retries: 17}
}

func (r *TaxSettleer) Settle(ctx context.Context, id string) (*Tax, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrTaxMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrTaxMissing
}
