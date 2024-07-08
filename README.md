# Universal Shell Config (USC)

A program to generate `.zshrc`, `.bashrc`, and `Profile.ps1` shell configs from a PKL config `<file>`

This program is currently in a usable, but very messy state and I plan to "improve" the pkl Config format in the future with breaking changes.

## Usage

```sh
$ usc.sh -i <file> [options]
>
```

Options

- `-o=<path>` Output directory that the script will use
- `--install` Replaces the system `.bashrc`, `.zshrc` and `Profile.ps1` configs with the generated configs

## How to run tests

```sh
$ go test ./...
>
```
