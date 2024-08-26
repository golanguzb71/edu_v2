package repository

import (
	"database/sql"
	"edu_v2/graph/model"
	"log"
)

type CollectionRepository struct {
	db *sql.DB
}

func NewCollectionRepository(db *sql.DB) *CollectionRepository {
	return &CollectionRepository{db: db}
}

func (r *CollectionRepository) Create(collection *model.Collection) error {
	_, err := r.db.Exec("INSERT INTO collections (title, questions) VALUES ($1 , $2)", collection.Title, collection.Questions)
	if err != nil {
		log.Printf("Error inserting collection: %v", err)
		return err
	}
	return nil
}

func (r *CollectionRepository) Get(id string) (*model.Collection, error) {
	var collection model.Collection
	err := r.db.QueryRow(`SELECT id , title , questions , created_at FROM collections where id=$1`, id).Scan(&collection.ID, &collection.Title, &collection.Questions, &collection.CreatedAt)
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
	rows, err := r.db.Query(`SELECT id, title, questions, created_at FROM collections`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var collections []*model.Collection
	for rows.Next() {
		var collect model.Collection
		err = rows.Scan(&collect.ID, &collect.Title, &collect.Questions, &collect.CreatedAt)
		if err != nil {
			return nil, err
		}
		collections = append(collections, &collect)
	}
	return collections, nil
}
