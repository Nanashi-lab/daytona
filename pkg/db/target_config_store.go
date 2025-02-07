// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package db

import (
	"context"

	"gorm.io/gorm"

	"github.com/daytonaio/daytona/pkg/models"
	"github.com/daytonaio/daytona/pkg/stores"
)

type TargetConfigStore struct {
	IStore
}

func NewTargetConfigStore(store IStore) (stores.TargetConfigStore, error) {
	err := store.AutoMigrate(&models.TargetConfig{})
	if err != nil {
		return nil, err
	}

	return &TargetConfigStore{store}, nil
}

func (s *TargetConfigStore) List(ctx context.Context, allowDeleted bool) ([]*models.TargetConfig, error) {
	tx := s.GetTransaction(ctx)

	targetConfigs := []*models.TargetConfig{}
	tx = processTargetConfigFilters(tx, nil, allowDeleted).Find(&targetConfigs)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return targetConfigs, nil
}

func (s *TargetConfigStore) Find(ctx context.Context, idOrName string, allowDeleted bool) (*models.TargetConfig, error) {
	tx := s.GetTransaction(ctx)

	targetConfig := &models.TargetConfig{}
	tx = processTargetConfigFilters(tx, &idOrName, allowDeleted).First(targetConfig)

	if tx.Error != nil {
		if IsRecordNotFound(tx.Error) {
			return nil, stores.ErrTargetConfigNotFound
		}
		return nil, tx.Error
	}

	return targetConfig, nil
}

func (s *TargetConfigStore) Save(ctx context.Context, targetConfig *models.TargetConfig) error {
	tx := s.GetTransaction(ctx)

	tx = tx.Save(targetConfig)
	return tx.Error
}

func processTargetConfigFilters(tx *gorm.DB, idOrName *string, allowDeleted bool) *gorm.DB {
	if idOrName != nil {
		tx = tx.Where("id = ? OR name = ?", *idOrName, *idOrName)
	}

	if !allowDeleted {
		tx = tx.Where("deleted = ?", false)
	}

	return tx
}
