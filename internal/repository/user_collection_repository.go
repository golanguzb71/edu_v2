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
		rows, err := r.db.Query(`SELECT answer_field, collection_id, created_at FROM user_collection WHERE user_id=$1 order by created_at desc LIMIT $2 OFFSET $3`, studentId, size, offset)
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

	_ = r.db.QueryRow(`SELECT id FROM users WHERE chat_id=$1`, chatId).Scan(&studentId)

	rows, err := r.db.Query(`SELECT answer_field, collection_id , created_at FROM user_collection WHERE user_id=$1 order by created_at desc LIMIT $2 OFFSET $3`, studentId, size, offset)
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
		_ = r.db.QueryRow(`SELECT answer_field FROM answers WHERE collection_id=$1`, UCT.CollectionID).Scan(pq.Array(&trueAnswers))
		trueCount := 0
		falseCount := 0
		var answerFields []*model.AnswerField

		maxLength := len(trueAnswers)
		if len(studentAnswers) > maxLength {
			maxLength = len(studentAnswers)
		}

		for i := 0; i < maxLength; i++ {
			var isTrue bool
			var studentAnswer, trueAnswer string

			if i < len(trueAnswers) {
				trueAnswer = trueAnswers[i]
			}
			if i < len(studentAnswers) {
				studentAnswer = studentAnswers[i]
			}

			if i < len(trueAnswers) {
				normalizedTrueAnswer := strings.ToLower(strings.TrimSpace(trueAnswer))
				if i < len(studentAnswers) {
					normalizedStudentAnswer := strings.ToLower(strings.TrimSpace(studentAnswer))
					if normalizedTrueAnswer == normalizedStudentAnswer {
						isTrue = true
						trueCount++
					} else {
						isTrue = false
						falseCount++
					}
				} else {
					isTrue = false
					falseCount++
				}

				answerFields = append(answerFields, &model.AnswerField{
					StudentAnswer: &studentAnswer,
					TrueAnswer:    trueAnswer,
					IsTrue:        &isTrue,
				})
			} else {
				answerFields = append(answerFields, &model.AnswerField{
					StudentAnswer: nil,
					TrueAnswer:    trueAnswer,
					IsTrue:        new(bool),
				})
				falseCount++
			}
		}

		UCT.AnswerField = answerFields
		UCT.TrueCount = &trueCount
		UCT.FalseCount = &falseCount

		level := determineLevel(trueCount, len(trueAnswers))

		if level != "" {
			var group model.Group
			_ = r.db.QueryRow(`SELECT id, name, teacher_name, level, start_time, started_date, days_week, created_at FROM groups WHERE level=$1`, level).
				Scan(&group.ID, &group.Name, &group.TeacherName, &group.Level, &group.StartAt, &group.StartedDate, &group.DaysWeek, &group.CreatedAt)
			UCT.RequestGroup = append(UCT.RequestGroup, &group)
		}
		UCTL = append(UCTL, &UCT)
	}

	return UCTL, nil
}

func determineLevel(trueCount, totalQuestions int) string {
	if totalQuestions == 0 {
		return ""
	}

	percentage := float64(trueCount) / float64(totalQuestions) * 100

	switch {
	case percentage <= 20:
		return "BEGINNER"
	case percentage > 20 && percentage <= 40:
		return "ELEMENTARY"
	case percentage > 40 && percentage <= 60:
		return "PRE_INTERMEDIATE"
	case percentage > 60 && percentage <= 80:
		return "INTERMEDIATE"
	case percentage > 80:
		return "UPPER_INTERMEDIATE"
	default:
		return ""
	}
}
