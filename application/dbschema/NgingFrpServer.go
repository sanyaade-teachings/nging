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

type Slice_NgingFrpServer []*NgingFrpServer

func (s Slice_NgingFrpServer) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFrpServer) RangeRaw(fn func(m *NgingFrpServer) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFrpServer) GroupBy(keyField string) map[string][]*NgingFrpServer {
	r := map[string][]*NgingFrpServer{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingFrpServer{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingFrpServer) KeyBy(keyField string) map[string]*NgingFrpServer {
	r := map[string]*NgingFrpServer{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingFrpServer) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingFrpServer) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingFrpServer) FromList(data interface{}) Slice_NgingFrpServer {
	values, ok := data.([]*NgingFrpServer)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingFrpServer{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingFrpServer FRP服务器设置
type NgingFrpServer struct {
	base    factory.Base
	objects []*NgingFrpServer

	Id                uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name              string `db:"name" bson:"name" comment:"名称" json:"name" xml:"name"`
	Disabled          string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	TcpMux            string `db:"tcp_mux" bson:"tcp_mux" comment:"" json:"tcp_mux" xml:"tcp_mux"`
	Addr              string `db:"addr" bson:"addr" comment:"" json:"addr" xml:"addr"`
	Port              uint   `db:"port" bson:"port" comment:"" json:"port" xml:"port"`
	UdpPort           uint   `db:"udp_port" bson:"udp_port" comment:"" json:"udp_port" xml:"udp_port"`
	KcpPort           uint   `db:"kcp_port" bson:"kcp_port" comment:"" json:"kcp_port" xml:"kcp_port"`
	ProxyAddr         string `db:"proxy_addr" bson:"proxy_addr" comment:"" json:"proxy_addr" xml:"proxy_addr"`
	VhostHttpPort     uint   `db:"vhost_http_port" bson:"vhost_http_port" comment:"" json:"vhost_http_port" xml:"vhost_http_port"`
	VhostHttpTimeout  uint64 `db:"vhost_http_timeout" bson:"vhost_http_timeout" comment:"" json:"vhost_http_timeout" xml:"vhost_http_timeout"`
	VhostHttpsPort    uint   `db:"vhost_https_port" bson:"vhost_https_port" comment:"" json:"vhost_https_port" xml:"vhost_https_port"`
	LogFile           string `db:"log_file" bson:"log_file" comment:"" json:"log_file" xml:"log_file"`
	LogWay            string `db:"log_way" bson:"log_way" comment:"console or file" json:"log_way" xml:"log_way"`
	LogLevel          string `db:"log_level" bson:"log_level" comment:"" json:"log_level" xml:"log_level"`
	LogMaxDays        uint   `db:"log_max_days" bson:"log_max_days" comment:"" json:"log_max_days" xml:"log_max_days"`
	Token             string `db:"token" bson:"token" comment:"" json:"token" xml:"token"`
	AuthTimeout       uint64 `db:"auth_timeout" bson:"auth_timeout" comment:"" json:"auth_timeout" xml:"auth_timeout"`
	SubdomainHost     string `db:"subdomain_host" bson:"subdomain_host" comment:"" json:"subdomain_host" xml:"subdomain_host"`
	MaxPortsPerClient int64  `db:"max_ports_per_client" bson:"max_ports_per_client" comment:"" json:"max_ports_per_client" xml:"max_ports_per_client"`
	MaxPoolCount      uint   `db:"max_pool_count" bson:"max_pool_count" comment:"" json:"max_pool_count" xml:"max_pool_count"`
	HeartBeatTimeout  uint   `db:"heart_beat_timeout" bson:"heart_beat_timeout" comment:"" json:"heart_beat_timeout" xml:"heart_beat_timeout"`
	UserConnTimeout   uint   `db:"user_conn_timeout" bson:"user_conn_timeout" comment:"" json:"user_conn_timeout" xml:"user_conn_timeout"`
	DashboardAddr     string `db:"dashboard_addr" bson:"dashboard_addr" comment:"" json:"dashboard_addr" xml:"dashboard_addr"`
	DashboardPort     uint   `db:"dashboard_port" bson:"dashboard_port" comment:"" json:"dashboard_port" xml:"dashboard_port"`
	DashboardUser     string `db:"dashboard_user" bson:"dashboard_user" comment:"" json:"dashboard_user" xml:"dashboard_user"`
	DashboardPwd      string `db:"dashboard_pwd" bson:"dashboard_pwd" comment:"" json:"dashboard_pwd" xml:"dashboard_pwd"`
	AllowPorts        string `db:"allow_ports" bson:"allow_ports" comment:"" json:"allow_ports" xml:"allow_ports"`
	Extra             string `db:"extra" bson:"extra" comment:"" json:"extra" xml:"extra"`
	Plugins           string `db:"plugins" bson:"plugins" comment:"启用插件(半角逗号分隔)" json:"plugins" xml:"plugins"`
	Uid               uint   `db:"uid" bson:"uid" comment:"" json:"uid" xml:"uid"`
	GroupId           uint   `db:"group_id" bson:"group_id" comment:"" json:"group_id" xml:"group_id"`
	Created           uint   `db:"created" bson:"created" comment:"" json:"created" xml:"created"`
	Updated           uint   `db:"updated" bson:"updated" comment:"" json:"updated" xml:"updated"`
}

// - base function

func (a *NgingFrpServer) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingFrpServer) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingFrpServer) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingFrpServer) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingFrpServer) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingFrpServer) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingFrpServer) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingFrpServer) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingFrpServer) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingFrpServer) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingFrpServer) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingFrpServer) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingFrpServer) Objects() []*NgingFrpServer {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingFrpServer) XObjects() Slice_NgingFrpServer {
	return Slice_NgingFrpServer(a.Objects())
}

