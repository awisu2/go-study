package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	fmt.Println(StrToTime("1234/12/23T12:34:56+09:00")) 
	fmt.Println(StrToTime("1234/12/23T12:34:56-07:00")) 
	fmt.Println(StrToTime("1234/12/23T12:34:56+0900")) 
	fmt.Println(StrToTime("1234/12/23T12:34:56 +09:00")) 
	fmt.Println(StrToTime("1234-12-23T12:34:56+0900")) 

	fmt.Println(StrToTime("  1234/12/23T12:34:56 +09:00  ")) 

	fmt.Println(StrToTime("1234/12/23")) 
	fmt.Println(StrToTime("1234/12")) 

	fmt.Println(StrToTime("12:34:56")) 
	fmt.Println(StrToTime("12:34")) 
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