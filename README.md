# binshell

If you have a use case where you want to embed a shell script inside a binary this may be useful for you.

# Requirements

* golang
* [packr](https://github.com/gobuffalo/packr)

# Usage

* Edit main.go and edit the `scriptName` constant and the `box := packr.NewBox("./scripts")` line
* Run `packr`
* Run `packr build .`

This will create a binary called *binshell*. Rename it if necessary.

__Important__: Everything inside the folder *scriptFolder* will be embedded inside the binary! Only keep the minimum necessary inside that folder!

# Example:
```
$ ls /tmp/scripts/
hello.sh
$ cat /tmp/scripts/hello.sh
#!/bin/bash

echo "Hello World! I got $@"


$ packr build .
$ ./binshell aaa b cde -v 123
Hello World! I got aaa b cde -v 123
```