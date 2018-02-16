# binshell

If you have a use case where you want to embed a shell script inside a binary this may be useful for you.

# Requirements

* golang
* [packr](https://github.com/gobuffalo/packr)

# Usage

* Add desired script to the scripts folder (has to be this folder)
* Run `./generate.sh scripts/script_name.sh`

This will create a binary called *binshell*. Rename it if necessary.

__Important__: Everything inside the folder *scripts/* will be embedded inside the binary! Only keep the minimum necessary inside that folder!

# Example:
```
$ cat scripts/sample.sh
#!/bin/bash

echo "Hello from Shell Script! I got $@"
$ ./generate.sh scripts/sample.sh
$ ./binshell some options
Hello from Shell Script! I got some options
```
