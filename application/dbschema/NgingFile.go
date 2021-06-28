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

type Slice_NgingFile []*NgingFile

func (s Slice_NgingFile) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFile) RangeRaw(fn func(m *NgingFile) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFile) GroupBy(keyField string) map[string][]*NgingFile {
	r := map[string][]*NgingFile{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingFile{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingFile) KeyBy(keyField string) map[string]*NgingFile {
	r := map[string]*NgingFile{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingFile) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingFile) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingFile) FromList(data interface{}) Slice_NgingFile {
	values, ok := data.([]*NgingFile)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingFile{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingFile 文件表
type NgingFile struct {
	base    factory.Base
	objects []*NgingFile

	Id         uint64 `db:"id,omitempty,pk" bson:"id,omitempty" comment:"文件ID" json:"id" xml:"id"`
	OwnerType  string `db:"owner_type" bson:"owner_type" comment:"用户类型" json:"owner_type" xml:"owner_type"`
	OwnerId    uint64 `db:"owner_id" bson:"owner_id" comment:"用户ID" json:"owner_id" xml:"owner_id"`
	Name       string `db:"name" bson:"name" comment:"原始文件名" json:"name" xml:"name"`
	SaveName   string `db:"save_name" bson:"save_name" comment:"保存名称" json:"save_name" xml:"save_name"`
	SavePath   string `db:"save_path" bson:"save_path" comment:"文件保存路径" json:"save_path" xml:"save_path"`
	ViewUrl    string `db:"view_url" bson:"view_url" comment:"查看链接" json:"view_url" xml:"view_url"`
	Ext        string `db:"ext" bson:"ext" comment:"文件后缀" json:"ext" xml:"ext"`
	Mime       string `db:"mime" bson:"mime" comment:"文件mime类型" json:"mime" xml:"mime"`
	Type       string `db:"type" bson:"type" comment:"文件类型" json:"type" xml:"type"`
	Size       uint64 `db:"size" bson:"size" comment:"文件大小" json:"size" xml:"size"`
	Width      uint   `db:"width" bson:"width" comment:"宽度(像素)" json:"width" xml:"width"`
	Height     uint   `db:"height" bson:"height" comment:"高度(像素)" json:"height" xml:"height"`
	Dpi        uint   `db:"dpi" bson:"dpi" comment:"分辨率" json:"dpi" xml:"dpi"`
	Md5        string `db:"md5" bson:"md5" comment:"文件md5" json:"md5" xml:"md5"`
	StorerName string `db:"storer_name" bson:"storer_name" comment:"文件保存位置" json:"storer_name" xml:"storer_name"`
	StorerId   string `db:"storer_id" bson:"storer_id" comment:"位置ID" json:"storer_id" xml:"storer_id"`
	Created    uint   `db:"created" bson:"created" comment:"上传时间" json:"created" xml:"created"`
	Updated    uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	Sort       int64  `db:"sort" bson:"sort" comment:"排序" json:"sort" xml:"sort"`
	Status     int    `db:"status" bson:"status" comment:"状态(1-已审核/0-未审核)" json:"status" xml:"status"`
	CategoryId uint   `db:"category_id" bson:"category_id" comment:"分类ID" json:"category_id" xml:"category_id"`
	Tags       string `db:"tags" bson:"tags" comment:"标签" json:"tags" xml:"tags"`
	Subdir     string `db:"subdir" bson:"subdir" comment:"子目录" json:"subdir" xml:"subdir"`
	UsedTimes  uint   `db:"used_times" bson:"used_times" comment:"被使用的次数" json:"used_times" xml:"used_times"`
}

// - base function

func (a *NgingFile) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingFile) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingFile) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingFile) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingFile) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingFile) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingFile) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingFile) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingFile) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingFile) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingFile) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingFile) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingFile) Objects() []*NgingFile {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingFile) XObjects() Slice_NgingFile {
	return Slice_NgingFile(a.Objects())
}

func (a *NgingFile) NewObjects() factory.Ranger {
	return &Slice_NgingFile{}
}

func (a *NgingFile) InitObjects() *[]*NgingFile {
	a.objects = []*NgingFile{}
	return &a.objects
}

func (a *NgingFile) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingFile) Short_() string {
	return "nging_file"
}

func (a *NgingFile) Struct_() string {
	return "NgingFile"
}

func (a *NgingFile) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingFile) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingFile) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	base := a.base
	if !a.base.Eventable() {
		err = a.Param(mw, args...).SetRecv(a).One()
		a.base = base
		return
	}
	queryParam := a.Param(mw, args...).SetRecv(a)
	if err = DBI.FireReading(a, queryParam); err != nil {
		return
	}
	err = queryParam.One()
	a.base = base
	if err == nil {
		err = DBI.FireReaded(a, queryParam)
	}
	return
}

