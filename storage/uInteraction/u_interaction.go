package uInteraction

import (
	comp "MusicMesh/Discovery-MusicMesh/generate/composition"
	"MusicMesh/Discovery-MusicMesh/generate/user"
	pb "MusicMesh/Discovery-MusicMesh/generate/user_interactions"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func (u *UserInteraction) Create(ctx context.Context, in *pb.UserInteraction, opts ...grpc.CallOption) (*pb.Void, error) {
	us, err := u.User.GetByID(ctx, &user.UserId{Id: in.UserId})
	if err != nil {
		return nil, fmt.Errorf("GetUserByID Error: %v", err)
	} else if us.UserID == "" {
		return nil, fmt.Errorf("User Not Found")
	}

	com, err := u.Composition.GetByID(ctx, &comp.CompositionId{CompositionID: in.CompositionId})
	if err != nil {
		return nil, fmt.Errorf("GetCompositionByID Error: %v", err)
	}
	if com.CompositionID == "" {
		return nil, fmt.Errorf("Composition Not Found")
	}

	_, err = u.DB.Exec(`INSERT INTO user_interactions(user_id, composition_id, interaction_type)
VALUES ($1, $2, $3)`, in.UserId, in.CompositionId, in.InteractionType)
	return &pb.Void{}, err
}

func (u *UserInteraction) Get(ctx context.Context, in *pb.Void, opts ...grpc.CallOption) (*pb.UserInteractionSRes, error) {
	rows, err := u.DB.Query(`SELECT id, user_id, composition_id, interaction_type, created_at, updated_at
		where deleted_at = 0`)
	if err != nil {
		return nil, fmt.Errorf("Query in userInteraction Error: %v", err)
	}
	defer rows.Close()

	userInteractions := []*pb.UserInteractionRes{}
	for rows.Next() {
		user := pb.UserInteractionRes{}
		err = rows.Scan(&user.Id, &user.UserId, &user.CompositionId, user.InteractionType, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("Scan in userInteraction Error: %v", err)
		}
		userInteractions = append(userInteractions, &user)
	}
	return &pb.UserInteractionSRes{UserInteraction: userInteractions}, nil
}

func (u *UserInteraction) GetById(ctx context.Context, in *pb.UserInteractionId, opts ...grpc.CallOption) (*pb.UserInteractionRes, error) {
	row := u.DB.QueryRow(`SELECT id, user_id, composition_id, interaction_type, created_at, updated_at 
                          FROM user_interactions 
                          WHERE id = $1 AND deleted_at IS NULL`, in.UserInteractionId)

	user := pb.UserInteractionRes{}
	err := row.Scan(&user.Id, &user.UserId, &user.CompositionId, &user.InteractionType, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("Error fetching UserInteraction by ID: %v", err)
	}

	return &user, nil
}

func (u *UserInteraction) Update(ctx context.Context, in *pb.UserInteractionRes, opts ...grpc.CallOption) (*pb.Void, error) {
	_, err := u.DB.Exec(`UPDATE user_interactions 
                         SET user_id = $1, composition_id = $2, interaction_type = $3
                         WHERE id = $4 AND deleted_at = 0`,
		in.UserId, in.CompositionId, in.InteractionType, in.Id)
	if err != nil {
		return nil, fmt.Errorf("Error updating UserInteraction: %v", err)
	}

	return &pb.Void{}, nil
}

func (u *UserInteraction) Delete(ctx context.Context, in *pb.UserInteractionId, opts ...grpc.CallOption) (*pb.Void, error) {
	_, err := u.DB.Exec(`update table_name set
 deleted_at = date_part('epoch', current_timestamp)::INT
where id = $1 and deleted_at = 0`, in.UserInteractionId)
	if err != nil {
		return nil, fmt.Errorf("Error deleting UserInteraction: %v", err)
	}

	return &pb.Void{}, nil
}
