package database

import (
	"context"
	"database/sql"
	"fmt"

	"errors"

	apisv1 "github.com/unkeyed/unkey/apps/agent/gen/proto/apis/v1"
)

func (db *database) FindApi(ctx context.Context, apiId string) (*apisv1.Api, bool, error) {

	model, err := db.read().FindApi(ctx, apiId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("unable to find api: %w", err)
	}

	api, err := transformApiModelToEntity(model)
	if err != nil {
		return nil, true, fmt.Errorf("unable to transform api model to entity: %w", err)
	}
	return api, true, nil
}