func (a *NgingFrpServer) NewObjects() factory.Ranger {
	return &Slice_NgingFrpServer{}
}

func (a *NgingFrpServer) InitObjects() *[]*NgingFrpServer {
	a.objects = []*NgingFrpServer{}
	return &a.objects
}

func (a *NgingFrpServer) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingFrpServer) Short_() string {
	return "nging_frp_server"
}

func (a *NgingFrpServer) Struct_() string {
	return "NgingFrpServer"
}

func (a *NgingFrpServer) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingFrpServer) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingFrpServer) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingFrpServer) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingFrpServer:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpServer(*v))
		case []*NgingFrpServer:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpServer(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFrpServer) GroupBy(keyField string, inputRows ...[]*NgingFrpServer) map[string][]*NgingFrpServer {
	var rows Slice_NgingFrpServer
	if len(inputRows) > 0 {
		rows = Slice_NgingFrpServer(inputRows[0])
	} else {
		rows = Slice_NgingFrpServer(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingFrpServer) KeyBy(keyField string, inputRows ...[]*NgingFrpServer) map[string]*NgingFrpServer {
	var rows Slice_NgingFrpServer
	if len(inputRows) > 0 {
		rows = Slice_NgingFrpServer(inputRows[0])
	} else {
		rows = Slice_NgingFrpServer(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingFrpServer) AsKV(keyField string, valueField string, inputRows ...[]*NgingFrpServer) param.Store {
	var rows Slice_NgingFrpServer
	if len(inputRows) > 0 {
		rows = Slice_NgingFrpServer(inputRows[0])
	} else {
		rows = Slice_NgingFrpServer(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingFrpServer) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingFrpServer:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpServer(*v))
		case []*NgingFrpServer:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpServer(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFrpServer) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.TcpMux) == 0 {
		a.TcpMux = "Y"
	}
	if len(a.Addr) == 0 {
		a.Addr = "0.0.0.0"
	}
	if len(a.ProxyAddr) == 0 {
		a.ProxyAddr = "0.0.0.0"
	}
	if len(a.LogFile) == 0 {
		a.LogFile = "console"
	}
	if len(a.LogWay) == 0 {
		a.LogWay = "console"
	}
	if len(a.LogLevel) == 0 {
		a.LogLevel = "info"
	}
	if len(a.DashboardAddr) == 0 {
		a.DashboardAddr = "0.0.0.0"
	}
	if len(a.DashboardUser) == 0 {
		a.DashboardUser = "admin"
	}
	if len(a.DashboardPwd) == 0 {
		a.DashboardPwd = "admin"
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

func (a *NgingFrpServer) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.TcpMux) == 0 {
		a.TcpMux = "Y"
	}
	if len(a.Addr) == 0 {
		a.Addr = "0.0.0.0"
	}
	if len(a.ProxyAddr) == 0 {
		a.ProxyAddr = "0.0.0.0"
	}
	if len(a.LogFile) == 0 {
		a.LogFile = "console"
	}
	if len(a.LogWay) == 0 {
		a.LogWay = "console"
	}
	if len(a.LogLevel) == 0 {
		a.LogLevel = "info"
	}
	if len(a.DashboardAddr) == 0 {
		a.DashboardAddr = "0.0.0.0"
	}
	if len(a.DashboardUser) == 0 {
		a.DashboardUser = "admin"
	}
	if len(a.DashboardPwd) == 0 {
		a.DashboardPwd = "admin"
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

func (a *NgingFrpServer) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingFrpServer) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	if val, ok := kvset["tcp_mux"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["tcp_mux"] = "Y"
		}
	}
	if val, ok := kvset["addr"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["addr"] = "0.0.0.0"
		}
	}
	if val, ok := kvset["proxy_addr"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["proxy_addr"] = "0.0.0.0"
		}
	}
	if val, ok := kvset["log_file"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["log_file"] = "console"
		}
	}
	if val, ok := kvset["log_way"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["log_way"] = "console"
		}
	}
	if val, ok := kvset["log_level"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["log_level"] = "info"
		}
	}
	if val, ok := kvset["dashboard_addr"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["dashboard_addr"] = "0.0.0.0"
		}
	}
	if val, ok := kvset["dashboard_user"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["dashboard_user"] = "admin"
		}
	}
	if val, ok := kvset["dashboard_pwd"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["dashboard_pwd"] = "admin"
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

func (a *NgingFrpServer) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.TcpMux) == 0 {
			a.TcpMux = "Y"
		}
		if len(a.Addr) == 0 {
			a.Addr = "0.0.0.0"
		}
		if len(a.ProxyAddr) == 0 {
			a.ProxyAddr = "0.0.0.0"
		}
		if len(a.LogFile) == 0 {
			a.LogFile = "console"
		}
		if len(a.LogWay) == 0 {
			a.LogWay = "console"
		}
		if len(a.LogLevel) == 0 {
			a.LogLevel = "info"
		}
		if len(a.DashboardAddr) == 0 {
			a.DashboardAddr = "0.0.0.0"
		}
		if len(a.DashboardUser) == 0 {
			a.DashboardUser = "admin"
		}
		if len(a.DashboardPwd) == 0 {
			a.DashboardPwd = "admin"
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
		if len(a.TcpMux) == 0 {
			a.TcpMux = "Y"
		}
		if len(a.Addr) == 0 {
			a.Addr = "0.0.0.0"
		}
		if len(a.ProxyAddr) == 0 {
			a.ProxyAddr = "0.0.0.0"
		}
		if len(a.LogFile) == 0 {
			a.LogFile = "console"
		}
		if len(a.LogWay) == 0 {
			a.LogWay = "console"
		}
		if len(a.LogLevel) == 0 {
			a.LogLevel = "info"
		}
		if len(a.DashboardAddr) == 0 {
			a.DashboardAddr = "0.0.0.0"
		}
		if len(a.DashboardUser) == 0 {
			a.DashboardUser = "admin"
		}
		if len(a.DashboardPwd) == 0 {
			a.DashboardPwd = "admin"
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

func (a *NgingFrpServer) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingFrpServer) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingFrpServer) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingFrpServer) Reset() *NgingFrpServer {
	a.Id = 0
	a.Name = ``
	a.Disabled = ``
	a.TcpMux = ``
	a.Addr = ``
	a.Port = 0
	a.UdpPort = 0
	a.KcpPort = 0
	a.ProxyAddr = ``
	a.VhostHttpPort = 0
	a.VhostHttpTimeout = 0
	a.VhostHttpsPort = 0
	a.LogFile = ``
	a.LogWay = ``
	a.LogLevel = ``
	a.LogMaxDays = 0
	a.Token = ``
	a.AuthTimeout = 0
	a.SubdomainHost = ``
	a.MaxPortsPerClient = 0
	a.MaxPoolCount = 0
	a.HeartBeatTimeout = 0
	a.UserConnTimeout = 0
	a.DashboardAddr = ``
	a.DashboardPort = 0
	a.DashboardUser = ``
	a.DashboardPwd = ``
	a.AllowPorts = ``
	a.Extra = ``
	a.Plugins = ``
	a.Uid = 0
	a.GroupId = 0
	a.Created = 0
	a.Updated = 0
	return a
}

