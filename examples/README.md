# Examples

This directory contains examples that are mostly used for documentation, but can also be run/tested manually via the Terraform CLI.

## Run the example
From inside this directory:

```bash
terraform init
terraform plan -out theplan
terraform apply theplan
```

## Display the resource from the state

```bash
terraform state show sample_sample.example
```

## Remove the example

```bash
terraform destroy
```
