package godb
import (
	"database/sql"
)

//构造MySQL请求
type MySQL struct {
	Host string
	Username string
	Password string
	Dbname string
	
	db *sql.DB
}

//连接数据库
func (mysql *MySQL)connect()error{
	db, err := sql.Open("mysql",mysql.Username+":"+mysql.Password+"@"+mysql.Host+"/"+mysql.Dbname)
	if err != nil {
		return err
	} else {
		mysql.db=db
		if err:=mysql.db.Ping();err!=nil{
			mysql.db.Close()
			return err
		}
		return nil
	}
}

//断开连接
func (mysql *MySQL)Disconnect(){
	if mysql.db!=nil{
		mysql.db.Close()
	}
}

//请求所有数据
func (mysql *MySQL)QueryAll(query string, args ...interface{})(sql.Rows,error){
	if mysql.db==nil {
		if err:=mysql.connect();err!=nil{
			return nil,err
		}
	}
	stmt,err :=mysql.db.Prepare(query)
	if err!=nil{
		return nil,err
	}
	
	rows, err :=stmt.Query(args)
	if err!=nil{
		return nil,err
	}
	
	return rows,nil
}

//请求一行数据
func (mysql *MySQL)QueryRow(query string, args ...interface{})(sql.Row,error){
	if mysql.db==nil {
		if err := mysql.connect(); err!=nil {
			return nil,err
		}
	}
	stmt,err := mysql.db.Prepare(query)
	if err!=nil {
		return nil, err
	}
	
	row,err := stmt.QueryRow(args)
	if err!=nil {
		return nil, err
	}
	
	return row,nil
}

//执行SQL语句
func (mysql *MySQL)Execute(query string, args ...interface{})(sql.Result,error){
	if mysql.db==nil {
		if err := mysql.connect(); err!=nil {
			return 0,err
		}
	}
	stmt,err:=mysql.db.Prepare()
	if err!=nil {
		return nil, err
	}

	res, err := stmt.Exec(args)
	if err != nil {
		return 0, err
	}

	return res, nil
}
