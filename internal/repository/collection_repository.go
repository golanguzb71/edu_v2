package repository

import (
	"database/sql"
	"edu_v2/graph/model"
	"edu_v2/internal/utils"
	"errors"
	"log"
)

type CollectionRepository struct {
	db *sql.DB
}

func NewCollectionRepository(db *sql.DB) *CollectionRepository {
	return &CollectionRepository{db: db}
}

func (r *CollectionRepository) Create(collection *model.Collection) (*int, error) {
	var value int
	err := r.db.QueryRow("INSERT INTO collections (title, questions) VALUES ($1 , $2) RETURNING id", collection.Title, collection.Questions).Scan(&value)
	if err != nil {
		log.Printf("Error inserting collection: %v", err)
		return nil, err
	}
	return &value, nil
}

func (r *CollectionRepository) Get(id string) (*model.Collection, error) {
	var collection model.Collection
	err := r.db.QueryRow(`SELECT id , title , questions , created_at ,is_active FROM collections where id=$1`, id).Scan(&collection.ID, &collection.Title, &collection.Questions, &collection.CreatedAt, &collection.IsActive)
	if err != nil {
		return nil, err
	}

	return &collection, nil
}

func (r *CollectionRepository) Update(collection *model.Collection) error {
	//_, err := r.db.Exec(
	//	"UPDATE collections SET title = $1, questions_url = $2 WHERE id = $3",
	//	collection.Title,
	//	pq.Array(collection.ImageURL),
	//	collection.ID,
	//)
	//if err != nil {
	//	log.Printf("Error updating collection: %v", err)
	//	return err
	//}
	//return nil
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
	rows, err := r.db.Query(`SELECT id, title, questions, created_at , is_active FROM collections`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var collections []*model.Collection
	for rows.Next() {
		var collect model.Collection
		err = rows.Scan(&collect.ID, &collect.Title, &collect.Questions, &collect.CreatedAt, &collect.IsActive)
		if err != nil {
			return nil, err
		}
		collections = append(collections, &collect)
	}
	return collections, nil
}

func (r *CollectionRepository) UpdateCollectionActive(id string) error {
	var checker bool
	_ = r.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM collections where id=$1)`, id).Scan(&checker)
	if !checker {
		return errors.New("collection not found")
	}
	_, err := r.db.Exec(`WITH update_old as (
    UPDATE collections
    SET is_active=false
    WHERE is_active=true
)
UPDATE collections SET is_active=true where id=$1;`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *CollectionRepository) GetCollectionActive() (*model.Collection, error) {
	var collection model.Collection
	err := r.db.QueryRow(`SELECT id, title, questions, created_at, is_active FROM collections where is_active=true`).Scan(&collection.ID, &collection.Title, &collection.Questions, &collection.CreatedAt, &collection.IsActive)
	if err != nil {
		utils.SendMessage(err.Error(), 6805374430)
		return nil, err
	}
	return &collection, nil
}
