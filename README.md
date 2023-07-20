# csbcu

(c)-(s)harp (b)inary (c)leaning (u)til is a small convenience CLI tool that can search for c# project folders and delete the obj and bin folders.

This tool was created due to having issues with VS 2022 not always cleaning the bin and obj folders when cleaning solution and doing it manually is tedious.

# Install

```bash
$ go install github.com/michaelbui99/csbcu
```

# Usage

## Clean all projects recusively from current working directory

```bash
$ csbcu list -q | csbcu clean
```

## Clean projects relative to current working directory

```bash
$ csbcu clean ./project1 ./project2 ./nested/project3
```

## Available commands, usage and flags

```bash
$ csbcu -help
```

```bash
$ csbcu <COMMAND> -help
```
