package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const mysqlDsn = ""

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open(mysqlDsn))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(g.GenerateModelAs("user", "User"))

	// Generate the code
	g.Execute()
}
