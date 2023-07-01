package lvn

import "time"

type lvnTime time.Time

func Time(t time.Time) lvnTime {
	return lvnTime(t)
}

func TimeInTashkent() time.Time {
	return time.Now().In(time.FixedZone("Tashkent", 5*60*60))
}

// returns start of the day in the time's location
func (lt lvnTime) StartOfTheDay() time.Time {
	t := time.Time(lt)
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// returns end of the day, without nanoseconds (23:59:59.000)
func (lt lvnTime) EndOfTheDay() time.Time {
	t := time.Time(lt)
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}

// returns start of the week in the time's location
func (lt lvnTime) StartOfTheWeek() time.Time {
	t := time.Time(lt)
	year, month, day := t.Date()
	return time.Date(year, month, day-int(t.Weekday())+1, 0, 0, 0, 0, t.Location())
}

// returns end of the week, without nanoseconds (23:59:59.000)
func (lt lvnTime) EndOfTheWeek() time.Time {
	t := time.Time(lt)
	year, month, day := t.Date()
	return time.Date(year, month, day-int(t.Weekday())+7, 23, 59, 59, 0, t.Location())
}

// returns start of the month in the time's location
func (lt lvnTime) StartOfTheMonth() time.Time {
	t := time.Time(lt)
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

// returns end of the month, without nanoseconds (23:59:59.000)
func (lt lvnTime) EndOfTheMonth() time.Time {
	t := time.Time(lt)
	year, month, _ := t.Date()
	return time.Date(year, month+1, 0, 23, 59, 59, 0, t.Location())
}

// returns start of the year in the time's location
func (lt lvnTime) StartOfTheYear() time.Time {
	t := time.Time(lt)
	year, _, _ := t.Date()
	return time.Date(year, 1, 1, 0, 0, 0, 0, t.Location())
}

// returns end of the year, without nanoseconds (23:59:59.000)
func (lt lvnTime) EndOfTheYear() time.Time {
	t := time.Time(lt)
	year, _, _ := t.Date()
	return time.Date(year+1, 0, 31, 23, 59, 59, 0, t.Location())
}
