package service

import (
	"strconv"
	"strings"
)

func parseMonthYear(date string) int {
	parts := strings.Split(date, "-")
	if len(parts) != 2 {
		return 0
	}
	m, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return y*12 + m
}

func CalculateOverlap(subStart, subEnd, reqStart, reqEnd string) int {
	startS := parseMonthYear(subStart)
	endS := 999999
	if subEnd != "" {
		endS = parseMonthYear(subEnd)
	}

	startR := parseMonthYear(reqStart)
	endR := parseMonthYear(reqEnd)

	overlapStart := startS
	if startR > startS {
		overlapStart = startR
	}

	overlapEnd := endS
	if endR < endS {
		overlapEnd = endR
	}

	if overlapStart > overlapEnd {
		return 0
	}
	return overlapEnd - overlapStart + 1
}
