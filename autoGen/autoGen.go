package main

import (
	"gorm.io/gen"

	"github.com/cdlavacudeg/go-budget-guardian/database"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext,
	})

	gormdb := database.InitDB()

	g.UseDB(gormdb)

	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
