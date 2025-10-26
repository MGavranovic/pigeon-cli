# pigeon-cli

A minimal yet extensible custom CLI tool inspired by Linux-style commands, built in Go.

## Features

- Linux-inspired command structure (`ls`, `cd`, `cat`, `touch`, `exit`, etc.)
- Colorized output with [`fatih/color`](https://github.com/fatih/color)
- File operations, tree listing, and word/byte counting
- History tracking with status reporting

## Installation

Clone and build:

```bash
git clone https://github.com/MGavranovic/pigeon-cli.git
cd pigeon-cli
go build -o pigeon-cli ./src
./pigeon-cli
```

## Project Structure

```bash
pigeon-cli/
├── internal/
│   └── cmd/              # Command implementations
│       ├── ls.go
│       ├── cd.go
│       ├── ...
│   └── autocomplete/     # Autocomplete pkg
│   └── autocomplete/     # Autocomplete pkg
│   └── inputpkg/         # Terminal drawing
│   └── helpers/          # Helpers pkg
├── main.go               # Main entry point
├── go.mod
└── README.md
```

## Roadmap

Below are planned or in-progress features for **pigeon-cli**.

### Completed

- [x] `ls`
      Lists all files and directories in the current working directory.
- [x] `cd`
      Changes the current working directory.
- [x] `cl`
      Clears the terminal screen.
- [x] `cp`
      Copies files from one location to another.
- [x] `cat`
      Displays the contents of a file.
- [x] `touch`
      Creates an empty file or updates the timestamp of an existing file.
- [x] `wc`
      Prints the byte size and word count of a file.
- [x] `tree`
      Recursively prints a tree structure of files and folders.
- [x] `help`
      Displays a list of available commands and their descriptions.
- [x] `history`
      Shows a list of previously run commands with success status.
- [x] `mv`
      Moves a file to a new location.
- [x] `rm`
      Deletes the specified file.
- [x] `rn`
      Renames a file.
- [x] `grep`
      Searches for a string in a file and highlights the matches.
- [x] Autocomplete (TAB-suggestions + arrow key navigation)

### In Progress
- [ ] Zip/upload utility
- [ ] Revamping args logic

### Planned

- [ ] Command aliases
- [ ] Config file for themes/settings/shortcuts
- [ ] File search
- [ ] Cross-platform terminal support (Windows/macOS/Linux)
- [ ] Setup a pipleine for running the build process for the app.

## Contributions

PRs, issues, and suggestions welcome.

## License

[MIT](./LICENSE) License. Built by Milos Gavranovic.

