CREATE TABLE inventory (
  id BIGINT PRIMARY KEY,
  sku TEXT NOT NULL,
  qty INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX inventory_sku_idx ON inventory (sku);
