package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
)

type AnswerRepository struct {
	db *sql.DB
}

func NewAnswerRepository(db *sql.DB) *AnswerRepository {
	return &AnswerRepository{db: db}
}

func (r *AnswerRepository) CreateAnswer(collectionId *string, answers []*string, isUpdated *bool) error {
	answersArray := pq.Array(answers)

	if isUpdated != nil && *isUpdated {
		var isHave bool
		err := r.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM answers WHERE collection_id=$1)`, *collectionId).Scan(&isHave)
		if err != nil {
			return fmt.Errorf("error checking for existing answers: %v", err)
		}
		if !isHave {
			return errors.New("cannot update because there are no existing answers for this collection")
		}
		_, err = r.db.Exec(`UPDATE answers SET answer_field=$1  WHERE collection_id=$2`, answersArray, *collectionId)
		if err != nil {
			return fmt.Errorf("error updating answers: %v", err)
		}
	} else {
		_, err := r.db.Exec(`INSERT INTO answers (collection_id, answer_field) VALUES ($1, $2)`, *collectionId, answersArray)
		if err != nil {
			return fmt.Errorf("error inserting answers: %v", err)
		}
	}

	return nil
}

func (r *AnswerRepository) DeleteAnswer(collectionId *string) error {
	_, err := r.db.Exec(`DELETE FROM answers where collection_id=$1`, collectionId)
	if err != nil {
		return err
	}
	return nil
}
