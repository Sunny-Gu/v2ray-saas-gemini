package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Represents the structure of the traffic report expected by node-service
type TrafficReport struct {
	UserID     uint    `json:"user_id"`
	UploadGB   float64 `json:"upload_gb"`
	DownloadGB float64 `json:"download_gb"`
}

// Represents the structure of the user traffic stats from V2Ray's API
type V2RayUserStat struct {
	Email string `json:"email"` // V2Ray identifies users by email in inbound settings
	Uplink   int64  `json:"uplink"`
	Downlink int64  `json:"downlink"`
}

// This function simulates fetching traffic stats from V2Ray's API.
// In a real implementation, this would make an RPC call to the V2Ray API.
func fetchV2RayTrafficStats() ([]V2RayUserStat, error) {
	log.Println("Fetching traffic stats from V2Ray API...")
	// --- SIMULATED DATA ---
	// In a real scenario, you would parse the user ID from the email tag,
	// e.g., "user-123-traffic@yourdomain.com" -> UserID: 123
	stats := []V2RayUserStat{
		{Email: "user-1-traffic", Uplink: 1024 * 1024 * 100, Downlink: 1024 * 1024 * 500}, // 100MB up, 500MB down
		{Email: "user-2-traffic", Uplink: 1024 * 1024 * 200, Downlink: 1024 * 1024 * 800}, // 200MB up, 800MB down
	}
	// --- END SIMULATED DATA ---
	return stats, nil
}

// This function reports the collected traffic data to the node-service.
func reportTrafficToNodeService(reports []TrafficReport) error {
	if len(reports) == 0 {
		log.Println("No traffic to report.")
		return nil
	}

	jsonData, err := json.Marshal(reports)
	if err != nil {
		return err
	}

	nodeServiceURL := "http://localhost:8082/api/v1/node/traffic"
	resp, err := http.Post(nodeServiceURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to report traffic. Status: %s", resp.Status)
	} else {
		log.Printf("Successfully reported traffic for %d users.", len(reports))
	}
	return nil
}

// startTrafficReporter initializes a ticker to periodically report traffic.
func startTrafficReporter() {
	log.Println("Starting traffic reporter...")
	// Report traffic every 1 minute
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		stats, err := fetchV2RayTrafficStats()
		if err != nil {
			log.Printf("Error fetching V2Ray stats: %v", err)
			continue
		}

		var reports []TrafficReport
		for _, stat := range stats {
			// This is a simplified way to get userID from email.
			// A robust implementation would use regex or structured tags.
			var userID uint
			if _, err := fmt.Sscanf(stat.Email, "user-%d-traffic", &userID); err != nil {
				continue // Skip if email format is wrong
			}

			reports = append(reports, TrafficReport{
				UserID:     userID,
				UploadGB:   float64(stat.Uplink) / (1024 * 1024 * 1024),   // Convert bytes to GB
				DownloadGB: float64(stat.Downlink) / (1024 * 1024 * 1024), // Convert bytes to GB
			})
		}

		if err := reportTrafficToNodeService(reports); err != nil {
			log.Printf("Error reporting traffic: %v", err)
		}
	}
}
