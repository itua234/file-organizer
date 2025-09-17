# File Organizer

A simple command-line utility written in Go to organize files in a directory by either their file extension or modification date.

## Features

-   **Organize by Extension**: Moves files into subdirectories named after their file extension (e.g., `images/`, `pdfs/`).
-   **Organize by Date**: Moves files into a `Year/Month` directory structure based on their last modified date (e.g., `2023/January/`).
-   **Dry Run Mode**: Preview file movements without making any changes to the file system.

## Installation

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/your-username/your-repo-name.git](https://github.com/your-username/your-repo-name.git)
    cd your-repo-name
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