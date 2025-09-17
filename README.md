# File Organizer

A simple command-line utility written in Go to organize files in a directory by either their file extension or modification date.

## Features

-   **Organize by Extension**: Moves files into subdirectories named after their file extension (e.g., `images/`, `pdfs/`).
-   **Organize by Date**: Moves files into a `Year/Month` directory structure based on their last modified date (e.g., `2023/January/`).
-   **Dry Run Mode**: Preview file movements without making any changes to the file system.

## Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/itua234/file-organizer.git
    cd file-organizer
    ```
2.  **Build the executable:**
    ```bash
    go build -o organizer.exe
    ```
    This will create an executable file named `organizer` (or `organizer.exe` on Windows).

## Usage

Run the executable from your terminal with the required flags.

```bash
./organizer -path <directory> -mode <organize-mode> -dry-run


### Flags

-   `-path` (string): The path to the directory you want to organize. Defaults to `./files`.
-   `-mode` (string): The organization method. Can be `extension` or `date`. Defaults to `extension`.
-   `-dry-run` (boolean): If present, the program will only print the actions it would take without moving any files.

### Examples

**Organize files by extension in the default `./files` directory:**
```bash
./organizer.exe