func (a *NgingFrpServer) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["Name"] = a.Name
		r["Disabled"] = a.Disabled
		r["TcpMux"] = a.TcpMux
		r["Addr"] = a.Addr
		r["Port"] = a.Port
		r["UdpPort"] = a.UdpPort
		r["KcpPort"] = a.KcpPort
		r["ProxyAddr"] = a.ProxyAddr
		r["VhostHttpPort"] = a.VhostHttpPort
		r["VhostHttpTimeout"] = a.VhostHttpTimeout
		r["VhostHttpsPort"] = a.VhostHttpsPort
		r["LogFile"] = a.LogFile
		r["LogWay"] = a.LogWay
		r["LogLevel"] = a.LogLevel
		r["LogMaxDays"] = a.LogMaxDays
		r["Token"] = a.Token
		r["AuthTimeout"] = a.AuthTimeout
		r["SubdomainHost"] = a.SubdomainHost
		r["MaxPortsPerClient"] = a.MaxPortsPerClient
		r["MaxPoolCount"] = a.MaxPoolCount
		r["HeartBeatTimeout"] = a.HeartBeatTimeout
		r["UserConnTimeout"] = a.UserConnTimeout
		r["DashboardAddr"] = a.DashboardAddr
		r["DashboardPort"] = a.DashboardPort
		r["DashboardUser"] = a.DashboardUser
		r["DashboardPwd"] = a.DashboardPwd
		r["AllowPorts"] = a.AllowPorts
		r["Extra"] = a.Extra
		r["Plugins"] = a.Plugins
		r["Uid"] = a.Uid
		r["GroupId"] = a.GroupId
		r["Created"] = a.Created
		r["Updated"] = a.Updated
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "Name":
			r["Name"] = a.Name
		case "Disabled":
			r["Disabled"] = a.Disabled
		case "TcpMux":
			r["TcpMux"] = a.TcpMux
		case "Addr":
			r["Addr"] = a.Addr
		case "Port":
			r["Port"] = a.Port
		case "UdpPort":
			r["UdpPort"] = a.UdpPort
		case "KcpPort":
			r["KcpPort"] = a.KcpPort
		case "ProxyAddr":
			r["ProxyAddr"] = a.ProxyAddr
		case "VhostHttpPort":
			r["VhostHttpPort"] = a.VhostHttpPort
		case "VhostHttpTimeout":
			r["VhostHttpTimeout"] = a.VhostHttpTimeout
		case "VhostHttpsPort":
			r["VhostHttpsPort"] = a.VhostHttpsPort
		case "LogFile":
			r["LogFile"] = a.LogFile
		case "LogWay":
			r["LogWay"] = a.LogWay
		case "LogLevel":
			r["LogLevel"] = a.LogLevel
		case "LogMaxDays":
			r["LogMaxDays"] = a.LogMaxDays
		case "Token":
			r["Token"] = a.Token
		case "AuthTimeout":
			r["AuthTimeout"] = a.AuthTimeout
		case "SubdomainHost":
			r["SubdomainHost"] = a.SubdomainHost
		case "MaxPortsPerClient":
			r["MaxPortsPerClient"] = a.MaxPortsPerClient
		case "MaxPoolCount":
			r["MaxPoolCount"] = a.MaxPoolCount
		case "HeartBeatTimeout":
			r["HeartBeatTimeout"] = a.HeartBeatTimeout
		case "UserConnTimeout":
			r["UserConnTimeout"] = a.UserConnTimeout
		case "DashboardAddr":
			r["DashboardAddr"] = a.DashboardAddr
		case "DashboardPort":
			r["DashboardPort"] = a.DashboardPort
		case "DashboardUser":
			r["DashboardUser"] = a.DashboardUser
		case "DashboardPwd":
			r["DashboardPwd"] = a.DashboardPwd
		case "AllowPorts":
			r["AllowPorts"] = a.AllowPorts
		case "Extra":
			r["Extra"] = a.Extra
		case "Plugins":
			r["Plugins"] = a.Plugins
		case "Uid":
			r["Uid"] = a.Uid
		case "GroupId":
			r["GroupId"] = a.GroupId
		case "Created":
			r["Created"] = a.Created
		case "Updated":
			r["Updated"] = a.Updated
		}
	}
	return r
}

