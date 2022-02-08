package main

import "cinema-search/internal/runner"

const configDir = "./config/"

func main() {
	//db := runner.Start(configDir)
	//fmt.Println(db)
	//cr := domain.CreateMovie{Id: 1}
	//movie := domain.InitMovie(cr)
	//fmt.Println(*movie)
	runner.Start(configDir)
	//cfg, err := config.New("./config/")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(cfg.Postgres.Name + " " + cfg.Postgres.Username + " " +  cfg.Postgres.Password)
	//dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	//	DB_USER, DB_PASSWORD, DB_NAME)
	//fmt.Println(dbInfo)
	//fmt.Println("Application starts here")
}
