package server_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	authenticationv1 "github.com/unkeyed/unkey/apps/agent/gen/proto/authentication/v1"

	"github.com/unkeyed/unkey/apps/agent/pkg/hash"
	"github.com/unkeyed/unkey/apps/agent/pkg/testutil"

	"github.com/stretchr/testify/require"
	"github.com/unkeyed/unkey/apps/agent/pkg/uid"
)

func TestDeleteKey(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	resources := testutil.SetupResources(t)

	key := &authenticationv1.Key{
		KeyId:       uid.Key(),
		KeyAuthId:   resources.UserKeyAuth.KeyAuthId,
		WorkspaceId: resources.UserWorkspace.WorkspaceId,
		Hash:        hash.Sha256(uid.New(16, "test")),
		CreatedAt:   time.Now().UnixMilli(),
	}
	err := resources.Database.InsertKey(ctx, key)
	require.NoError(t, err)

	srv := testutil.NewServer(t, resources)
	testutil.Json[any](t, srv.App, testutil.JsonRequest{
		Method:     "DELETE",
		Path:       fmt.Sprintf("/v1/keys/%s", key.KeyId),
		Bearer:     resources.UserRootKey,
		StatusCode: 200,
	})

	key, found, err := resources.Database.FindKeyById(ctx, key.KeyId)
	require.NoError(t, err)
	require.True(t, found)
	require.NotNil(t, key.DeletedAt)
}
