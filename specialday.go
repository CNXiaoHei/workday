package workday

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type SpecialDay struct {
	Date string
	Work bool
	Desc string
}

func GetSpecialDayMap() map[string]SpecialDay {
	specialDayMap := make(map[string]SpecialDay)

	file, err := os.OpenFile("D:/Projects/Go/workday/special_days.csv", os.O_RDONLY, 0666)
	if err != nil {
		return specialDayMap
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF {
			break
		}
		if err != nil || line == "" {
			break
		}
		spl := strings.Split(line, ",")
		s := SpecialDay{}
		s.Date = spl[0]
		s.Desc = spl[1]
		if spl[2] == "1" {
			s.Work = false
		} else {
			s.Work = true
		}
		specialDayMap[s.Date] = s
	}
	return specialDayMap
}
