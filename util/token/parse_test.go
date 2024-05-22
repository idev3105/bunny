package tokenutil_test

import (
	"context"
	"fmt"
	"testing"

	tokenutil "org.idev.bunny/backend/util/token"
)

func TestParse(t *testing.T) {
	ctx := context.TODO()
	// Replace new token to test
	tokenRaw := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ2Mk5fQ1J1M0J0Q2hfYkZzNGVkbXVZODRWY2R6NDhhYlpTVTVQdm9wVTlnIn0.eyJleHAiOjE3MTM4OTIyNTYsImlhdCI6MTcxMzg5MjE5NiwiYXV0aF90aW1lIjoxNzEzODkxMjA3LCJqdGkiOiI5ZGE0MTVjMi03Y2YyLTQwOWQtYjdkMi03ZWU5ZDljYTY5MWIiLCJpc3MiOiJodHRwOi8vMC4wLjAuMDo4MDgwL2F1dGgvcmVhbG1zL21hc3RlciIsImF1ZCI6InNlY3VyaXR5LWFkbWluLWNvbnNvbGUiLCJzdWIiOiJhOWE5N2ViZC1lNDk0LTQ0NzctOTgzNC1lNWQzNTkzNGI1NWMiLCJ0eXAiOiJJRCIsImF6cCI6InNlY3VyaXR5LWFkbWluLWNvbnNvbGUiLCJub25jZSI6ImE5NDQ4NzgzLTgwNDYtNDM0Zi04MzVmLWQ4MDlhNzNjZmVmOSIsInNlc3Npb25fc3RhdGUiOiJmMjlhM2I3OS1kNjVjLTQxNDAtYmU3NS03NDRhZmJiNmY4NmQiLCJhdF9oYXNoIjoieXpzb1UwbFhhbG9zaEI5Vklmc2poZyIsImFjciI6IjAiLCJzaWQiOiJmMjlhM2I3OS1kNjVjLTQxNDAtYmU3NS03NDRhZmJiNmY4NmQiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluIn0.BGuI9SSJ3ZV5Qds1Zyg5hJRST6f3gMzU9YgKAsgSPvlm4BtLS3VhKtD_kzrGumylCQCmIhMonR2OqTuv5y6Kvub_WVOO8BHNMmSzbQx7668m2NKaFD04RIdr_o0vsbcy5M4sUKL-oeZnLuRTgrJd4PnVZHgYoXOJGxjPGdFVbkphOwC8UhywL3aL6it_nngjfVXSzeFqoKepW_2tkhJe9TcbpFg-7Zl8OQrQTBdavb79749_5XtsidzLsJug7AgrMr9kVsnXxoyZ69YgI3FZzoU0DJfgCHWBtnhd-7G3ZP5Yyhqg4nTsUQueoxqqz4RzYWUxf6IPNnlGbpskVY33cw"
	jwkUrl := "http://0.0.0.0:8080/auth/realms/master/protocol/openid-connect/certs"
	token, err := tokenutil.Parse(ctx, tokenRaw, jwkUrl)
	if err != nil {
		t.Fatalf("Parse token failed: %v", err)
	}
	sub, isOk := token.Get("sub")
	if !isOk {
		t.Fatalf("Parse token failed: token does not contain 'sub' field")
	}
	fmt.Println(sub)
}
