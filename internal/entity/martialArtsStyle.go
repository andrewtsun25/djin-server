package entity

import "cloud.google.com/go/firestore"

type MartialArtsStyleEntity struct {
	Biography     []string                 `firestore:"biography"`
	BlackBeltRank int64                    `firestore:"blackBeltRank"`
	Description   string                   `firestore:"description"`
	LogoUrl       string                   `firestore:"logoUrl"`
	MediaCaption  string                   `firestore:"mediaCaption"`
	MediaUrl      string                   `firestore:"mediaUrl"`
	Name          string                   `firestore:"name"`
	Studios       []*firestore.DocumentRef `firestore:"studios"`
	Type          string                   `firestore:"type"`
}
