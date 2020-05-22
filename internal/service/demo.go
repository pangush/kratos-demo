package service

import "context"

func (s *Service) Demo(ctx context.Context) error {
	err := s.dao.Ping(ctx)
	return err
}
