#!/usr/bin/env bash
#
# Usage:
#   ./deploy.sh <LOCATION> <PROJECT_NAME>
#
# Example:
#   ./deploy.sh eastus myproject
#
# This script will deploy resources to 'stg' and 'prd' environments.

# Exit immediately if a command exits with a non-zero status
set -e

LOCATION="$1"
PROJECT_NAME="$2"

if [[ -z "$LOCATION" || -z "$PROJECT_NAME" ]]; then
  echo "Usage: $0 <LOCATION> <PROJECT_NAME>"
  echo "Example: $0 eastus myproject"
  exit 1
fi

# Define the environments you want to deploy to
ENVIRONMENTS=("stg" "prd")

for ENVIRONMENT in "${ENVIRONMENTS[@]}"; do

  echo "=========================================================="
  echo "Starting deployment for environment: ${ENVIRONMENT}"
  echo "=========================================================="

  # Create resource group name using the naming convention
  RG_NAME="rg-${LOCATION}-${PROJECT_NAME}-${ENVIRONMENT}"
  # Create storage account name using the naming convention
  STORAGE_ACCOUNT_NAME="satf${LOCATION}${PROJECT_NAME}${ENVIRONMENT}"
  # Container name
  CONTAINER_NAME="tfstate-init"
  # Tags (add more key-value pairs if you want)
  TAGS="Owner=YourName Project=${PROJECT_NAME} Environment=${ENVIRONMENT}"

  echo "---------------------------------------------------"
  echo "Creating Resource Group: ${RG_NAME}"
  echo "---------------------------------------------------"
  az group create \
    --name "${RG_NAME}" \
    --location "${LOCATION}" \
    --tags ${TAGS}

  echo "---------------------------------------------------"
  echo "Creating Storage Account: ${STORAGE_ACCOUNT_NAME}"
  echo "---------------------------------------------------"
  az storage account create \
    --name "${STORAGE_ACCOUNT_NAME}" \
    --resource-group "${RG_NAME}" \
    --location "${LOCATION}" \
    --sku Standard_ZRS \
    --kind StorageV2 \
    --allow-blob-public-access false \
    --min-tls-version TLS1_2 \
    --tags ${TAGS}

  echo "---------------------------------------------------"
  echo "Enabling Versioning on the Storage Account"
  echo "---------------------------------------------------"
  az storage account blob-service-properties update \
    --account-name "${STORAGE_ACCOUNT_NAME}" \
    --resource-group "${RG_NAME}" \
    --enable-versioning true

  echo "---------------------------------------------------"
  echo "Creating Container: ${CONTAINER_NAME}"
  echo "---------------------------------------------------"
  az storage container create \
    --name "${CONTAINER_NAME}" \
    --account-name "${STORAGE_ACCOUNT_NAME}" \
    --auth-mode login \
    --public-access off

  echo "=========================================================="
  echo "Deployment for environment ${ENVIRONMENT} completed!"
  echo "Resource Group: ${RG_NAME}"
  echo "Storage Account: ${STORAGE_ACCOUNT_NAME}"
  echo "Container: ${CONTAINER_NAME}"
  echo "=========================================================="
  echo ""
done

echo "All deployments are finished!"
