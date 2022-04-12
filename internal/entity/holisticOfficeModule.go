package entity

type HolisticOfficeModuleEntity struct {
	Components map[string]string `firestore:"components"`
	Name       string            `firestore:"name"`
}
