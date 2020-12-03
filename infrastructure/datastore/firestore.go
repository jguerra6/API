package datastore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type firestoreDB struct{}

//define the projectId
const (
	projectID string = "bootcamp-d79dd"
)

//NewFirestoreDB will return the db object
func NewFirestoreDB() Database {
	return &firestoreDB{}
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

func (*firestoreDB) GetAll(tableName string) ([]map[string]interface{}, error) {
	ctx := context.Background()
	client := createClient(ctx)

	docs := []map[string]interface{}{}
	iter := client.Collection(tableName).Documents(ctx)

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to get all elements from %v. \nMessage: %v", tableName, err)
			return nil, err
		}

		dict := map[string]interface{}{}
		for key, value := range doc.Data() {

			//valType := reflect.ValueOf(value).Kind()
			//valType := reflect.TypeOf(value).(string)
			switch value.(type) {
			case string:
				dict[key] = value.(string)
			case int64:
				dict[key] = value.(int64)
			case float64:
				dict[key] = value.(float64)
			}

			//fmt.Printf("key %v has value of %v and a type of %T\n", key, value, value)

		}
		dict["id"] = doc.Ref.ID

		docs = append(docs, dict)

	}
	//fmt.Println(docs)
	//fmt.Println(&docs)
	return docs, nil

}

func (*firestoreDB) GetItemByID(tableName string, id string) (map[string]interface{}, error) {
	ctx := context.Background()
	client := createClient(ctx)

	dsnap, err := client.Collection(tableName).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	m := dsnap.Data()
	m["id"] = dsnap.Ref.ID

	return m, nil
}
func (*firestoreDB) DeleteItem(tableName string, id string) error {
	ctx := context.Background()
	client := createClient(ctx)

	_, err := client.Collection(tableName).Doc(id).Delete(ctx)
	if err != nil {
		return err
	}

	return nil

}
func (*firestoreDB) UpdateItem(tableName string, id string) {

}
func (*firestoreDB) AddItem(tableName string, item map[string]interface{}) (map[string]interface{}, error) {
	ctx := context.Background()

	client := createClient(ctx)
	ref, _, err := client.Collection(tableName).Add(ctx, item)

	if err != nil {
		log.Fatal("Failed adding item: ", err)
		return nil, err
	}

	defer client.Close()

	item["id"] = ref.ID

	return item, nil
}
