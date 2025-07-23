package formatteddates

import (
	"time"
)

func GetCurrencyFormattedDate() (string, error) {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return "", err
	}

	currencyDate := time.Now().In(loc)

	layout := "2006-01-02 15:04:05"

	parsedTime := currencyDate.Format(layout)

	return parsedTime, nil
}
