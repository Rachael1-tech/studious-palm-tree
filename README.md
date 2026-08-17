## Email Sender (Go)

A simple Go program that sends an email via SMTP using Gmail, with credentials and addresses supplied through environment variables (nothing sensitive is hardcoded).

### Requirements
- Go installed (check with `go version`)
- A Gmail account with 2-Step Verification enabled
- A Gmail App Password ([generate one here](https://myaccount.google.com/apppasswords))

### Setup

1. Install dependencies:
```bash
   go mod tidy
```

2. Set the required environment variables:

   **Git Bash / macOS / Linux:**
```bash
   export SMTP_FROM="youremail@gmail.com"
   export SMTP_TO="recipient@example.com"
   export SMTP_PASSWORD="your16charapppassword"
```

   **PowerShell:**
```powershell
   $env:SMTP_FROM="youremail@gmail.com"
   $env:SMTP_TO="recipient@example.com"
   $env:SMTP_PASSWORD="your16charapppassword"
```

3. Run the program:
```bash
   go run main.go
```

### Notes
- Environment variables only persist for the current terminal session — you'll need to re-set them if you open a new terminal.
- Never commit real credentials or `.env` files containing secrets.