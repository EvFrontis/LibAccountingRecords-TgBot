package google_drive_test

import (
	"os"
	"testing"

	google "github.com/EvFrontis/LibAccountingRecords-TgBot/pkg/google_drive"
	"github.com/joho/godotenv"
)

func TestAppendToFile(t *testing.T) {
	godotenv.Load()
	fileID := os.Getenv("FILE_ID") // ID of the file to be modified

	got := google.AppendToFile(fileID, []byte("Test. "))
	if got != nil {
		t.Errorf("AppendToFile(%s, 'Test. ') = %d; want nil", fileID, got)
	}
}
