package controller

import (
	"context"
	"wrapper_for_rusprofile/internal/api"
	"wrapper_for_rusprofile/internal/usecase"
)

type InnHandler struct{ innUseCase usecase.InnUseCase }

func (h *InnHandler) GetByInn(ctx context.Context, req *api.GetInformationRequest) (*api.GetInformationResponse, error) {
	fullInfoInn, err := h.innUseCase.GetFullInformationInn(ctx, req.GetInn())
	if err != nil {
		return &api.GetInformationResponse{}, err
	}
	return &api.GetInformationResponse{Inn: fullInfoInn.Inn, Kpp: fullInfoInn.Kpp, Fio: fullInfoInn.Fio, Company: fullInfoInn.Company}, nil
}
