# **working with only v2 secrets at this moment**
# Do a search through vault secrets by a search string
## keys "-k" or "-l" need to be place in begin
```
Usage of go-vault-search:
  -k	search secret key instead secret value
  -l	show only list of vaults in path
  -p string
    	path to vault secret start searching (default "kv/")
  -s string
    	what to search
  -v	version
```

## example
search by secret value "-s"
```
❯ go-vault-search -s 124

kv/TEST - third = qwe124
kv/TEST2 - third = 124
kv/TEST23 - third = 124

found 3
```
search by secret value "-s" in "-p" path
```
❯ go-vault-search -s 125 -p kv/TEST2

kv/TEST2/tt - foru = 125

found 1
```
search by secret key "-k -s" in "-p" path
```
❯ go-vault-search -k -s for -p kv/

kv/TEST2/tt - foru = 125
kv/TEST3/t2/qwe - for = 122
kv/TEST3/tt/qwe - foru = 122

found 3
```
list folders/files in "-p" path
```
❯ go-vault-search -l -p kv/TEST2

kv/TEST2/tt

found 1

```
list folders/files in default path "kv/"
defined in main.go "vaultPath := flag.String("p", "kv/", "path to vault secret start searching")"
```
❯ go-vault-search -l

found dirs in kv/:
	TEST
	TEST2
	TEST2/
	TEST23
	TEST3/

found 5
```