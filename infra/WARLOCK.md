<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# infra

Terraform infrastructure configuration provisioning the S3 bucket used to store the project's build artifacts.

## Files

- `main.tf` (210 B) — Terraform config declaring the region variable, the aws_s3_bucket resource artifacts (bucket warlock-test-artifacts), and output bucket_name.

## Structure

- output bucket_name reads its value from aws_s3_bucket.artifacts.bucket
- aws_s3_bucket.artifacts is configured independently of the region variable

## Rules

- variable region defaults to eu-west-1
- the artifacts bucket name is fixed as warlock-test-artifacts

## Where to look

- where is the S3 bucket for build artifacts defined → `main.tf` `aws_s3_bucket.artifacts`
- what AWS region does infra deploy to → `main.tf` `region`
- what output exposes the bucket name → `main.tf` `bucket_name`