func (a *NgingFile) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingFile:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFile(*v))
		case []*NgingFile:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFile(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFile) GroupBy(keyField string, inputRows ...[]*NgingFile) map[string][]*NgingFile {
	var rows Slice_NgingFile
	if len(inputRows) > 0 {
		rows = Slice_NgingFile(inputRows[0])
	} else {
		rows = Slice_NgingFile(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingFile) KeyBy(keyField string, inputRows ...[]*NgingFile) map[string]*NgingFile {
	var rows Slice_NgingFile
	if len(inputRows) > 0 {
		rows = Slice_NgingFile(inputRows[0])
	} else {
		rows = Slice_NgingFile(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingFile) AsKV(keyField string, valueField string, inputRows ...[]*NgingFile) param.Store {
	var rows Slice_NgingFile
	if len(inputRows) > 0 {
		rows = Slice_NgingFile(inputRows[0])
	} else {
		rows = Slice_NgingFile(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingFile) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingFile:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFile(*v))
		case []*NgingFile:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFile(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFile) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.OwnerType) == 0 {
		a.OwnerType = "user"
	}
	if len(a.Type) == 0 {
		a.Type = "image"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingFile) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.OwnerType) == 0 {
		a.OwnerType = "user"
	}
	if len(a.Type) == 0 {
		a.Type = "image"
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

func (a *NgingFile) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingFile) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["owner_type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["owner_type"] = "user"
		}
	}
	if val, ok := kvset["type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["type"] = "image"
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

func (a *NgingFile) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.OwnerType) == 0 {
			a.OwnerType = "user"
		}
		if len(a.Type) == 0 {
			a.Type = "image"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.OwnerType) == 0 {
			a.OwnerType = "user"
		}
		if len(a.Type) == 0 {
			a.Type = "image"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
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

func (a *NgingFile) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingFile) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingFile) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingFile) Reset() *NgingFile {
	a.Id = 0
	a.OwnerType = ``
	a.OwnerId = 0
	a.Name = ``
	a.SaveName = ``
	a.SavePath = ``
	a.ViewUrl = ``
	a.Ext = ``
	a.Mime = ``
	a.Type = ``
	a.Size = 0
	a.Width = 0
	a.Height = 0
	a.Dpi = 0
	a.Md5 = ``
	a.StorerName = ``
	a.StorerId = ``
	a.Created = 0
	a.Updated = 0
	a.Sort = 0
	a.Status = 0
	a.CategoryId = 0
	a.Tags = ``
	a.Subdir = ``
	a.UsedTimes = 0
	return a
}

func (a *NgingFile) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["OwnerType"] = a.OwnerType
		r["OwnerId"] = a.OwnerId
		r["Name"] = a.Name
		r["SaveName"] = a.SaveName
		r["SavePath"] = a.SavePath
		r["ViewUrl"] = a.ViewUrl
		r["Ext"] = a.Ext
		r["Mime"] = a.Mime
		r["Type"] = a.Type
		r["Size"] = a.Size
		r["Width"] = a.Width
		r["Height"] = a.Height
		r["Dpi"] = a.Dpi
		r["Md5"] = a.Md5
		r["StorerName"] = a.StorerName
		r["StorerId"] = a.StorerId
		r["Created"] = a.Created
		r["Updated"] = a.Updated
		r["Sort"] = a.Sort
		r["Status"] = a.Status
		r["CategoryId"] = a.CategoryId
		r["Tags"] = a.Tags
		r["Subdir"] = a.Subdir
		r["UsedTimes"] = a.UsedTimes
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "OwnerType":
			r["OwnerType"] = a.OwnerType
		case "OwnerId":
			r["OwnerId"] = a.OwnerId
		case "Name":
			r["Name"] = a.Name
		case "SaveName":
			r["SaveName"] = a.SaveName
		case "SavePath":
			r["SavePath"] = a.SavePath
		case "ViewUrl":
			r["ViewUrl"] = a.ViewUrl
		case "Ext":
			r["Ext"] = a.Ext
		case "Mime":
			r["Mime"] = a.Mime
		case "Type":
			r["Type"] = a.Type
		case "Size":
			r["Size"] = a.Size
		case "Width":
			r["Width"] = a.Width
		case "Height":
			r["Height"] = a.Height
		case "Dpi":
			r["Dpi"] = a.Dpi
		case "Md5":
			r["Md5"] = a.Md5
		case "StorerName":
			r["StorerName"] = a.StorerName
		case "StorerId":
			r["StorerId"] = a.StorerId
		case "Created":
			r["Created"] = a.Created
		case "Updated":
			r["Updated"] = a.Updated
		case "Sort":
			r["Sort"] = a.Sort
		case "Status":
			r["Status"] = a.Status
		case "CategoryId":
			r["CategoryId"] = a.CategoryId
		case "Tags":
			r["Tags"] = a.Tags
		case "Subdir":
			r["Subdir"] = a.Subdir
		case "UsedTimes":
			r["UsedTimes"] = a.UsedTimes
		}
	}
	return r
}

