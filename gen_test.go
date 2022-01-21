package ydb_go_genproto

import (
	"testing"

	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Cms"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Coordination"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Discovery"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Experimental"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Export"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Import"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Issue"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Monitoring"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Operations"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_PersQueue_V1"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_RateLimiter"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Scheme"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Scripting"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Table"
	_ "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_TableStats"
)

func TestCompatibility(_ *testing.T) {
}
