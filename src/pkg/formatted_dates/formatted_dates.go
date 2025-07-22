package formatteddates

import (
	"time"
)

func GetCurrencyFormattedDate() string {
	currencyDate := time.Now()

	layout := "2006-01-02 15:04:05"

	parsedTime := currencyDate.Format(layout)

	return parsedTime
}
