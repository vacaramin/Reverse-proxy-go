# Reverse Proxy Server

This repository contains a Go program for a reverse proxy server that handles requests to specific endpoints by forwarding them to a target URL, modifying the response body if necessary, and serving the modified response back to the client.

## Installation

Ensure you have Go installed on your system. Clone this repository and navigate to the directory containing the `main.go` file.

## Usage

Run the Go program by executing the

```bash
go run main.go
```

command. The server will start listening on port 8080 by default.

## Endpoints

### `/login`

The `/login` endpoint handles successful login requests and returns an HTML response indicating successful login. It sets the necessary CORS headers to allow cross-origin requests.

### `/` (default)

The default endpoint redirects requests to the specified URI `/app/documents`. It sets the necessary CORS headers and forwards the request to the target URL `https://beta.frase.io`.

## Dependencies

This program utilizes the following standard Go packages:

- `net/http` for handling HTTP requests and responses
- `bytes`, `fmt`, `io`, `log`, `strings` for various utility functions
