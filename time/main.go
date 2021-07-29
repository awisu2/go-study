package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	sampleGet()
	sampleCalc()
	sampleStrToTime()
	sampleLocation()
}

func sampleGet() {
	fmt.Println("---------- sampleGet")
	t := time.Now()
	fmt.Println(t.Year()) // 2021
	fmt.Println(t.YearDay()) // 210
	fmt.Println(t.Month()) // July
	fmt.Println(t.Weekday()) // Thursday
	fmt.Println(t.Day()) //29
	fmt.Println(t.Hour()) // 8
	fmt.Println(t.Minute()) // 28
	fmt.Println(t.Second()) // 30
	fmt.Println(t.Nanosecond()) // 701755200
	fmt.Println(t.Zone()) // JST 32400

	fmt.Println(t.Format("2006-01-02T15:04:05Z07:00"))
	fmt.Println(t.Format(time.RFC3339))
}

func sampleCalc() {
	fmt.Println("---------- sampleCalc")
	t := time.Now()
	durationS := time.Second * 30
	durationH := time.Hour * 30
	durationMc := time.Microsecond * 30
	durationMi := time.Millisecond * 30
	durationM := time.Minute * 30
	durationN := time.Nanosecond * 30

	duration := durationS + durationH + durationMc + durationMi + durationM + durationN
	tInc := t.Add(duration)
	tDec := t.Add(-duration)
	fmt.Println(tInc, tDec)
	// 引数よりも前であるか
	if !tInc.Before(tDec) {
		fmt.Println("tInc is not before tDec")
	}
	// 引数よりも後であるか
	if tInc.After(tDec) {
		fmt.Println("tInc is after tDec")
	}
	// = ではない
	if !(tInc.After(tInc) || tInc.Before(tInc)){
		fmt.Println("same time is not hit befor after")
	}

	// locationが異なってもOK
	tLoc := tInc.In(time.UTC)
	if tInc.Equal(tLoc) {
		fmt.Println("equal time")
	}
}

func sampleLocation() {
	fmt.Println("---------- sampleLocation")
	t, _ := StrToTime("2000/12/23T12:34:56+09:00", nil)
	// 1234-12-23 12:34:56 +0900
	fmt.Println(t)

	// UTCにする
	// 見た目上の時間はlocationにあわせて調整される
	// 1234-12-23 03:34:56 +0000 UTC
	t = t.UTC()
	fmt.Println("UTC:", t)

	// 単にlocationを変換
	// 見た目上の時間はlocationにあわせて調整される
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err == nil {
		fmt.Println("location:", loc)
		tIn := t.In(loc)
		fmt.Printf(`In: "%v" > "%v"` + "\n", t, tIn)
	}

	// 自前でlocationを作る というかFixなのでlocationを修正するといったほうが正しそう
	loc = time.FixedZone("JPT", 9*60*60)
	fmt.Println("location:", loc)
	tIn := t.In(loc)
	fmt.Println("FixedZone:", tIn)
	fmt.Println(tIn.Zone())


	// defaultのlocationを変更
	if loc, err := time.LoadLocation("Africa/Cairo"); err == nil {
		time.Local = loc
		fmt.Println("Local Africa/Cairo", time.Now())
	}

	if loc, err := time.LoadLocation("Asia/Tokyo"); err == nil {
		time.Local = loc
		fmt.Println("Local Asia/Tokyo", time.Now())

		// parseを行うと、localは無視されてUTCになる
		// StrToTimeの内部処理はparseを行うとlocalは強制的にUTCになる
		t, err := time.Parse("2006-01-02T15:04:05", "1234-12-23T12:34:56")
		fmt.Println("Local Asia/Tokyo on parse:", t, err) // 1234-12-23 12:34:56 +0000 UTC <nil>

		// ParseInLocationを利用することで、localを反映できる
		t, err = time.ParseInLocation("2006-01-02T15:04:05", "1234-12-23T12:34:56", time.Local)
		fmt.Println("Local Asia/Tokyo on parseInLocation:", t, err) // 1234-12-23 12:34:56 +0918 LMT <nil>

		// zone情報がある場合はそちらが優先される
		t, err = time.ParseInLocation("2006-01-02T15:04:05-07:00", "1234-12-23T12:34:56+02:00", time.Local)
		fmt.Println("Local Asia/Tokyo on parseInLocation with zone:", t, err) // 1234-12-23 12:34:56 +0200 +0200 <nil>
	}
}

func sampleStrToTime() {
	fmt.Println("---------- sampleStrToTime")
	fmt.Println(StrToTime("1234/12/23T12:34:56+09:00", nil))
	fmt.Println(StrToTime("1234/12/23T12:34:56-07:00", nil)) 
	fmt.Println(StrToTime("1234/12/23T12:34:56+0900", nil)) 
	fmt.Println(StrToTime("1234/12/23T12:34:56 +09:00", nil)) 
	fmt.Println(StrToTime("1234-12-23T12:34:56+0900", nil)) 
	fmt.Println(StrToTime("1234/12/23T12:34:56 +09", nil)) 

	fmt.Println(StrToTime("  1234/12/23T12:34:56 +09:00  ", nil)) 

	fmt.Println(StrToTime("1234/12/23T12:34:56", nil)) 


	fmt.Println(StrToTime("1234/12/23", nil)) 
	fmt.Println(StrToTime("1234/12", nil)) 

	fmt.Println(StrToTime("12:34:56", nil)) 
	fmt.Println(StrToTime("12:34", nil)) 
}

// 文字列をtimeへ変換
var zoneRe = regexp.MustCompile(`[+-][0-9:]*$`)
var spacesRe = regexp.MustCompile(` +`)

type StrToTimeOption struct {
	loc *time.Location
}
// convert string to time.Time
func StrToTime(s string, opt *StrToTimeOption) (time.Time, error) {
	// optから値を取得
	loc := time.Local
	if opt != nil {
		if opt.loc != nil {
			loc = opt.loc
		}
	}

	// チェックしやすいように文字変換
	s = strings.TrimSpace(s)

	// zone情報の取得
	zone := string(zoneRe.FindString(s))
	if zone != "" {
		s = strings.Replace(s, zone, "", -1)
		s = strings.TrimSpace(s)
	}

	s = strings.Replace(s, "/", "-", -1)
	if strings.Contains(s, "T") {
		s = spacesRe.ReplaceAllString(s, "")
	} else {
		s = spacesRe.ReplaceAllString(s, "T")
	}

	layout := ""

	// date
	n := strings.Count(s, "-")
	
	switch n {
	case 2:
		layout += "2006-01-02"
	case 1:
		layout += "2006-01"
	}

	// sepalater
	if strings.Contains(s, "T") {
		layout += "T"
	}

	// time
	n = strings.Count(s, ":")
	switch n {
	case 2:
		layout += "15:04:05"
	case 1:
		layout += "15:04"
	}

	// zone
	if zone != "" {
		s += " " + zone
		if len(zone) == 3 {
			layout += " Z07"
		} else {
			if strings.Contains(zone, ":") {
				layout += " Z07:00"
			} else {
				layout += " Z0700"
			}
		}
	}

	// fmt.Println(layout, s)
	return time.ParseInLocation(layout, s, loc)
}