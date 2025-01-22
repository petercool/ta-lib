package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	defaultBinanceBaseURL = "https://api.binance.com/api/v3"
	defaultYahooBaseURL   = "https://query1.finance.yahoo.com/v8/finance"
)

// OHLCV represents a single candlestick data point
type OHLCV struct {
	Time   int64
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

// DataFeed defines the interface for different data sources
type DataFeed interface {
	// GetData retrieves OHLCV data for the specified time range
	GetData(startTime, endTime time.Time) ([]OHLCV, error)
}

// CSVFeed implements DataFeed for CSV file data source
type CSVFeed struct {
	FilePath string
}

// BinanceFeed implements DataFeed for Binance API
type BinanceFeed struct {
	Symbol   string
	Interval string
	BaseURL  string // Optional base URL
}

// YahooFeed implements DataFeed for Yahoo Finance API
type YahooFeed struct {
	Symbol   string
	Interval string
	BaseURL  string // Optional base URL
}

// NewCSVFeed creates a new CSV data feed
func NewCSVFeed(filePath string) *CSVFeed {
	return &CSVFeed{FilePath: filePath}
}

// NewBinanceFeed creates a new Binance data feed
func NewBinanceFeed(symbol, interval string) *BinanceFeed {
	return &BinanceFeed{
		Symbol:   symbol,
		Interval: interval,
		BaseURL:  defaultBinanceBaseURL,
	}
}

// NewBinanceFeedWithBaseURL creates a new Binance data feed with custom base URL
func NewBinanceFeedWithBaseURL(symbol, interval, baseURL string) *BinanceFeed {
	return &BinanceFeed{
		Symbol:   symbol,
		Interval: interval,
		BaseURL:  baseURL,
	}
}

// NewYahooFeed creates a new Yahoo Finance data feed
func NewYahooFeed(symbol, interval string) *YahooFeed {
	return &YahooFeed{
		Symbol:   symbol,
		Interval: interval,
		BaseURL:  defaultYahooBaseURL,
	}
}

// NewYahooFeedWithBaseURL creates a new Yahoo Finance data feed with custom base URL
func NewYahooFeedWithBaseURL(symbol, interval, baseURL string) *YahooFeed {
	return &YahooFeed{
		Symbol:   symbol,
		Interval: interval,
		BaseURL:  baseURL,
	}
}

// GetData implements DataFeed interface for CSVFeed
func (f *CSVFeed) GetData(startTime, endTime time.Time) ([]OHLCV, error) {
	// Open CSV file
	file, err := os.Open(f.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	// Create CSV reader
	reader := csv.NewReader(file)

	// Read header
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %w", err)
	}

	// Validate header
	if len(header) < 6 || header[0] != "time" {
		return nil, fmt.Errorf("invalid CSV format: expected time,open,high,low,close,volume")
	}

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV records: %w", err)
	}

	// Parse records into OHLCV structs
	data := make([]OHLCV, 0, len(records))
	for _, record := range records {
		timestamp, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			continue
		}

		open, _ := strconv.ParseFloat(record[1], 64)
		high, _ := strconv.ParseFloat(record[2], 64)
		low, _ := strconv.ParseFloat(record[3], 64)
		close, _ := strconv.ParseFloat(record[4], 64)
		volume, _ := strconv.ParseFloat(record[5], 64)

		data = append(data, OHLCV{
			Time:   timestamp,
			Open:   open,
			High:   high,
			Low:    low,
			Close:  close,
			Volume: volume,
		})
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("no data found in CSV file")
	}

	return data, nil
}

