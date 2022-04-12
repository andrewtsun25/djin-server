package db

import grpcEntity "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/entity/v1"

const (
	MartialArtsStyleType_HDGD    = "hdgd"
	MartialArtsStyleType_ITF_TKD = "itfTkd"
	MartialArtsStyleType_WT_TKD  = "wtTkd"
)

// TODO: delete these bridge maps once djin-server is deployed
var MartialArtsStyleTypeDbTypeToProtoMap = map[string]string{
	MartialArtsStyleType_HDGD:    grpcEntity.MartialArtsStyle_MartialArtsStyleType_name[int32(grpcEntity.MartialArtsStyle_MARTIAL_ARTS_STYLE_TYPE_HDGD)],
	MartialArtsStyleType_ITF_TKD: grpcEntity.MartialArtsStyle_MartialArtsStyleType_name[int32(grpcEntity.MartialArtsStyle_MARTIAL_ARTS_STYLE_TYPE_ITF_TKD)],
	MartialArtsStyleType_WT_TKD:  grpcEntity.MartialArtsStyle_MartialArtsStyleType_name[int32(grpcEntity.MartialArtsStyle_MARTIAL_ARTS_STYLE_TYPE_WT)],
	"":                           grpcEntity.MartialArtsStyle_MartialArtsStyleType_name[int32(grpcEntity.MartialArtsStyle_MARTIAL_ARTS_STYLE_TYPE_UNSPECIFIED)],
}
