package utils

import (
	"log"
	"time"
)

func LogOperation(operation string, id int) {
	log.Printf("[%s] Operation: %s, AccountID: %d\n", time.Now().Format(time.RFC3339), operation, id)
}
