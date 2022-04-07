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
	HbvResearchPapersCollection       = "hbvResearchPapers"
	HolisticOfficeLinksCollection     = "holisticOfficeLinks"
	HolisticOfficeModulesCollection   = "holisticOfficeModules"
	MartialArtsStudiosCollection      = "martialArtsStudios"
	MartialArtsStylesCollection       = "martialArtsStyles"
	MusicInstrumentsCollection        = "musicInstruments"
	MusicScoresCollection             = "musicScores"
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
			Description:      dbEmployment.Description,
			Domains:          dbEmployment.Domains,
			EndDate:          timestamppb.New(dbEmployment.EndDate),
			MediaUrl:         dbEmployment.MediaUrl,
			Organization:     organization,
			Responsibilities: dbEmployment.Responsibilities,
			Role:             dbEmployment.Role,
			Skills:           dbEmployment.Skills,
			SkillTypes:       dbEmployment.SkillTypes,
			StartDate:        timestamppb.New(dbEmployment.StartDate),
			Type:             grpcEntity.Employment_JobType(grpcEntity.Employment_JobType_value[EmploymentDbTypeToProtoMap[dbEmployment.JobType]]),
		}
		grpcEmployments = append(grpcEmployments, grpcEmployment)
	}
	return grpcEmployments, nil
}

// ListHBVResearchPapers

func (f *FirestoreDB) ListHbvResearchPapers(ctx context.Context) ([]*grpcEntity.HbvResearchPaper, error) {
	var grpcHbvResearchPapers []*grpcEntity.HbvResearchPaper
	iter := f.client.Collection(HbvResearchPapersCollection).Documents(ctx)
	for {
		hbvResearchDoc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		dbHbvResearchPaper := &dbEntity.HbvResearchPaperEntity{}
		if err = hbvResearchDoc.DataTo(dbHbvResearchPaper); err != nil {
			return []*grpcEntity.HbvResearchPaper{}, err
		}
		organization, err := f.GetOrganizationById(ctx, dbHbvResearchPaper.Organization.ID)
		if err != nil {
			return []*grpcEntity.HbvResearchPaper{}, err
		}
		grpcHbvResearchPaper := &grpcEntity.HbvResearchPaper{
			Id:               hbvResearchDoc.Ref.ID,
			Description:      dbHbvResearchPaper.Description,
			EndDate:          timestamppb.New(dbHbvResearchPaper.EndDate),
			MediaUrl:         dbHbvResearchPaper.MediaUrl,
			Name:             dbHbvResearchPaper.Name,
			Organization:     organization,
			PaperUrl:         dbHbvResearchPaper.PaperUrl,
			Responsibilities: dbHbvResearchPaper.Responsibilities,
			Skills:           dbHbvResearchPaper.Skills,
			StartDate:        timestamppb.New(dbHbvResearchPaper.StartDate),
		}
		grpcHbvResearchPapers = append(grpcHbvResearchPapers, grpcHbvResearchPaper)
	}
	return grpcHbvResearchPapers, nil
}

// Holistic Office

func (f *FirestoreDB) ListHolisticOfficeLinks(ctx context.Context) ([]*grpcEntity.HolisticOfficeLink, error) {
	var grpcHolisticOfficeLinks []*grpcEntity.HolisticOfficeLink
	iter := f.client.Collection(HolisticOfficeLinksCollection).Documents(ctx)
	for {
		holisticOfficeLinkDoc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		dbHolisticOfficeLink := &dbEntity.HolisticOfficeLinkEntity{}
		if err = holisticOfficeLinkDoc.DataTo(dbHolisticOfficeLink); err != nil {
			return []*grpcEntity.HolisticOfficeLink{}, err
		}
		grpcHolisticOfficeLink := &grpcEntity.HolisticOfficeLink{
			Id:   holisticOfficeLinkDoc.Ref.ID,
			Name: dbHolisticOfficeLink.Name,
			Type: grpcEntity.HolisticOfficeLink_HolisticOfficeLinkType(grpcEntity.HolisticOfficeLink_HolisticOfficeLinkType_value[HolisticOffficeLinkDbTypeToProtoMap[dbHolisticOfficeLink.Type]]),
			Url:  dbHolisticOfficeLink.Url,
		}
		grpcHolisticOfficeLinks = append(grpcHolisticOfficeLinks, grpcHolisticOfficeLink)
	}
	return grpcHolisticOfficeLinks, nil
}

func (f *FirestoreDB) ListHolisticOfficeModules(ctx context.Context) ([]*grpcEntity.HolisticOfficeModule, error) {
	var grpcHolisticOfficeModules []*grpcEntity.HolisticOfficeModule
	iter := f.client.Collection(HolisticOfficeModulesCollection).Documents(ctx)
	for {
		holisticOfficeModuleDoc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		dbHolisticOfficeModule := &dbEntity.HolisticOfficeModuleEntity{}
		if err = holisticOfficeModuleDoc.DataTo(dbHolisticOfficeModule); err != nil {
			return []*grpcEntity.HolisticOfficeModule{}, err
		}
		grpcHolisticOfficeModule := &grpcEntity.HolisticOfficeModule{
			Id:         holisticOfficeModuleDoc.Ref.ID,
			Components: dbHolisticOfficeModule.Components,
			Name:       dbHolisticOfficeModule.Name,
		}
		grpcHolisticOfficeModules = append(grpcHolisticOfficeModules, grpcHolisticOfficeModule)
	}
	return grpcHolisticOfficeModules, nil
}

