# hdkey

hdkey derives arbitrary Decred hierarchical deterministic (HD) keys
from a provided public or private key.  Keys in hexadecimal format are
assumed to be master private keys.

## Examples

Derive simnet keys with path `m/44'/1'/0'/0/7` from a hex seed:

```
$ hdkey -net=sim -key b280922d2cffda44648346412c5ec97f429938105003730414f10b01e1402eac -path "m/44'/1'/0'/0/7"
xpriv: sprvZLzf9mDaUVqRxa6Ktn56xUZJqhLjLUFxwFMocYrLfqa1FNRHvLTCt7obcAVDoKSUBbhGgsDY4Jm6zX5SawxdT1tMtToDLi6VKZNPU9vjHst
xpub:  spubVZz1ZGkUJsPjB4Anzoc7KcW3PjBDjvypJUHQQwFxEB6z8AkSTsmTRv85TRyBqWfgTEq8vxZ8gY5P9LPjKXs5bdJB5UcVyXRe2Ae84hViwz3
addr:  Sso52TPnorVkSaRYzHmi4FgU8F5BFEDZsiK
```

Derive simnet SLIP0044 coin type keys from a hex seed:

```
$ ./hdkey -net=sim -key b280922d2cffda44648346412c5ec97f429938105003730414f10b01e1402eac -path "m/44'/1'/0'"
xpriv: sprvZHKvNdvPAcEp4uVjKCAL6y3BdexiVa4T8HDF7hAZZCSs1FmiXz6wpLJHqBjrD3xum3vuDHV4L4wKs11Zt7KU4GKrYRE78SvqQib5wQfgP7h
xpub:  spubVWKGn9TGzyo7HPaCRDhLU6yvBgoCu2nJVW8qv5aB7Xyqt46s5XRCN8cmgW2RApGdFyJgaL2iiGFRZLuL2KgXN3zpTG9CYE4dWpVmKnjdSwJ
addr:  SsXKWc2DZkz9hcY7xEPiZ6QdjSHts2DehEw
```

Derive relative simnet path `0/7` using second example's coin type
xpub (result is same key as from first example):

```
$ hdkey -net=sim -key spubVWKGn9TGzyo7HPaCRDhLU6yvBgoCu2nJVW8qv5aB7Xyqt46s5XRCN8cmgW2RApGdFyJgaL2iiGFRZLuL2KgXN3zpTG9CYE4dWpVmKnjdSwJ -path "0/7"
xpub:  spubVZz1ZGkUJsPjB4Anzoc7KcW3PjBDjvypJUHQQwFxEB6z8AkSTsmTRv85TRyBqWfgTEq8vxZ8gY5P9LPjKXs5bdJB5UcVyXRe2Ae84hViwz3
addr:  Sso52TPnorVkSaRYzHmi4FgU8F5BFEDZsiK
```

## Security notice

hdkey is designed to be a developer's tool and private key security is
not a primary focus.  In particular, private keys are passed as
process arguments and may be visable to all users on the machine.
