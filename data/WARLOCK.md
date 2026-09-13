<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# data

Data directory holding the inventory dataset and its database schema, backing an inventory tracking system keyed by id and sku.

## Files

- `inventory.json` (1.6 MB) — JSON array of inventory records, each with id, sku (SKU-xxxxxxx), and qty, ~150k entries.
- `logo.bin` (16.0 KB) — not text; name and size only
- `schema.sql` (159 B) — SQL DDL defining the inventory table (id, sku, qty) and the inventory_sku_idx index on sku.

## Structure

- inventory.json's records correspond to rows of the inventory table defined in schema.sql.
- inventory_sku_idx indexes the sku column that inventory.json entries populate.

## Rules

- id is declared PRIMARY KEY in the inventory table.
- sku is NOT NULL in the inventory table.
- qty is INTEGER NOT NULL DEFAULT 0 in the inventory table.

## Where to look

- what fields does an inventory record have → `schema.sql` `inventory`
- sample or bulk inventory data → `inventory.json` `sku`
- how sku lookups are indexed → `schema.sql` `inventory_sku_idx`
