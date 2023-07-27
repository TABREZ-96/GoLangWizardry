package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Define the struct for the PurchaseEvent in the Mable format
type PurchaseEvent struct {
	// Add the necessary fields as per the Mable format
	EventName string `json:"eventName"`
	// Add other required fields...
}

// Function to send the Purchase Event to the Facebook test pixel
func sendEventToFacebook(pixelID, accessKey, testEventCode string, event PurchaseEvent) error {
	// Serialize the event to JSON
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Apply normalization as per FB API requirements (Remove white spaces and convert to lowercase)
	normalizedPayload := bytes.ToLower(bytes.ReplaceAll(payload, []byte(" "), []byte("")))

	// Get the current Unix timestamp (in seconds)
	timestamp := time.Now().Unix()

	// Generate the sha256 HMAC signature using the accessKey as the key
	h := hmac.New(sha256.New, []byte(accessKey))
	h.Write([]byte(fmt.Sprintf("%d.%s", timestamp, normalizedPayload)))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	// Create the HTTP request to send the event to Facebook test pixel
	url := fmt.Sprintf("https://graph.facebook.com/v17.0/%s/events", pixelID)
	req, err := http.NewRequest("POST", url, bytes.NewReader(normalizedPayload))
	if err != nil {
		return err
	}

	// Set the required headers for the Facebook API request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-FB-Event-Time", fmt.Sprintf("%d", timestamp))
	req.Header.Set("X-FB-Pixel-Id", pixelID)
	req.Header.Set("X-FB-Test-Event-Code", testEventCode)
	req.Header.Set("X-FB-Debug", "true")
	req.Header.Set("X-FB-Data-Hmac-Sha256", signature)

	// Send the HTTP request to Facebook
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Validate the response from Facebook (optional, you can choose to handle different response codes accordingly)

	return nil
}

// Handler for the API endpoint
func handlePurchaseEvent(w http.ResponseWriter, r *http.Request) {
	var event PurchaseEvent
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Access the pixelId, accessKey, and testEventCode from the request payload or configuration
	pixelID := "1300740743881123"
	accessKey := "EAAJT0IZCxLsABOx90t71rhO6QZCe7leyygMD0Y6ua7o4bQKpSuGiZCA6Y4ASEfYGVenWavsBxCNJ04DtwMRrb6hhvpMaZBhNRmYm8Qo3dFFrs8e1WCQohQjG4M5uf3wg921AzF1XVZAxoZBG3fPsmCHtI1CXzKzQx1y4Vsv3qGZAeieLmGA4GriDZAK0ymlG1ENQbZAV9A5nnjA6XItEj9ZBZCbE0oeRwZDZD"
	testEventCode := "TEST29466"

	// Send the event to Facebook test pixel
	err = sendEventToFacebook(pixelID, accessKey, testEventCode, event)
	if err != nil {
		http.Error(w, "Error sending event to Facebook", http.StatusInternalServerError)
		return
	}

	// Respond with a success message or any other useful response
	response := map[string]interface{}{

		"message": "Event sent successfully to Facebook test pixel",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/send/purchase", handlePurchaseEvent).Methods("POST")

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", router)
}
