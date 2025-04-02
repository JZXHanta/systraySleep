# Systray Sleep

A system tray application that allows you to schedule your computer to enter sleep mode after a specified time period. This tool is particularly useful for managing media playback sessions and preventing unnecessary system operation.

## Features
- System tray interface for easy access
- Configurable sleep timer options (30 minutes, 1 hour, 2 hours)
- Minimal resource usage
- Background operation

## Current Status
- **Known Issue**: The application currently triggers system hibernation instead of sleep mode. This may be related to system privileges or hardware configuration.

## Planned Features
- Real-time countdown timer display on hover/click
- Last-minute cancellation prompt
- Additional timer customization options
- System tray icon status indicator

## How to Run

### Prerequisites
- Go 1.16 or later
- Windows operating system (for sleep functionality)

### Building from Source
1. Clone the repository:
   ```bash
   git clone https://github.com/systraySleep.git
   cd systraySleep
   ```
2. Build the application:
   ```bash
   go build
   ```
3. Run the executable:
   ```bash
   ./systraySleep
   ```

### Development Mode
To run directly from source:
```bash
go run .
```

## Usage
1. Launch the application
2. Access the system tray menu
3. Select your desired sleep timer
4. The system will automatically enter sleep mode after the selected duration

## Development
This project is under active development. Contributions and suggestions are welcome.