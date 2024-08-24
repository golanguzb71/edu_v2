package repository

import (
	"context"
	"database/sql"
	"edu_v2/graph/model"
	"errors"
	"github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"strings"
)

type UserCollectionRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewUserCollectionRepository(db *sql.DB, rdb *redis.Client) *UserCollectionRepository {
	return &UserCollectionRepository{db: db, rdb: rdb}
}
func (r *UserCollectionRepository) GetStudentTestExams(code *string, studentId *string, page *int, size *int) ([]*model.UserCollectionTestExams, error) {
	if page == nil || *page < 1 {
		p := 1
		page = &p
	}
	if size == nil || *size < 1 {
		s := 10
		size = &s
	}

	offset := *size * (*page - 1)

	if studentId != nil {
		if code != nil && *code != "KEY_ADM" {
			return nil, errors.New("forbidden: you don't have access to see this")
		}
		rows, err := r.db.Query(`SELECT answer_field, collection_id, created_at FROM user_collection WHERE user_id=$1 LIMIT $2 OFFSET $3`, studentId, size, offset)
		if err != nil {
			return nil, err
		}
		return studentTestExamsForStudent(rows, r)
	}

	chatId, err := r.rdb.Get(context.Background(), *code).Result()
	if err != nil {
		return nil, errors.New("error while getting student information; please update key by Telegram bot => https://t.me/codevanbot")
	}
	if chatId == "" {
		return nil, errors.New("error while getting student information; please update key by Telegram bot => https://t.me/codevanbot")
	}

	err = r.db.QueryRow(`SELECT id FROM users WHERE chat_id=$1`, chatId).Scan(&studentId)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(`SELECT answer_field, collection_id FROM user_collection WHERE user_id=$1 LIMIT $2 OFFSET $3`, studentId, size, offset)
	if err != nil {
		return nil, err
	}
	return studentTestExamsForStudent(rows, r)
}

func studentTestExamsForStudent(rows *sql.Rows, r *UserCollectionRepository) ([]*model.UserCollectionTestExams, error) {
	var UCTL []*model.UserCollectionTestExams
	for rows.Next() {
		var UCT model.UserCollectionTestExams
		var studentAnswers []string

		err := rows.Scan(pq.Array(&studentAnswers), &UCT.CollectionID, &UCT.CreatedAt)
		if err != nil {
			return nil, err
		}

		var trueAnswers []string
		err = r.db.QueryRow(`SELECT answer_field FROM answers WHERE collection_id=$1`, UCT.CollectionID).Scan(pq.Array(&trueAnswers))
		if err != nil {
			return nil, err
		}

		trueCount := 0
		falseCount := 0

		for i := range trueAnswers {
			var isTrue bool
			normalizedTrueAnswer := strings.ToLower(strings.TrimSpace(trueAnswers[i]))
			normalizedStudentAnswer := strings.ToLower(strings.TrimSpace(studentAnswers[i]))

			if i < len(studentAnswers) && normalizedTrueAnswer == normalizedStudentAnswer {
				isTrue = true
				trueCount++
			} else {
				isTrue = false
				falseCount++
			}

			UCT.AnswerField = append(UCT.AnswerField, &model.AnswerField{
				StudentAnswer: &studentAnswers[i],
				TrueAnswer:    trueAnswers[i],
				IsTrue:        &isTrue,
			})
		}
		UCT.TrueCount = &trueCount
		UCT.FalseCount = &falseCount

		level := determineLevel(trueCount)

		if level != "" {
			var group model.Group
			err := r.db.QueryRow(`SELECT id, name, teacher_name, level, start_time, started_date, days_week, created_at FROM groups WHERE level=$1`, level).
				Scan(&group.ID, &group.Name, &group.TeacherName, &group.Level, &group.StartAt, &group.StartedDate, &group.DaysWeek, &group.CreatedAt)
			if err != nil {
				return nil, err
			}

			UCT.RequestGroup = append(UCT.RequestGroup, &group)
		}
		UCTL = append(UCTL, &UCT)
	}

	return UCTL, nil
}

func determineLevel(trueCount int) string {
	switch {
	case trueCount >= 0 && trueCount <= 10:
		return "BEGINNER"
	case trueCount > 10 && trueCount <= 20:
		return "ELEMENTARY"
	case trueCount > 20 && trueCount <= 30:
		return "PRE_INTERMEDIATE"
	case trueCount > 30 && trueCount <= 40:
		return "INTERMEDIATE"
	case trueCount > 40 && trueCount <= 50:
		return "UPPER_INTERMEDIATE"
	case trueCount > 50 && trueCount <= 60:
		return "ADVANCED"
	case trueCount > 60 && trueCount <= 70:
		return "PROFICIENT"
	default:
		return ""
	}
}
