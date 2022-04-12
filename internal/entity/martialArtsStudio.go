package entity

import "time"

type MartialArtsStudioEntity struct {
	City      string    `firestore:"city"`
	JoinDate  time.Time `firestore:"joinDate"`
	LogoUrl   string    `firestore:"logoUrl"`
	Name      string    `firestore:"name"`
	StudioUrl string    `firestore:"studioUrl"`
}