// GetData implements DataFeed interface for BinanceFeed
func (f *BinanceFeed) GetData(startTime, endTime time.Time) ([]OHLCV, error) {
	// Construct URL
	url := fmt.Sprintf(
		"%s/klines?symbol=%s&interval=%s&startTime=%d&endTime=%d&limit=1000",
		f.BaseURL,
		f.Symbol,
		f.Interval,
		startTime.UnixMilli(),
		endTime.UnixMilli(),
	)

	// Make request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Binance data: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Binance response: %w", err)
	}

	// Parse response
	var klines [][]interface{}
	if err := json.Unmarshal(body, &klines); err != nil {
		return nil, fmt.Errorf("failed to parse Binance response: %w", err)
	}

	// Convert to OHLCV format
	data := make([]OHLCV, len(klines))
	for i, k := range klines {
		openTime := int64(k[0].(float64))
		open, _ := strconv.ParseFloat(k[1].(string), 64)
		high, _ := strconv.ParseFloat(k[2].(string), 64)
		low, _ := strconv.ParseFloat(k[3].(string), 64)
		close, _ := strconv.ParseFloat(k[4].(string), 64)
		volume, _ := strconv.ParseFloat(k[5].(string), 64)

		data[i] = OHLCV{
			Time:   openTime / 1000, // Convert from milliseconds to seconds
			Open:   open,
			High:   high,
			Low:    low,
			Close:  close,
			Volume: volume,
		}
	}

	return data, nil
}

// GetData implements DataFeed interface for YahooFeed
func (f *YahooFeed) GetData(startTime, endTime time.Time) ([]OHLCV, error) {
	// Convert interval to Yahoo format
	yahooInterval := convertIntervalToYahoo(f.Interval)

	// Construct URL
	url := fmt.Sprintf(
		"%s/chart/%s?period1=%d&period2=%d&interval=%s",
		f.BaseURL,
		f.Symbol,
		startTime.Unix(),
		endTime.Unix(),
		yahooInterval,
	)

	// Make request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Yahoo data: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Yahoo response: %w", err)
	}

	// Parse response
	var yahooResp struct {
		Chart struct {
			Result []struct {
				Timestamp  []int64 `json:"timestamp"`
				Indicators struct {
					Quote []struct {
						Open   []float64 `json:"open"`
						High   []float64 `json:"high"`
						Low    []float64 `json:"low"`
						Close  []float64 `json:"close"`
						Volume []float64 `json:"volume"`
					} `json:"quote"`
				} `json:"indicators"`
			} `json:"result"`
		} `json:"chart"`
	}

	if err := json.Unmarshal(body, &yahooResp); err != nil {
		return nil, fmt.Errorf("failed to parse Yahoo response: %w", err)
	}

	// Check if we have any data
	if len(yahooResp.Chart.Result) == 0 {
		return nil, fmt.Errorf("no data available for %s", f.Symbol)
	}

	result := yahooResp.Chart.Result[0]
	quotes := result.Indicators.Quote[0]

	// Convert to OHLCV format
	data := make([]OHLCV, len(result.Timestamp))
	for i, ts := range result.Timestamp {
		data[i] = OHLCV{
			Time:   ts,
			Open:   quotes.Open[i],
			High:   quotes.High[i],
			Low:    quotes.Low[i],
			Close:  quotes.Close[i],
			Volume: quotes.Volume[i],
		}
	}

	return data, nil
}

// GetOHLCVSlices converts OHLCV data to separate slices
func GetOHLCVSlices(data []OHLCV) ([]float64, []float64, []float64, []float64, []float64) {
	open := make([]float64, len(data))
	high := make([]float64, len(data))
	low := make([]float64, len(data))
	close := make([]float64, len(data))
	volume := make([]float64, len(data))

	for i, d := range data {
		open[i] = d.Open
		high[i] = d.High
		low[i] = d.Low
		close[i] = d.Close
		volume[i] = d.Volume
	}

	return open, high, low, close, volume
}

// convertIntervalToYahoo converts standard interval notation to Yahoo format
func convertIntervalToYahoo(interval string) string {
	switch interval {
	case "1m":
		return "1m"
	case "5m":
		return "5m"
	case "15m":
		return "15m"
	case "30m":
		return "30m"
	case "1h":
		return "1h"
	case "1d":
		return "1d"
	case "1w":
		return "1wk"
	case "1M":
		return "1mo"
	default:
		return "1d"
	}
}
