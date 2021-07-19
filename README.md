# prepend

## Motivation

There is no single shell builtin to prepend stuff to files.
Instead you have to resort to sed and awk trickery to make it work.

## Installation

```
make install
```

## Usage

```shell
echo "This will be prepended to the file" | prepend filename
```

## License

This project is licensed under the GPL-3 license.
