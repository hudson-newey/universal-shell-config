# Universal Shell Config (USC)

A program to generate `.zshrc`, `.bashrc`, and `Profile.ps1` shell configs from a PKL config `<file>`

## Usage

```sh
$ usc.sh -i <file> [options]
>
```

Options

- `-o=<path>` Output directory that the script will use
- `--install` Replaces the system `.bashrc`, `.zshrc` and `Profile.ps1` configs with the generated configs
