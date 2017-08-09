package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Post struct {
	ID      bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Slug    string        `bson:",omitempty" json:"slug"`
	Title   string        `bson:",omitempty" json:"title"`
	Content string        `bson:",omitempty" json:"content"`
	Updated time.Time     `bson:",omitempty" json:"updated"`
	Created time.Time     `bson:",omitempty" json:"created"`
}

func init() {
	ensureIndexes(&Post{})
}

func (p *Post) Collection() string {
	return "posts"
}

func (p *Post) Unique() bson.M {
	return bson.M{"slug": p.Slug}
}

func (p *Post) Indexes() []mgo.Index {
	slug := mgo.Index{
		Unique:   true,
		DropDups: true,
		Key: []string{
			"slug",
		},
	}
	return []mgo.Index{
		slug,
	}
}

func (p *Post) BeforeCreate() {
	p.ID = bson.NewObjectId()
	p.Slug = genUniqSlug(8)
	p.Created = time.Now()
	p.Updated = p.Created
}

func (p *Post) BeforeUpdate() {
	p.Updated = time.Now()
}

func genUniqSlug(n int) string {
	var post *Post

	slug := RandSeq(n)

	Where(post, slug).One(&post)

	if post != nil {
		return genUniqSlug(n)
	}

	return slug
}
