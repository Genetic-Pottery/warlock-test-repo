package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrShipmentMissing is returned when no shipment row backs the id.
var ErrShipmentMissing = errors.New("shipment: not found")

type ShipmentApplyer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentApplyer(store *ShipmentStore) *ShipmentApplyer {
	return &ShipmentApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *ShipmentApplyer) Apply(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentResolveer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentResolveer(store *ShipmentStore) *ShipmentResolveer {
	return &ShipmentResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *ShipmentResolveer) Resolve(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentCompacter struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentCompacter(store *ShipmentStore) *ShipmentCompacter {
	return &ShipmentCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *ShipmentCompacter) Compact(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentValidateer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentValidateer(store *ShipmentStore) *ShipmentValidateer {
	return &ShipmentValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *ShipmentValidateer) Validate(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentProjecter struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentProjecter(store *ShipmentStore) *ShipmentProjecter {
	return &ShipmentProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *ShipmentProjecter) Project(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentReconcileer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentReconcileer(store *ShipmentStore) *ShipmentReconcileer {
	return &ShipmentReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *ShipmentReconcileer) Reconcile(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentEmiter struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentEmiter(store *ShipmentStore) *ShipmentEmiter {
	return &ShipmentEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *ShipmentEmiter) Emit(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentSettleer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentSettleer(store *ShipmentStore) *ShipmentSettleer {
	return &ShipmentSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *ShipmentSettleer) Settle(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentApplyer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentApplyer(store *ShipmentStore) *ShipmentApplyer {
	return &ShipmentApplyer{store: store, clock: time.Now, retries: 10}
}

func (r *ShipmentApplyer) Apply(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentResolveer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentResolveer(store *ShipmentStore) *ShipmentResolveer {
	return &ShipmentResolveer{store: store, clock: time.Now, retries: 11}
}

func (r *ShipmentResolveer) Resolve(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentCompacter struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentCompacter(store *ShipmentStore) *ShipmentCompacter {
	return &ShipmentCompacter{store: store, clock: time.Now, retries: 12}
}

func (r *ShipmentCompacter) Compact(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentValidateer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentValidateer(store *ShipmentStore) *ShipmentValidateer {
	return &ShipmentValidateer{store: store, clock: time.Now, retries: 13}
}

func (r *ShipmentValidateer) Validate(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentProjecter struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentProjecter(store *ShipmentStore) *ShipmentProjecter {
	return &ShipmentProjecter{store: store, clock: time.Now, retries: 14}
}

func (r *ShipmentProjecter) Project(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentReconcileer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentReconcileer(store *ShipmentStore) *ShipmentReconcileer {
	return &ShipmentReconcileer{store: store, clock: time.Now, retries: 15}
}

func (r *ShipmentReconcileer) Reconcile(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentEmiter struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentEmiter(store *ShipmentStore) *ShipmentEmiter {
	return &ShipmentEmiter{store: store, clock: time.Now, retries: 16}
}

func (r *ShipmentEmiter) Emit(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentSettleer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentSettleer(store *ShipmentStore) *ShipmentSettleer {
	return &ShipmentSettleer{store: store, clock: time.Now, retries: 17}
}

func (r *ShipmentSettleer) Settle(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentApplyer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentApplyer(store *ShipmentStore) *ShipmentApplyer {
	return &ShipmentApplyer{store: store, clock: time.Now, retries: 18}
}

func (r *ShipmentApplyer) Apply(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentResolveer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentResolveer(store *ShipmentStore) *ShipmentResolveer {
	return &ShipmentResolveer{store: store, clock: time.Now, retries: 19}
}

func (r *ShipmentResolveer) Resolve(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentCompacter struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentCompacter(store *ShipmentStore) *ShipmentCompacter {
	return &ShipmentCompacter{store: store, clock: time.Now, retries: 20}
}

func (r *ShipmentCompacter) Compact(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentValidateer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentValidateer(store *ShipmentStore) *ShipmentValidateer {
	return &ShipmentValidateer{store: store, clock: time.Now, retries: 21}
}

func (r *ShipmentValidateer) Validate(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentProjecter struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentProjecter(store *ShipmentStore) *ShipmentProjecter {
	return &ShipmentProjecter{store: store, clock: time.Now, retries: 22}
}

func (r *ShipmentProjecter) Project(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentReconcileer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentReconcileer(store *ShipmentStore) *ShipmentReconcileer {
	return &ShipmentReconcileer{store: store, clock: time.Now, retries: 23}
}

func (r *ShipmentReconcileer) Reconcile(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentEmiter struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentEmiter(store *ShipmentStore) *ShipmentEmiter {
	return &ShipmentEmiter{store: store, clock: time.Now, retries: 24}
}

func (r *ShipmentEmiter) Emit(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentSettleer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentSettleer(store *ShipmentStore) *ShipmentSettleer {
	return &ShipmentSettleer{store: store, clock: time.Now, retries: 25}
}

func (r *ShipmentSettleer) Settle(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentApplyer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentApplyer(store *ShipmentStore) *ShipmentApplyer {
	return &ShipmentApplyer{store: store, clock: time.Now, retries: 26}
}

func (r *ShipmentApplyer) Apply(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentResolveer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentResolveer(store *ShipmentStore) *ShipmentResolveer {
	return &ShipmentResolveer{store: store, clock: time.Now, retries: 27}
}

func (r *ShipmentResolveer) Resolve(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentCompacter struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentCompacter(store *ShipmentStore) *ShipmentCompacter {
	return &ShipmentCompacter{store: store, clock: time.Now, retries: 28}
}

func (r *ShipmentCompacter) Compact(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentValidateer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentValidateer(store *ShipmentStore) *ShipmentValidateer {
	return &ShipmentValidateer{store: store, clock: time.Now, retries: 29}
}

func (r *ShipmentValidateer) Validate(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentProjecter struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentProjecter(store *ShipmentStore) *ShipmentProjecter {
	return &ShipmentProjecter{store: store, clock: time.Now, retries: 30}
}

func (r *ShipmentProjecter) Project(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentReconcileer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentReconcileer(store *ShipmentStore) *ShipmentReconcileer {
	return &ShipmentReconcileer{store: store, clock: time.Now, retries: 31}
}

func (r *ShipmentReconcileer) Reconcile(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentEmiter struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentEmiter(store *ShipmentStore) *ShipmentEmiter {
	return &ShipmentEmiter{store: store, clock: time.Now, retries: 32}
}

func (r *ShipmentEmiter) Emit(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentSettleer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentSettleer(store *ShipmentStore) *ShipmentSettleer {
	return &ShipmentSettleer{store: store, clock: time.Now, retries: 33}
}

func (r *ShipmentSettleer) Settle(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentApplyer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentApplyer(store *ShipmentStore) *ShipmentApplyer {
	return &ShipmentApplyer{store: store, clock: time.Now, retries: 34}
}

func (r *ShipmentApplyer) Apply(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}

type ShipmentResolveer struct {
	store   *ShipmentStore
	clock   func() time.Time
	retries int
}

func NewShipmentResolveer(store *ShipmentStore) *ShipmentResolveer {
	return &ShipmentResolveer{store: store, clock: time.Now, retries: 35}
}

func (r *ShipmentResolveer) Resolve(ctx context.Context, id string) (*Shipment, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrShipmentMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrShipmentMissing
}
