package worker

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Worker struct {
	pool *pgxpool.Pool
}

func NewWorker(pool *pgxpool.Pool) *Worker {
	return &Worker{pool: pool}
}

func (w *Worker) Start(ctx context.Context, count int) {

	for i := 0; i < count; i++ {

		go w.workerLoop(ctx, i)
	}
}

func (w *Worker) workerLoop(ctx context.Context, id int) {

	log.Println("Worker started:", id)

	for {

		select {

		case <-ctx.Done():
			log.Println("Worker shutting down:", id)
			return

		default:

			jobID, email, err := w.fetchJob(ctx)

			if err != nil {

				time.Sleep(2 * time.Second)
				continue
			}

			log.Println("Worker", id, "processing", email)

			// simulate verification for now
			time.Sleep(1 * time.Second)

			err = w.completeJob(ctx, jobID)

			if err != nil {

				log.Println("error completing job:", err)
			}
		}
	}
}


func (w *Worker) fetchJob(ctx context.Context) (string, string, error) {

	tx, err := w.pool.Begin(ctx)
	if err != nil {
		return "", "", err
	}

	defer tx.Rollback(ctx)

	var jobID string
	var email string

	err = tx.QueryRow(
		ctx,
		`
		SELECT id, email
		FROM jobs
		WHERE status='pending'
		ORDER BY created_at
		LIMIT 1
		FOR UPDATE SKIP LOCKED
		`,
	).Scan(&jobID, &email)

	if err != nil {
		return "", "", err
	}

	_, err = tx.Exec(
		ctx,
		`
		UPDATE jobs
		SET status='processing',
		    updated_at = NOW()
		WHERE id=$1
		`,
		jobID,
	)

	if err != nil {
		return "", "", err
	}

	err = tx.Commit(ctx)

	if err != nil {
		return "", "", err
	}

	return jobID, email, nil
}


func (w *Worker) completeJob(ctx context.Context, jobID string) error {

	_, err := w.pool.Exec(
		ctx,
		`
		UPDATE jobs
		SET status='completed',
		    updated_at = NOW()
		WHERE id=$1
		`,
		jobID,
	)

	return err
}

func (w *Worker) failJob(ctx context.Context, jobID string) error {

	_, err := w.pool.Exec(
		ctx,
		`
		UPDATE jobs
		SET status='failed',
		    attempts = attempts + 1,
		    updated_at = NOW()
		WHERE id=$1
		`,
		jobID,
	)

	return err
}