package entity

type OrganizationEntity struct {
	Name    string `firestore:"name"`
	LogoUrl string `firestore:"logoUrl"`
}