func (a *NgingFrpServer) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "tcp_mux":
			a.TcpMux = param.AsString(value)
		case "addr":
			a.Addr = param.AsString(value)
		case "port":
			a.Port = param.AsUint(value)
		case "udp_port":
			a.UdpPort = param.AsUint(value)
		case "kcp_port":
			a.KcpPort = param.AsUint(value)
		case "proxy_addr":
			a.ProxyAddr = param.AsString(value)
		case "vhost_http_port":
			a.VhostHttpPort = param.AsUint(value)
		case "vhost_http_timeout":
			a.VhostHttpTimeout = param.AsUint64(value)
		case "vhost_https_port":
			a.VhostHttpsPort = param.AsUint(value)
		case "log_file":
			a.LogFile = param.AsString(value)
		case "log_way":
			a.LogWay = param.AsString(value)
		case "log_level":
			a.LogLevel = param.AsString(value)
		case "log_max_days":
			a.LogMaxDays = param.AsUint(value)
		case "token":
			a.Token = param.AsString(value)
		case "auth_timeout":
			a.AuthTimeout = param.AsUint64(value)
		case "subdomain_host":
			a.SubdomainHost = param.AsString(value)
		case "max_ports_per_client":
			a.MaxPortsPerClient = param.AsInt64(value)
		case "max_pool_count":
			a.MaxPoolCount = param.AsUint(value)
		case "heart_beat_timeout":
			a.HeartBeatTimeout = param.AsUint(value)
		case "user_conn_timeout":
			a.UserConnTimeout = param.AsUint(value)
		case "dashboard_addr":
			a.DashboardAddr = param.AsString(value)
		case "dashboard_port":
			a.DashboardPort = param.AsUint(value)
		case "dashboard_user":
			a.DashboardUser = param.AsString(value)
		case "dashboard_pwd":
			a.DashboardPwd = param.AsString(value)
		case "allow_ports":
			a.AllowPorts = param.AsString(value)
		case "extra":
			a.Extra = param.AsString(value)
		case "plugins":
			a.Plugins = param.AsString(value)
		case "uid":
			a.Uid = param.AsUint(value)
		case "group_id":
			a.GroupId = param.AsUint(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		}
	}
}

