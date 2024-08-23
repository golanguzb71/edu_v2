package repository

import (
	"database/sql"
	"edu_v2/graph/model"
	"github.com/lib/pq"
	"log"
)

type CollectionRepository struct {
	db *sql.DB
}

func NewCollectionRepository(db *sql.DB) *CollectionRepository {
	return &CollectionRepository{db: db}
}

// Create inserts a new collection into the database.
func (r *CollectionRepository) Create(collection *model.Collection) error {
	_, err := r.db.Exec("INSERT INTO collections (title, questions_url) VALUES ($1 , $2)", collection.Title, pq.Array(collection.ImageURL))
	if err != nil {
		log.Printf("Error inserting collection: %v", err)
		return err
	}
	return nil
}

func (r *CollectionRepository) Get(id string) (*model.Collection, error) {
	var collection model.Collection
	var questionUrl pq.StringArray
	err := r.db.QueryRow(`SELECT id , title , questions_url , created_at FROM collections where id=$1`, id).Scan(&collection.ID, &collection.Title, &questionUrl, &collection.CreatedAt)
	if err != nil {
		return nil, err
	}
	collection.ImageURL = questionUrl
	return &collection, nil
}

func (r *CollectionRepository) Update(collection *model.Collection) error {
	_, err := r.db.Exec(
		"UPDATE collections SET title = $1, questions_url = $2 WHERE id = $3",
		collection.Title,
		pq.Array(collection.ImageURL),
		collection.ID,
	)
	if err != nil {
		log.Printf("Error updating collection: %v", err)
		return err
	}
	return nil
}
func (r *CollectionRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM collections WHERE id = $1", id)
	if err != nil {
		log.Printf("Error deleting collection: %v", err)
		return err
	}
	return nil
}
func (r *CollectionRepository) GetCollections() ([]*model.Collection, error) {
	rows, err := r.db.Query(`SELECT id, title, questions_url, created_at FROM collections`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var collections []*model.Collection
	for rows.Next() {
		var collect model.Collection
		var questionsURL pq.StringArray
		err = rows.Scan(&collect.ID, &collect.Title, &questionsURL, &collect.CreatedAt)
		if err != nil {
			return nil, err
		}
		collect.ImageURL = questionsURL
		collections = append(collections, &collect)
	}
	return collections, nil
}
