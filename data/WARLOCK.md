<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# data

Data directory holding the inventory database schema and its associated generated data and binary asset files.

## Files

- `inventory.json` (1.6 MB) — Large generated JSON data file (~1.7MB) holding inventory records; not meant to be read directly, consult for raw inventory data only.
- `logo.bin` (16.0 KB) — logo.bin: binary asset (16384 bytes), likely a logo image resource, not source or documentation text.
- `schema.sql` (159 B) — SQL DDL defining the inventory table (id, sku, qty) and the inventory_sku_idx index on sku.

## Structure

- Large generated JSON data file (~1.7MB) holding inventory records; not meant to be read directly, consult for raw inventory data only.
- logo.bin: binary asset (16384 bytes), likely a logo image resource, not source or documentation text.
- SQL DDL defining the inventory table (id, sku, qty) and the inventory_sku_idx index on sku.
