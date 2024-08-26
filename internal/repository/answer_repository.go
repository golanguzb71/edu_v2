package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type AnswerRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewAnswerRepository(db *sql.DB, rdb *redis.Client) *AnswerRepository {
	return &AnswerRepository{db: db, rdb: rdb}
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

func (r *AnswerRepository) CreateStudentAnswer(collectionId string, answers []*string, code string) error {
	chatId, err := r.rdb.Get(context.TODO(), code).Result()
	if err != nil {
		return errors.New("error while getting user information please update your code => https://t.me/codevanbot")
	}
	var userId int
	err = r.db.QueryRow(`SELECT id FROM users WHERE chat_id=$1`, chatId).Scan(&userId)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(`INSERT INTO user_collection(user_id, answer_field, collection_id) values ($1,$2,$3)`, userId, pq.Array(answers), collectionId)
	if err != nil {
		return err
	}
	return nil
}
