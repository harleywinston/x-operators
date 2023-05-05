package services

import (
	"github.com/harleywinston/x-operators/pkg/models"
)

type SyncerServices struct{}

func (s *SyncerServices) getMasterState() (models.MasterState, error) {
	var res models.MasterState

	return res, nil
}

// func (s *SyncerServices) getXUIState() (models, error) {
// 	xuiService := service.SetupServices{}
// 	res, err := xuiService.ListInbounds()
// 	if err != nil {
// 		return models.XUIState{}, err
// 	}
// 	return models.XUIState{}, nil
// }

func (s *SyncerServices) Sync() error {
	return nil
}
