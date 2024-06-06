package line

import (
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/domain/service"
)

type Service struct {
	repo service.LineService
}

func NewLineService(cfg config.Config) *Service {

	return &Service{repo: service.NewLineService(cfg)}
}

//// CreateLan handles the logic for creating a LAN.
//func (s *LineService) CreateLine(ctx context.Context, dto CreateLanDTO) (line.Lan, error) {
//	// Convert DTO to domain entity
//	lan := line.NewLan(dto.Name, dto.CIDR)
//	if err := s.repo.Save(ctx, lan); err != nil {
//		return lan, err
//	}
//	return lan, nil
//}
