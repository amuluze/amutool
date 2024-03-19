// Package task
// Date: 2023/7/19 15:30
// Author: Amu
// Description:
package task

import (
	"time"
)

const (
	DurationMS uint8 = iota + 1
	DurationMonth
	DurationYear
)

type Plan struct {
	startTime    *time.Time
	duration     int64
	durationType uint8
	lastTime     *time.Time
}

func NewPlan(duration int64, startTime *time.Time, durationType uint8) *Plan {
	if startTime == nil {
		nowTime := time.Now()
		startTime = &nowTime
	}
	return &Plan{
		startTime:    startTime,
		duration:     duration,
		durationType: durationType,
	}
}

func (p *Plan) GetFirstDuration() time.Duration {
	return p.getDuration()
}

func (p *Plan) getDuration() time.Duration {
	return time.Until(p.getNextTime())
}

func (p *Plan) getNextTime() (nextTime time.Time) {
	if p.startTime == nil {
		nextTime = p.getNetTimeWithTime(time.Now())
		return
	}

	if p.lastTime != nil {
		nextTime = p.getNetTimeWithTime(*p.lastTime)
	} else {
		nextTime = *p.startTime
	}

	nowTime := time.Now()
	for {
		if nextTime.Equal(nowTime) || nextTime.After(nowTime) {
			break
		}

		if p.duration > 0 {
			nextTime = p.getNetTimeWithTime(nextTime)
		} else {
			nextTime = nowTime
		}
	}

	p.lastTime = &nextTime
	return
}

func (p *Plan) getNetTimeWithTime(nowTime time.Time) (nextTime time.Time) {
	switch p.durationType {
	case DurationMonth:
		return AddMonth(nowTime, int(p.duration))
	case DurationYear:
		return nowTime.AddDate(int(p.duration), 0, 0)
	default:
		return nowTime.Add(time.Duration(p.duration) * time.Millisecond)
	}
}

var thirtyOneDayMonths = []time.Month{time.January, time.March, time.May, time.July, time.August, time.October, time.December}

func AddMonth(addTime time.Time, duration int) time.Time {
	year, month, day := addTime.Date()
	intMonth := int(month) + duration

	year += intMonth / 12
	month = time.Month(intMonth % 12)
	if day > 28 {
		if day == 29 {
			if month == time.February && (year%4) != 0 {
				month += 1
			}

		} else if day == 30 {
			if month == time.February {
				month += 1
			}
		} else {
			for _, thirtyOneDayMonth := range thirtyOneDayMonths {
				if month <= thirtyOneDayMonth {
					month = thirtyOneDayMonth
					break
				}
			}
		}
	}

	return time.Date(year, month, day, addTime.Hour(), addTime.Minute(), addTime.Second(), addTime.Nanosecond(), addTime.Location())
}
