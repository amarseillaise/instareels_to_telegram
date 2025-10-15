package services

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/amarseillaise/instareels_to_telegram/services"
	"github.com/stretchr/testify/assert"
)

func TestGetReel(t *testing.T) {
	// Test setup
	shortcode := "DNvqPxyWDZp"
	tempPath := filepath.Join("temp", shortcode)
	err := os.MkdirAll(tempPath, 0755)
	assert.NoError(t, err)
	defer os.RemoveAll(tempPath)

	// Create test files
	descContent := "Test description"
	err = os.WriteFile(filepath.Join(tempPath, "desc.txt"), []byte(descContent), 0644)
	assert.NoError(t, err)

	dummyVideo := []byte("dummy video content")
	err = os.WriteFile(filepath.Join(tempPath, "video.mp4"), dummyVideo, 0644)
	assert.NoError(t, err)

	t.Run("successful reel retrieval", func(t *testing.T) {
		reel, err := services.GetReel(shortcode)
		assert.NoError(t, err)
		assert.Equal(t, descContent, reel.Description)
		assert.NotNil(t, reel.Video)
	})

	t.Run("non-existent shortcode", func(t *testing.T) {
		reel, err := services.GetReel("non-existent")
		assert.Error(t, err)
		assert.Empty(t, reel.Description)
		assert.Empty(t, reel.Video)
	})
}
