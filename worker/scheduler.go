package worker

import (
	"context"
	"fmt"
	"log"

	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
)

// Function to delete expired sessions
func deleteExpiredSessions(ctx context.Context, repository db.DatabaseContract) {
	err := repository.DeleteSession(ctx)
	if err != nil {
		log.Printf("Error deleting expired sessions: %v\n", err)
		return
	}
	fmt.Println("Expired sessions deleted successfully.")
}

// Function to set up the scheduler
func SetupScheduler(ctx context.Context, repository db.DatabaseContract) {
	c := cron.New()

	_, err := c.AddFunc("0 */12 * * *", func() {
		deleteExpiredSessions(ctx, repository)
	})
	if err != nil {
		log.Fatalf("Error scheduling cron job: %v\n", err)
	}

	c.Start()

	// keep program running
	select {}
}
