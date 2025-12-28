**spilx** is a simple and fast CLI tool written in **Go** to split large text files into smaller parts based on a specified number of lines. It supports both direct file input and standard input (stdin/pipe).

## Features
- ⚡ Fast splitting of large files.
- 📥 Supports both direct file arguments and piping (`stdin`).
- 🔢 Customizable line count per output file.
- 📝 Auto-generates output filenames (e.g., `file_part_1.txt`).

## Installation

Make sure you have [Go](https://go.dev/dl/) installed. Run the following command to build the tool:

```bash
go install -v github.com/l1ackerronin/spilx@latest
```

## Usage

You can use the tool in two ways:

### 1. Direct File Input
Provide the filename and the number of lines per file:

```bash
./splitter <filename> <lines_per_file>
```

**Example:**
```bash
./splitter hackerone.txt 10000
```
*This will create `hackerone_part_1.txt`, `hackerone_part_2.txt`, etc.*

### 2. Standard Input (Pipe)
Pipe the output of another command into `splitter`. In this mode, output files will be named `stdin_part_N.txt`.

```bash
cat <filename> | ./splitter <lines_per_file>
```


# Author
- **GitHub**: [@l1ackerronin](https://github.com/l1ackerronin)
- **Twitter**: [@l1ackerronin](https://x.com/l1ackerronin)
- **Linkedin**: [@l1ackerronin](https://www.linkedin.com/in/l1ackerronin)
- **Instagram**: [@l1ackerronin](https://www.instagram.com/l1ackerronin)
- **Email**: [l1ackerronin@gmail.com](mailto:l1ackerronin@gmail.com)
