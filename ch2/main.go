package main

func main() {
	_, err = query(db, "select * from abc where 1 = 1");
	if err != nil {
		fmt.Printf("original error: %T %v", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
		os.Exits(1)
	}
}

func query(db *sql.DB, sql string) ([]interface{}, error) {
    DB, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	args, err := DB.QueryRow("select * from user where id=1").Scan(user.id)
	if err != nil {
		e = errorcs.Wrap(err, "sql ErrNoRows")
	    return nil, e
	}
	return args, nil
}