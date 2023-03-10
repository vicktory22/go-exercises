package expenses

import "errors"

type Record struct {
	Day      int
	Amount   float64
	Category string
}

type DaysPeriod struct {
	From int
	To   int
}

func Filter(in []Record, predicate func(Record) bool) []Record {
	filteredRecords := []Record{}

	for _, record := range in {
		if !predicate(record) {
			continue
		}

		filteredRecords = append(filteredRecords, record)
	}

	return filteredRecords
}

func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		return p.From <= record.Day && p.To >= record.Day
	}
}

func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
		return c == record.Category
	}
}

func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	filteredRecords := Filter(in, ByDaysPeriod(p))
	totalAmount := 0.0

	for _, record := range filteredRecords {
		totalAmount += record.Amount
	}

	return totalAmount
}

func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	byCategory := Filter(in, ByCategory(c))

	if len(byCategory) == 0 {
		return 0.0, errors.New("not a category")
	}

	byCategoryAndPeriod := Filter(byCategory, ByDaysPeriod(p))

	if len(byCategoryAndPeriod) == 0 {
		return 0.0, nil
	}

	totalAmount := 0.0

	for _, record := range byCategoryAndPeriod {
		totalAmount += record.Amount
	}

	return totalAmount, nil
}
