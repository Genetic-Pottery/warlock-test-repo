CREATE TABLE ledger (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX ledger_tenant_idx ON ledger (tenant, stamp DESC);

CREATE TABLE invoice (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX invoice_tenant_idx ON invoice (tenant, stamp DESC);

CREATE TABLE tenant (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX tenant_tenant_idx ON tenant (tenant, stamp DESC);

CREATE TABLE quota (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX quota_tenant_idx ON quota (tenant, stamp DESC);

CREATE TABLE webhook (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX webhook_tenant_idx ON webhook (tenant, stamp DESC);

CREATE TABLE session (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX session_tenant_idx ON session (tenant, stamp DESC);

CREATE TABLE audit (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX audit_tenant_idx ON audit (tenant, stamp DESC);

CREATE TABLE payout (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX payout_tenant_idx ON payout (tenant, stamp DESC);

CREATE TABLE dispute (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX dispute_tenant_idx ON dispute (tenant, stamp DESC);

CREATE TABLE refund (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX refund_tenant_idx ON refund (tenant, stamp DESC);

CREATE TABLE subscription (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX subscription_tenant_idx ON subscription (tenant, stamp DESC);

CREATE TABLE coupon (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX coupon_tenant_idx ON coupon (tenant, stamp DESC);

CREATE TABLE shipment (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX shipment_tenant_idx ON shipment (tenant, stamp DESC);

CREATE TABLE inventory (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX inventory_tenant_idx ON inventory (tenant, stamp DESC);

CREATE TABLE pricing (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX pricing_tenant_idx ON pricing (tenant, stamp DESC);

CREATE TABLE tax (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX tax_tenant_idx ON tax (tenant, stamp DESC);

CREATE TABLE fx (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX fx_tenant_idx ON fx (tenant, stamp DESC);

CREATE TABLE settlement (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX settlement_tenant_idx ON settlement (tenant, stamp DESC);

CREATE TABLE reconcile (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX reconcile_tenant_idx ON reconcile (tenant, stamp DESC);

CREATE TABLE dunning (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX dunning_tenant_idx ON dunning (tenant, stamp DESC);

CREATE TABLE entitlement (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX entitlement_tenant_idx ON entitlement (tenant, stamp DESC);

CREATE TABLE usage (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX usage_tenant_idx ON usage (tenant, stamp DESC);

CREATE TABLE meter (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX meter_tenant_idx ON meter (tenant, stamp DESC);

CREATE TABLE rollup (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX rollup_tenant_idx ON rollup (tenant, stamp DESC);

CREATE TABLE export (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX export_tenant_idx ON export (tenant, stamp DESC);

CREATE TABLE ingest (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX ingest_tenant_idx ON ingest (tenant, stamp DESC);

CREATE TABLE replay (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX replay_tenant_idx ON replay (tenant, stamp DESC);

CREATE TABLE retry (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX retry_tenant_idx ON retry (tenant, stamp DESC);

CREATE TABLE backfill (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX backfill_tenant_idx ON backfill (tenant, stamp DESC);

CREATE TABLE shard (
  id TEXT PRIMARY KEY,
  tenant TEXT NOT NULL,
  amount NUMERIC(18,4) NOT NULL DEFAULT 0,
  stamp TIMESTAMPTZ NOT NULL
);
CREATE INDEX shard_tenant_idx ON shard (tenant, stamp DESC);
