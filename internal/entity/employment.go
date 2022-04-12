package entity

import (
	"cloud.google.com/go/firestore"
	"time"
)

type EmploymentEntity struct {
	Description      string                 `firestore:"description"`
	Domains          []string               `firestore:"domains"`
	EndDate          time.Time              `firestore:"endDate"`
	JobType          string                 `firestore:"jobType"`
	MediaUrl         string                 `firestore:"mediaUrl"`
	Organization     *firestore.DocumentRef `firestore:"organization"`
	Responsibilities []string               `firestore:"responsibilities"`
	Role             string                 `firestore:"role"`
	SkillTypes       []string               `firestore:"skillTypes"`
	Skills           []string               `firestore:"skills"`
	StartDate        time.Time              `firestore:"startDate"`
}
