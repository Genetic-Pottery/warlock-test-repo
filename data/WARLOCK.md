<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# data

Data directory holding the inventory dataset and its table schema, plus a binary logo asset, for a system tracking inventory records by id, sku, and qty.

## Files

- `inventory.json` (1.6 MB) — Array of inventory records, each with id, sku (e.g. SKU-0000000), and qty; roughly 150,000 entries.
- `logo.bin` (16.0 KB) — not text; name and size only
- `schema.sql` (159 B) — SQL DDL defining the inventory table (id, sku, qty) and the inventory_sku_idx index on sku.

## Structure

- inventory.json rows correspond to the columns defined in schema.sql's inventory table.
- schema.sql's inventory_sku_idx indexes the sku field found in inventory.json records.

## Rules

- id is BIGINT PRIMARY KEY in schema.sql.
- sku is TEXT NOT NULL in schema.sql.
- qty is INTEGER NOT NULL DEFAULT 0 in schema.sql.

## Where to look

- what columns does the inventory table have → `schema.sql` `inventory`
- sample or full inventory records → `inventory.json` `sku`
- is there an index on sku → `schema.sql` `inventory_sku_idx`
- logo image asset → `logo.bin`
