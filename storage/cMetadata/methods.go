package cMetadata

import (
	"MusicMesh/Discovery-MusicMesh/generate/composition"
	pb "MusicMesh/Discovery-MusicMesh/generate/composition_metadata"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func (m *CMetadata) Create(ctx context.Context, in *pb.CompositionMetadata, opts ...grpc.CallOption) (*pb.Void, error) {
	comp, err := m.Composition.GetByID(ctx, &composition.CompositionId{CompositionID: in.CompositionId})
	if err != nil {
		return nil, fmt.Errorf("Error checking composition Id: %v", err)
	}
	if comp == nil {
		return nil, fmt.Errorf("Composition does not exist")
	}

	_, err = m.DB.Exec("INSERT INTO composition_metadata (composition_id, genre, tags) VALUES ($1, $2, $3)",
		in.CompositionId, in.Genre, in.Tags)
	return &pb.Void{}, err
}
