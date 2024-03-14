package service

import "filmlibr/internal/entity"

type ActorService struct {
}

func NewActorService() *ActorService {
	return &ActorService{}
}

func (s *ActorService) GetAll() ([]entity.ActorWithFilms, error) {
	return nil, nil
}

func (s *ActorService) Add(a entity.Actor) (entity.Actor, error) {
	return entity.Actor{}, nil
}

func (s *ActorService) GetByID(id int) (entity.Actor, error) {
	return entity.Actor{}, nil
}

func (s *ActorService) UpdateCompletely(a entity.Actor) (entity.Actor, error) {
	return entity.Actor{}, nil
}

func (s *ActorService) UpdatePartial(a entity.Actor) (entity.Actor, error) {
	return entity.Actor{}, nil
}

func (s *ActorService) Delete(id int) error {
	return nil
}
