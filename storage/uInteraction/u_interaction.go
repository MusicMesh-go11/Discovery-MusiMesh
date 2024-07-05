package uInteraction

import (
	c "MusicMesh/Discovery-MusicMesh/generate/composition"
	u "MusicMesh/Discovery-MusicMesh/generate/user"
	pb "MusicMesh/Discovery-MusicMesh/generate/user_interactions"
	"database/sql"
	"google.golang.org/grpc"
)

type UserInteraction struct {
	pb.UnimplementedUserInteractionsServiceServer
	User        u.UserServiceClient
	Composition c.CompositionServiceClient
	DB          *sql.DB
}

func NewUserInteraction(db *sql.DB, UserConn *grpc.ClientConn, CompConn *grpc.ClientConn) *UserInteraction {
	user := u.NewUserServiceClient(UserConn)
	composition := c.NewCompositionServiceClient(CompConn)
	return &UserInteraction{User: user, Composition: composition, DB: db}
}
