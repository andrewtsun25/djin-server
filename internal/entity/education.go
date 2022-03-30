package entity

import (
	"cloud.google.com/go/firestore"
	"time"
)

type EducationEntity struct {
	Department         string                 `firebase:"department"`
	Description        string                 `firebase:"description"`
	EndDate            time.Time              `firebase:"endDate"`
	GPA                float64                `firebase:"gpa"`
	Major              string                 `firebase:"major"`
	Minors             []string               `firebase:"minor"`
	Organization       *firestore.DocumentRef `firebase:"organization"`
	ResidentialCollege string                 `firebase:"residentialCollege"`
	StartDate          time.Time              `firebase:"startDate"`
	SyllabusUrls       map[string]string      `firebase:"syllabusUrls"`
	Type               string                 `firebase:"type"`
}
