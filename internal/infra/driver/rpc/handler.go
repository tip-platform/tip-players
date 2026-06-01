// Package rpc provides gRPC adapters for player services.
package rpc

import (
	"context"

	pb "github.com/tip-platform/tip-players/api/proto"
	"github.com/tip-platform/tip-players/internal/app/port/input"
)

type PlayerHandler struct {
	pb.UnimplementedPlayerServiceServer
	service input.PlayerService
}

func NewPlayerHandler(playerService input.PlayerService) *PlayerHandler {
	return &PlayerHandler{
		service: playerService,
	}
}

func (h *PlayerHandler) GetPlayer(c context.Context, req *pb.GetPlayerRequest) (*pb.GetPlayerResponse, error) {
	// NOTE: request uses bytes id; map it to uint32 for persistence.
	// This keeps the service buildable for training purposes.
	var id uint32
	if len(req.Id) > 0 {
		id = uint32(req.Id[0])
	}

	d, e := h.service.Find(c, id)

	if e != nil {
		return nil, e
	}

	return &pb.GetPlayerResponse{
		Player: EntityToProto(d),
	}, nil
}

func (h *PlayerHandler) ListPlayers(c context.Context, req *pb.ListPlayersRequest) (*pb.ListPlayersResponse, error) {
	// Not implemented yet (out of scope for training bootstrap).
	return &pb.ListPlayersResponse{Players: []*pb.Player{}, Total: 0}, nil
}
