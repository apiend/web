package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {
	http.HandleFunc("/", catchAll)
	http.HandleFunc("/posts/", CORS(posts))
	http.ListenAndServe(":4567", nil)
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	handleError(w, 404)
}

func posts(w http.ResponseWriter, r *http.Request) {
	segs := GetSegs(r)
	size := len(segs)

	switch {
	case size == 2:
		postsWithID(w, r)
		return
	case size > 2:
		handleError(w, 404)
		return
	}

	switch r.Method {
	case "GET":
		postsGet(w, r)
		return
	case "POST":
		postsPost(w, r)
		return
	}

	handleError(w, 404)
}

func postsWithID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		postsShow(w, r)
		return
	case "DELETE":
		postsDelete(w, r)
		return
	case "PUT":
		postsUpdate(w, r)
		return
	}

	handleError(w, 404)
}

func postsGet(w http.ResponseWriter, r *http.Request) {
	posts := []*Post{}

	All(&Post{}).All(&posts)

	JSON(w, posts, 200)
}

func postsPost(w http.ResponseWriter, r *http.Request) {
	post := &Post{}

	json.NewDecoder(r.Body).Decode(&post)

	post.BeforeCreate()

	if err := Insert(post); err != nil {
		handleError(w, 400)
		return
	}

	JSON(w, post, 201)
}

func postsUpdate(w http.ResponseWriter, r *http.Request) {
	post := &Post{Slug: GetSeg(r, 2)}

	json.NewDecoder(r.Body).Decode(&post)

	post.BeforeUpdate()

	_, err := Update(post)

	switch {
	case err == mgo.ErrNotFound:
		handleError(w, 404)
		return
	case err != nil:
		handleError(w, 400)
		return
	}

	JSON(w, post, 200)
}

func postsShow(w http.ResponseWriter, r *http.Request) {
	post := &Post{Slug: GetSeg(r, 2)}

	err := Find(post).One(&post)

	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, post, 200)
}

func postsDelete(w http.ResponseWriter, r *http.Request) {
	err := Delete(&Post{
		Slug: GetSeg(r, 2),
	})

	if err == mgo.ErrNotFound {
		handleError(w, 404)
	}
}

func handleError(w http.ResponseWriter, code int) {
	JSON(w, Error{codes[code], code}, code)
}
