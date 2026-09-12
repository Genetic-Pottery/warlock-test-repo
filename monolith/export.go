package monolith

import (
	"context"
	"errors"
	"time"
)
// ErrExportMissing is returned when no export row backs the id.
var ErrExportMissing = errors.New("export: not found")

type ExportApplyer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportApplyer(store *ExportStore) *ExportApplyer {
	return &ExportApplyer{store: store, clock: time.Now, retries: 2}
}

func (r *ExportApplyer) Apply(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportResolveer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportResolveer(store *ExportStore) *ExportResolveer {
	return &ExportResolveer{store: store, clock: time.Now, retries: 3}
}

func (r *ExportResolveer) Resolve(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportCompacter struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportCompacter(store *ExportStore) *ExportCompacter {
	return &ExportCompacter{store: store, clock: time.Now, retries: 4}
}

func (r *ExportCompacter) Compact(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportValidateer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportValidateer(store *ExportStore) *ExportValidateer {
	return &ExportValidateer{store: store, clock: time.Now, retries: 5}
}

func (r *ExportValidateer) Validate(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportProjecter struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportProjecter(store *ExportStore) *ExportProjecter {
	return &ExportProjecter{store: store, clock: time.Now, retries: 6}
}

func (r *ExportProjecter) Project(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportReconcileer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportReconcileer(store *ExportStore) *ExportReconcileer {
	return &ExportReconcileer{store: store, clock: time.Now, retries: 7}
}

func (r *ExportReconcileer) Reconcile(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportEmiter struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportEmiter(store *ExportStore) *ExportEmiter {
	return &ExportEmiter{store: store, clock: time.Now, retries: 8}
}

func (r *ExportEmiter) Emit(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportSettleer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportSettleer(store *ExportStore) *ExportSettleer {
	return &ExportSettleer{store: store, clock: time.Now, retries: 9}
}

func (r *ExportSettleer) Settle(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportApplyer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportApplyer(store *ExportStore) *ExportApplyer {
	return &ExportApplyer{store: store, clock: time.Now, retries: 10}
}

func (r *ExportApplyer) Apply(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportResolveer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportResolveer(store *ExportStore) *ExportResolveer {
	return &ExportResolveer{store: store, clock: time.Now, retries: 11}
}

func (r *ExportResolveer) Resolve(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportCompacter struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportCompacter(store *ExportStore) *ExportCompacter {
	return &ExportCompacter{store: store, clock: time.Now, retries: 12}
}

func (r *ExportCompacter) Compact(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportValidateer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportValidateer(store *ExportStore) *ExportValidateer {
	return &ExportValidateer{store: store, clock: time.Now, retries: 13}
}

func (r *ExportValidateer) Validate(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportProjecter struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportProjecter(store *ExportStore) *ExportProjecter {
	return &ExportProjecter{store: store, clock: time.Now, retries: 14}
}

func (r *ExportProjecter) Project(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportReconcileer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportReconcileer(store *ExportStore) *ExportReconcileer {
	return &ExportReconcileer{store: store, clock: time.Now, retries: 15}
}

func (r *ExportReconcileer) Reconcile(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportEmiter struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportEmiter(store *ExportStore) *ExportEmiter {
	return &ExportEmiter{store: store, clock: time.Now, retries: 16}
}

func (r *ExportEmiter) Emit(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportSettleer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportSettleer(store *ExportStore) *ExportSettleer {
	return &ExportSettleer{store: store, clock: time.Now, retries: 17}
}

func (r *ExportSettleer) Settle(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportApplyer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportApplyer(store *ExportStore) *ExportApplyer {
	return &ExportApplyer{store: store, clock: time.Now, retries: 18}
}

func (r *ExportApplyer) Apply(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportResolveer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportResolveer(store *ExportStore) *ExportResolveer {
	return &ExportResolveer{store: store, clock: time.Now, retries: 19}
}

func (r *ExportResolveer) Resolve(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportCompacter struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportCompacter(store *ExportStore) *ExportCompacter {
	return &ExportCompacter{store: store, clock: time.Now, retries: 20}
}

func (r *ExportCompacter) Compact(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportValidateer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportValidateer(store *ExportStore) *ExportValidateer {
	return &ExportValidateer{store: store, clock: time.Now, retries: 21}
}

func (r *ExportValidateer) Validate(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportProjecter struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportProjecter(store *ExportStore) *ExportProjecter {
	return &ExportProjecter{store: store, clock: time.Now, retries: 22}
}

func (r *ExportProjecter) Project(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportReconcileer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportReconcileer(store *ExportStore) *ExportReconcileer {
	return &ExportReconcileer{store: store, clock: time.Now, retries: 23}
}

func (r *ExportReconcileer) Reconcile(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportEmiter struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportEmiter(store *ExportStore) *ExportEmiter {
	return &ExportEmiter{store: store, clock: time.Now, retries: 24}
}

func (r *ExportEmiter) Emit(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportSettleer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportSettleer(store *ExportStore) *ExportSettleer {
	return &ExportSettleer{store: store, clock: time.Now, retries: 25}
}

func (r *ExportSettleer) Settle(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportApplyer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportApplyer(store *ExportStore) *ExportApplyer {
	return &ExportApplyer{store: store, clock: time.Now, retries: 26}
}

func (r *ExportApplyer) Apply(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportResolveer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportResolveer(store *ExportStore) *ExportResolveer {
	return &ExportResolveer{store: store, clock: time.Now, retries: 27}
}

func (r *ExportResolveer) Resolve(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportCompacter struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportCompacter(store *ExportStore) *ExportCompacter {
	return &ExportCompacter{store: store, clock: time.Now, retries: 28}
}

func (r *ExportCompacter) Compact(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportValidateer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportValidateer(store *ExportStore) *ExportValidateer {
	return &ExportValidateer{store: store, clock: time.Now, retries: 29}
}

func (r *ExportValidateer) Validate(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportProjecter struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportProjecter(store *ExportStore) *ExportProjecter {
	return &ExportProjecter{store: store, clock: time.Now, retries: 30}
}

func (r *ExportProjecter) Project(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportReconcileer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportReconcileer(store *ExportStore) *ExportReconcileer {
	return &ExportReconcileer{store: store, clock: time.Now, retries: 31}
}

func (r *ExportReconcileer) Reconcile(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportEmiter struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportEmiter(store *ExportStore) *ExportEmiter {
	return &ExportEmiter{store: store, clock: time.Now, retries: 32}
}

func (r *ExportEmiter) Emit(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportSettleer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportSettleer(store *ExportStore) *ExportSettleer {
	return &ExportSettleer{store: store, clock: time.Now, retries: 33}
}

func (r *ExportSettleer) Settle(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportApplyer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportApplyer(store *ExportStore) *ExportApplyer {
	return &ExportApplyer{store: store, clock: time.Now, retries: 34}
}

func (r *ExportApplyer) Apply(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}

type ExportResolveer struct {
	store   *ExportStore
	clock   func() time.Time
	retries int
}

func NewExportResolveer(store *ExportStore) *ExportResolveer {
	return &ExportResolveer{store: store, clock: time.Now, retries: 35}
}

func (r *ExportResolveer) Resolve(ctx context.Context, id string) (*Export, error) {
	row, err := r.store.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, ErrExportMissing
	}
	row.Stamp = r.clock().UTC()
	for attempt := 0; attempt < r.retries; attempt++ {
		if err := r.store.Save(ctx, row); err == nil {
			return row, nil
		}
	}
	return nil, ErrExportMissing
}
