package db

import grpcEntity "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/entity/v1"

const (
	Education_Coding = "coding"
	Education_Music  = "music"
)

// TODO: delete these bridge maps once djin-server is deployed
var EducationDbTypeToProtoMap = map[string]string{
	Education_Coding: grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_CODING)],
	Education_Music:  grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_MUSIC)],
	"":               grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_UNSPECIFIED)],
}

var EducationProtoToDbTypeMap = map[string]string{
	grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_CODING)]:      Education_Coding,
	grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_MUSIC)]:       Education_Music,
	grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_UNSPECIFIED)]: "",
}
