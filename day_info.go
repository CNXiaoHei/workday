package workday

import "time"

type DateStore interface {
	Search(date string) DayInfo
}

type MemoryDateStore struct {
	store map[string]SpecialDay
}

func NewMemoryDateStore() *MemoryDateStore {
	m := &MemoryDateStore{}
	m.store = GetSpecialDayMap()
	return m
}

func (m *MemoryDateStore) Search(date string) DayInfo {
	dayInfo := DayInfo{}
	parse, err := time.Parse("2006-01-02", date)
	if err != nil {
		return dayInfo
	}

	specialDay, ok := m.store[date]
	dayInfo.Date = date
	if ok {
		dayInfo.Weekday = parse.Weekday().String()
		dayInfo.Work = specialDay.Work
		dayInfo.Desc = specialDay.Desc
	} else {
		switch parse.Weekday() {
		case time.Monday:
			dayInfo.Weekday = time.Monday.String()
			dayInfo.Work = true
			dayInfo.Desc = "星期一"
		case time.Tuesday:
			dayInfo.Weekday = time.Tuesday.String()
			dayInfo.Work = true
			dayInfo.Desc = "星期二"
		case time.Wednesday:
			dayInfo.Weekday = time.Wednesday.String()
			dayInfo.Work = true
			dayInfo.Desc = "星期三"
		case time.Thursday:
			dayInfo.Weekday = time.Thursday.String()
			dayInfo.Work = true
			dayInfo.Desc = "星期四"
		case time.Friday:
			dayInfo.Weekday = time.Friday.String()
			dayInfo.Work = true
			dayInfo.Desc = "星期四"
		case time.Saturday:
			dayInfo.Weekday = time.Saturday.String()
			dayInfo.Work = false
			dayInfo.Desc = "星期六"
		case time.Sunday:
			dayInfo.Weekday = time.Sunday.String()
			dayInfo.Work = false
			dayInfo.Desc = "星期日"
		}
	}
	return dayInfo
}

type DayInfo struct {
	Date    string
	Weekday string
	Work    bool
	Desc    string
}

func GetDayInfo(date string) DayInfo {
	dayInfo := DayInfo{}
	specialDayMap := GetSpecialDayMap()
	specialDay, ok := specialDayMap[date]

	parse, err := time.Parse("2006-01-02", date)
	if err != nil {
		return dayInfo
	}
	dayInfo.Date = date
	if ok {
		dayInfo.Weekday = parse.Weekday().String()
		dayInfo.Work = specialDay.Work
		dayInfo.Desc = specialDay.Desc
	} else {
		switch parse.Weekday() {
		case time.Monday:
			dayInfo.Weekday = time.Monday.String()
			dayInfo.Work = true
			dayInfo.Desc = "星期一"
		case time.Tuesday:
			dayInfo.Weekday = time.Tuesday.String()
			dayInfo.Work = true
			dayInfo.Desc = "星期二"
		case time.Wednesday:
			dayInfo.Weekday = time.Wednesday.String()
			dayInfo.Work = true
			dayInfo.Desc = "星期三"
		case time.Thursday:
			dayInfo.Weekday = time.Thursday.String()
			dayInfo.Work = true
			dayInfo.Desc = "星期四"
		case time.Friday:
			dayInfo.Weekday = time.Friday.String()
			dayInfo.Work = true
			dayInfo.Desc = "星期四"
		case time.Saturday:
			dayInfo.Weekday = time.Saturday.String()
			dayInfo.Work = false
			dayInfo.Desc = "星期六"
		case time.Sunday:
			dayInfo.Weekday = time.Sunday.String()
			dayInfo.Work = false
			dayInfo.Desc = "星期日"
		}
	}
	return dayInfo
}
