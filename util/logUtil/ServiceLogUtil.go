package util

import (
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

func ServiceLoggerInfo(clientId string, lastPayment time.Time, totalAmount float64, message string) {
	var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.Info().
		Str("class", "SERVICE").
		Str("client", clientId).
		Str("lastPayment", lastPayment.String()).
		Str("totalAmount", strconv.Itoa(int(totalAmount))).
		Msg(message)
}

func ServiceLoggerError(clientId string, amountRequest float64, totalAmount float64, message string) {
	var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.Error().
		Str("class", "SERVICE").
		Str("amountRequest", clientId).
		Str("lastPayment", strconv.Itoa(int(amountRequest))).
		Str("totalAmount", strconv.Itoa(int(totalAmount))).
		Msg(message)
}
