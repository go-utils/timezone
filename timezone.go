package timezone

import "time"

// LocationJST - returns jst location
func LocationJST() *time.Location {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		jst = time.FixedZone("Asia/Tokyo", 9*60*60)
	}

	return jst
}

// LocationUTC - returns the utc location
func LocationUTC() *time.Location {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		utc = time.FixedZone("UTC", 0)
	}

	return utc
}

// InUTC - returns time with utc converted arguments
func InUTC(t time.Time) time.Time {
	return t.In(LocationUTC())
}

// InJST - returns time with jst conversion of arguments
func InJST(t time.Time) time.Time {
	return t.In(LocationJST())
}

// NowUTC - returns the current date and time converted to utc
func NowUTC() time.Time {
	return InUTC(time.Now())
}

// NowJST - returns the current date and time converted to jst
func NowJST() time.Time {
	return InJST(time.Now())
}
