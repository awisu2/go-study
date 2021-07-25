package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	sampleStrToTime()
	sampleLocation()
}

func sampleStrToTime() {
	fmt.Println("---------- sampleStrToTime")
	fmt.Println(StrToTime("1234/12/23T12:34:56+09:00")) 
	fmt.Println(StrToTime("1234/12/23T12:34:56-07:00")) 
	fmt.Println(StrToTime("1234/12/23T12:34:56+0900")) 
	fmt.Println(StrToTime("1234/12/23T12:34:56 +09:00")) 
	fmt.Println(StrToTime("1234-12-23T12:34:56+0900")) 

	fmt.Println(StrToTime("  1234/12/23T12:34:56 +09:00  ")) 

	fmt.Println(StrToTime("1234/12/23T12:34:56")) 


	fmt.Println(StrToTime("1234/12/23")) 
	fmt.Println(StrToTime("1234/12")) 

	fmt.Println(StrToTime("12:34:56")) 
	fmt.Println(StrToTime("12:34")) 
}

func sampleLocation() {
	fmt.Println("---------- sampleLocation")
	t, _ := StrToTime("2000/12/23T12:34:56+09:00")
	// 1234-12-23 12:34:56 +0900
	fmt.Println(t)

	// UTCにする
	// 見た目上の時間はlocationにあわせて調整される
	// 1234-12-23 03:34:56 +0000 UTC
	t = t.UTC()
	fmt.Println(t)

	// 単にlocationを変換
	// 見た目上の時間は変わらない
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err == nil {
		fmt.Println(t.In(loc))
		fmt.Println(t)
	}

	// 自前でlocationを作る というかFixなのでlocationを修正するといったほうが正しそう
	loc = time.FixedZone("JPT", 9*60*60)
	fmt.Println(t.In(loc))

}

// 文字列をtimeへ変換
var zoneRe = regexp.MustCompile(`[+-][0-9:]*$`)
var spacesRe = regexp.MustCompile(` +`)

// convert string to time.Time
func StrToTime(s string) (time.Time, error) {
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
			layout += " -07"
		} else {
			if strings.Contains(zone, ":") {
				layout += " -07:00"
			} else {
				layout += " -0700"
			}
		}
	}

	// fmt.Println(layout, s)
	return time.Parse(layout, s)
}