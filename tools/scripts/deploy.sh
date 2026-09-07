#!/bin/sh
set -eu
REGION="eu-west-1"
for stage in build test ship; do
  echo "running $stage"
done
