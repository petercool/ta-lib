package main

import (
	"fmt"
	"log"
	"time"

	"github.com/petercool/ta-lib/go/ta-lib/utils"
)

func main() {
	// Example 1: Use CSV Feed
	fmt.Println("Reading data from CSV...")
	csvFeed := utils.NewCSVFeed("../../go/ta-lib/tests/testdata/data_0.csv")
	csvData, err := csvFeed.GetData(time.Time{}, time.Time{}) // Empty time values since we don't need filtering
	if err != nil {
		log.Fatal("Error reading CSV data:", err)
	}

	fmt.Printf("Retrieved %d candlesticks from CSV\n", len(csvData))
	for i, candle := range csvData[len(csvData)-10:] { // Show the last 10 candles
		fmt.Printf("Candle %d: Time=%s, Open=%.2f, High=%.2f, Low=%.2f, Close=%.2f, Volume=%.2f\n",
			i+1,
			time.Unix(candle.Time, 0).Format("2006-01-02 15:04:05"),
			candle.Open,
			candle.High,
			candle.Low,
			candle.Close,
			candle.Volume,
		)
	}

	// Example 2: Use Binance Feed with custom base URL (optional)
	fmt.Println("\nFetching BTC/USDT data from Binance...")
	binanceFeed := utils.NewBinanceFeed("BTCUSDT", "1d")
	// Or with custom base URL:
	// binanceFeed := utils.NewBinanceFeedWithBaseURL("BTCUSDT", "1d", "https://api.binance.us/api/v3")
	endTime := time.Now()
	startTime := endTime.Add(-7 * 24 * time.Hour) // Last 7 days

	btcData, err := binanceFeed.GetData(startTime, endTime)
	if err != nil {
		log.Fatal("Error fetching Binance data:", err)
	}

	fmt.Printf("Retrieved %d BTC/USDT candlesticks\n", len(btcData))
	for i, candle := range btcData {
		fmt.Printf("Candle %d: Time=%s, Open=%.2f, High=%.2f, Low=%.2f, Close=%.2f, Volume=%.2f\n",
			i+1,
			time.Unix(candle.Time, 0).Format("2006-01-02 15:04:05"),
			candle.Open,
			candle.High,
			candle.Low,
			candle.Close,
			candle.Volume,
		)
	}

	// Example 3: Use Yahoo Feed with custom base URL (optional)
	fmt.Println("\nFetching AAPL data from Yahoo Finance...")
	yahooFeed := utils.NewYahooFeed("AAPL", "1d")
	// Or with custom base URL:
	// yahooFeed := utils.NewYahooFeedWithBaseURL("AAPL", "1d", "https://custom.yahoo.api/v8/finance")
	endTime = time.Now()
	startTime = endTime.Add(-30 * 24 * time.Hour) // Last 30 days

	aaplData, err := yahooFeed.GetData(startTime, endTime)
	if err != nil {
		log.Fatal("Error fetching Yahoo data:", err)
	}

	fmt.Printf("Retrieved %d AAPL candlesticks\n", len(aaplData))
	for i, candle := range aaplData {
		fmt.Printf("Candle %d: Time=%s, Open=%.2f, High=%.2f, Low=%.2f, Close=%.2f, Volume=%.2f\n",
			i+1,
			time.Unix(candle.Time, 0).Format("2006-01-02 15:04:05"),
			candle.Open,
			candle.High,
			candle.Low,
			candle.Close,
			candle.Volume,
		)
	}

	// Convert to slices for technical analysis
	open, high, low, close, volume := utils.GetOHLCVSlices(aaplData)
	fmt.Printf("\nData ready for technical analysis:\n")
	fmt.Printf("Open: %v\n", open[:5]) // Show first 5 values
	fmt.Printf("High: %v\n", high[:5])
	fmt.Printf("Low: %v\n", low[:5])
	fmt.Printf("Close: %v\n", close[:5])
	fmt.Printf("Volume: %v\n", volume[:5])
}
