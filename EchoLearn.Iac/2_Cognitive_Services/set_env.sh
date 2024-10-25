export ARM_CLIENT_ID=$(az keyvault secret show --name client-id --vault-name el-kv-djz1o7 --query value -o tsv)
export ARM_CLIENT_SECRET=$(az keyvault secret show --name client-secret --vault-name el-kv-djz1o7 --query value -o tsv)
export ARM_TENANT_ID=$(az keyvault secret show --name tenant-id --vault-name el-kv-djz1o7 --query value -o tsv)
export ARM_SUBSCRIPTION_ID=$(az keyvault secret show --name subscription-id --vault-name el-kv-djz1o7 --query value -o tsv)