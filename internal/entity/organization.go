package entity

type Organization struct {
	Name    string `firestore:"name"`
	LogoUrl string `firestore:"logoUrl"`
}
