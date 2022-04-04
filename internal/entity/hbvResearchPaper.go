package entity

import (
	"cloud.google.com/go/firestore"
	"time"
)

type HbvResearchPaperEntity struct {
	Description      string                 `firestore:"description"`
	EndDate          time.Time              `firestore:"endDate"`
	MediaUrl         string                 `firestore:"mediaUrl"`
	Name             string                 `firestore:"name"`
	Organization     *firestore.DocumentRef `firestore:"organization"`
	Responsibilities []string               `firestore:"responsibilities"`
	Skills           []string               `firestore:"skills"`
	StartDate        time.Time              `firestore:"startDate"`
	PaperUrl         string                 `firestore:"url"`
}
