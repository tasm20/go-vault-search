# go-vault-search

A command-line tool to search through HashiCorp Vault KV v2 secrets by key or value.

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

## Usage

```
go-vault-search [flags]

Flags:
  -s string
        what to search (required unless using -l)
  -p string
        path to vault secret start searching (default "kv/")
  -k    search secret key instead of secret value
  -l    show only list of vaults in path
  -cat  search folder or file
  -plain
        output in plain text format (not json as default output)
  -v    version
```

### Important Notes
- The `-k` and `-l` flags should be placed at the beginning
- The `-s` flag (search item) should be at the end when using multiple search terms
- Supports multiple search terms: `-s term1 term2 term3`

## Examples

### Search by secret value
```bash
❯ go-vault-search -s 124

kv/TEST - third = qwe124
kv/TEST2 - third = 124
kv/TEST23 - third = 124

found 3
```

### Search in specific path
```bash
❯ go-vault-search -p kv/TEST2 -s 125

kv/TEST2/tt - foru = 125

found 1
```

### Search by secret key
```bash
❯ go-vault-search -p kv/ -k -s fo

kv/TEST2/tt - foru = 125
kv/TEST3/t2/qwe - for = 122
kv/TEST3/tt/qwe - foru = 122

found 3
```

### List folders/files in path
```bash
❯ go-vault-search -p kv/TEST2 -l

in kv/TEST2 was found:
    tt

found 1
```

### List folders/files in default path
```bash
❯ go-vault-search -l

found dirs in kv/:
    TEST
    TEST2
    TEST2/
    TEST23
    TEST3/

found 5
```

### Plain text output (no colors)
```bash
❯ go-vault-search -plain -s 124 

kv/TEST - third = qwe124
kv/TEST2 - third = 124
kv/TEST23 - third = 124

found 3
```

### Multiple search terms
```bash
❯ go-vault-search -s password token secret

# Searches for all three terms across vault secrets
```

## Limitations

- **Works with Vault KV v2 secrets only** - KV v1 is not supported
- Requires valid Vault authentication token
