# Go Reverse Proxy

This repository contains a Go implementation of a reverse proxy originally written in PHP. The reverse proxy is designed to forward incoming requests to a specified URI while modifying certain aspects of the request and response.

## Functionality

The Go reverse proxy performs the following tasks:

1. **Request Handling**:
   - Parses incoming requests.
   - Redirects requests to a specified URI (`/app/documents`) if the requested URI is `/`.
   - Otherwise, forwards requests to the specified endpoint (`https://beta.frase.io` concatenated with the requested URI).

2. **Header Handling**:
   - Extracts and forwards specific headers from the incoming request to the target server.
   - Adds custom headers such as `Cookie` and `Origin` to the forwarded request.

3. **Request Modification**:
   - Modifies the request method based on the incoming request method (`GET`, `POST`, etc.).
   - Attaches the payload (`POST` data) from the incoming request to the forwarded request.

4. **Response Modification**:
   - Modifies the response received from the target server before forwarding it to the client.
   - Replaces occurrences of `beta.frase.io` with `spyfu.host` in the response HTML.
   - Injects a CSS style into the response HTML to hide specific elements.
   - Sets the appropriate `Content-Type` header for the response.

5. **Client-Side Scripting**:
   - Injects JavaScript code into the response HTML to perform client-side operations.
   - Sets items in the `localStorage` and `sessionStorage`.
   - Redirects the client to a specific URL after a timeout.

## Usage

To use this reverse proxy:

1. Ensure you have Go installed on your system.
2. Clone this repository.
3. Navigate to the repository directory.
4. Build the Go code using the appropriate command.
5. Run the compiled binary, providing any necessary configuration or environment variables.

## Additional Notes

- This reverse proxy is designed for specific use cases and may require modifications for broader applications.
- Be mindful of the legal and ethical considerations when using this reverse proxy, especially when modifying request and response data.
- Ensure compliance with relevant terms of service and privacy policies of the target server and any intermediary services.

For detailed implementation and customization instructions, refer to the source code and relevant documentation within this repository.

