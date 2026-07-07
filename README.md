# go-vault-search

A command-line tool to search HashiCorp Vault KV v2 secrets by key, value, or path.

## Prerequisites

Before using this tool, you need:
- HashiCorp Vault with KV v2 secrets engine
- Valid Vault credentials set as environment variables:
  ```bash
  export VAULT_ADDR='https://vault.example.com'
  export VAULT_TOKEN='your-vault-token'
  ```

## Installation

### From Source
```bash
go build -o go-vault-search
```

### Using Build Script
```bash
./make_binary_files.sh
```

This creates binaries for multiple platforms:
- `go-vault-search_amd64_linux`
- `go-vault-search_arm64_darwin`

The build script is self-contained: it can be run from any directory and writes
the binaries into the repository directory.

## Usage

```
go-vault-search [flags]

Flags:
  -s string
        what to search (required unless using -l)
  -p string
        Vault KV path to start searching (default "kv/")
  -k    search secret key instead of secret value
  -l    list folders/files in path
  -cat  search folder/file paths instead of secret values
  -plain
        output secret values as plain text instead of JSON-style formatting
  -v    version
```

### Important Notes
- Requires `VAULT_ADDR` and `VAULT_TOKEN` to be set.
- Supports multiple search terms: `-s term1 term2 term3`.
- `-k` searches secret keys. Without `-k`, the app searches secret values.
- `-cat` searches Vault folder/file paths, including the exact path passed with `-p`.
- `-l` lists the contents of a single path and does not require `-s`.

## Examples

### Search by secret value
```bash
❯ go-vault-search -s 124

Found in: /apps/example-api
        third: "qwe124"

Found in: /apps/example-worker
        third: "124"
```

### Search in specific path
```bash
❯ go-vault-search -p kv/apps/example-worker -s 125

Found in: /apps/example-worker/config
        port: "125"
```

### Search by secret key
```bash
❯ go-vault-search -p kv/ -k -s API

Found in: /apps/example-api
        API_TOKEN: "redacted"

Found in: /apps/example-worker
        API_URL: "https://api.example.com"
```

### List folders/files in path
```bash
❯ go-vault-search -p kv/apps -l

in kv/apps was found:
    example-api
    example-worker
```

### List folders/files in default path
```bash
❯ go-vault-search -l

in kv/ was found:
    apps/
    platform/
    shared/
```

### Search folder/file paths
```bash
❯ go-vault-search -p kv/apps/example-api/config -cat -s example-api

kv/apps/example-api/config
```

### Plain text output (no colors)
```bash
❯ go-vault-search -plain -s 124

Found in: /apps/example-api
        third: qwe124

Found in: /apps/example-worker
        third: 124
```

### Multiple search terms
```bash
❯ go-vault-search -s password token secret

# Searches for all three terms across vault secrets
```

### Environment variable example
```bash
VAULT_ADDR=https://vault.example.com VAULT_TOKEN=your-vault-token \
  go-vault-search -p kv/apps/example-api -k -s API_TOKEN
```

## Limitations

- **Works with Vault KV v2 secrets only** - KV v1 is not supported
- Requires valid Vault authentication token
