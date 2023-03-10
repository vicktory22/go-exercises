package booking

import "time"

func Schedule(date string) time.Time {
	time, err := time.Parse("1/2/2006 15:04:05", date)

	if err != nil {
		panic(err)
	}

	return time
}

func HasPassed(date string) bool {
	dateToCheck, err := time.Parse("January 2, 2006 15:04:05", date)

	if err != nil {
		panic(err)
	}

	return dateToCheck.Before(time.Now())
}

func IsAfternoonAppointment(date string) bool {
	dateToCheck, err := time.Parse("Monday, January 2, 2006 15:04:05", date)

	if err != nil {
		panic(err)
	}

	return dateToCheck.Hour() >= 12 && dateToCheck.Hour() < 18
}

func Description(date string) string {
	dateToCheck, err := time.Parse("1/2/2006 15:04:05", date)

	if err != nil {
		panic(err)
	}

	return "You have an appointment on " + dateToCheck.Format("Monday, January 2, 2006, at 15:04.")
}

func AnniversaryDate() time.Time {
	currDate := time.Now()

	anniversary, _ := time.Parse("1/2/2006", "9/15/"+currDate.Format("2006"))

	return anniversary
}
