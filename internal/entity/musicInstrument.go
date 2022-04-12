package entity

type MusicInstrumentEntity struct {
	MediaUrl string `firestore:"mediaUrl"`
	Name     string `firestore:"name"`
	Type     string `firestore:"type"`
}
