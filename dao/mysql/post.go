package mysql

import (
	"bluebell1/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post (post_id, title, content, author_id, community_id) VALUES (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	if err != nil {
		return
	}
	return
}

func GetPostByID(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id,author_id,community_id,title,content,creat_time from post where post_id=?`
	err = db.Get(post, sqlStr, pid)
	return
}
func GetUserByID(authorid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id,username from user where user_id=?`
	if err = db.Get(user, sqlStr, authorid); err != nil {
		return
	}
	return
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id, title,content,author_id,community_id,creat_time from post order by creat_time desc limit ?,?`
	posts = make([]*models.Post, 0)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

// 根据给定的id列表查询帖子数据
func GetPostListByIDs(ids []string) (posts []*models.Post, err error) {
	sqlStr := `select post_id, title,content,author_id,community_id,creat_time from post where post_id in (?) 
				order by  find_in_set(post_id,?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&posts, query, args...)
	return
}
