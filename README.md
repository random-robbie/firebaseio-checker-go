# Firebase Database Security Checker

A Go-based security audit tool designed to identify publicly accessible Firebase Realtime Database instances with misconfigured security rules. This tool helps security researchers and developers discover exposed databases during penetration testing and security assessments.

## Features

- **Concurrent Scanning**: Efficiently check multiple Firebase databases simultaneously
- **Configurable Parameters**: Customize concurrency, timeouts, and output directories
- **Colored Output**: Clear visual feedback for scan results
- **JSON Export**: Automatically save accessible database contents to JSON files
- **Error Handling**: Robust error handling with detailed logging
- **Comment Support**: Add comments to your input file with `#`

## Prerequisites

- Go 1.16 or higher
- Internet connection for scanning remote Firebase instances

## Installation

### Clone the Repository

```bash
git clone https://github.com/random-robbie/firebaseio-checker-go.git
cd firebaseio-checker-go
```

### Install Dependencies

Dependencies are managed using Go modules. Simply run:

```bash
go mod download
```

### Build the Application

```bash
go build -o firebaseio-checker
```

## Usage

### Basic Usage

Create a file named `list.txt` containing Firebase database URLs (one per line, without `https://`):

```
example-db.firebaseio.com
another-db.firebaseio.com
test-database.firebaseio.com
```

Run the scanner:

```bash
./firebaseio-checker
```

### Advanced Options

The tool supports several command-line flags for customization:

```bash
./firebaseio-checker [options]
```

**Available Flags:**

- `-input <file>`: Input file containing Firebase URLs (default: `list.txt`)
- `-output <dir>`: Output directory for JSON files (default: `./cfg/`)
- `-concurrency <num>`: Number of concurrent requests (default: `10`)
- `-timeout <seconds>`: HTTP request timeout in seconds (default: `30`)
- `-colors <bool>`: Enable or disable colored output (default: `true`)

**Example:**

```bash
./firebaseio-checker -input targets.txt -output ./results/ -concurrency 20 -timeout 15
```

## Output

- **Accessible Databases**: JSON files are saved to the output directory (default: `./cfg/`)
- **File Naming**: Files are named after the database name (e.g., `example-db.json`)
- **Console Output**: Real-time status updates with color-coded results:
  - 🟢 Green: Database accessible and saved
  - 🟡 Yellow: Access denied (properly configured)
  - 🔴 Red: Database not found or request failed

[![Example Screenshot](https://i.postimg.cc/52ngnsqp/Capture.png)](https://postimg.cc/75JSLMRT)

## Ethical Usage

**⚠️ IMPORTANT:** This tool is designed for authorized security testing only. Always ensure you have explicit permission before scanning Firebase databases. Unauthorized access to computer systems is illegal.

**Intended Use Cases:**
- Bug bounty programs
- Authorized penetration testing
- Security audits of your own applications
- Educational purposes in controlled environments

## Dependencies

- [github.com/logrusorgru/aurora](https://github.com/logrusorgru/aurora) - Terminal colors
- [github.com/remeh/sizedwaitgroup](https://github.com/remeh/sizedwaitgroup) - Concurrency control

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

Copyright @random_robbie (c) 2018

## Infrastructure Recommendations

For efficient scanning of large target lists, consider using a VPS with good network connectivity:

[![DigitalOcean Referral Badge](https://web-platforms.sfo2.cdn.digitaloceanspaces.com/WWW/Badge%201.svg)](https://www.digitalocean.com/?refcode=e22bbff5f6f1&utm_campaign=Referral_Invite&utm_medium=Referral_Program&utm_source=badge)

## Author

Created by [@random_robbie](https://github.com/random-robbie)
