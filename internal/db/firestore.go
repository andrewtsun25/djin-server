package db

import (
	"cloud.google.com/go/firestore"
	"context"
	dbEntity "djin-server/internal/entity"
	firebase "firebase.google.com/go"
	grpcEntity "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/entity/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
)

const (
	EducationsCollection              = "educations"
	EmploymentsCollection             = "employments"
	OrganizationsCollection           = "organizations"
	StudentOrganizationsSubCollection = "studentOrganizations"
)

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

// Educations

func (f *FirestoreDB) ListEducationsByType(ctx context.Context, eduType grpcEntity.EducationType) ([]*grpcEntity.Education, error) {
	iter := f.client.Collection(EducationsCollection).Where("type", "==", EducationProtoToDbTypeMap[eduType.String()]).Documents(ctx)
	var grpcEducations []*grpcEntity.Education
	for {
		educationDoc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		dbEducation := &dbEntity.EducationEntity{}
		if err = educationDoc.DataTo(dbEducation); err != nil {
			return []*grpcEntity.Education{}, nil
		}
		organization, err := f.GetOrganizationById(ctx, dbEducation.Organization.ID)
		if err != nil {
			return []*grpcEntity.Education{}, nil
		}
		studentOrganizations, err := f.ListStudentOrganizationsForEducation(ctx, educationDoc.Ref.ID)
		if err != nil {
			return []*grpcEntity.Education{}, nil
		}
		grpcEducation := &grpcEntity.Education{
			Id:                   educationDoc.Ref.ID,
			Department:           nil,
			Description:          dbEducation.Description,
			EndDate:              timestamppb.New(dbEducation.EndDate),
			Gpa:                  dbEducation.GPA,
			Major:                dbEducation.Major,
			Minors:               dbEducation.Minors,
			Organization:         organization,
			ResidentialCollege:   nil,
			StartDate:            timestamppb.New(dbEducation.StartDate),
			StudentOrganizations: studentOrganizations,
			SyllabusUrls:         dbEducation.SyllabusUrls,
			Type:                 grpcEntity.EducationType(grpcEntity.EducationType_value[EducationDbTypeToProtoMap[dbEducation.Type]]),
		}
		if dbEducation.Department != "" {
			grpcEducation.Department = wrapperspb.String(dbEducation.Department)
		}
		if dbEducation.ResidentialCollege != "" {
			grpcEducation.ResidentialCollege = wrapperspb.String(dbEducation.ResidentialCollege)
		}
		grpcEducations = append(grpcEducations, grpcEducation)
	}
	return grpcEducations, nil
}

func (f *FirestoreDB) ListStudentOrganizationsForEducation(ctx context.Context, educationId string) ([]*grpcEntity.StudentOrganization, error) {
	var grpcStudentOrganizations []*grpcEntity.StudentOrganization
	iter := f.client.Collection(OrganizationsCollection).Doc(educationId).Collection(StudentOrganizationsSubCollection).Documents(ctx)
	for {
		studentOrganizationDoc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		dbStudentOrganization := &dbEntity.StudentOrganizationEntity{}
		if err = studentOrganizationDoc.DataTo(dbStudentOrganization); err != nil {
			return []*grpcEntity.StudentOrganization{}, err
		}
		grpcStudentOrganization := &grpcEntity.StudentOrganization{
			Id:          studentOrganizationDoc.Ref.ID,
			Biography:   dbStudentOrganization.Biography,
			Description: dbStudentOrganization.Description,
			Name:        dbStudentOrganization.Name,
		}
		grpcStudentOrganizations = append(grpcStudentOrganizations, grpcStudentOrganization)
	}
	return grpcStudentOrganizations, nil
}

// Employments

func (f *FirestoreDB) ListEmployments(ctx context.Context) ([]*grpcEntity.Employment, error) {
	var grpcEmployments []*grpcEntity.Employment
	iter := f.client.Collection(EmploymentsCollection).Documents(ctx)
	for {
		employmentDoc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		dbEmployment := &dbEntity.EmploymentEntity{}
		if err = employmentDoc.DataTo(dbEmployment); err != nil {
			return []*grpcEntity.Employment{}, err
		}
		organization, err := f.GetOrganizationById(ctx, dbEmployment.Organization.ID)
		if err != nil {
			return []*grpcEntity.Employment{}, err
		}
		grpcEmployment := &grpcEntity.Employment{
			Id:               employmentDoc.Ref.ID,
			Organization:     organization,
			MediaUrl:         dbEmployment.MediaUrl,
			Role:             dbEmployment.Role,
			StartDate:        timestamppb.New(dbEmployment.StartDate),
			EndDate:          timestamppb.New(dbEmployment.EndDate),
			Description:      dbEmployment.Description,
			Responsibilities: dbEmployment.Responsibilities,
			Skills:           dbEmployment.Skills,
			SkillTypes:       dbEmployment.SkillTypes,
			Domains:          dbEmployment.Domains,
			Type:             grpcEntity.Employment_JobType(grpcEntity.Employment_JobType_value[EmploymentDbTypeToProtoMap[dbEmployment.JobType]]),
		}
		grpcEmployments = append(grpcEmployments, grpcEmployment)
	}
	return grpcEmployments, nil
}

// Organizations

func (f *FirestoreDB) GetOrganizationById(ctx context.Context, id string) (*grpcEntity.Organization, error) {
	organizationDoc, err := f.client.Collection(OrganizationsCollection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	dbOrganization := &dbEntity.OrganizationEntity{}
	if err = organizationDoc.DataTo(dbOrganization); err != nil {
		return nil, err
	}
	grpcOrganization := &grpcEntity.Organization{
		Id:      organizationDoc.Ref.ID,
		Name:    dbOrganization.Name,
		LogoUrl: nil,
	}
	if dbOrganization.LogoUrl != "" {
		grpcOrganization.LogoUrl = wrapperspb.String(dbOrganization.LogoUrl)
	}
	return grpcOrganization, nil
}
