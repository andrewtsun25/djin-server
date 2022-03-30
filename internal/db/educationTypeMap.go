package db

import grpcEntity "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/entity/v1"

// TODO: delete these bridge maps once djin-server is deployed
var EducationDbTypeToProtoMap = map[string]string{
	"coding": grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_CODING)],
	"music":  grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_MUSIC)],
	"":       grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_UNSPECIFIED)],
}

var EducationProtoToDbTypeMap = map[string]string{
	grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_CODING)]:      "coding",
	grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_MUSIC)]:       "music",
	grpcEntity.EducationType_name[int32(grpcEntity.EducationType_EDUCATION_TYPE_UNSPECIFIED)]: "",
}
