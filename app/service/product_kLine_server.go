package service

import (
	"math"
	"math/rand"
	"time"
)

// Trade represents a single trade data point

type TradingData struct {
	Time   string  `json:"time"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
}

func generateFakeTradeData(numDataPoints int, speed time.Duration, targetPrice float64, currentPrice float64, fluctuation float64) []TradingData {
	// Use the current time as a starting point
	startTime := time.Now()

	// Calculate the time interval between consecutive data points
	timeInterval := speed

	// Create a slice to store trade data
	trades := make([]TradingData, numDataPoints)

	// Generate fake trade data
	for i := 0; i < numDataPoints; i++ {
		// Calculate the time for this data point
		tradeTime := startTime.Add(time.Duration(i) * timeInterval)

		// Calculate the opening price for this data point based on the target price
		openPrice := currentPrice + math.Min(targetPrice-currentPrice, (targetPrice-currentPrice)*(float64(i)/float64(numDataPoints)))

		// Calculate the price for this data point with fluctuation
		timeDiff := float64(i) * speed.Seconds() / 60 // Convert to minutes
		price := openPrice + fluctuation*math.Sin(2*math.Pi*timeDiff/60)

		// Generate a fake volume
		volume := rand.Float64() * 1000 // Assume volume varies between 0 and 1000

		// Create a Trade object and add it to the slice
		trades[i] = TradingData{
			Time:   tradeTime.Format("2006-01-02 15:04"), // Format time as "yyyy-mm-dd hh:mm"
			Open:   openPrice,
			High:   price + rand.Float64()*2,
			Low:    price - rand.Float64()*2,
			Close:  price + 0.01, // Ensure close price is different by 0.01
			Volume: volume,
		}
	}

	// Ensure the closing price of the last data point does not exceed the target price
	lastIndex := numDataPoints - 1
	if trades[lastIndex].Close > targetPrice {
		trades[lastIndex].Close = targetPrice - 0.01
	}

	return trades
}

func GenerateDataTradingView(fluctuation float64, speedMinutes int, targetPrice float64, currentPrice float64) []TradingData {
	// Parameters for generating fake trade data
	numDataPoints := 0
	if fluctuation > targetPrice {
		numDataPoints = rand.Intn(3) + 1
	} else {
		numDataPoints = rand.Intn(100) + 50
	}

	speed := time.Duration(speedMinutes) * time.Minute

	// Generate fake trade data
	trades := generateFakeTradeData(numDataPoints, speed, targetPrice, currentPrice, fluctuation)

	return trades
}

func GenerateTradingDataS(startDate time.Time, endDate time.Time, initialPrice float64, fluctuationPercentage float64, upProbability float64, downProbability float64) []TradingData {
	rand.Seed(time.Now().UnixNano())

	minutes := int(endDate.Sub(startDate).Minutes())
	data := make([]TradingData, minutes)

	currentPrice := initialPrice       // Initialize the current price
	currentVolume := rand.Int63n(1000) // Initial volume

	for i := 0; i < minutes; i++ {
		currentDate := startDate.Add(time.Duration(i) * time.Minute)
		data[i].Time = currentDate.Format("2006-01-02 15:04")

		// Calculate the fluctuation based on the current close price and the percentage
		fluctuation := currentPrice * fluctuationPercentage

		// Generate a random number between 0 and 1
		randomValue := rand.Float64()

		// Choose a sign for the fluctuation based on the specified probabilities
		sign := 1.0 // Default to up
		if randomValue < downProbability {
			sign = -1.0 // Down
		} else if randomValue >= (1 - upProbability) {
			sign = 1.0 // Up
		}

		// Calculate the new close price based on the fluctuation and the current close price
		close := currentPrice + (fluctuation * sign)

		// Set the open price as the current close price
		data[i].Open = currentPrice

		// Update low and high prices
		if close < currentPrice {
			data[i].Low = close
			data[i].High = currentPrice
		} else {
			data[i].Low = currentPrice
			data[i].High = close
		}

		// Assign calculated values to the trading data
		data[i].Close = close

		// Update the current price for the next iteration
		currentPrice = close

		// Generate a random volume change (up to 10% of the current volume)
		volumeChange := float64(currentVolume) * rand.Float64() * 0.1 // Up to 10% change
		if rand.Float64() < 0.5 {                                     // 50% chance for negative change
			volumeChange *= -1
		}
		currentVolume += int64(volumeChange)
		if currentVolume < 0 {
			currentVolume = 0
		}
		data[i].Volume = float64(currentVolume)
	}

	return data
}
