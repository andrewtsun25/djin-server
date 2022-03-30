package db

import (
	"cloud.google.com/go/firestore"
	"context"
	dbEntity "djin-server/internal/entity"
	firebase "firebase.google.com/go"
	grpcEntity "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/entity/v1"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
)

const OrganizationsCollection = "organizations"

type FirestoreDB struct {
	client *firestore.Client
}

func NewFirestoreDB(ctx context.Context, serviceAccount option.ClientOption) *FirestoreDB {
	app, err := firebase.NewApp(ctx, nil, serviceAccount)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return &FirestoreDB{
		client: client,
	}
}

func (f *FirestoreDB) Close() error {
	return f.client.Close()
}

// Organizations

func (f *FirestoreDB) GetOrganizationById(ctx context.Context, id string) (*grpcEntity.Organization, error) {
	docSnapshot, err := f.client.Collection(OrganizationsCollection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	dbOrganization := &dbEntity.Organization{}
	if err = docSnapshot.DataTo(dbOrganization); err != nil {
		return nil, err
	}
	grpcOrganization := &grpcEntity.Organization{
		Id:      id,
		Name:    dbOrganization.Name,
		LogoUrl: nil,
	}
	if dbOrganization.LogoUrl != "" {
		grpcOrganization.LogoUrl = wrapperspb.String(dbOrganization.LogoUrl)
	}
	return grpcOrganization, nil
}
