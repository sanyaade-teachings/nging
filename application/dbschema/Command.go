// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_Command []*Command

func (s Slice_Command) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_Command) RangeRaw(fn func(m *Command) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_Command) GroupBy(keyField string) map[string][]*Command {
	r := map[string][]*Command{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*Command{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_Command) KeyBy(keyField string) map[string]*Command {
	r := map[string]*Command{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_Command) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_Command) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

// Command 指令集
type Command struct {
	base    factory.Base
	objects []*Command

	Id            uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name          string `db:"name" bson:"name" comment:"名称" json:"name" xml:"name"`
	Description   string `db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Command       string `db:"command" bson:"command" comment:"命令" json:"command" xml:"command"`
	WorkDirectory string `db:"work_directory" bson:"work_directory" comment:"工作目录" json:"work_directory" xml:"work_directory"`
	Env           string `db:"env" bson:"env" comment:"环境变量(一行一个，格式为：var1=val1)" json:"env" xml:"env"`
	Created       uint   `db:"created" bson:"created" comment:"添加时间" json:"created" xml:"created"`
	Updated       uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	Disabled      string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	Remote        string `db:"remote" bson:"remote" comment:"是否执行远程SSH命令" json:"remote" xml:"remote"`
	SshAccountId  uint   `db:"ssh_account_id" bson:"ssh_account_id" comment:"SSH账号ID" json:"ssh_account_id" xml:"ssh_account_id"`
}

// - base function

func (a *Command) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *Command) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *Command) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *Command) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *Command) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *Command) Context() echo.Context {
	return a.base.Context()
}

func (a *Command) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *Command) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *Command) Namer() func(string) string {
	return a.base.Namer()
}

func (a *Command) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *Command) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *Command) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *Command) Objects() []*Command {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *Command) XObjects() Slice_Command {
	return Slice_Command(a.Objects())
}

func (a *Command) NewObjects() factory.Ranger {
	return &Slice_Command{}
}

func (a *Command) InitObjects() *[]*Command {
	a.objects = []*Command{}
	return &a.objects
}

func (a *Command) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *Command) Short_() string {
	return "command"
}

func (a *Command) Struct_() string {
	return "Command"
}

func (a *Command) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *Command) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *Command) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	base := a.base
	err := a.Param(mw, args...).SetRecv(a).One()
	a.base = base
	return err
}

func (a *Command) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
}

func (a *Command) GroupBy(keyField string, inputRows ...[]*Command) map[string][]*Command {
	var rows Slice_Command
	if len(inputRows) > 0 {
		rows = Slice_Command(inputRows[0])
	} else {
		rows = Slice_Command(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *Command) KeyBy(keyField string, inputRows ...[]*Command) map[string]*Command {
	var rows Slice_Command
	if len(inputRows) > 0 {
		rows = Slice_Command(inputRows[0])
	} else {
		rows = Slice_Command(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *Command) AsKV(keyField string, valueField string, inputRows ...[]*Command) param.Store {
	var rows Slice_Command
	if len(inputRows) > 0 {
		rows = Slice_Command(inputRows[0])
	} else {
		rows = Slice_Command(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *Command) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
}

func (a *Command) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Remote) == 0 {
		a.Remote = "N"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *Command) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Remote) == 0 {
		a.Remote = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *Command) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *Command) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	if val, ok := kvset["remote"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["remote"] = "N"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Update()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *Command) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.Remote) == 0 {
			a.Remote = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.Remote) == 0 {
			a.Remote = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *Command) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Delete()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *Command) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *Command) Reset() *Command {
	a.Id = 0
	a.Name = ``
	a.Description = ``
	a.Command = ``
	a.WorkDirectory = ``
	a.Env = ``
	a.Created = 0
	a.Updated = 0
	a.Disabled = ``
	a.Remote = ``
	a.SshAccountId = 0
	return a
}

func (a *Command) AsMap() param.Store {
	r := param.Store{}
	r["Id"] = a.Id
	r["Name"] = a.Name
	r["Description"] = a.Description
	r["Command"] = a.Command
	r["WorkDirectory"] = a.WorkDirectory
	r["Env"] = a.Env
	r["Created"] = a.Created
	r["Updated"] = a.Updated
	r["Disabled"] = a.Disabled
	r["Remote"] = a.Remote
	r["SshAccountId"] = a.SshAccountId
	return r
}

func (a *Command) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "description":
			a.Description = param.AsString(value)
		case "command":
			a.Command = param.AsString(value)
		case "work_directory":
			a.WorkDirectory = param.AsString(value)
		case "env":
			a.Env = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "remote":
			a.Remote = param.AsString(value)
		case "ssh_account_id":
			a.SshAccountId = param.AsUint(value)
		}
	}
}

func (a *Command) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Id":
			a.Id = param.AsUint(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "Description":
			a.Description = param.AsString(vv)
		case "Command":
			a.Command = param.AsString(vv)
		case "WorkDirectory":
			a.WorkDirectory = param.AsString(vv)
		case "Env":
			a.Env = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "Remote":
			a.Remote = param.AsString(vv)
		case "SshAccountId":
			a.SshAccountId = param.AsUint(vv)
		}
	}
}

func (a *Command) AsRow() param.Store {
	r := param.Store{}
	r["id"] = a.Id
	r["name"] = a.Name
	r["description"] = a.Description
	r["command"] = a.Command
	r["work_directory"] = a.WorkDirectory
	r["env"] = a.Env
	r["created"] = a.Created
	r["updated"] = a.Updated
	r["disabled"] = a.Disabled
	r["remote"] = a.Remote
	r["ssh_account_id"] = a.SshAccountId
	return r
}

func (a *Command) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *Command) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
