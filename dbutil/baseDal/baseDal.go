package baseDAL

import (
	"database/sql"
	"log"

	"github.com/zhouweigithub/goutil/logutil"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type DbQuery struct {
	db               *sql.DB
	sqlType          string
	connectionString string
}

// 使用默认数据库连接字符串，生成新的实例
func (p *DbQuery) New(dbType string, connString string) *DbQuery {
	return &DbQuery{
		sqlType:          dbType,
		connectionString: connString,
	}
}

//var sqlType = common.GetConfigs().DbType
//var connectionString = common.GetConfigs().ConnectionString
//var db *mysql.DB

// 打开数据库连接
func (p *DbQuery) OpenDB() {
	if p.db == nil || p.db.Stats().OpenConnections == 0 {
		var err error
		p.db, err = sql.Open(p.sqlType, p.connectionString)
		if err != nil {
			log.Println("open database error: ", err.Error())
			logutil.Error(err.Error())
		} else {
			log.Println("open database success")
		}
	}
}

// 关闭数据库连接
func (p *DbQuery) CloseDB() {
	if p.db != nil {
		p.db.Close()
		log.Println("close database success")
	}
}

// 需要手动关闭数据库连接
func (p *DbQuery) ExeNonQuery(sql string) int64 {
	res, err := p.db.Exec(sql)
	if err != nil {
		errMst := "数据库操作异常！sql: " + sql + "\r\n" + err.Error()
		log.Println(errMst)
		logutil.Error(errMst)
		return 0
	}
	rowsAffected, _ := res.RowsAffected()
	return rowsAffected
}

// 需要手动关闭数据库连接
func (p *DbQuery) Query(sql string) *sql.Rows {
	rows, err := p.db.Query(sql)
	if err != nil {
		errMst := "数据库操作异常！sql: " + sql + "\r\n" + err.Error()
		log.Println(errMst)
		logutil.Error(errMst)
		return nil
	}
	return rows
}

// 需要手动关闭数据库连接
func (p *DbQuery) QueryRow(sql string) *sql.Row {
	row := p.db.QueryRow(sql)
	return row
}
