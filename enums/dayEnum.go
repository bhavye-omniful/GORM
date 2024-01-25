package enums

type Day string
type DayType string

const (
	Monday    Day = "monday"
	Tuesday   Day = "tuesday"
	Wednesday Day = "wednesday"
	Thursday  Day = "thursday"
	Friday    Day = "friday"
	Saturday  Day = "saturday"
	Sunday    Day = "sunday"
)

const (
	Working DayType = "working"
	Holiday DayType = "holiday"
)

var StringToEnumDay = map[string]Day{
	"monday":    Monday,
	"tuesday":   Tuesday,
	"wednesday": Wednesday,
	"thursday":  Thursday,
	"friday":    Friday,
	"saturday":  Saturday,
	"sunday":    Sunday,
}

var stringToEnumDayType = map[string]DayType{
	"working": Working,
	"holiday": Holiday,
}
