# Smart Health App - Go Backend

A Go backend service that extracts nutrition facts from images using Google Cloud Vision API.

## Prerequisites

- Go 1.21 or later
- Google Cloud account with Vision API enabled
- Service account key file for Google Cloud Vision API

## Setup

1. **Set up Google Cloud credentials:**
   ```bash
   export GOOGLE_APPLICATION_CREDENTIALS="/path/to/your/service-account-key.json"
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the server:**
   ```bash
   go run main.go
   ```

   The server will start on port 5000 by default. Set the `PORT` environment variable to use a different port.

## API Endpoints

### POST /process-image

Extracts nutrition facts from an uploaded image.

**Request:**
- Content-Type: `multipart/form-data`
- Body: Form field `File` containing the image file

**Response:**
```json
{
  "Calories": "200",
  "Total Fat": "8",
  "Protein": "5",
  "Total Carbohydrate": "24",
  ...
}
```

## Project Structure

```
backend/
├── main.go           # Entry point, HTTP server setup
├── go.mod            # Go module definition
├── handlers/
│   └── image.go      # HTTP request handlers
└── utils/
    └── nutrition.go  # Text processing and nutrition extraction
```

## Building for Production

```bash
go build -o server main.go
./server
```

## Environment Variables

- `PORT` - Server port (default: 5000)
- `GOOGLE_APPLICATION_CREDENTIALS` - Path to Google Cloud service account key file

## CORS Configuration

The server is configured to accept requests from:
- http://localhost:3000
- https://cal-qulate.vercel.app
- https://calqulate.ayushsharma.site
