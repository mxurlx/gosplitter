# Gosplitter

A command-line tool written in Go for splitting large files into smaller chunks and merging them back together.

## Usage

`gosplitter <command> [flags]`

Available commands: `split`, `merge`

## Common Flags

| Flag     | Short | Description        |
|----------|-------|--------------------|
| `--help`   | `-h`  | Show help message  |
| `--version`| `-v`  | Show version       |

## Split Command

Splits a file into multiple chunks.

`gosplitter split [flags] <input>`

| Flag        | Short | Default       | Description              |
|-------------|-------|---------------|--------------------------|
| `--output`   | `-o`  | `input_name` | Output directory         |
| `--suffix`   | `-s`  | `part`       | Chunk suffix             |
| `--chunksize`| `-c`  | `4096`       | Chunk size in bytes      |

**Example:**

```bash
gosplitter split -o output -s chunk mylargefile.txt
```

This will split `mylargefile.txt` into chunks with the suffix "chunk" and place them in the `output` directory, using a chunk size of 4096 bytes.

## Merge Command

Merges multiple file chunks back into a single file.

`gosplitter merge [flags] <input>`

| Flag       | Short | Description       |
|------------|-------|-------------------|
| `--output` | `-o`  | Output file       |

**Example:**

```bash
gosplitter merge -o mergedfile.txt input
```

This will merge all the files in the `input` directory into a single file named `mergedfile.txt`. Files are sorted based on their chunk number embedded within the filename.
