package service

import (
	"context"
	"wrapper_for_rusprofile/internal/entity"
	"wrapper_for_rusprofile/pkg/client/rusprofile"
)

type InnService struct {
	client rusprofile.Client
}

func (i InnService) GetByInn(ctx context.Context, inn string) (entity.FullInfoInn, error) {
	i.client = rusprofile.New()
	infoInn, err := i.client.GetFullInformationByInn(inn)
	if err != nil || len(infoInn) == 0 {
		return entity.FullInfoInn{}, err
	}
	return entity.FullInfoInn{Inn: infoInn[0].Inn, Kpp: infoInn[0].Kpp, Company: infoInn[0].Company, Fio: infoInn[0].Fio}, nil
}
