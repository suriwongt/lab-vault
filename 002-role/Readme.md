## Keys Management

1. Open directory
```bash
cd ~/...
```

2. Generate some 16 digits password
```bash
openssl rand -base64 24 | head -c 16; echo
```

3. Create POLICY
```bash
vault policy write admin admin-policy.hcl
vault policy write read read-policy.hcl
```

4. Generate some 16 digits password
```bash
openssl rand -base64 24 | head -c 16; echo
```

5. Create USER
```bash
vault write auth/userpass/users/admin policies="admin" password="<YOUR-PASSWORD>"
```

6. Login 
```bash
vault login -method=userpass username=admin
```

7. List all values in path 
```bash
vault secrets enable -path=<KV_PATH> kv
vault kv list <KV_PATH>
```

8. List all values in path kv/data/customers/customer1/els
```bash
vault kv get <KV_PATH>
```

9. Generate some access token for customer2
```bash
openssl rand -base64 24 | head -c 24; echo
```

10. Set access token to customer2 using generated access token
```bash
vault kv put <KV_PATH> access_token=<YOUR_ACCESS_TOKEN_VALUE>
vault kv put secret/test/example access_token=yX4CGae2mDGxxCjE5kp2Bz3R
```

11. Generate some 16 digits password
```bash
openssl rand -base64 24 | head -c 16; echo
```

12. Create USER Read
```bash
vault write auth/userpass/users/read policies="read" password="<YOUR-PASSWORD>"
```