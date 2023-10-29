package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	apisv1 "github.com/unkeyed/unkey/apps/agent/gen/proto/apis/v1"
	authenticationv1 "github.com/unkeyed/unkey/apps/agent/gen/proto/authentication/v1"
	"github.com/unkeyed/unkey/apps/agent/pkg/cache"

	"github.com/unkeyed/unkey/apps/agent/pkg/logging"
	"github.com/unkeyed/unkey/apps/agent/pkg/testutil"
	"github.com/unkeyed/unkey/apps/agent/pkg/tracing"
)

func TestRootCreateKey_Simple(t *testing.T) {
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
	})

	buf := bytes.NewBufferString(fmt.Sprintf(`{
		"name":"simple",
		"forWorkspaceId":"%s"
		}`, resources.UserWorkspace.WorkspaceId))

	req := httptest.NewRequest("POST", "/v1/internal/rootkeys", buf)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", srv.unkeyAppAuthToken))
	req.Header.Set("Content-Type", "application/json")

	res, err := srv.app.Test(req)
	require.NoError(t, err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	require.Equal(t, res.StatusCode, 200)

	createRootKeyResponse := CreateRootKeyResponse{}
	err = json.Unmarshal(body, &createRootKeyResponse)
	require.NoError(t, err)

	require.NotEmpty(t, createRootKeyResponse.Key)
	require.NotEmpty(t, createRootKeyResponse.KeyId)

	foundKey, found, err := resources.Database.FindKeyById(ctx, createRootKeyResponse.KeyId)
	require.NoError(t, err)
	require.True(t, found)

	require.Equal(t, createRootKeyResponse.KeyId, foundKey.KeyId)
}

func TestRootCreateKey_WithExpiry(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	resources := testutil.SetupResources(t)

	srv := New(Config{
		Logger:            logging.NewNoop(),
		KeyCache:          cache.NewNoopCache[*authenticationv1.Key](),
		ApiCache:          cache.NewNoopCache[*apisv1.Api](),
		Database:          resources.Database,
		Tracer:            tracing.NewNoop(),
		UnkeyAppAuthToken: "supersecret",
		UnkeyWorkspaceId:  resources.UnkeyWorkspace.WorkspaceId,
		UnkeyApiId:        resources.UnkeyApi.ApiId,
	})

	buf := bytes.NewBufferString(fmt.Sprintf(`{
		"name":"simple",
		"forWorkspaceId":"%s",
		"expires": %d
		}`, resources.UserWorkspace.WorkspaceId, time.Now().Add(time.Hour).UnixMilli()))

	req := httptest.NewRequest("POST", "/v1/internal/rootkeys", buf)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", srv.unkeyAppAuthToken))
	req.Header.Set("Content-Type", "application/json")

	res, err := srv.app.Test(req)
	require.NoError(t, err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	require.Equal(t, 200, res.StatusCode)

	createRootKeyResponse := CreateRootKeyResponse{}
	err = json.Unmarshal(body, &createRootKeyResponse)
	require.NoError(t, err)

	require.NotEmpty(t, createRootKeyResponse.Key)
	require.NotEmpty(t, createRootKeyResponse.KeyId)

	foundKey, found, err := resources.Database.FindKeyById(ctx, createRootKeyResponse.KeyId)
	require.True(t, found)
	require.NoError(t, err)
	require.Equal(t, createRootKeyResponse.KeyId, foundKey.KeyId)
	require.GreaterOrEqual(t, len(createRootKeyResponse.Key), 30)
	require.True(t, strings.HasPrefix(createRootKeyResponse.Key, "unkey_"))
	require.True(t, time.UnixMilli(foundKey.GetExpires()).After(time.Now()))

	require.Equal(t, authenticationv1.RatelimitType_RATELIMIT_TYPE_FAST, foundKey.Ratelimit.Type)
	require.Equal(t, int32(100), foundKey.Ratelimit.Limit)
	require.Equal(t, int32(10), foundKey.Ratelimit.RefillRate)
	require.Equal(t, int32(1000), foundKey.Ratelimit.RefillInterval)

}
