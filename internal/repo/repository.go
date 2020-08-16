package repo

import (
	"github.com/go-sql-driver/mysql"
	"log"

	"github.com/jinzhu/gorm"
	proConf "github.com/widrik/pr/internal/config"
	"github.com/widrik/pr/internal/entities"
)

type Repository struct {
	DB     *gorm.DB
}

func GetRepo(proConf *proConf.Configuration) *Repository {
	repo := &Repository{}
	repo.initDB(proConf)

	return repo
}

func (repo *Repository) initDB(proConf *proConf.Configuration) {
	config := mysql.NewConfig()
	config.Addr = proConf.Database.Host
	config.DBName = proConf.Database.Name
	config.User = proConf.Database.User
	config.Passwd = proConf.Database.Password
	config.Net = "tcp"
	config.Collation = "utf8mb4_unicode_ci"
	config.Params = getParams()
	config.ParseTime = true

	db, err := gorm.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal("connection failed")
	}
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")

	repo.DB = db
}

func Migrate(repo *Repository) {
	if err := repo.DB.AutoMigrate(&entities.Banner{}, &entities.Slot{}, &entities.SocialGroup{}, &entities.Stats{}).Error; err != nil {
		log.Fatal("migration failed")
	}

	slots := []*entities.Slot{
		{
			Description: "Правый",
		},
		{
			Description: "Левый",
		},
		{
			Description: "Нижний большой",
		},
	}

	socialGroups := []*entities.SocialGroup{
		{
			Description: "Дети",
		},
		{
			Description: "Женщины",
		},
		{
			Description: "Мужчины",
		},
	}

	for _, slot := range slots {
		if err := repo.DB.Save(slot).Error; err != nil {
			log.Fatal(err)
		}
	}

	for _, socialGroup := range socialGroups {
		if err := repo.DB.Save(socialGroup).Error; err != nil {
			log.Fatal(err)
		}
	}
}

func getParams()  map[string]string {
	params := make(map[string]string)
	params["charset"] = "utf8mb4"

	return params
}