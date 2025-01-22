package main

import (
	"fmt"
	"log"

	"github.com/petercool/ta-lib/go/ta-lib/indicators"
	"github.com/petercool/ta-lib/go/ta-lib/tests/testdata"
)

func main() {
	// Get test data
	_, high, low, close, volume, err := testdata.GetOHLCVSlices()
	if err != nil {
		log.Fatal("Error loading test data:", err)
	}

	// Calculate SMA
	fmt.Println("\nSimple Moving Average (SMA):")
	sma, err := indicators.SMA(close, 5)
	if err != nil {
		log.Fatal("Error calculating SMA:", err)
	}
	fmt.Printf("SMA Values: %v\n", sma.Values)

	// Calculate EMA
	fmt.Println("\nExponential Moving Average (EMA):")
	ema, err := indicators.EMA(close, 5)
	if err != nil {
		log.Fatal("Error calculating EMA:", err)
	}
	fmt.Printf("EMA Values: %v\n", ema.Values)

	// Calculate RSI
	fmt.Println("\nRelative Strength Index (RSI):")
	rsi, err := indicators.RSI(close, 14)
	if err != nil {
		log.Fatal("Error calculating RSI:", err)
	}
	fmt.Printf("RSI Values: %v\n", rsi.Values)

	// Calculate MACD
	fmt.Println("\nMoving Average Convergence Divergence (MACD):")
	macd, err := indicators.MACD(close, 12, 26, 9)
	if err != nil {
		log.Fatal("Error calculating MACD:", err)
	}
	fmt.Printf("MACD Line: %v\n", macd.Values)
	fmt.Printf("Signal Line: %v\n", macd.MACDSignal)
	fmt.Printf("Histogram: %v\n", macd.MACDHist)

	// Calculate Bollinger Bands
	fmt.Println("\nBollinger Bands:")
	bb, err := indicators.BBANDS(close, 20, 2.0, 2.0)
	if err != nil {
		log.Fatal("Error calculating Bollinger Bands:", err)
	}
	fmt.Printf("Middle Band: %v\n", bb.Values)
	fmt.Printf("Upper Band: %v\n", bb.UpperBand)
	fmt.Printf("Lower Band: %v\n", bb.LowerBand)

	// Calculate ATR
	fmt.Println("\nAverage True Range (ATR):")
	atr, err := indicators.ATR(high, low, close, 14)
	if err != nil {
		log.Fatal("Error calculating ATR:", err)
	}
	fmt.Printf("ATR Values: %v\n", atr.Values)

	// Calculate Stochastic Oscillator
	fmt.Println("\nStochastic Oscillator:")
	stoch, err := indicators.STOCH(high, low, close, 5, 3, 3)
	if err != nil {
		log.Fatal("Error calculating Stochastic:", err)
	}
	fmt.Printf("Slow K: %v\n", stoch.SlowK)
	fmt.Printf("Slow D: %v\n", stoch.SlowD)

	// Calculate OBV
	fmt.Println("\nOn Balance Volume (OBV):")
	obv, err := indicators.OBV(close, volume)
	if err != nil {
		log.Fatal("Error calculating OBV:", err)
	}
	fmt.Printf("OBV Values: %v\n", obv.Values)

	// Calculate ADX
	fmt.Println("\nAverage Directional Index (ADX):")
	adx, err := indicators.ADX(high, low, close, 14)
	if err != nil {
		log.Fatal("Error calculating ADX:", err)
	}
	fmt.Printf("ADX Values: %v\n", adx.Values)
	fmt.Printf("+DI Values: %v\n", adx.PlusDI)
	fmt.Printf("-DI Values: %v\n", adx.MinusDI)

	// Calculate CCI
	fmt.Println("\nCommodity Channel Index (CCI):")
	cci, err := indicators.CCI(high, low, close, 14)
	if err != nil {
		log.Fatal("Error calculating CCI:", err)
	}
	fmt.Printf("CCI Values: %v\n", cci.Values)

	// Calculate Williams %R
	fmt.Println("\nWilliams %R:")
	willr, err := indicators.WILLR(high, low, close, 14)
	if err != nil {
		log.Fatal("Error calculating Williams %R:", err)
	}
	fmt.Printf("Williams %%R Values: %v\n", willr)

	// Calculate MFI
	fmt.Println("\nMoney Flow Index (MFI):")
	mfi, err := indicators.MFI(high, low, close, volume, 14)
	if err != nil {
		log.Fatal("Error calculating MFI:", err)
	}
	fmt.Printf("MFI Values: %v\n", mfi)

	// Calculate ROC
	fmt.Println("\nRate of Change (ROC):")
	roc, err := indicators.ROC(close, 10)
	if err != nil {
		log.Fatal("Error calculating ROC:", err)
	}
	fmt.Printf("ROC Values: %v\n", roc)

	// Calculate StochRSI
	fmt.Println("\nStochastic RSI:")
	stochRsi, stochRsiSignal, err := indicators.STOCHRSI(close, 14, 14, 3, 3)
	if err != nil {
		log.Fatal("Error calculating StochRSI:", err)
	}
	fmt.Printf("StochRSI: %v\n", stochRsi)
	fmt.Printf("StochRSI Signal: %v\n", stochRsiSignal)

	// Calculate APO
	fmt.Println("\nAbsolute Price Oscillator (APO):")
	apo, err := indicators.APO(close, 12, 26)
	if err != nil {
		log.Fatal("Error calculating APO:", err)
	}
	fmt.Printf("APO Values: %v\n", apo.Values)
}
