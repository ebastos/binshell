# binshell

If you have a use case where you want to embed a shell script inside a binary this may be useful for you.

# Requirements

* golang
* [packr](https://github.com/gobuffalo/packr)

# Usage

* Edit main.go and edit the `scriptFolder` and `scriptName` constants
* Run `packr build .`

This will create a binary called *binshell*. Rename it if necessary.

__Important__: Everything inside the folder *scriptFolder* will be embedded inside the binary! Only keep the minimum necessary inside that folder!