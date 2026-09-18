<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# infra

Terraform infrastructure configuration for the project, declaring cloud resources such as the artifacts S3 bucket and the region they deploy into.

## Files

- `main.tf` (210 B) — Terraform config declaring region variable (default eu-west-1), aws_s3_bucket.artifacts (bucket warlock-test-artifacts), and bucket_name output.

## Structure

- Terraform config declaring region variable (default eu-west-1), aws_s3_bucket.artifacts (bucket warlock-test-artifacts), and bucket_name output.
