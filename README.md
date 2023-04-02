# **working with only v2 secrets at this moment**
# Do a search through vault secrets by a search string
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
```
❯ ./go-vault-get -s 124
kv/TEST - third = 124
kv/TEST2 - third = 124
kv/TEST23 - third = 124
kv/localtest - localvar = local124
duration  166ns
```
&
```
❯ ./go-vault-get -s 122 -p kv/TEST2
kv/TEST2/tt - foru = 122
duration  167ns
```
&
```
❯ ./go-vault-get -s qwe -p kv/TEST2 -k
kv/TEST2/qwe
duration  375ns
```
&
```
❯ go-vault-search -l -p kv/TEST2

kv/TEST2/tt

found 1 keys


❯ go-vault-search -l

kv/TEST
kv/TEST2
kv/TEST2/
kv/TEST23
kv/TEST3/
```