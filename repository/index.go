package repository

import (
	"github.com/myrachanto/power/httperors" 
	"github.com/myrachanto/power/model"
	"gorm.io/driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)
var (
	IndexRepo indexRepo = indexRepo{}

) 

//Layout ...
const (
	Layout   = "2006-01-02"
	layoutUS = "January 2, 2006"
)
type Db struct {
	DbType     string `mapstructure:"DbType"`
	DbName     string `mapstructure:"DbName"`
	DbUsername string `mapstructure:"DbUsername"`
	DbPassword string `mapstructure:"DbPassword"`
	DbPort     string `mapstructure:"DbPort"`
}

func LoaddbConfig() (db Db, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&db)
	return
}

///curtesy to gorm
type indexRepo struct{}
func init() {
	sdb, ers := LoaddbConfig()
	if ers != nil {
		panic("failed to logg")
	}
	GormDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       sdb.DbUsername + ":" + sdb.DbPassword + "@tcp(127.0.0.1:3306)/" + sdb.DbName + "?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                                                                     // default size for string fields
		DisableDatetimePrecision:  true,                                                                                                                    // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                                                                    // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                                                    // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                                                   // auto configure based on currently MySQL version
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return 
	}
	GormDB.AutoMigrate(&model.Customer{})
	GormDB.AutoMigrate(&model.User{})
	GormDB.AutoMigrate(&model.Auth{})
	GormDB.AutoMigrate(&model.Workspace{})
	GormDB.AutoMigrate(&model.Project{})
	GormDB.AutoMigrate(&model.Info{})
	return
}
func (indexRepo indexRepo) Getconnected() (*gorm.DB, *httperors.HttpError) {
	sdb, ers := LoaddbConfig()
	if ers != nil {
		return nil, httperors.NewNotFoundError("Something went wrong with viper --db!")
	}
	GormDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       sdb.DbUsername + ":" + sdb.DbPassword + "@tcp(127.0.0.1:3306)/" + sdb.DbName + "?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                                                                     // default size for string fields
		DisableDatetimePrecision:  true,                                                                                                                    // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                                                                    // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                                                    // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                                                   // auto configure based on currently MySQL version
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, httperors.NewNotFoundError("Something went wrong with viper --db!")
	}
	return GormDB, nil
}
func (indexRepo indexRepo) DbClose(GormDB *gorm.DB) {
	// defer GormDB.Close()
	sqlDB, err := GormDB.DB()
	if err != nil {
		return
	}
	sqlDB.Close()
}