package entity

type HolisticOfficeLinkEntity struct {
	Name string `firestore:"name"`
	Type string `firestore:"type"`
	Url  string `firestore:"url"`
}
