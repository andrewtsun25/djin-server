package entity

import (
	"cloud.google.com/go/firestore"
	"time"
)

type ProjectEntity struct {
	Description      string                 `firebase:"description"`
	Disclaimer       string                 `firebase:"disclaimer"`
	Domains          []string               `firebase:"domains"`
	EndDate          time.Time              `firebase:"endDate"`
	MediaUrl         string                 `firebase:"mediaUrl"`
	Name             string                 `firebase:"name"`
	Organization     *firestore.DocumentRef `firebase:"organization"`
	ProjectUrls      map[string]string      `firebase:"projectUrls"`
	Responsibilities []string               `firebase:"responsibilities"`
	SkillTypes       []string               `firebase:"skillTypes"`
	Skills           []string               `firebase:"skills"`
	StartDate        time.Time              `firebase:"startDate"`
}
