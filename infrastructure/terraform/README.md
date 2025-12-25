# Terraform Infrastructure

AWS infrastructure provisioning for VAP Platform.

## Resources

- VPC with public/private subnets
- RDS PostgreSQL
- ElastiCache Redis
- S3 buckets for recordings
- EKS Kubernetes cluster
- ALB Ingress Controller

## Usage

```bash
# Initialize
terraform init

# Plan
terraform plan -var-file=environments/production.tfvars

# Apply
terraform apply -var-file=environments/production.tfvars
```
