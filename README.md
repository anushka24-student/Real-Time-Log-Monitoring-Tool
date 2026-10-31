# LogPulse 
## Real-Time Log Monitoring & Alert System

LogPulse is a real-time log monitoring tool built using **Go (Golang)**. It monitors multiple log files concurrently, detects errors, tracks error counts, generates threshold-based alerts, and exports summary reports.

## Features

- Monitor multiple log files simultaneously.
- Detect `ERROR`, `FAILURE`, and `CRITICAL` messages.
- Track error counts by file, type, and time.
- Generate alerts when errors reach a defined threshold.
- Export reports in CSV and JSON formats.
- Handle missing or inaccessible log files gracefully.
- Concurrent monitoring using goroutines and channels.

## Tech Stack

- **Language:** Go
- **Concurrency:** Goroutines and Channels
- **File Handling:** Go standard library
- **Report Formats:** CSV and JSON

## Project Structure

```text
logpulse-go/
├── main.go
├── go.mod
├── sample-logs/
│   ├── app.log
│   └── server.log
├── summary.csv
├── summary.json
└── README.md
```

## Getting Started

1. Clone the repository

```bash
git clone https://github.com/YOUR_USERNAME/logpulse-go.git
cd logpulse-go
```

2. Run the project

```bash
go run main.go
```

3. Stop monitoring

Press `Ctrl + C` to stop the program and generate summary reports.

## Configuration

The error threshold and monitoring interval can be adjusted in `main.go`:

```go
const (
    ERROR_THRESHOLD = 5
    POLL_INTERVAL   = 500 * time.Millisecond
)
```

## Future Improvements

- Live monitoring dashboard.
- REST API integration.
- Configurable alert thresholds.
- Docker support.
- Unit and concurrency testing.

