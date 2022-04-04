package db

import grpcEntity "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/entity/v1"

const (
	Employment_Internship = "Employment_Internship"
	Employment_FullTime   = "Full-Time"
)

// TODO: delete this bridge map once djin-server is deployed
var EmploymentDbTypeToProtoMap = map[string]string{
	Employment_Internship: grpcEntity.Employment_JobType_name[int32(grpcEntity.Employment_JOB_TYPE_INTERNSHIP)],
	Employment_FullTime:   grpcEntity.Employment_JobType_name[int32(grpcEntity.Employment_JOB_TYPE_FULL_TIME)],
	"":                    grpcEntity.Employment_JobType_name[int32(grpcEntity.Employment_JOB_TYPE_UNSPECIFIED)],
}
