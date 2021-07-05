// Package dates 天数和周期数计算
package dates

const DaysInWeek = 7

// WeeksToDays 计算多少周有多少天
func WeeksToDays(weeks int) int {
	return weeks * DaysInWeek
}

// DaysToWeeks 将天数转成有多少周
func DaysToWeeks(days int) float64 {
	return float64(days) / float64(DaysInWeek)
}


