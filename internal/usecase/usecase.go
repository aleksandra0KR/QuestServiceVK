package usecase

import (
	"vk/internal/repository"
)

type UseCase struct {
	UserUseCase
	QuestUseCase
	SubquestUseCase
}

func NewUseCase(repository *repository.Repository) *UseCase {
	return &UseCase{
		UserUseCase:     NewUserUseCaseImplementation(repository.UserRepository),
		QuestUseCase:    NewQuestUseCaseImplementation(repository.QuestRepository),
		SubquestUseCase: NewSubquestUseCaseImplementation(repository.SubquestRepository)}
}
