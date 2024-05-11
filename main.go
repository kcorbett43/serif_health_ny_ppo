package main

import (
	"fmt"
	"net/http"
	"compress/gzip"
	"encoding/json"
	"io"
	"strings"
	"os"
	"time"
)


type Plan struct {
	PlanName       string `json:"plan_name"`
	PlanIDType     string `json:"plan_id_type"`
	PlanID         string `json:"plan_id"`
	PlanMarketType string `json:"plan_market_type"`
}

type FileLocation struct {
	Description string `json:"description"`
	Location    string `json:"location"`
}

type ReportingStructure struct {
	ReportingPlans     []Plan          `json:"reporting_plans"`
	InNetworkFiles     []FileLocation  `json:"in_network_files"`
	AllowedAmountFile  *FileLocation   `json:"allowed_amount_file"`
}

type ReportingMetaData struct {
	ReportingEntityName string               `json:"reporting_entity_name"`
	ReportingEntityType string               `json:"reporting_entity_type"`
	ReportingStructure  []ReportingStructure `json:"reporting_structure"`
	Version 						string 							 `json:"version"`
}

func main() {

	// Create file to write URLs to 
	folder := "./ny_ppo_output"

	// Format yyyyMMdd hhmmss
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("%s/file_%s.txt", folder, timestamp)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return 
	}
	defer file.Close()

	// Download anthem file
	anthemMrfMay2024Url := "https://antm-pt-prod-dataz-nogbd-nophi-us-east1.s3.amazonaws.com/anthem/2024-05-01_anthem_index.json.gz"
	resp, err := http.Get(anthemMrfMay2024Url)
	if err != nil {
		fmt.Println("Error downloading Anthem MRF for May 2024:", err)
		return 
	}
	defer resp.Body.Close()

	// Decompress gzip content
	gzReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return 
	}
	defer gzReader.Close()


	// Decode 
	decoder := json.NewDecoder(gzReader)

	for {
		token, err := decoder.Token()
		if err == io.EOF {
				break
		}
		if err != nil {
				continue
		}

		// Check for reporting structure
		if token == "reporting_structure" {
			// Expect the array to start
			_, err := decoder.Token()
			if err != nil {
				continue
			}
			
			// Decode each as it comes
			for decoder.More() {
				var reportingStructure ReportingStructure
				err := decoder.Decode(&reportingStructure)
				if err != nil {
					continue
				}

				for _, fileLocation := range reportingStructure.InNetworkFiles {
					err := decoder.Decode(&fileLocation)
					if err != nil {
						continue
					}

					// Get location for analysis, if ny ppo then write to file
					location := fileLocation.Location
					if isNewYorkPPO(location) {
						file.WriteString(fmt.Sprintf("%s\n", location))
					}
				}
			}

			// Read the end array token
			_, err = decoder.Token()
			if err != nil {
				fmt.Println("Error reading array end:", err)
				return
			}
		}
	}	
	
}


// Look for URLs that have the "39" identifier (eg: 39B0, 39F0)
func isNewYorkPPO(url string) bool {
	urlParts := strings.Split(url, "/")
	if len(urlParts) < 1 {
		return false
	}

	endPart := urlParts[len(urlParts)-1]
	standardPart := strings.Split(endPart, ".json")
	if len(standardPart) < 1 {
		return false
	}

	return strings.Contains(standardPart[0], "_39")
}

