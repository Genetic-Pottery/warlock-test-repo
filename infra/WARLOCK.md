<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# infra

Terraform infrastructure configuration for the project, provisioning the AWS S3 bucket used for build artifacts.

## Files

- `main.tf` (210 B) — Defines the region variable (default eu-west-1), the aws_s3_bucket resource artifacts (bucket warlock-test-artifacts), and the bucket_name output.

## Structure

- bucket_name output derives from aws_s3_bucket.artifacts.bucket
- aws_s3_bucket.artifacts is created in the region set by the region variable

## Rules

- region variable defaults to eu-west-1
- artifacts bucket name is fixed to warlock-test-artifacts

## Where to look

- what AWS resources does this project provision → `main.tf` `aws_s3_bucket`
- where is the artifacts bucket name set → `main.tf` `warlock-test-artifacts`
- what region is infrastructure deployed to → `main.tf` `region`
