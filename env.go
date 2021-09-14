package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetEnvTable(ctx *context.Context) table.Table {

	env := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := env.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)

	info.SetTable("env").SetTitle("Env").SetDescription("Env")

	formList := env.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)

	formList.SetTable("env").SetTitle("Env").SetDescription("Env")

	return env
}
