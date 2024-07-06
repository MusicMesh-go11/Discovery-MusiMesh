package uInteraction

import (
	pb "MusicMesh/Discovery-MusicMesh/generate/user_interactions"
	"context"
	"database/sql"
	"fmt"
	"log"
)

func (u *UserInteraction) Create(ctx context.Context, in *pb.UserInteraction) (*pb.Void, error) {
	// Insert user interaction
	_, err := u.DB.ExecContext(ctx, `INSERT INTO user_interactions(user_id, composition_id, interaction_type)
		VALUES ($1, $2, $3)`, in.UserId, in.CompositionId, in.InteractionType)
	if err != nil {
		log.Printf("Database Insert Error: %v", err)
		return nil, fmt.Errorf("Database Insert Error: %w", err)
	}
	return &pb.Void{}, nil
}

func (u *UserInteraction) Get(ctx context.Context, in *pb.Void) (*pb.UserInteractionSRes, error) {
	rows, err := u.DB.Query(`SELECT id, user_id, composition_id, interaction_type, created_at, updated_at
		FROM user_interactions
		WHERE deleted_at IS NULL`) // Fix: Use IS NULL to check for non-deleted entries
	if err != nil {
		return nil, fmt.Errorf("Query in userInteraction Error: %w", err)
	}
	defer rows.Close()

	userInteractions := []*pb.UserInteractionRes{}
	for rows.Next() {
		user := pb.UserInteractionRes{}
		err = rows.Scan(&user.Id, &user.UserId, &user.CompositionId, &user.InteractionType, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("Scan in userInteraction Error: %w", err)
		}
		userInteractions = append(userInteractions, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Row iteration error: %w", err)
	}

	return &pb.UserInteractionSRes{UserInteraction: userInteractions}, nil
}

func (u *UserInteraction) GetById(ctx context.Context, in *pb.UserInteractionId) (*pb.UserInteractionRes, error) {
	row := u.DB.QueryRow(`SELECT id, user_id, composition_id, interaction_type, created_at, updated_at 
                          FROM user_interactions 
                          WHERE id = $1 AND deleted_at = 0`, in.UserInteractionId)

	user := pb.UserInteractionRes{}
	err := row.Scan(&user.Id, &user.UserId, &user.CompositionId, &user.InteractionType, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("UserInteraction not found")
		}
		return nil, fmt.Errorf("Error fetching UserInteraction by ID: %w", err)
	}

	return &user, nil
}

func (u *UserInteraction) Update(ctx context.Context, in *pb.UserInteractionRes) (*pb.Void, error) {
	_, err := u.DB.Exec(`UPDATE user_interactions 
                         SET user_id = $1, composition_id = $2, interaction_type = $3, updated_at = now()
                         WHERE id = $4 AND deleted_at = 0`,
		in.UserId, in.CompositionId, in.InteractionType, in.Id)
	if err != nil {
		return nil, fmt.Errorf("Error updating UserInteraction: %w", err)
	}

	return &pb.Void{}, nil
}

func (u *UserInteraction) Delete(ctx context.Context, in *pb.UserInteractionId) (*pb.Void, error) {
	_, err := u.DB.Exec(`UPDATE user_interactions
                         SET deleted_at = extract(epoch from now())::BIGINT
                         WHERE id = $1 AND deleted_at = 0`, in.UserInteractionId)
	if err != nil {
		return nil, fmt.Errorf("Error deleting UserInteraction: %w", err)
	}

	return &pb.Void{}, nil
}