// Martial Arts

func (f *FirestoreDB) GetMartialArtsStyleById(ctx context.Context, id string) (*grpcEntity.MartialArtsStyle, error) {
	martialArtsStyleDoc, err := f.client.Collection(MartialArtsStylesCollection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	dbMartialArtsStyle := &dbEntity.MartialArtsStyleEntity{}
	if err = martialArtsStyleDoc.DataTo(dbMartialArtsStyle); err != nil {
		return nil, err
	}
	var studios []*grpcEntity.MartialArtsStudio
	for _, studioRef := range dbMartialArtsStyle.Studios {
		studioId := studioRef.ID
		studio, err := f.GetMartialArtsStudioById(ctx, studioId)
		if err != nil {
			return nil, err
		}
		studios = append(studios, studio)
	}
	grpcMartialArtsStyle := &grpcEntity.MartialArtsStyle{
		Id:                  martialArtsStyleDoc.Ref.ID,
		BiographyParagraphs: dbMartialArtsStyle.Biography,
		BlackBeltRank:       dbMartialArtsStyle.BlackBeltRank,
		Description:         dbMartialArtsStyle.Description,
		LogoUrl:             dbMartialArtsStyle.LogoUrl,
		MediaCaption:        dbMartialArtsStyle.MediaCaption,
		MediaUrl:            dbMartialArtsStyle.MediaUrl,
		Name:                dbMartialArtsStyle.Name,
		Studios:             studios,
		Type:                grpcEntity.MartialArtsStyle_MartialArtsStyleType(grpcEntity.MartialArtsStyle_MartialArtsStyleType_value[MartialArtsStyleTypeDbTypeToProtoMap[dbMartialArtsStyle.Type]]),
	}
	return grpcMartialArtsStyle, nil
}

func (f *FirestoreDB) GetMartialArtsStudioById(ctx context.Context, id string) (*grpcEntity.MartialArtsStudio, error) {
	martialArtsStudioDoc, err := f.client.Collection(MartialArtsStudiosCollection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	dbMartialArtsStudio := &dbEntity.MartialArtsStudioEntity{}
	if err = martialArtsStudioDoc.DataTo(dbMartialArtsStudio); err != nil {
		return nil, err
	}
	grpcMartialArtsStudio := &grpcEntity.MartialArtsStudio{
		Id:        "",
		Name:      dbMartialArtsStudio.Name,
		LogoUrl:   dbMartialArtsStudio.LogoUrl,
		StudioUrl: dbMartialArtsStudio.StudioUrl,
		City:      dbMartialArtsStudio.City,
		JoinDate:  timestamppb.New(dbMartialArtsStudio.JoinDate),
	}
	return grpcMartialArtsStudio, nil
}

// Music

func (f *FirestoreDB) ListMusicInstruments(ctx context.Context) ([]*grpcEntity.Instrument, error) {
	var grpcInstruments []*grpcEntity.Instrument
	iter := f.client.Collection(MusicInstrumentsCollection).Documents(ctx)
	for {
		musicInstrumentDoc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		dbMusicInstrument := &dbEntity.MusicInstrumentEntity{}
		if err = musicInstrumentDoc.DataTo(dbMusicInstrument); err != nil {
			return []*grpcEntity.Instrument{}, err
		}
		grpcInstrument := &grpcEntity.Instrument{
			Id:       musicInstrumentDoc.Ref.ID,
			MediaUrl: dbMusicInstrument.MediaUrl,
			Name:     dbMusicInstrument.Name,
		}
		grpcInstruments = append(grpcInstruments, grpcInstrument)
	}
	return grpcInstruments, nil
}

func (f *FirestoreDB) ListMusicScores(ctx context.Context) ([]*grpcEntity.MusicScore, error) {
	var grpcMusicScores []*grpcEntity.MusicScore
	iter := f.client.Collection(MusicScoresCollection).Documents(ctx)
	for {
		musicScoreDoc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		dbMusicScore := &dbEntity.MusicScoreEntity{}
		if err = musicScoreDoc.DataTo(dbMusicScore); err != nil {
			return []*grpcEntity.MusicScore{}, err
		}
		grpcMusicScore := &grpcEntity.MusicScore{
			Id:       musicScoreDoc.Ref.ID,
			Date:     timestamppb.New(dbMusicScore.Date),
			Name:     dbMusicScore.Name,
			Sections: dbMusicScore.Sections,
			TrackUrl: dbMusicScore.TrackUrl,
		}
		grpcMusicScores = append(grpcMusicScores, grpcMusicScore)
	}
	return grpcMusicScores, nil
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
		LogoUrl: nil,
		Name:    dbOrganization.Name,
	}
	if dbOrganization.LogoUrl != "" {
		grpcOrganization.LogoUrl = wrapperspb.String(dbOrganization.LogoUrl)
	}
	return grpcOrganization, nil
}
