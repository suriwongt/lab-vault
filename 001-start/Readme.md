1. Start
```
    docker-compose up -d
    docker-compose down --rmi -all
```

2. Unseal Vault
```bash
vault operator unseal
```

3. Login
```bash
vault login
```

**Use root token**

4. Enable the userpass authentication method
```bash
vault auth enable userpass
```

5. List auth methods
```bash
vault auth list
```

6. Enable kv secret engine
```bash
vault secrets enable -path=<new-path> kv
```
7. enable version
```bash
vault kv enable-versioning <new-path> 
```

7. List secret engine
```bash
vault secrets list
```

