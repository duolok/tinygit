# TinyGit

Small version control system written in Go, inspired by Git.

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/tinygit.git

# Build the project
cd tinygit
go build -o tinygit

# Make it executable (for Unix-based systems)
chmod +x tinygit
```

## Basic Commands

| Command | Description |
|---------|-------------|
| `--config {NAME}` | Set the author name for commits |
| `--add {FILE_NAME}` | Track a file |
| `--commit {MESSAGE}` | Save changes to tracked files with a commit message |
| `--log` | Display all commits |
| `--checkout {COMMIT_HASH}` | Switch to a specific commit |
| `--show-commit {COMMIT_HASH}` | Display metadata for a specific commit |
| `--tracked-files` | Show all currently tracked files |
| `--help` | Display help information |

## Command Details

### Configure Author

```bash
./tinygit --config {NAME}
```

Sets the author name for your commits.

### Track Files

```bash
./tinygit --add {FILE_NAME}
```

Adds a file to the list of tracked files.

### Commit Changes

```bash
./tinygit --commit "Your commit message"
```

Creates a new commit with the current state of all tracked files. The commit includes the provided message and the configured author name.

### View Commit History

```bash
./tinygit --log
```

Displays a list of all commits, including their hash, author, date, and message.

### Switch Between Commits

```bash
./tinygit --checkout {COMMIT_HASH}
```

Restores the state of all files to a previous commit.

### Show Commit Metadata

```bash
./tinygit --show-commit {COMMIT_HASH}
```

Displays details of a specific commit, including the author, date, and message.

### Show Tracked Files

```bash
./tinygit --tracked-files
```

Lists all files currently being tracked by TinyGit.

### Get Help

```bash
./tinygit --help
```

Displays a list of available commands and their descriptions.

## Contributing

If you'd like to contribute to TinyGit, feel free to submit issues or pull requests on the repository.

## License

This project is licensed under the MIT License.
