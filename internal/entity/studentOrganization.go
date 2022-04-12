package entity

type StudentOrganizationEntity struct {
	Biography   string `firestore:"biography"`
	Description string `firestore:"description"`
	Name        string `firestore:"name"`
}
