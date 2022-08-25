package action

import (
	"time"
	"unicode"

	"github.com/rsteube/carapace"
	_time "github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/rsteube/carapace/pkg/style"
)

func ActionDate() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch(
			carapace.ActionValuesDescribed(
				"monday", DateForWeekday("monday").Format("01-02-2006"),
				"tuesday", DateForWeekday("tuesday").Format("01-02-2006"),
				"wednesday", DateForWeekday("wednesday").Format("01-02-2006"),
				"thursday", DateForWeekday("thursday").Format("01-02-2006"),
				"friday", DateForWeekday("friday").Format("01-02-2006"),
				"saturday", DateForWeekday("saturday").Format("01-02-2006"),
				"sunday", DateForWeekday("sunday").Format("01-02-2006"),
				"today", DateForWeekday("today").Format("01-02-2006"),
				"yesterday", DateForWeekday("yesterday").Format("01-02-2006"),
			).Style(style.Blue),
		)

		if len(c.CallbackValue) > 0 && unicode.IsDigit([]rune(c.CallbackValue)[0]) {
			batch = append(batch, _time.ActionDate())
		}

		return batch.ToA()
	})
}

func ActionDuration() carapace.Action {
	return carapace.ActionValues(
		"0.25", "0.5", "0.75", "1.0", "1.25", "1.5", "1.75", "2.0",
		"2.25", "2.5", "2.75", "3.0", "3.25", "3.5", "3.75", "4.0",
		"4.25", "4.5", "4.75", "5.0", "5.25", "5.5", "5.75", "6.0",
		"6.25", "6.5", "6.75", "7.0", "7.25", "7.5", "7.75", "8.0",
		"9.25", "9.5", "9.75",
	)
}

func DateForWeekday(s string) time.Time {
	today := time.Now()
	switch s {
	case "today":
		return today
	case "yesterday":
		return today.Add(-24 * time.Hour)
	case "monday":
		return LastWeekday(time.Monday)
	case "tuesday":
		return LastWeekday(time.Tuesday)
	case "wednesday":
		return LastWeekday(time.Wednesday)
	case "thursday":
		return LastWeekday(time.Thursday)
	case "friday":
		return LastWeekday(time.Friday)
	case "saturday":
		return LastWeekday(time.Saturday)
	case "sunday":
		return LastWeekday(time.Sunday)
	default:
		return today
	}
}

func LastWeekday(w time.Weekday) time.Time {
	d := time.Now()
	for i := 0; d.Weekday() != w && i < 7; i++ {
		d = d.Add(-24 * time.Hour)
	}
	return d
}
