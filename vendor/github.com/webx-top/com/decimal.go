package com

import (
	"fmt"
	"math"
	"strings"
)

func NumberTrim(number string, precision int, separator ...string) string {
	s := number
	if len(s) == 0 {
		if precision <= 0 {
			return `0`
		}
		return `0.` + strings.Repeat(`0`, precision)
	}
	p := strings.LastIndex(s, `.`)
	if p < 0 {
		if precision > 0 {
			s += `.` + strings.Repeat(`0`, precision)
		}
		return numberWithSeparator(s, separator...)
	}
	if precision <= 0 {
		return numberWithSeparator(s[0:p], separator...)
	}
	r := s[p+1:]
	if len(r) >= precision {
		return numberWithSeparator(s[0:p]+`.`+r[0:precision], separator...)
	}
	return numberWithSeparator(s, separator...)
}

func numberWithSeparator(r string, separator ...string) string {
	d := `,`
	if len(separator) > 0 {
		d = separator[0]
		if len(d) == 0 {
			return r
		}
	}
	p := strings.LastIndex(r, `.`)
	var (
		i int
		v string
	)
	size := len(r)
	if p <= 0 {
		i = size
	} else {
		i = p
		v = r[i:]
	}
	j := int(math.Ceil(float64(i) / float64(3)))
	s := make([]string, j)
	for i > 0 && j > 0 {
		j--
		start := i - 3
		if start < 0 {
			start = 0
		}
		s[j] = r[start:i]
		i = start
	}
	return strings.Join(s, d) + v
}

func NumberFormat(number interface{}, precision int, separator ...string) string {
	r := fmt.Sprintf(`%.*f`, precision, Float64(number))
	return numberWithSeparator(r, separator...)
}
