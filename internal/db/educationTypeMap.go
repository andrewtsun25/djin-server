package db

import grpcEntity "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/entity/v1"

const (
	Coding = "coding"
	Music  = "music"
)

// TODO: delete these bridge maps once djin-server is deployed
var EducationDbTypeToProtoMap = map[string]string{
	Coding: grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_CODING)],
	Music:  grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_MUSIC)],
	"":     grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_UNSPECIFIED)],
}

var EducationProtoToDbTypeMap = map[string]string{
	grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_CODING)]:      Coding,
	grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_MUSIC)]:       Music,
	grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_UNSPECIFIED)]: "",
}
