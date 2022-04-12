package db

import grpcEntity "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/entity/v1"

const (
	HolisticOfficeLink_Code          = "HolisticOfficeLink_Code"
	HolisticOfficeLink_Documentation = "HolisticOfficeLink_Documentation"
)

// TODO: delete this bridge map once djin-server is deployed
var HolisticOffficeLinkDbTypeToProtoMap = map[string]string{
	HolisticOfficeLink_Code:          grpcEntity.HolisticOfficeLink_HolisticOfficeLinkType_name[int32(grpcEntity.HolisticOfficeLink_HOLISTIC_OFFICE_LINK_TYPE_CODE)],
	HolisticOfficeLink_Documentation: grpcEntity.HolisticOfficeLink_HolisticOfficeLinkType_name[int32(grpcEntity.HolisticOfficeLink_HOLISTIC_OFFICE_LINK_TYPE_DOCUMENTATION)],
	"":                               grpcEntity.HolisticOfficeLink_HolisticOfficeLinkType_name[int32(grpcEntity.HolisticOfficeLink_HOLISTIC_OFFICE_LINK_TYPE_UNSPECIFIED)],
}
