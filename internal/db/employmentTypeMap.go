package db

import grpcEntity "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/entity/v1"

const (
	Internship = "Internship"
	FullTime   = "Full-Time"
)

// TODO: delete this bridge map once djin-server is deployed
var EmploymentDbTypeToProtoMap = map[string]string{
	Internship: grpcEntity.Employment_JobType_name[int32(grpcEntity.Employment_JOB_TYPE_INTERNSHIP)],
	FullTime:   grpcEntity.Employment_JobType_name[int32(grpcEntity.Employment_JOB_TYPE_FULL_TIME)],
	"":         grpcEntity.Employment_JobType_name[int32(grpcEntity.Employment_JOB_TYPE_UNSPECIFIED)],
}