func (a *NgingFrpServer) Set(key interface{}, value ...interface{}) {
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
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "TcpMux":
			a.TcpMux = param.AsString(vv)
		case "Addr":
			a.Addr = param.AsString(vv)
		case "Port":
			a.Port = param.AsUint(vv)
		case "UdpPort":
			a.UdpPort = param.AsUint(vv)
		case "KcpPort":
			a.KcpPort = param.AsUint(vv)
		case "ProxyAddr":
			a.ProxyAddr = param.AsString(vv)
		case "VhostHttpPort":
			a.VhostHttpPort = param.AsUint(vv)
		case "VhostHttpTimeout":
			a.VhostHttpTimeout = param.AsUint64(vv)
		case "VhostHttpsPort":
			a.VhostHttpsPort = param.AsUint(vv)
		case "LogFile":
			a.LogFile = param.AsString(vv)
		case "LogWay":
			a.LogWay = param.AsString(vv)
		case "LogLevel":
			a.LogLevel = param.AsString(vv)
		case "LogMaxDays":
			a.LogMaxDays = param.AsUint(vv)
		case "Token":
			a.Token = param.AsString(vv)
		case "AuthTimeout":
			a.AuthTimeout = param.AsUint64(vv)
		case "SubdomainHost":
			a.SubdomainHost = param.AsString(vv)
		case "MaxPortsPerClient":
			a.MaxPortsPerClient = param.AsInt64(vv)
		case "MaxPoolCount":
			a.MaxPoolCount = param.AsUint(vv)
		case "HeartBeatTimeout":
			a.HeartBeatTimeout = param.AsUint(vv)
		case "UserConnTimeout":
			a.UserConnTimeout = param.AsUint(vv)
		case "DashboardAddr":
			a.DashboardAddr = param.AsString(vv)
		case "DashboardPort":
			a.DashboardPort = param.AsUint(vv)
		case "DashboardUser":
			a.DashboardUser = param.AsString(vv)
		case "DashboardPwd":
			a.DashboardPwd = param.AsString(vv)
		case "AllowPorts":
			a.AllowPorts = param.AsString(vv)
		case "Extra":
			a.Extra = param.AsString(vv)
		case "Plugins":
			a.Plugins = param.AsString(vv)
		case "Uid":
			a.Uid = param.AsUint(vv)
		case "GroupId":
			a.GroupId = param.AsUint(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		}
	}
}

