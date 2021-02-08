package models

import (
	blackfriday "github.com/russross/blackfriday/v2"
	"html/template"
	"time"
)

// Post db table
type Post struct {
	ID          int
	Title       string `form:"title" json:"title" xml:"title"  binding:"required"`
	Author      int
	AuthorName  string
	Content     string `form:"content" json:"content" xml:"content"  binding:"required"`
	HTMLContent template.HTML
	CreatedAt   string
	Status      int `form:"status" json:"status" xml:"status"`
}

// GetPosts post list
func GetPosts(num int, page int, isAll bool) ([]Post, error) {
	posts := []Post{}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * num
	sql := "select id, post_title, post_created_at from tiny_posts where post_status=1 order by id desc limit ?, ?"
	if isAll {
		sql = "select id, post_title, post_created_at from tiny_posts order by id desc limit ?, ?"
	}
	rows, err := Db.Query(sql, offset, num)
	if err != nil {
		return posts, err
	}
	defer rows.Close()
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.ID, &post.Title, &post.CreatedAt)
		if err != nil {
			return posts, err
		}
		t, _ := time.Parse("2006-01-02 15:04:05", post.CreatedAt)
		post.CreatedAt = t.Format("2006-01-02")
		posts = append(posts, post)
	}
	return posts, err
}

// GetPost return post info
func GetPost(id int) Post {
	post := Post{}
	Db.QueryRow("select id, post_title, post_content, post_created_at, post_status from tiny_posts where id=?", id).
		Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.Status)

	post.HTMLContent = template.HTML(blackfriday.Run([]byte(post.Content)))
	t, _ := time.Parse("2006-01-02 15:04:05", post.CreatedAt)
	post.CreatedAt = t.Format("2006-01-02")

	return post
}

// GetAboutPost get about post
func GetAboutPost() Post {
	post := Post{}
	Db.QueryRow("select id, post_title, post_content, post_created_at from tiny_posts where post_title=? and post_status=1", "about").
		Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt)

	post.HTMLContent = template.HTML(blackfriday.Run([]byte(post.Content)))
	t, _ := time.Parse("2006-01-02 15:04:05", post.CreatedAt)
	post.CreatedAt = t.Format("2006-01-02")

	return post
}

// GetPostTotal get total
func GetPostTotal(isAll bool) int {
	total := 0
	sql := "select count(*) as total from tiny_posts where post_status=1"
	if isAll {
		sql = "select count(*) as total from tiny_posts"
	}
	Db.QueryRow(sql).Scan(&total)
	return total
}

func UpdatePost(post Post) error {
	stmt, err := Db.Prepare("update tiny_posts set post_title = ?, post_content = ?, post_author = ?, post_status=?, post_updated_at=? where id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	stmt.Exec(post.Title, post.Content, post.Author, post.Status, time.Now().Format("2006-01-02 15:04:05"), post.ID)
	return nil
}

func AddPost(post Post) int {
	sql := "insert into tiny_posts (post_title, post_content, post_author, post_status, post_created_at) values (?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(post.Title, post.Content, post.Author, post.Status, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		panic(err)
	}

	pid, _ := res.LastInsertId()
	post.ID = int(pid)
	return post.ID
}

func DeletePost(id int) {
	stmt, err := Db.Prepare("delete from tiny_posts where id = ? limit 1")
	if err != nil {
		return
	}
	defer stmt.Close()
	stmt.Exec(id)
}
