package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sync"
	"time"
	"marketApi/pkg/setting"
	"marketApi/pkg/snowflakelib"
)

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open(setting.DB.Type,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&parseTime=True&loc=Local",
			setting.DB.User,
			setting.DB.Password,
			setting.DB.Host,
			setting.DB.Port,
			setting.DB.Database,
			setting.DB.Charset,
			setting.DB.Collation))
	if err != nil {
		log.Panicf("数据库连接失败:%s", err)
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	log.Println("数据库链接成功")
}

func CloseDB() {
	defer db.Close()
}

type Model struct {
	ID        uint64     `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (m *Model) BeforeSave(scope *gorm.Scope) error {
	err := scope.SetColumn("ID", snowflakelib.Uint64ID())
	return err
}

func (m *Model) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("ID", snowflakelib.Uint64ID())
	return err
}

type queryWhere = map[string]interface{}

type Query struct {
	Table string
	Where queryWhere
}

func (q *Query) Map() map[string]interface{} {
	return q.Where
}

type Transaction struct {
	db       *gorm.DB
	Tx       *gorm.DB
	rollBack bool
	commit   bool
	lock     sync.Mutex
}

func NewTransaction() *Transaction {
	return &Transaction{
		db: db,
		Tx: db.Begin(),
	}
}

func (t *Transaction) Rollback() {
	t.lock.Lock()
	t.lock.Unlock()

	if t.rollBack == false && t.commit == false {
		t.Tx.Rollback()
		t.rollBack = true
	}
}

func (t *Transaction) Commit() {
	t.lock.Lock()
	t.lock.Unlock()

	if t.commit == false && t.rollBack == false {
		t.Tx.Commit()
		t.commit = true
	}
}
