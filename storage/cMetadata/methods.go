package cMetadata

import (
	pb "MusicMesh/Discovery-MusicMesh/generate/composition_metadata"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func (m *CmetadataRepo) Create(ctx context.Context, in *pb.CompositionMetadata) (*pb.Void, error) {

	query := "INSERT INTO composition_metadata (composition_id, genre, tags) VALUES ($1, $2, $3)"
	_, err := m.DB.ExecContext(ctx, query, in.CompositionId, in.Genre, in.Tags)
	if err != nil {
		return nil, fmt.Errorf("failed to insert composition metadata: %v", err)
	}
	return &pb.Void{}, nil
}

// GetTrending retrieves the top 10 compositions ordered by listen count
func (m *CmetadataRepo) GetTrending(ctx context.Context, in *pb.Void) (*pb.CompositionsRes, error) {
	query := "SELECT composition_id, genre, tags, listen_count, like_count FROM composition_metadata ORDER BY listen_count DESC LIMIT 10"
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch trending compositions: %v", err)
	}
	defer rows.Close()

	var compositions []*pb.CompositionRes
	for rows.Next() {
		var comp pb.CompositionRes
		if err := rows.Scan(&comp.CompositionId, &comp.Genre, &comp.Tags, &comp.ListenCount, &comp.LikeCount); err != nil {
			return nil, fmt.Errorf("failed to scan composition row: %v", err)
		}
		compositions = append(compositions, &comp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return &pb.CompositionsRes{Compositions: compositions}, nil
}

// GetRecommended retrieves recommended compositions for a user based on interactions
func (m *CmetadataRepo) GetRecommended(ctx context.Context, in *pb.UserId) (*pb.CompositionsRes, error) {
	query := `
		SELECT cm.composition_id, cm.genre, cm.tags, cm.listen_count, cm.like_count
		FROM composition_metadata cm
		JOIN user_interactions ui ON cm.composition_id = ui.composition_id
		WHERE ui.user_id = $1
		ORDER BY cm.like_count DESC`
	rows, err := m.DB.QueryContext(ctx, query, in.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch recommended compositions: %v", err)
	}
	defer rows.Close()

	var compositions []*pb.CompositionRes
	for rows.Next() {
		var comp pb.CompositionRes
		if err := rows.Scan(&comp.CompositionId, &comp.Genre, &comp.Tags, &comp.ListenCount, &comp.LikeCount); err != nil {
			return nil, fmt.Errorf("failed to scan composition row: %v", err)
		}
		compositions = append(compositions, &comp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return &pb.CompositionsRes{Compositions: compositions}, nil
}

// GetByGenre retrieves compositions filtered by genre
func (m *CmetadataRepo) GetByGenre(ctx context.Context, in *pb.GenreRequest, opts ...grpc.CallOption) (*pb.CompositionsRes, error) {
	query := "SELECT composition_id, genre, tags, listen_count, like_count FROM composition_metadata WHERE genre = $1 ORDER BY listen_count DESC"
	rows, err := m.DB.QueryContext(ctx, query, in.Genre)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch compositions by genre: %v", err)
	}
	defer rows.Close()

	var compositions []*pb.CompositionRes
	for rows.Next() {
		var comp pb.CompositionRes
		if err := rows.Scan(&comp.CompositionId, &comp.Genre, &comp.Tags, &comp.ListenCount, &comp.LikeCount); err != nil {
			return nil, fmt.Errorf("failed to scan composition row: %v", err)
		}
		compositions = append(compositions, &comp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return &pb.CompositionsRes{Compositions: compositions}, nil
}

// Update modifies existing composition metadata in the database
func (m *CmetadataRepo) Update(ctx context.Context, in *pb.CompositionRes, opts ...grpc.CallOption) (*pb.Void, error) {
	query := "UPDATE composition_metadata SET genre = $1, tags = $2, listen_count = $3, like_count = $4 WHERE composition_id = $5"
	_, err := m.DB.ExecContext(ctx, query, in.Genre, in.Tags, in.ListenCount, in.LikeCount, in.CompositionId)
	if err != nil {
		return nil, fmt.Errorf("failed to update composition metadata: %v", err)
	}
	return &pb.Void{}, nil
}

// Delete removes a composition metadata record from the database
func (m *CmetadataRepo) Delete(ctx context.Context, in *pb.CompositionMetadataId, opts ...grpc.CallOption) (*pb.Void, error) {
	query := "DELETE FROM composition_metadata WHERE composition_id = $1"
	_, err := m.DB.ExecContext(ctx, query, in.MetadataId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete composition metadata: %v", err)
	}
	return &pb.Void{}, nil
}
