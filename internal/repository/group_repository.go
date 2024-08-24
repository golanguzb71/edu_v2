package repository

import (
	"database/sql"
	"edu_v2/graph/model"
	"fmt"
	"log"
	"time"
)

type GroupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) Create(group *model.Group) error {
	_, err := r.db.Exec("INSERT INTO groups (name, teacher_name, level, start_time, started_date, days_week) VALUES ($1, $2, $3,$4,$5,$6)",
		group.Name, group.TeacherName, group.Level, group.StartAt, group.StartedDate, group.DaysWeek)
	if err != nil {
		log.Printf("Error inserting group: %v", err)
		return err
	}
	return nil
}

func (r *GroupRepository) Get(id *string, orderLevel *bool) ([]*model.Group, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if id != nil {
		sql := `SELECT id, name, teacher_name, level, start_time, started_date, days_week, created_at FROM groups WHERE id = $1`
		rows, err = r.db.Query(sql, id)
	} else {
		sql := `SELECT id, name, teacher_name, level, start_time, started_date, days_week, created_at FROM groups`
		if orderLevel != nil {
			sql += ` order by level`
		}
		rows, err = r.db.Query(sql)
	}

	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var groups []*model.Group

	for rows.Next() {
		var (
			id          int
			name        string
			teacherName string
			level       model.GroupLevel
			startAt     string
			startDate   string
			daysWeek    model.DaysWeek
			createdAt   time.Time
		)

		if err := rows.Scan(&id, &name, &teacherName, &level, &startAt, &startDate, &daysWeek, &createdAt); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		group := &model.Group{
			ID:          fmt.Sprintf("%d", id),
			Name:        name,
			TeacherName: teacherName,
			Level:       level,
			StartAt:     startAt,
			StartedDate: startDate,
			DaysWeek:    daysWeek,
			CreatedAt:   createdAt.Format(time.RFC3339),
		}

		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return groups, nil
}

func (r *GroupRepository) Update(group *model.Group) error {
	_, err := r.db.Exec("UPDATE groups SET name = $1, teacher_name = $2, level = $3 , days_week=$4,start_time=$5,started_date=$6 WHERE id = $7",
		group.Name, group.TeacherName, group.Level, group.DaysWeek, group.StartAt, group.StartedDate, group.ID)
	if err != nil {
		log.Printf("Error updating group: %v", err)
		return err
	}
	return nil
}

// Delete removes a group from the database.
func (r *GroupRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM groups WHERE id = $1", id)
	if err != nil {
		log.Printf("Error deleting group: %v", err)
		return err
	}
	return nil
}
