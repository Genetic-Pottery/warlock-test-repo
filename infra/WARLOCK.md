<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# infra

Terraform configuration provisioning AWS infrastructure, currently a single S3 bucket for artifacts.

## Files

- `main.tf` (210 B) — Terraform config defining region variable, aws_s3_bucket.artifacts (bucket "warlock-test-artifacts"), and bucket_name output.

## Structure

- Terraform config defining region variable, aws_s3_bucket.artifacts (bucket "warlock-test-artifacts"), and bucket_name output.
