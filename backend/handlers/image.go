package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	vision "cloud.google.com/go/vision/v2/apiv1"
	visionpb "cloud.google.com/go/vision/v2/apiv1/visionpb"
	"smart-health-app/utils"
)

// ProcessImage handles the /process-image endpoint
func ProcessImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form with 10MB max memory
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		sendJSONError(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Get the uploaded file
	file, handler, err := r.FormFile("File")
	if err != nil {
		log.Printf("Error getting file: %v", err)
		sendJSONError(w, "No image file provided", http.StatusBadRequest)
		return
	}
	defer file.Close()

	log.Printf("Received image: %s", handler.Filename)

	// Read file content
	content, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		sendJSONError(w, "Failed to read image file", http.StatusInternalServerError)
		return
	}

	// Extract text using Google Cloud Vision
	extractedText, err := detectDocumentText(content)
	if err != nil {
		log.Printf("Error extracting text: %v", err)
		sendJSONError(w, "Failed to process image", http.StatusInternalServerError)
		return
	}

	// Parse nutrition facts from extracted text
	parsedData := utils.ParseNutritionFacts(extractedText)

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parsedData)
}

// detectDocumentText uses Google Cloud Vision API to extract text from image
func detectDocumentText(imageContent []byte) (string, error) {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	image := &visionpb.Image{Content: imageContent}
	request := &visionpb.BatchAnnotateImagesRequest{
		Requests: []*visionpb.AnnotateImageRequest{
			{
				Image: image,
				Features: []*visionpb.Feature{
					{Type: visionpb.Feature_DOCUMENT_TEXT_DETECTION},
				},
			},
		},
	}

	resp, err := client.BatchAnnotateImages(ctx, request)
	if err != nil {
		return "", err
	}

	if len(resp.Responses) == 0 || resp.Responses[0].FullTextAnnotation == nil {
		return "", nil
	}

	return resp.Responses[0].FullTextAnnotation.Text, nil
}

// sendJSONError sends an error response in JSON format
func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
