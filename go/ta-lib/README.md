# TA-Lib for Go

This is a pure Go implementation of the Technical Analysis Library (TA-Lib), ported from the original C implementation. It provides a comprehensive collection of technical analysis indicators and functions for financial market data analysis.

## Installation

```bash
go get github.com/petercool/ta-lib/go/ta-lib
```

## Features

- Pure Go implementation
- No CGO dependencies
- Thread-safe
- Implements all TA-Lib indicators
- Built-in data feeds for multiple sources
- Optimized for performance

## Data Feeds

The library includes built-in data feeds for various sources:

### CSV Feed

```go
import "github.com/petercool/ta-lib/go/ta-lib/utils"

// Create a CSV feed
feed := utils.NewCSVFeed("path/to/data.csv")
data, err := feed.GetData(time.Time{}, time.Time{})
```

### Binance Feed

```go
// Create a Binance feed with default API URL
feed := utils.NewBinanceFeed("BTCUSDT", "1d")

// Or with custom API URL
feed := utils.NewBinanceFeedWithBaseURL("BTCUSDT", "1d", "https://api.binance.us/api/v3")

// Fetch data for a specific time range
data, err := feed.GetData(startTime, endTime)
```

### Yahoo Finance Feed

```go
// Create a Yahoo feed with default API URL
feed := utils.NewYahooFeed("AAPL", "1d")

// Fetch data for a specific time range
data, err := feed.GetData(startTime, endTime)
```

### Converting Data for Analysis

```go
// Convert OHLCV data to slices for technical analysis
open, high, low, close, volume := utils.GetOHLCVSlices(data)
```

## Technical Indicators

The library includes implementations of all popular technical analysis indicators:

### Moving Averages

- Simple Moving Average (SMA)
- Exponential Moving Average (EMA)
- Weighted Moving Average (WMA)
- Double Exponential Moving Average (DEMA)
- Triple Exponential Moving Average (TEMA)
- Triangular Moving Average (TRIMA)

### Momentum Indicators

- Relative Strength Index (RSI)
- Moving Average Convergence Divergence (MACD)
- Stochastic Oscillator
- Rate of Change (ROC)
- Commodity Channel Index (CCI)
- Average Directional Index (ADX)

### Volatility Indicators

- Bollinger Bands
- Average True Range (ATR)
- Standard Deviation

### Volume Indicators

- On Balance Volume (OBV)
- Accumulation/Distribution Line
- Money Flow Index (MFI)

### Price Transform

- Average Price
- Median Price
- Typical Price
- Weighted Close Price

### Pattern Recognition

- Various candlestick patterns
- Support and resistance levels

## Example Usage

```go
import (
    "github.com/petercool/ta-lib/go/ta-lib/indicators"
    "github.com/petercool/ta-lib/go/ta-lib/utils"
)

// Fetch data from Binance
feed := utils.NewBinanceFeed("BTCUSDT", "1d")
data, err := feed.GetData(startTime, endTime)
if err != nil {
    log.Fatal(err)
}

// Convert to slices
_, high, low, close, _ := utils.GetOHLCVSlices(data)

// Calculate RSI
rsi, err := indicators.RSI(close, 14)
if err != nil {
    log.Fatal(err)
}

// Calculate Bollinger Bands
bb, err := indicators.BBANDS(close, 20, 2.0, 2.0)
if err != nil {
    log.Fatal(err)
}
```

## License

This project is licensed under the same terms as the original TA-Lib.
