package sumduration

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Execute(durations []string) string {
	var (
		sumDuration = make(map[string]int)
		totalSeconds int
	)

	for _, duration := range durations {
		hourMinSec := strings.Split(duration, ":")
		if len(hourMinSec) < 3 {
			hourMinSec = append([]string{"00"}, hourMinSec...)
		}

		for i, value := range hourMinSec {
			valueAsInt, err := strconv.Atoi(value)
			if err != nil {
				errors.New("error parsing value as integer")
			}

			switch i {
			case 0:
				totalSeconds += valueAsInt*60*60
			case 1:
				totalSeconds += valueAsInt*60
			case 2:
				totalSeconds += valueAsInt
			}
		}
	}

	hours := float64(totalSeconds)/3600
	minutes := (hours - math.Floor(hours)) * 60
	seconds := (minutes - math.Floor(minutes)) * 60

	sumDuration["hours"] = int(math.Floor(hours))
	sumDuration["minutes"] = int(math.Floor(minutes))
	sumDuration["seconds"] = int(math.Round(seconds))

	return fmt.Sprintf("%d:%02d:%02d", sumDuration["hours"], sumDuration["minutes"], sumDuration["seconds"])
}