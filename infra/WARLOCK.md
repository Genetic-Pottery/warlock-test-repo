<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# infra

Terraform infrastructure definition provisioning the S3 bucket used to store build artifacts for this project.

## Files

- `main.tf` (210 B) — Defines variable region (default eu-west-1), resource aws_s3_bucket.artifacts (bucket warlock-test-artifacts), and output bucket_name.

## Structure

- output bucket_name reads from aws_s3_bucket.artifacts.bucket, which is created in the same file

## Rules

- region defaults to eu-west-1
- bucket name is fixed to warlock-test-artifacts

## Where to look

- what AWS region is used → `main.tf` `region`
- name of the S3 bucket for artifacts → `main.tf` `aws_s3_bucket`
- terraform output values → `main.tf` `bucket_name`
