<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# data

Holds the inventory data assets: a generated JSON inventory dataset, its SQL table schema, and an unrelated binary logo asset.

## Files

- `inventory.json` (1.6 MB) — Large generated JSON inventory dataset (~1.7MB); contents not loaded, structure and fields not inspected.
- `logo.bin` (16.0 KB) — logo.bin: opaque 16KB binary blob, likely a raster/icon asset; not text, no inspectable symbols.
- `schema.sql` (159 B) — Defines the inventory table (id, sku, qty) with a default qty of 0, plus inventory_sku_idx index on sku.

## Structure

- Large generated JSON inventory dataset (~1.7MB); contents not loaded, structure and fields not inspected.
- logo.bin: opaque 16KB binary blob, likely a raster/icon asset; not text, no inspectable symbols.
- Defines the inventory table (id, sku, qty) with a default qty of 0, plus inventory_sku_idx index on sku.
