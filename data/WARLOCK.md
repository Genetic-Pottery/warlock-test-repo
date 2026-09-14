<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# data

Holds the inventory data assets: the large JSON inventory dataset, a binary logo asset, and the SQL schema defining the inventory table and its index.

## Files

- `inventory.json` (1.6 MB) — Large JSON data file (1.7MB) holding inventory records; contents not loaded, consult directly for schema and entries.
- `logo.bin` (16.0 KB) — Binary logo asset (16384 bytes), not text; opened only when the raw image data itself is needed.
- `schema.sql` (159 B) — SQL schema defining the inventory table (id, sku, qty) and inventory_sku_idx index on sku.

## Structure

- Large JSON data file (1.7MB) holding inventory records; contents not loaded, consult directly for schema and entries.
- Binary logo asset (16384 bytes), not text; opened only when the raw image data itself is needed.
- SQL schema defining the inventory table (id, sku, qty) and inventory_sku_idx index on sku.
