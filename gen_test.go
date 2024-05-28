package ydb_go_genproto

import (
	"testing"

	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Auth"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Cms"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Coordination"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Discovery"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Export"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Formats"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Import"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Issue"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Monitoring"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Operations"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Query"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_RateLimiter"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Scheme"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Scripting"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Table"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_TableStats"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Topic"

	_ "github.com/ydb-platform/ydb-go-genproto/draft/protos/Ydb_DynamicConfig"
	_ "github.com/ydb-platform/ydb-go-genproto/draft/protos/Ydb_Maintenance"
	_ "github.com/ydb-platform/ydb-go-genproto/draft/protos/Ydb_ObjectStorage"
)

func TestCompatibility(_ *testing.T) {
}
