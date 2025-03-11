package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	migrationPath, err := filepath.Abs("./cmd/db/migrations")
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}

	if _, err := os.Stat(migrationPath); os.IsNotExist(err) {
		log.Fatalf("Migration path does not exist: %v", migrationPath)
	}

	// Run the migrate command
	//cmd := exec.Command("migrate", "-path", migrationPath, "-database", cfg.Inst().DatabaseDSN, "-verbose", "up")
	//log.Println(cmd.Args)
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//
	//err = cmd.Run()
	//if err != nil {
	//	log.Fatalf("Error running migrate command: %v", err)
	//}
}
