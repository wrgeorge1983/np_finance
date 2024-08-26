package db

import (
	"github.com/ostafen/clover/v2"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
)

type DB struct {
	db *clover.DB
}

func (db *DB) Close() error {
	return db.db.Close()
}

type Docable interface {
	Doc() *document.Document
}

type User struct {
	ID       string `json:"id" clover:"id"`
	Username string `json:"username" clover:"username"`
}

func (u *User) Doc() *document.Document {
	doc := document.NewDocument()
	doc.Set("id", u.ID)
	doc.Set("username", u.Username)
	return doc
}

type Run struct {
	Title     string            `json:"title" clover:"title"`
	UserID    string            `json:"user_id" clover:"user_id"`
	Inputs    map[string]string `json:"inputs" clover:"inputs"`
	Schedule  string            `json:"schedule" clover:"schedule"`
	Worksheet string            `json:"worksheet" clover:"worksheet"`
	Outputs   map[string]string `json:"outputs" clover:"outputs"`
}

func (r *Run) Doc() *document.Document {
	doc := document.NewDocument()
	doc.Set("title", r.Title)
	doc.Set("user_id", r.UserID)
	doc.Set("inputs", r.Inputs)
	doc.Set("schedule", r.Schedule)
	doc.Set("worksheet", r.Worksheet)
	doc.Set("outputs", r.Outputs)
	return doc
}

func NewDB(filename string) *DB {
	db, err := clover.Open(filename)

	if err != nil {
		panic(err)
	}

	thisDb := &DB{db: db}
	err = thisDb.Init()
	if err != nil {
		panic(err)
	}

	return thisDb
}

func (db *DB) insert(collectionName string, doc Docable) error {
	return db.db.Insert(collectionName, doc.Doc())
}

func (db *DB) InsertUser(u *User) error {
	return db.insert("users", u)
}

func (db *DB) InsertRun(r *Run) error {
	return db.insert("runs", r)
}

func (db *DB) AllUsers() ([]User, error) {
	q := query.NewQuery("users")
	docs, err := db.db.FindAll(q)
	if err != nil {
		return nil, err
	}

	users := make([]User, 0, len(docs))
	for _, doc := range docs {
		u := &User{}
		err := doc.Unmarshal(u)
		if err != nil {
			return nil, err
		}
		users = append(users, *u)
	}

	return users, nil
}

func (db *DB) AllRuns() ([]Run, error) {
	q := query.NewQuery("runs")
	docs, err := db.db.FindAll(q)
	if err != nil {
		return nil, err
	}

	runs := make([]Run, 0, len(docs))
	for _, doc := range docs {
		r := &Run{}
		err := doc.Unmarshal(r)
		if err != nil {
			return nil, err
		}
		runs = append(runs, *r)
	}

	return runs, nil
}

func (db *DB) Init() error {

	for _, name := range []string{"users", "runs"} {
		if has, err := db.db.HasCollection(name); !has {
			if err != nil {
				panic(err)
			}
			err = db.db.CreateCollection(name)
			if err != nil {
				panic(err)
			}
		}
	}

	users := []*User{
		{ID: "1", Username: "alice"},
		{ID: "2", Username: "bob"},
		{ID: "3", Username: "charlie"},
	}

	runs := []*Run{
		{
			Title:     "Run 1",
			UserID:    "1",
			Inputs:    map[string]string{"input1": "value1", "input2": "value2"},
			Schedule:  "schedule1",
			Worksheet: "worksheet1",
			Outputs:   map[string]string{"output1": "value1", "output2": "value2"},
		},
		{
			Title:     "Run 2",
			UserID:    "2",
			Inputs:    map[string]string{"input1": "value1", "input2": "value2"},
			Schedule:  "schedule2",
			Worksheet: "worksheet2",
			Outputs:   map[string]string{"output1": "value1", "output2": "value2"},
		},
		{
			Title:     "Run 3",
			UserID:    "3",
			Inputs:    map[string]string{"input1": "value1", "input2": "value2"},
			Schedule:  "schedule3",
			Worksheet: "worksheet3",
			Outputs:   map[string]string{"output1": "value1", "output2": "value2"},
		},
	}

	for _, u := range users {
		err := db.InsertUser(u)
		if err != nil {
			return err
		}
	}

	for _, r := range runs {
		err := db.InsertRun(r)
		if err != nil {
			return err
		}
	}

	return nil

}

func (db *DB) Purge() error {
	collections, err := db.db.ListCollections()
	if err != nil {
		return err
	}
	for _, name := range collections {
		err := db.db.DropCollection(name)
		if err != nil {
			return err
		}
	}
	return nil
}
