package persistence_test

import (
	"database/sql"
	"testing"
	"time"

	"blog/internal/domain/entity"
	"blog/internal/infrastructure/persistence"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMySQLPostRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建mock数据库连接失败: %v", err)
	}
	defer db.Close()

	conn := &persistence.MySQLConnection{DB: db}
	repo := persistence.NewMySQLPostRepository(conn)

	now := time.Now()
	post := &entity.Post{
		Title:     "测试标题",
		Content:   "测试内容",
		Author:    "测试作者",
		CreatedAt: now,
		UpdatedAt: now,
	}

	mock.ExpectExec("INSERT INTO posts").
		WithArgs(post.Title, post.Content, post.Author, post.CreatedAt, post.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(post)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), post.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLPostRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建mock数据库连接失败: %v", err)
	}
	defer db.Close()

	conn := &persistence.MySQLConnection{DB: db}
	repo := persistence.NewMySQLPostRepository(conn)

	now := time.Now()
	expectedPost := &entity.Post{
		ID:        1,
		Title:     "测试标题",
		Content:   "测试内容",
		Author:    "测试作者",
		CreatedAt: now,
		UpdatedAt: now,
	}

	rows := sqlmock.NewRows([]string{"id", "title", "content", "author", "created_at", "updated_at"}).
		AddRow(expectedPost.ID, expectedPost.Title, expectedPost.Content, expectedPost.Author, expectedPost.CreatedAt, expectedPost.UpdatedAt)

	mock.ExpectQuery("SELECT (.+) FROM posts WHERE id = ?").
		WithArgs(1).
		WillReturnRows(rows)

	post, err := repo.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedPost.ID, post.ID)
	assert.Equal(t, expectedPost.Title, post.Title)
	assert.Equal(t, expectedPost.Content, post.Content)
	assert.Equal(t, expectedPost.Author, post.Author)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLPostRepository_GetByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建mock数据库连接失败: %v", err)
	}
	defer db.Close()

	conn := &persistence.MySQLConnection{DB: db}
	repo := persistence.NewMySQLPostRepository(conn)

	mock.ExpectQuery("SELECT (.+) FROM posts WHERE id = ?").
		WithArgs(999).
		WillReturnError(sql.ErrNoRows)

	post, err := repo.GetByID(999)
	assert.Error(t, err)
	assert.Nil(t, post)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLPostRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建mock数据库连接失败: %v", err)
	}
	defer db.Close()

	conn := &persistence.MySQLConnection{DB: db}
	repo := persistence.NewMySQLPostRepository(conn)

	post := &entity.Post{
		ID:      1,
		Title:   "更新标题",
		Content: "更新内容",
	}

	mock.ExpectExec("UPDATE posts SET title = \\?, content = \\?, updated_at = \\? WHERE id = \\?").
		WithArgs(post.Title, post.Content, sqlmock.AnyArg(), post.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Update(post)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLPostRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建mock数据库连接失败: %v", err)
	}
	defer db.Close()

	conn := &persistence.MySQLConnection{DB: db}
	repo := persistence.NewMySQLPostRepository(conn)

	mock.ExpectExec("DELETE FROM posts WHERE id = \\?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Delete(1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
