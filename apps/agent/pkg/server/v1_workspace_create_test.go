package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	apisv1 "github.com/unkeyed/unkey/apps/agent/gen/proto/apis/v1"
	authenticationv1 "github.com/unkeyed/unkey/apps/agent/gen/proto/authentication/v1"
	"github.com/unkeyed/unkey/apps/agent/pkg/cache"
	"github.com/unkeyed/unkey/apps/agent/pkg/events"
	"github.com/unkeyed/unkey/apps/agent/pkg/logging"
	"github.com/unkeyed/unkey/apps/agent/pkg/services/keys"
	"github.com/unkeyed/unkey/apps/agent/pkg/services/workspaces"
	"github.com/unkeyed/unkey/apps/agent/pkg/testutil"
	"github.com/unkeyed/unkey/apps/agent/pkg/tracing"
	"github.com/unkeyed/unkey/apps/agent/pkg/uid"
)

func TestCreateWorkspace_Simple(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	resources := testutil.SetupResources(t)

	srv := New(Config{
		Logger:            logging.NewNoop(),
		KeyCache:          cache.NewNoopCache[*authenticationv1.Key](),
		ApiCache:          cache.NewNoopCache[*apisv1.Api](),
		Database:          resources.Database,
		Tracer:            tracing.NewNoop(),
		UnkeyWorkspaceId:  resources.UnkeyWorkspace.WorkspaceId,
		UnkeyApiId:        resources.UnkeyApi.ApiId,
		UnkeyAppAuthToken: "supersecret",
		WorkspaceService:  workspaces.New(workspaces.Config{Database: resources.Database}),
		KeyService: keys.New(keys.Config{
			Database: resources.Database,
			Events:   events.NewNoop(),
		}),
	})

	tenantId := uid.New(16, "")
	res := CreateWorkspaceResponseV1{}
	testutil.Json(t, srv.app, testutil.JsonRequest{
		Method: "POST",
		Path:   "/v1/workspaces.createWorkspace",
		Body: fmt.Sprintf(`{
			"name":"simple",
			"tenantId": "%s"
			}`, tenantId),
		Bearer:     srv.unkeyAppAuthToken,
		Response:   &res,
		StatusCode: 200,
	})
	t.Logf("res: %+v", res)

	require.NotEmpty(t, res.Id)

	foundWorkspace, found, err := resources.Database.FindWorkspace(ctx, res.Id)
	require.NoError(t, err)
	require.True(t, found)

	require.Equal(t, res.Id, foundWorkspace.WorkspaceId)
	require.Equal(t, "simple", foundWorkspace.Name)
	require.Equal(t, tenantId, foundWorkspace.TenantId)
}
