package task4

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Route struct {
	ID         int
	StartPoint string
	EndPoint   string
}

type Preference struct {
	ID         int
	UserID     int
	mode       string
	outputmode string
	tactics    string
}

type Alias struct {
	ID       int
	UserID   int
	Location string
	Alias    string
}

func dbTest(db *sql.DB) {
	// 插入数据
	InsertRoute(db, Route{ID: 1, StartPoint: "A", EndPoint: "B"})
	InsertRoute(db, Route{ID: 2, StartPoint: "B", EndPoint: "C"})

	InsertPreference(db, Preference{ID: 1, UserID: 1, mode: "Option A", outputmode: "Option B", tactics: "Option C"})

	InsertAlias(db, Alias{ID: 1, UserID: 1, Location: "Work", Alias: "Office"})
	InsertAlias(db, Alias{ID: 2, UserID: 1, Location: "Home", Alias: "My House"})

	// 查询数据
	fmt.Println("查询路线数据：")
	routes := QueryRoutes(db)
	for _, route := range routes {
		fmt.Printf("ID: %d, StartPoint: %s, EndPoint: %s\n", route.ID, route.StartPoint, route.EndPoint)
	}

	fmt.Println("\n查询用户偏好数据：")
	preferences := QueryPreferences(db)
	for _, preference := range preferences {
		fmt.Printf("ID: %d, UserID: %d, mode: %s, outputmode: %s, tactics: %s\n", preference.ID, preference.UserID, preference.mode, preference.outputmode, preference.tactics)
	}

	fmt.Println("\n查询地点别名数据：")
	aliases := QueryAliases(db)
	for _, alias := range aliases {
		fmt.Printf("ID: %d, UserID: %d, Location: %s, Alias: %s\n", alias.ID, alias.UserID, alias.Location, alias.Alias)
	}
}

func DbManager() {
	// 连接到 MySQL 数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/task4")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 创建表
	createTables(db)
	CreateUser(db, "default")
	// 测试
	dbTest(db)
}

func createTables(db *sql.DB) {
	createRoutesTableSQL := `CREATE TABLE IF NOT EXISTS routes (
		id INT AUTO_INCREMENT PRIMARY KEY,
		start_point VARCHAR(255) NOT NULL,
		end_point VARCHAR(255) NOT NULL
	);`
	creatwUsersTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL
	);`
	createPreferencesTableSQL := `CREATE TABLE IF NOT EXISTS preferences (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		mode VARCHAR(255),
		outputmode VARCHAR(255),
		tactics VARCHAR(255),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	createAliasesTableSQL := `CREATE TABLE IF NOT EXISTS aliases (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		location VARCHAR(255) NOT NULL,
		alias VARCHAR(255) NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	_, err := db.Exec(createRoutesTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(creatwUsersTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(createPreferencesTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(createAliasesTableSQL)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateUser(db *sql.DB, name string) (int64, error) {
	// 准备插入语句
	insertUserSQL := "INSERT INTO users (name) VALUES (?)"

	// 执行插入操作
	result, err := db.Exec(insertUserSQL, name)
	if err != nil {
		return 0, err
	}

	// 获取插入后的用户ID
	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}