func (a *NgingFrpServer) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["name"] = a.Name
		r["disabled"] = a.Disabled
		r["tcp_mux"] = a.TcpMux
		r["addr"] = a.Addr
		r["port"] = a.Port
		r["udp_port"] = a.UdpPort
		r["kcp_port"] = a.KcpPort
		r["proxy_addr"] = a.ProxyAddr
		r["vhost_http_port"] = a.VhostHttpPort
		r["vhost_http_timeout"] = a.VhostHttpTimeout
		r["vhost_https_port"] = a.VhostHttpsPort
		r["log_file"] = a.LogFile
		r["log_way"] = a.LogWay
		r["log_level"] = a.LogLevel
		r["log_max_days"] = a.LogMaxDays
		r["token"] = a.Token
		r["auth_timeout"] = a.AuthTimeout
		r["subdomain_host"] = a.SubdomainHost
		r["max_ports_per_client"] = a.MaxPortsPerClient
		r["max_pool_count"] = a.MaxPoolCount
		r["heart_beat_timeout"] = a.HeartBeatTimeout
		r["user_conn_timeout"] = a.UserConnTimeout
		r["dashboard_addr"] = a.DashboardAddr
		r["dashboard_port"] = a.DashboardPort
		r["dashboard_user"] = a.DashboardUser
		r["dashboard_pwd"] = a.DashboardPwd
		r["allow_ports"] = a.AllowPorts
		r["extra"] = a.Extra
		r["plugins"] = a.Plugins
		r["uid"] = a.Uid
		r["group_id"] = a.GroupId
		r["created"] = a.Created
		r["updated"] = a.Updated
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "name":
			r["name"] = a.Name
		case "disabled":
			r["disabled"] = a.Disabled
		case "tcp_mux":
			r["tcp_mux"] = a.TcpMux
		case "addr":
			r["addr"] = a.Addr
		case "port":
			r["port"] = a.Port
		case "udp_port":
			r["udp_port"] = a.UdpPort
		case "kcp_port":
			r["kcp_port"] = a.KcpPort
		case "proxy_addr":
			r["proxy_addr"] = a.ProxyAddr
		case "vhost_http_port":
			r["vhost_http_port"] = a.VhostHttpPort
		case "vhost_http_timeout":
			r["vhost_http_timeout"] = a.VhostHttpTimeout
		case "vhost_https_port":
			r["vhost_https_port"] = a.VhostHttpsPort
		case "log_file":
			r["log_file"] = a.LogFile
		case "log_way":
			r["log_way"] = a.LogWay
		case "log_level":
			r["log_level"] = a.LogLevel
		case "log_max_days":
			r["log_max_days"] = a.LogMaxDays
		case "token":
			r["token"] = a.Token
		case "auth_timeout":
			r["auth_timeout"] = a.AuthTimeout
		case "subdomain_host":
			r["subdomain_host"] = a.SubdomainHost
		case "max_ports_per_client":
			r["max_ports_per_client"] = a.MaxPortsPerClient
		case "max_pool_count":
			r["max_pool_count"] = a.MaxPoolCount
		case "heart_beat_timeout":
			r["heart_beat_timeout"] = a.HeartBeatTimeout
		case "user_conn_timeout":
			r["user_conn_timeout"] = a.UserConnTimeout
		case "dashboard_addr":
			r["dashboard_addr"] = a.DashboardAddr
		case "dashboard_port":
			r["dashboard_port"] = a.DashboardPort
		case "dashboard_user":
			r["dashboard_user"] = a.DashboardUser
		case "dashboard_pwd":
			r["dashboard_pwd"] = a.DashboardPwd
		case "allow_ports":
			r["allow_ports"] = a.AllowPorts
		case "extra":
			r["extra"] = a.Extra
		case "plugins":
			r["plugins"] = a.Plugins
		case "uid":
			r["uid"] = a.Uid
		case "group_id":
			r["group_id"] = a.GroupId
		case "created":
			r["created"] = a.Created
		case "updated":
			r["updated"] = a.Updated
		}
	}
	return r
}

func (a *NgingFrpServer) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingFrpServer) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}
