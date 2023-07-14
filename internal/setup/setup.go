package setup

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func Setup() {
	_, b, _, _ := runtime.Caller(0)
	fname := filepath.Join(filepath.Dir(b), "..", "..", "config", ".env")
	err := godotenv.Load(fname)
	if err != nil {
		log.Fatal("Error in loading env file, place the file in config dir")
	}
}
