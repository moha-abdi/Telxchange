package merchant

import (
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/moha-abdi/telxchange/api"
	"github.com/moha-abdi/telxchange/config"
	"github.com/moha-abdi/telxchange/internal/database"
	"github.com/moha-abdi/telxchange/internal/exchange/network"
	"github.com/moha-abdi/telxchange/pkg/utils/logger"
	"go.uber.org/zap"
)

var log = logger.GetLogger()

func GenerateMerchants(maxMerchants int, maxGoroutines int) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Couldn't load .env file")
	}
	defer log.Sync()
	username := os.Getenv("TELXCHANGE_USERNAME")
	password := os.Getenv("TELXCHANGE_PASSWORD")
	deviceID := os.Getenv("TELXCHANGE_DEVICE_ID")
	log.Info(
		"Loaded account Info -> ",
		zap.String("Username", username),
		zap.String("DeviceID", deviceID),
	)
	apiCli := api.NewApiClient(config.MFS_PROXY_DEFAULT, username, deviceID)
	_, err := apiCli.Login(password)
	if err != nil {
		log.Fatal("Login Error: ", zap.Error(err))
	}
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal("Couldn't connect to database:", zap.Error(err))
	}

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a channel to limit the number of concurrent goroutines
	semaphore := make(chan struct{}, maxGoroutines) // Adjust this number based on your needs

	for i := 300000; i < 300000+maxMerchants && i < 400000; i++ {
		wg.Add(1)
		semaphore <- struct{}{} // Acquire a slot in the semaphore

		go func(partnerID int) {
			defer wg.Done()
			defer func() { <-semaphore }() // Release the semaphore slot

			retries := 0
			maxRetries := 3
			for retries < maxRetries {
				partner, apiErr := apiCli.GetPartnerInfo(network.Zaad, strconv.Itoa(partnerID))
				if apiErr != nil {
					switch apiErr.Code {
					case "ER0605", "-1":
						retries++
						if retries < maxRetries {
							log.Debug(
								"Rate limit or timeout error. Retrying...",
								zap.String("ErrorCode", apiErr.Code),
								zap.Int("PartnerID", partnerID),
								zap.Int("Retry", retries),
							)
							time.Sleep(
								time.Duration(rand.Intn(5)+1) * time.Second,
							) // Random delay between 1-5 seconds
							continue
						}
						log.Warn(
							"Max retries reached for partner",
							zap.String("ErrorCode", apiErr.Code),
							zap.Int("PartnerID", partnerID),
						)
					default:
						log.Debug(
							"Couldn't get PartnerInfo",
							zap.Error(err),
							zap.Int("PartnerID", partnerID),
						)
					}
					return
				}

				log.Info(
					"Found merchant",
					zap.Int("PartnerID", partnerID),
					zap.String("Name", partner.Name),
					zap.String("Status", partner.Status),
				)

				// Add to database
				errNew := db.AddPartner(partner)
				if errNew != nil {
					log.Error(
						"Failed to add partner to database",
						zap.Error(errNew),
						zap.Int("PartnerID", partnerID),
					)
				}
				break // Success, exit the retry loop
			}
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
