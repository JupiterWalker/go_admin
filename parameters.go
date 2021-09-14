package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"time"
)

func GetParametersTable(ctx *context.Context) table.Table {

	parameters := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := parameters.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Param_key", "param_key", db.Varchar)
	info.AddField("Param_value", "param_value", db.Varchar)
	info.AddField("Param_type", "param_type", db.Varchar)
	info.AddField("Env", "env", db.Varchar).FieldFilterable()
	info.AddField("Domain", "domain", db.Varchar).FieldFilterable()
	info.AddField("Update_user", "update_user", db.Varchar)
	info.AddField("Last_update", "last_update", db.Datetime)
	info.AddField("Create_at", "create_at", db.Datetime)

	info.SetTable("parameters").SetTitle("Parameters").SetDescription("Parameters")

	formList := parameters.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Param_key", "param_key", db.Varchar, form.Text)
	formList.AddField("Param_value", "param_value", db.Varchar, form.Text)
	formList.AddField("Param_type", "param_type", db.Varchar, form.Text)

	env := formList.AddField("Env", "env", db.Varchar, form.SelectSingle)
	// 单选的选项，text代表显示内容，value代表对应值
	env.FieldOptions(types.FieldOptions{
		{Text: "dev", Value: "dev"},
		{Text: "qa1", Value: "qa1"},
		{Text: "qa2", Value: "qa2"},
		{Text: "prod", Value: "prod"},
	}).
		// 设置默认值
		FieldDefault("dev")

	domain := formList.AddField("Domain", "domain", db.Varchar, form.SelectSingle)
	domain.FieldOptions(types.FieldOptions{
		{Text: "lh-svc", Value: "lh-svc"},
		{Text: "lh-flow", Value: "lh-flow"},
	}).
		// 设置默认值
		FieldDefault("lh-svc")
	formList.AddField("Update_user", "update_user", db.Varchar, form.Text)
	lastUpdate := formList.AddField("Last_update", "last_update", db.Datetime, form.Datetime)
	//lastUpdate.FieldDisableWhenCreate()
	lastUpdate.FieldHideWhenCreate()
	lastUpdate.FieldHideWhenUpdate()
	lastUpdate.FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
		return time.Now()
	})
	createAt := formList.AddField("Create_at", "create_at", db.Datetime, form.Datetime)
	createAt.FieldDisableWhenUpdate()
	createAt.FieldHideWhenCreate()
	createAt.FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
		return time.Now()
	})
	formList.SetTable("parameters").SetTitle("Parameters").SetDescription("Parameters")

	return parameters
}
