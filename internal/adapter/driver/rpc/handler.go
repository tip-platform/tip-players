// Package rpc provides gRPC adapters for player services.
package rpc

import (
	"context"

	"github.com/tip-platform/tip-players/internal/port/input"
	pb "github.com/tip-platform/tip-players/proto"
)

type PlayerHandler struct {
	pb.UnimplementedPlayerServiceServer
	repo *PlayerRepositoryAdapter
}

func NewPlayerHandler(playerService input.PlayerService) *PlayerHandler {
	return &PlayerHandler{
		repo: NewPlayerRepositoryAdapter(playerService),
	}
}

func (h *PlayerHandler) GetPlayer(c context.Context, req *pb.GetPlayerRequest) (*pb.GetPlayerResponse, error) {
	// NOTE: request uses bytes id; map it to uint32 for persistence.
	// This keeps the service buildable for training purposes.
	var id uint32
	if len(req.Id) > 0 {
		id = uint32(req.Id[0])
	}

	d, e := h.repo.Find(c, id)

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
