package models

import (
	"strings"
)

type PriorityModel string

const (
	HighPriority   PriorityModel = "High"
	MediumPriority PriorityModel = "Medium"
	LowPriority    PriorityModel = "Low"
)

func GetPriority(s string) PriorityModel {
	switch strings.ToLower(s) {
	case "high":
		return HighPriority
	case "medium":
		return MediumPriority
	case "low":
		return LowPriority
	default:
		return "low"
	}
}
