package tasksmock

import (
	models "backend/db/tasks"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func InsertDB(db *gorm.DB) error {
	_, b, _, _ := runtime.Caller(0)
	fPath := filepath.Join(filepath.Dir(b), "mock.json")
	f, err := os.ReadFile(fPath)
	if err != nil {
		return fmt.Errorf("could not open task mock json file: %v, %s", err, fPath)
	}
	var ts []models.Task
	err = json.Unmarshal(f, &ts)
	if err != nil {
		return fmt.Errorf("err unmarshaling mock data: %v", err)
	}
	for _, t := range ts {
		err := db.Create(&t).Error
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				fmt.Println("database contains mock.json data")
				return nil
			} else {
				return fmt.Errorf("err inserting mock task: %v", err)
			}
		}
	}
	return nil
}
