// Package rpc provides gRPC adapters for player services.
package rpc

import (
	"github.com/tip-platform/tip-players/internal/domain/entity"
	pb "github.com/tip-platform/tip-players/proto"
)

// schemaToProto mapea de la entidad de dominio al mensaje generado por gRPC.
func EntityToProto(e entity.PlayerRecord) *pb.Player {
	return &pb.Player{
		Id:          "",
		ApiId:       e.APIID,
		Name:        e.Name,
		ShortName:   e.ShortName,
		CountryCode: e.CountryCode,
		CountryName: e.CountryName,
		Age:         uint32(e.Age),
		Plays:       e.Plays,
	}
}

// protoToSchema mapea del mensaje gRPC a la entidad de dominio.
func protoToEntity(p *pb.Player) entity.PlayerRecord {
	if p == nil {
		return entity.PlayerRecord{}
	}

	return entity.PlayerRecord{
		ID:          0,
		APIID:       p.ApiId,
		Name:        p.Name,
		ShortName:   p.ShortName,
		CountryCode: p.CountryCode,
		CountryName: p.CountryName,
		Age:         uint8(p.Age),
		Plays:       p.Plays,
	}
}
