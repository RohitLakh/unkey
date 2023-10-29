package keys

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	authenticationv1 "github.com/unkeyed/unkey/apps/agent/gen/proto/authentication/v1"
	"github.com/unkeyed/unkey/apps/agent/pkg/cache"
	"github.com/unkeyed/unkey/apps/agent/pkg/events"
	"github.com/unkeyed/unkey/apps/agent/pkg/testutil"
	"github.com/unkeyed/unkey/apps/agent/pkg/util"
)

func Test_CreateKey_Minimal(t *testing.T) {

	resources := testutil.SetupResources(t)

	svc := New(Config{
		Database: resources.Database,
		Events:   events.NewNoop(),
	})

	ctx := context.Background()

	res, err := svc.CreateKey(ctx, &authenticationv1.CreateKeyRequest{
		WorkspaceId: resources.UserApi.WorkspaceId,
		KeyAuthId:   resources.UserApi.GetKeyAuthId(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, res.Key)
	require.NotEmpty(t, res.KeyId)

	key, found, err := resources.Database.FindKeyById(ctx, res.KeyId)
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, res.KeyId, key.KeyId)
	require.Equal(t, resources.UserApi.WorkspaceId, key.WorkspaceId)
	require.Equal(t, resources.UserApi.GetKeyAuthId(), key.KeyAuthId)

}

func Test_CreateKey_WithExpiration(t *testing.T) {

	resources := testutil.SetupResources(t)

	svc := New(Config{
		Database: resources.Database,
		Events:   events.NewNoop(),
	})

	ctx := context.Background()

	expires := time.Now().Add(time.Hour).UnixMilli()
	res, err := svc.CreateKey(ctx, &authenticationv1.CreateKeyRequest{
		WorkspaceId: resources.UserApi.WorkspaceId,
		KeyAuthId:   resources.UserApi.GetKeyAuthId(),
		Expires:     util.Pointer(expires),
	})
	require.NoError(t, err)
	require.NotEmpty(t, res.Key)
	require.NotEmpty(t, res.KeyId)

	key, found, err := resources.Database.FindKeyById(ctx, res.KeyId)
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, res.KeyId, key.KeyId)
	require.Equal(t, resources.UserApi.WorkspaceId, key.WorkspaceId)
	require.Equal(t, resources.UserApi.GetKeyAuthId(), key.KeyAuthId)
	require.NotNil(t, key.Expires)
	require.Equal(t, expires, *key.Expires)

}

func Test_CreateKey_WithRemaining(t *testing.T) {

	resources := testutil.SetupResources(t)

	svc := New(Config{
		Database: resources.Database,
		Events:   events.NewNoop(),
		KeyCache: cache.NewNoopCache[*authenticationv1.Key](),
	})

	ctx := context.Background()

	res, err := svc.CreateKey(ctx, &authenticationv1.CreateKeyRequest{
		WorkspaceId: resources.UserApi.WorkspaceId,
		KeyAuthId:   resources.UserApi.GetKeyAuthId(),
		Remaining:   util.Pointer(int32(10)),
	})
	require.NoError(t, err)
	require.NotEmpty(t, res.Key)
	require.NotEmpty(t, res.KeyId)

	key, found, err := resources.Database.FindKeyById(ctx, res.KeyId)
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, res.KeyId, key.KeyId)
	require.Equal(t, resources.UserApi.WorkspaceId, key.WorkspaceId)
	require.Equal(t, resources.UserApi.GetKeyAuthId(), key.KeyAuthId)
	require.NotNil(t, key.Remaining)
	require.Equal(t, int32(10), *key.Remaining)

}
