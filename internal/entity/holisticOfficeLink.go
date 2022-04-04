package entity

type HolisticOfficeLink struct {
	Name string `firestore:"name"`
	Type string `firestore:"type"`
	Url  string `firestore:"url"`
}
