package cMetadata

import (
	pb "MusicMesh/Discovery-MusicMesh/generate/composition_metadata"
	"database/sql"
)

type CmetadataRepo struct {
	pb.UnimplementedCompositionMetadataServiceServer
	DB *sql.DB
}

func NewCmetadataRepo(db *sql.DB) *CmetadataRepo {
	return &CmetadataRepo{DB: db}
}
