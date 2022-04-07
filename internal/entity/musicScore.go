package entity

import "time"

type MusicScoreEntity struct {
	Date     time.Time         `firestore:"date"`
	Name     string            `firestore:"name"`
	Sections map[string]string `firestore:"sections"`
	TrackUrl string            `firestore:"trackUrl"`
}
