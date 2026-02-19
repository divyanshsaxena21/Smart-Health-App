# CalQulate - Smart Health App

An AI-powered nutrition assistant that instantly analyzes food labels from images. Upload a photo of any nutrition label and get a detailed breakdown of calories, fats, sugars, vitamins, and other key nutrients.

![Next.js](https://img.shields.io/badge/Next.js-15-black?logo=next.js)
![Go](https://img.shields.io/badge/Go-1.21-00ADD8?logo=go)
![Google Cloud](https://img.shields.io/badge/Google%20Cloud-Vision%20API-4285F4?logo=google-cloud)
![Tailwind CSS](https://img.shields.io/badge/Tailwind%20CSS-3.x-06B6D4?logo=tailwindcss)

## Features

- **⚡ Instant Analysis** - AI reads your label in seconds with no manual input required
- **📊 Nutritional Breakdown** - Get insights on calories, fat, sugar, protein, vitamins and more
- **🔐 Privacy First** - No data is stored; everything processes securely in real-time
- **🤖 AI-Powered Insights** - Uses Groq AI to provide personalized dietary advice

## Tech Stack

### Frontend
- **Next.js 15** - React framework with App Router
- **TypeScript** - Type-safe JavaScript
- **Tailwind CSS** - Utility-first CSS framework
- **Vercel** - Frontend hosting

### Backend
- **Go 1.21** - High-performance backend
- **Google Cloud Vision API** - OCR for text extraction from images
- **Render** - Backend hosting

## Project Structure

```
Smart-Health-App/
├── frontend/                 # Next.js frontend application
│   ├── app/                  # App Router pages
│   │   ├── page.tsx          # Home page
│   │   ├── upload/           # Upload & analyze page
│   │   └── about/            # About page
│   ├── components/           # Reusable React components
│   └── lib/                  # Utility functions
│
├── backend/                  # Go backend API
│   ├── main.go               # Entry point & server setup
│   ├── handlers/             # HTTP request handlers
│   │   └── image.go          # Image processing endpoint
│   └── utils/                # Utility functions
│       └── nutrition.go      # Nutrition text parsing
│
└── README.md                 # This file
```

## Getting Started

### Prerequisites

- Node.js 18+ and npm
- Go 1.21+
- Google Cloud account with Vision API enabled
- Service account key for Google Cloud Vision API

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Create a `.env` file from the example:
   ```bash
   cp .env.example .env
   ```

3. Add your Google Cloud credentials to `.env`:
   ```env
   GOOGLE_APPLICATION_CREDENTIALS=path/to/your-service-account-key.json
   PORT=5000
   ```

4. Install dependencies and run:
   ```bash
   go mod tidy
   go run main.go
   ```

The backend will start at `http://localhost:5000`

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Create a `.env` file:
   ```env
   NEXT_PUBLIC_API_URL=http://localhost:5000/process-image
   NEXT_PUBLIC_GROQ_API_URL=your_groq_api_key
   ```

4. Run the development server:
   ```bash
   npm run dev
   ```

The frontend will start at `http://localhost:3000`

## API Endpoints

### `POST /process-image`

Extracts nutrition facts from an uploaded food label image.

**Request:**
- Content-Type: `multipart/form-data`
- Body: Form field `File` containing the image

**Response:**
```json
{
  "Calories": "200",
  "Total Fat": "8",
  "Saturated Fat": "3",
  "Cholesterol": "30",
  "Sodium": "450",
  "Total Carbohydrate": "24",
  "Dietary Fiber": "2",
  "Sugars": "12",
  "Protein": "5",
  "Vitamin A": "10",
  "Vitamin C": "15",
  "Calcium": "20",
  "Iron": "8"
}
```

### `GET /health`

Health check endpoint for monitoring.

## Deployment

### Backend (Render)

1. Push code to GitHub
2. Create a new Web Service on [Render](https://render.com)
3. Connect your repository and select the `backend` folder
4. Add environment variable `GOOGLE_CREDENTIALS_JSON` with your full service account JSON content
5. Deploy

### Frontend (Vercel)

1. Push code to GitHub
2. Import project on [Vercel](https://vercel.com)
3. Select the `frontend` folder as the root
4. Add environment variables:
   - `NEXT_PUBLIC_API_URL` - Your Render backend URL
   - `NEXT_PUBLIC_GROQ_API_URL` - Your Groq API key
5. Deploy

## How It Works

1. **Upload** - User uploads a photo of a food nutrition label
2. **OCR** - Google Cloud Vision API extracts text from the image
3. **Parse** - Backend parses the extracted text using regex patterns to identify nutrients
4. **Display** - Frontend displays the structured nutrition data
5. **AI Insights** - Groq AI provides personalized dietary recommendations

## Environment Variables

### Backend
| Variable | Description |
|----------|-------------|
| `PORT` | Server port (default: 5000) |
| `GOOGLE_APPLICATION_CREDENTIALS` | Path to GCP service account key (local) |
| `GOOGLE_CREDENTIALS_JSON` | Full JSON content of service account key (production) |

### Frontend
| Variable | Description |
|----------|-------------|
| `NEXT_PUBLIC_API_URL` | Backend API URL for image processing |
| `NEXT_PUBLIC_GROQ_API_URL` | Groq API key for AI insights |

## License

MIT License - feel free to use this project for personal or commercial purposes.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
