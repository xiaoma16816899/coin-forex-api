package cache

import "time"

const (
	AGE_30_MIN    = time.Hour / 2
	AGE_ONE_HOUR  = time.Hour
	AGE_FOUR_HOUR = time.Hour * 4
	AGE_ONE_DAY   = time.Hour * 24
	AGE_ONE_WEEK  = time.Hour * 24 * 7
	AGE_ONE_MONTH = time.Hour * 24 * 7 * 30
)

const (
	PrefixLock              = "lock:"
	PrefixReq               = "req:"
	PrefixLatestGameResult  = "lastGameResult:gameId:"
	PrefixPredictGameResult = "PredictGameResult:resultNo:"
)

const (
	KeyHotGames      = "hotGames"
	KeyScrapingGames = "scrapingGames"
)

const (
	ChannelGameAfterCreatedOrUpdated = "ChannelGameAfterCreatedOrUpdated"
)
