package cMetadata

import (
	"MusicMesh/Discovery-MusicMesh/generate/composition"
	"database/sql"
	"google.golang.org/grpc"
)

type CMetadata struct {
	Composition composition.CompositionServiceClient
	DB          *sql.DB
}

func NewCMetadata(db *sql.DB, conn *grpc.ClientConn) *CMetadata {
	Client := composition.NewCompositionServiceClient(conn)
	return &CMetadata{DB: db, Composition: Client}
}
