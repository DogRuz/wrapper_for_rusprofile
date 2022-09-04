package usecase

import (
	"context"
	"log"
	"wrapper_for_rusprofile/internal/domain/service"
	"wrapper_for_rusprofile/internal/entity"
)

type InnUseCase struct {
	innService service.InnService
}

func (u InnUseCase) GetFullInformationInn(ctx context.Context, inn string) (entity.FullInfoInn, error) {
	fullInfoInn, err := u.innService.GetByInn(ctx, inn)
	if err != nil {
		log.Println(err)
		return entity.FullInfoInn{}, err
	}
	return fullInfoInn, err
}
