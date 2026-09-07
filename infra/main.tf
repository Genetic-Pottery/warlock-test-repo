variable "region" {
  type    = string
  default = "eu-west-1"
}

resource "aws_s3_bucket" "artifacts" {
  bucket = "warlock-test-artifacts"
}

output "bucket_name" {
  value = aws_s3_bucket.artifacts.bucket
}
