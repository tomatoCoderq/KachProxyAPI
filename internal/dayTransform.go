package internal

import (
	"strings"
)

const (
	January   = 1
	February  = 2
	March     = 3
	April     = 4
	May       = 5
	June      = 6
	July      = 7
	August    = 8
	September = 9
	October   = 10
	November  = 11
	December  = 12
)

func RemoveDay(fullDate string) int {
	if strings.HasPrefix(fullDate, "Янва") {
		return January
	} else if strings.HasPrefix(fullDate, "Февр") {
		return February
	} else if strings.HasPrefix(fullDate, "Март") {
		return March
	} else if strings.HasPrefix(fullDate, "Апре") {
		return April
	} else if strings.HasPrefix(fullDate, "Июнь") {
		return June
	} else if strings.HasPrefix(fullDate, "Июль") {
		return July
	} else if strings.HasPrefix(fullDate, "Авгу") {
		return August
	} else if strings.HasPrefix(fullDate, "Сент") {
		return September
	} else if strings.HasPrefix(fullDate, "Октя") {
		return October
	} else if strings.HasPrefix(fullDate, "Нояб") {
		return November
	} else if strings.HasPrefix(fullDate, "Дека") {
		return December
	} else {
		return May
	}
}

