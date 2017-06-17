package main

import (
	"fmt"
)

func main() {
	day_months := make(map[string]int)
	day_months["Jan"] = 31
	day_months["Feb"] = 28
	day_months["Mar"] = 31
	day_months["Apr"] = 30
	day_months["May"] = 31
	day_months["Jun"] = 30
	day_months["Jul"] = 31
	day_months["Aug"] = 31
	day_months["Sep"] = 30
	day_months["Oct"] = 31
	day_months["Nov"] = 30
	day_months["Dec"] = 31

	for month, days := range day_months {
		fmt.Printf("%s has %d days\n", month, days)
	}
}