func (a *NgingFile) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint64(value)
		case "owner_type":
			a.OwnerType = param.AsString(value)
		case "owner_id":
			a.OwnerId = param.AsUint64(value)
		case "name":
			a.Name = param.AsString(value)
		case "save_name":
			a.SaveName = param.AsString(value)
		case "save_path":
			a.SavePath = param.AsString(value)
		case "view_url":
			a.ViewUrl = param.AsString(value)
		case "ext":
			a.Ext = param.AsString(value)
		case "mime":
			a.Mime = param.AsString(value)
		case "type":
			a.Type = param.AsString(value)
		case "size":
			a.Size = param.AsUint64(value)
		case "width":
			a.Width = param.AsUint(value)
		case "height":
			a.Height = param.AsUint(value)
		case "dpi":
			a.Dpi = param.AsUint(value)
		case "md5":
			a.Md5 = param.AsString(value)
		case "storer_name":
			a.StorerName = param.AsString(value)
		case "storer_id":
			a.StorerId = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		case "sort":
			a.Sort = param.AsInt64(value)
		case "status":
			a.Status = param.AsInt(value)
		case "category_id":
			a.CategoryId = param.AsUint(value)
		case "tags":
			a.Tags = param.AsString(value)
		case "subdir":
			a.Subdir = param.AsString(value)
		case "used_times":
			a.UsedTimes = param.AsUint(value)
		}
	}
}

func (a *NgingFile) Set(key interface{}, value ...interface{}) {
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
			a.Id = param.AsUint64(vv)
		case "OwnerType":
			a.OwnerType = param.AsString(vv)
		case "OwnerId":
			a.OwnerId = param.AsUint64(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "SaveName":
			a.SaveName = param.AsString(vv)
		case "SavePath":
			a.SavePath = param.AsString(vv)
		case "ViewUrl":
			a.ViewUrl = param.AsString(vv)
		case "Ext":
			a.Ext = param.AsString(vv)
		case "Mime":
			a.Mime = param.AsString(vv)
		case "Type":
			a.Type = param.AsString(vv)
		case "Size":
			a.Size = param.AsUint64(vv)
		case "Width":
			a.Width = param.AsUint(vv)
		case "Height":
			a.Height = param.AsUint(vv)
		case "Dpi":
			a.Dpi = param.AsUint(vv)
		case "Md5":
			a.Md5 = param.AsString(vv)
		case "StorerName":
			a.StorerName = param.AsString(vv)
		case "StorerId":
			a.StorerId = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		case "Sort":
			a.Sort = param.AsInt64(vv)
		case "Status":
			a.Status = param.AsInt(vv)
		case "CategoryId":
			a.CategoryId = param.AsUint(vv)
		case "Tags":
			a.Tags = param.AsString(vv)
		case "Subdir":
			a.Subdir = param.AsString(vv)
		case "UsedTimes":
			a.UsedTimes = param.AsUint(vv)
		}
	}
}

func (a *NgingFile) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["owner_type"] = a.OwnerType
		r["owner_id"] = a.OwnerId
		r["name"] = a.Name
		r["save_name"] = a.SaveName
		r["save_path"] = a.SavePath
		r["view_url"] = a.ViewUrl
		r["ext"] = a.Ext
		r["mime"] = a.Mime
		r["type"] = a.Type
		r["size"] = a.Size
		r["width"] = a.Width
		r["height"] = a.Height
		r["dpi"] = a.Dpi
		r["md5"] = a.Md5
		r["storer_name"] = a.StorerName
		r["storer_id"] = a.StorerId
		r["created"] = a.Created
		r["updated"] = a.Updated
		r["sort"] = a.Sort
		r["status"] = a.Status
		r["category_id"] = a.CategoryId
		r["tags"] = a.Tags
		r["subdir"] = a.Subdir
		r["used_times"] = a.UsedTimes
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "owner_type":
			r["owner_type"] = a.OwnerType
		case "owner_id":
			r["owner_id"] = a.OwnerId
		case "name":
			r["name"] = a.Name
		case "save_name":
			r["save_name"] = a.SaveName
		case "save_path":
			r["save_path"] = a.SavePath
		case "view_url":
			r["view_url"] = a.ViewUrl
		case "ext":
			r["ext"] = a.Ext
		case "mime":
			r["mime"] = a.Mime
		case "type":
			r["type"] = a.Type
		case "size":
			r["size"] = a.Size
		case "width":
			r["width"] = a.Width
		case "height":
			r["height"] = a.Height
		case "dpi":
			r["dpi"] = a.Dpi
		case "md5":
			r["md5"] = a.Md5
		case "storer_name":
			r["storer_name"] = a.StorerName
		case "storer_id":
			r["storer_id"] = a.StorerId
		case "created":
			r["created"] = a.Created
		case "updated":
			r["updated"] = a.Updated
		case "sort":
			r["sort"] = a.Sort
		case "status":
			r["status"] = a.Status
		case "category_id":
			r["category_id"] = a.CategoryId
		case "tags":
			r["tags"] = a.Tags
		case "subdir":
			r["subdir"] = a.Subdir
		case "used_times":
			r["used_times"] = a.UsedTimes
		}
	}
	return r
}

func (a *NgingFile) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingFile) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}
