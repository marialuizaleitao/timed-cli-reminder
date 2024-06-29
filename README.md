# Golang CLI Reminder

A simple CLI reminder application written in Go that allows you to set reminders for a specific time and notifies you when it's time.

## Features
- Sets reminders for specific times.
- Notifications displayed using system alerts.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/marialuizaleitao/timed-cli-reminder.git
    ```

2. Navigate to the project directory:
    ```bash
    cd timed-cli-reminder
    ```

3. Build the application:
    ```bash
    go build -o reminder
    ```

## Usage

To set a reminder, use the following command:

```bash
./reminder <hh:mm> <text message>
```

<hh:mm>: The time for the reminder in 24-hour format.
<text message>: The reminder message to be displayed.

### Examples

Set a reminder for 14:30 with the message "Commit your changes":

```bash
./reminder 14:30 "Commit your changes"
```
Set a reminder for 09:00 with the message "Revisit the function":

```bash
./reminder 09:00 "Revisit the sum function"
```

## Dependencies
This application relies on the following Go packages:

- github.com/gen2brain/beeep: For displaying system notifications.
- github.com/olebedev/when: For parsing the time input.

### Install the dependencies using:

```bash
go get github.com/gen2brain/beeep
go get github.com/olebedev/when
```

## Notes
Ensure that the notification icon file (assets/information.png) exists in the specified path.
The application will notify you at the specified time using a system alert.
