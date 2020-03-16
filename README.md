# Vault
Taken from "Go Blueprints"

Simple HTTP and gRPC microservice using Go kit

## Usage

Run the vaultd service:

```
go run cmd/vaultd/main.go
```

Hash a password:

```bash
curl -XPOST -d'{"password":"MySecretPassword123"}' localhost:8080/hash
```

```json
{"hash":"$2a$10$L/Riz9xbgTBDn7F6uLInq.9Tr67PvBCmxzrLgemitnRM53ht7LGpC"}
```

Validate passwords with hashes:

```bash
curl -XPOST -d'{"password":"MySecretPassword123","hash":"$2a$10$L/Riz9xbgTBDn7F6uLInq.9Tr67PvBCmxzrLgemitnRM53ht7LGpC"}' localhost:8080/validate
```

```json
{"valid":true}
```

or if you get the password wrong:

```bash
curl -XPOST -d'{"password":"NOPE","hash":"$2a$10$L/Riz9xbgTBDn7F6uLInq.9Tr67PvBCmxzrLgemitnRM53ht7LGpC"}' localhost:8080/validate
```

```json
{"valid":false}
```