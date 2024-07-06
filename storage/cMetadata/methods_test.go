// cmetadata_test.go
package cMetadata

import (
	"testing"

	"MusicMesh/Discovery-MusicMesh/storage/postgres"
)

func setupTestRepo(t *testing.T) *CmetadataRepo {
	db, err := postgres.Conn()
	if err != nil {
		t.Fatalf("Failed to connect to the test database: %v", err)
	}

	// Clean up test database before each test
	//_, err = db.Exec("TRUNCATE composition_metadata RESTART IDENTITY CASCADE")
	//if err != nil {
	//	t.Fatalf("Failed to truncate tables: %v", err)
	//}

	return &CmetadataRepo{DB: db}
}

//func TestCmetadataRepo_Create(t *testing.T) {
//	repo := setupTestRepo(t)
//
//	testMetadata := &pb.CompositionMetadata{
//		CompositionId: "36be4369-6df1-444b-9b8c-1e52c36fbcf7",
//		Genre:         "rock",
//		Tags:          "tag1",
//	}
//
//	_, err := repo.Create(context.Background(), testMetadata)
//	assert.NoError(t, err)
//
//	var count int
//	err = repo.DB.QueryRow("SELECT COUNT(*) FROM composition_metadata WHERE composition_id = $1 AND genre = $2 AND tags = $3",
//		testMetadata.CompositionId, testMetadata.Genre, testMetadata.Tags).Scan(&count)
//	assert.NoError(t, err)
//	assert.Equal(t, 1, count)
//}
//
//func TestCmetadataRepo_GetTrending(t *testing.T) {
//	repo := setupTestRepo(t)
//
//	// Insert test data
//	_, err := repo.DB.Exec(`INSERT INTO composition_metadata (composition_id, genre, tags, listen_count, like_count)
//     VALUES ('36be4369-6df1-444b-9b8c-1e52c36fbcf9', 'rock', 'tag1', 100, 10),
//            ('36be4369-6df1-444b-9b8c-1e52c36fbcf8', 'pop', 'tag3', 200, 20)`)
//	assert.NoError(t, err)
//
//	trending, err := repo.GetTrending(context.Background(), &pb.Void{})
//	assert.NoError(t, err)
//	assert.Len(t, trending.Compositions, 2)
//	assert.Equal(t, "36be4369-6df1-444b-9b8c-1e52c36fbcf8", trending.Compositions[0].CompositionId)
//	assert.Equal(t, int64(200), trending.Compositions[0].ListenCount)
//}

//func TestCmetadataRepo_GetRecommended(t *testing.T) {
//	repo := setupTestRepo(t)
//
//	// Insert test data
//	_, err := repo.DB.Exec(`INSERT INTO composition_metadata (composition_id, genre, tags, listen_count, like_count)
//	VALUES ('36be4369-6df1-444b-9b8c-1e52c36fbcf9', 'rock', 'tag1', 100, 10),
//	      ('36be4369-6df1-444b-9b8c-1e52c36fbcf8', 'pop', 'tag3', 200, 20)`)
//	assert.NoError(t, err)
//
//	_, err = repo.DB.Exec(`INSERT INTO user_interactions (user_id, composition_id)
//	VALUES ('2da4ef9e-5cf4-4cbc-8c69-8b6e222d43f0', '36be4369-6df1-444b-9b8c-1e52c36fbcf9'), ('2da4ef9e-5cf4-4cbc-8c69-8b6e222d43f0', '36be4369-6df1-444b-9b8c-1e52c36fbcf8')`)
//	assert.NoError(t, err)
//
//	recommended, err := repo.GetRecommended(context.Background(), &pb.UserId{UserId: "2da4ef9e-5cf4-4cbc-8c69-8b6e222d43f0"})
//	assert.NoError(t, err)
//	assert.Len(t, recommended.Compositions, 2)
//	assert.Equal(t, "36be4369-6df1-444b-9b8c-1e52c36fbcf8", recommended.Compositions[0].CompositionId)
//	assert.Equal(t, int64(20), recommended.Compositions[0].LikeCount)
//}

//func TestCmetadataRepo_GetByGenre(t *testing.T) {
//	repo := setupTestRepo(t)
//
//	// Insert test data
//	_, err := repo.DB.Exec(`INSERT INTO composition_metadata (composition_id, genre, tags, listen_count, like_count)
//	 VALUES ('36be4369-6df1-444b-9b8c-1e52c36fbcf9', 'rock', 'tag1', 100, 10),
//	        ('36be4369-6df1-444b-9b8c-1e52c36fbcf8', 'rock', 'tag3', 200, 20),
//	        ('36be4369-6df1-444b-9b8c-1e52c36fbcf7', 'pop', 'tag4', 300, 30)`)
//	assert.NoError(t, err)
//
//	byGenre, err := repo.GetByGenre(context.Background(), &pb.GenreRequest{Genre: "rock"})
//	assert.NoError(t, err)
//	assert.Len(t, byGenre.Compositions, 2)
//	assert.Equal(t, "36be4369-6df1-444b-9b8c-1e52c36fbcf8", byGenre.Compositions[0].CompositionId)
//}

//func TestCmetadataRepo_Update(t *testing.T) {
//	repo := setupTestRepo(t)
//
//	// Insert test data
//	_, err := repo.DB.Exec(`INSERT INTO composition_metadata (composition_id, genre, tags, listen_count, like_count)
//       VALUES ('36be4369-6df1-444b-9b8c-1e52c36fbcf9', 'rock', 'tag1', 100, 10)`)
//	assert.NoError(t, err)
//
//	updatedMetadata := &pb.CompositionRes{
//		CompositionId: "36be4369-6df1-444b-9b8c-1e52c36fbcf9",
//		Genre:         "jazz",
//		Tags:          "newtag1",
//		ListenCount:   150,
//		LikeCount:     15,
//	}
//
//	_, err = repo.Update(context.Background(), updatedMetadata)
//	assert.NoError(t, err)
//
//	var genre string
//	var listenCount, likeCount int64
//	err = repo.DB.QueryRow(`SELECT genre, listen_count, like_count FROM composition_metadata WHERE composition_id = $1`,
//		"36be4369-6df1-444b-9b8c-1e52c36fbcf9").Scan(&genre, &listenCount, &likeCount)
//	assert.NoError(t, err)
//	assert.Equal(t, "jazz", genre)
//	assert.Equal(t, int64(150), listenCount)
//	assert.Equal(t, int64(15), likeCount)
//}

//func TestCmetadataRepo_Delete(t *testing.T) {
//	repo := setupTestRepo(t)
//
//	// Insert test data
//	_, err := repo.DB.Exec(`INSERT INTO composition_metadata (composition_id, genre, tags, listen_count, like_count)
//      VALUES ('36be4369-6df1-444b-9b8c-1e52c36fbcf1', 'rock', 'tag1', 100, 10)`)
//	assert.NoError(t, err)
//
//	_, err = repo.Delete(context.Background(), &pb.CompositionMetadataId{MetadataId: "36be4369-6df1-444b-9b8c-1e52c36fbcf1"})
//	assert.NoError(t, err)
//
//	var count int
//	err = repo.DB.QueryRow("SELECT COUNT(*) FROM composition_metadata WHERE composition_id = $1", "36be4369-6df1-444b-9b8c-1e52c36fbcf9").Scan(&count)
//	assert.NoError(t, err)
//	assert.Equal(t, 0, count)
//}
