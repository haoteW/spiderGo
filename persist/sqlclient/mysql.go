package sqlclient

import (
	"database/sql"
	"fmt"
	"good-spider/model"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/movieInfo?charset=utf8mb4")
	_, err := db.Exec(createTable)
	if err != nil {
		log.Fatal("sqlclient connect fail", err)
	}
}

// 插入数据
func Insert(movieInfo model.MovieInfo) error {
	insertMessage := fmt.Sprintf(insertData1,
		"2338160", "https://movie.douban.com/subject/2338160/", " ", "王树忱", " ", "动画/短片", "中国大陆", "无对白", "1988", "19分钟", "Feelings of Mountains and Waters", "无翼鸟")
	//movieInfo.ID, movieInfo.Director, movieInfo.Screenwriter, movieInfo.Starring, movieInfo.Type, movieInfo.Region, movieInfo.Language, movieInfo.ReleaseDate, movieInfo.Long, movieInfo.Sname, movieInfo.Title, movieInfo.URL)
	_, err := db.Exec(insertMessage)
	return err
}

// 修改数据
func Update(movieInfo model.MovieInfo) {
	db.Exec(updateData, movieInfo.Director, movieInfo.Screenwriter, movieInfo.Starring, movieInfo.Type, movieInfo.Region, movieInfo.Language, movieInfo.ReleaseDate, movieInfo.Long, movieInfo.Sname, movieInfo.Title, movieInfo.URL, movieInfo.ID)
}

// 查询数据
// fields 传入需要查询的字段 传空默认全字段查询
// conditions 传入查询条件
//
//		"where" 对应map[string]string field : val
//		"group" group by对应 string field
//	 	"order" order by对应 map[string]string
//	 		1、field请用,隔开 field:param1,param2
//	 		2、orderType可选 desc asc orderType:asc 为空默认asc升序
//		"limit" 对应string 0,20 page,pageSize
func Query(fields []string, conditions map[string]interface{}) {
	retField := strings.Join(fields, ", ")
	if retField == "" {
		retField = "*"
	}
	query := "SELECT " + retField + " FROM movieInfo"
	if whereCondition, ok := conditions["where"]; ok {
		query += buildWhereClause(whereCondition)
		query += buildClause(conditions)
	}
	_, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
}

// 支持事务回滚机制的批量数据插入
func MultiInsert(items ...model.MovieInfo) error {
	// 批量数据插入
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("tx begin err")
	}
	stmt, err := tx.Prepare(insertData)
	defer stmt.Close()
	for _, item := range items {
		_, err := stmt.Exec(item.ID, item.Director, item.Screenwriter, item.Starring, item.Type, item.Region, item.Language, item.ReleaseDate, item.Long, item.Sname, item.Title, item.URL)
		if err != nil {
			log.Printf("INSERT failed：%v", err)
			tx.Rollback()
			return fmt.Errorf("INSERT failed：%v", err)
		}
	}
	tx.Commit()
	return nil
}
