# ch

ch (short for check hash/create hash) is a tool to create file/string hashes and also check them.

Usage:
```
ch <mode> [arguments...] <algorithm>
```

Modes:
```
check   - check a file against a given hash
checks  - check a string against a given hash
create  - create a hash from a file
creates - create a hash from a string
```

Supported algorithms:
```
MD5
SHA1
SHA224
SHA256
SHA384
SHA512
```

Example:

Creating a hash:
```
ch creates Foo md5
```
```
MD5 of "Foo": 1356c67d7ad1638d816bfb822dd2c25d
```

Checking a hash:
```
ch checks Foo 1356c67d7ad1638d816bfb822dd2c25d md5
```
```
Ok!
```

```
ch checks NotFoo 1356c67d7ad1638d816bfb822dd2c25d md5
```
```
Not ok!
```
