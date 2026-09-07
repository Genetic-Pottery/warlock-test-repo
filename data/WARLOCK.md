<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# data

Data directory holding the inventory dataset and its database schema: a large JSON inventory dump and the SQL table definition it loads into.

## Files

- `inventory.json` (1.6 MB) — Array of inventory records, each with id, sku (e.g. SKU-0000000), and qty; ~150,000 entries.
- `logo.bin` (16.0 KB) — not text; name and size only
- `schema.sql` (159 B) — Defines the inventory table (id BIGINT PRIMARY KEY, sku TEXT NOT NULL, qty INTEGER NOT NULL DEFAULT 0) and the inventory_sku_idx index on sku.

## Structure

- inventory.json's records map field-for-field onto schema.sql's inventory table columns (id, sku, qty).

## Rules

- inventory.id is the primary key (BIGINT PRIMARY KEY).
- inventory.sku is NOT NULL and indexed via inventory_sku_idx.
- inventory.qty is NOT NULL DEFAULT 0.

## Where to look

- what columns does the inventory table have → `schema.sql` `inventory`
- sample or full inventory records → `inventory.json` `sku`
- is sku indexed for lookups → `schema.sql` `inventory_sku_idx`
