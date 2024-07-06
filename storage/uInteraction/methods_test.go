package uInteraction

import (
	pb "MusicMesh/Discovery-MusicMesh/generate/user_interactions"
	"MusicMesh/Discovery-MusicMesh/storage/postgres"
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupTestRepo(t *testing.T) *UserInteraction {
	db, err := postgres.Conn()
	if err != nil {
		t.Fatalf("Failed to connect to the test database: %v", err)
	}

	// Clean up test database before each test
	_, err = db.Exec("TRUNCATE composition_metadata, user_interactions RESTART IDENTITY CASCADE")
	if err != nil {
		t.Fatalf("Failed to truncate tables: %v", err)
	}

	return &UserInteraction{DB: db}
}

func teardownTestRepo(t *testing.T, db *sql.DB) {
	_, err := db.Exec("TRUNCATE composition_metadata, user_interactions RESTART IDENTITY CASCADE")
	if err != nil {
		t.Fatalf("Failed to truncate tables during teardown: %v", err)
	}
}

//func TestUserInteraction_Create(t *testing.T) {
//	repo := setupTestRepo(t)
//	//defer teardownTestRepo(t, repo.DB)
//
//	// Using valid mock IDs to test successful creation
//	testInteraction := &pb.UserInteraction{
//		UserId:          "2da4ef9e-5cf4-4cbc-8c69-8b6e222d43f0",
//		CompositionId:   "aac53533-6a60-4a0b-b9fa-fbff24f61f6b",
//		InteractionType: "like",
//	}
//
//	_, err := repo.Create(context.Background(), testInteraction)
//	assert.NoError(t, err)
//
//	var count int
//	err = repo.DB.QueryRow("SELECT COUNT(*) FROM user_interactions WHERE user_id = $1 AND composition_id = $2 AND interaction_type = $3",
//		testInteraction.UserId, testInteraction.CompositionId, testInteraction.InteractionType).Scan(&count)
//	assert.NoError(t, err)
//	assert.Equal(t, 1, count)
//}

//func TestUserInteraction_GetByID(t *testing.T) {
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db: %v", err)
//	}
//	defer db.Close()
//
//	repo := setupTestRepo(t)
//
//	interactionID := "3a576d57-e7f4-4567-92e6-c6a69c469cf6"
//
//	// Insert a test interaction into the user_interactions table if not already present
//	_, err = db.Exec(`INSERT INTO user_interactions (id, user_id, composition_id, interaction_type)
//        VALUES ($1, $2, $3, $4)
//        ON CONFLICT (id) DO NOTHING`,
//		interactionID, "2da4ef9e-5cf4-4cbc-8c69-8b6e222d43f0", "aac53533-6a60-4a0b-b9fa-fbff24f61f6b", "like")
//	if err != nil {
//		t.Fatalf("Failed to insert test data: %v", err)
//	}
//
//	interactionReq := pb.UserInteractionId{UserInteractionId: interactionID}
//
//	interactionRes, err := repo.GetById(context.Background(), &interactionReq)
//	assert.NoError(t, err)
//	assert.Equal(t, interactionID, interactionRes.Id)
//	assert.Equal(t, "2da4ef9e-5cf4-4cbc-8c69-8b6e222d43f0", interactionRes.UserId)
//	assert.Equal(t, "aac53533-6a60-4a0b-b9fa-fbff24f61f6b", interactionRes.CompositionId)
//	assert.Equal(t, "like", interactionRes.InteractionType)
//}

//func TestUserInteraction_Update(t *testing.T) {
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db: %v", err)
//	}
//	defer db.Close()
//
//	repo := &UserInteraction{DB: db}
//
//	interactionID := "3a576d57-e7f4-4567-92e6-c6a69c469cf6"
//
//	// Insert a test interaction into the user_interactions table if not already present
//	_, err = db.Exec(`INSERT INTO user_interactions (id, user_id, composition_id, interaction_type)
//        VALUES ($1, $2, $3, $4)
//        ON CONFLICT (id) DO NOTHING`,
//		interactionID, "2da4ef9e-5cf4-4cbc-8c69-8b6e222d43f0", "aac53533-6a60-4a0b-b9fa-fbff24f61f6b", "like")
//	if err != nil {
//		t.Fatalf("Failed to insert test data: %v", err)
//	}
//
//	updatedInteraction := &pb.UserInteractionRes{
//		Id:              interactionID,
//		UserId:          "2da4ef9e-5cf4-4cbc-8c69-8b6e222d43f0",
//		CompositionId:   "aac53533-6a60-4a0b-b9fa-fbff24f61f6b",
//		InteractionType: "love",
//	}
//
//	_, err = repo.Update(context.Background(), updatedInteraction)
//	assert.NoError(t, err)
//
//	var interactionType string
//	err = db.QueryRow("SELECT interaction_type FROM user_interactions WHERE id = $1", interactionID).Scan(&interactionType)
//	assert.NoError(t, err)
//	assert.Equal(t, "love", interactionType)
//}

func TestUserInteraction_Delete(t *testing.T) {
	db, err := postgres.Conn()
	if err != nil {
		t.Fatalf("Failed to open sql db: %v", err)
	}
	defer db.Close()

	repo := &UserInteraction{DB: db}

	interactionID := "3a576d57-e7f4-4567-92e6-c6a69c469cf6"

	// Insert a test interaction into the user_interactions table if not already present
	_, err = db.Exec(`INSERT INTO user_interactions (id, user_id, composition_id, interaction_type)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (id) DO NOTHING`,
		interactionID, "2da4ef9e-5cf4-4cbc-8c69-8b6e222d43f0", "aac53533-6a60-4a0b-b9fa-fbff24f61f6b", "like")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	interactionToDelete := &pb.UserInteractionId{UserInteractionId: interactionID}

	_, err = repo.Delete(context.Background(), interactionToDelete)
	assert.NoError(t, err)

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM user_interactions WHERE id = $1 and deleted_at != 0", interactionID).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}
