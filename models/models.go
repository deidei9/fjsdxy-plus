package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
	"time"
)

type Model struct {
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	UpdateTime time.Time `orm:"null;auto_now;type(datetime)" json:"-"`
	DeleteTime time.Time `orm:"null" json:"-"`
}

func init() {
	//初始化数据库
	runmode := beego.AppConfig.String("runmode")
	if runmode == "prod" {
		orm.Debug = false
	} else {
		orm.Debug = true
	}
	RegisterDB()
}

//注册数据库
func RegisterDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	models := []interface{}{
		NewUser(),
		NewStudent(),
		NewClasses(),
		NewCourse(),
		NewExam(),
		NewWechat(),
		NewWeek(),
	}
	orm.RegisterModelWithPrefix(beego.AppConfig.DefaultString("db::prefix", "fjsdxy_"), models...)
	dbUser := beego.AppConfig.String("db::user")
	dbPassword := beego.AppConfig.String("db::password")
	dbDatabase := beego.AppConfig.String("db::database")
	dbCharset := beego.AppConfig.String("db::charset")
	dbHost := beego.AppConfig.String("db::host")
	dbPort := beego.AppConfig.String("db::port")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", dbUser, dbPassword, dbHost, dbPort, dbDatabase, dbCharset)
	maxIdle := beego.AppConfig.DefaultInt("db::maxIdle", 50)
	maxConn := beego.AppConfig.DefaultInt("db::maxConn", 300)
	if err := orm.RegisterDataBase("default", "mysql", dbLink, maxIdle, maxConn); err != nil {
		panic(err)
	}
}

func DB() orm.Ormer {
	return orm.NewOrm()
}

//获取带表前缀的数据表
//@param            table               数据表
func GetTable(table string) string {
	prefix := beego.AppConfig.DefaultString("db::prefix", "fjsdxy_")
	if !strings.HasPrefix(table, prefix) {
		table = prefix + table
	}
	return table
}

//根据指定的表和id删除指定的记录，如果在删除记录的时候也删除记录中记录的文件，则不能调用该方法
//@param            table                   指定要删除记录的数据表
//@param            id                      要删除的记录的ID
//@return           affected                影响的记录数
//@return           err                     错误
func DelByIds(table string, id ...interface{}) (affected int64, err error) {
	return DB().QueryTable(GetTable(table)).Filter("Id__in", id...).Delete()
}

//根据指定的表和id条件更新表字段，不支持批量更新
//@param            table                   需要更新的表
//@param            field                   需要更新的字段
//@param            value                   需要更新的字段的值
//@param            id                      id条件
//@return           affected                影响的记录数
//@return           err                     错误
func UpdateByIds(table string, field string, value interface{}, id ...interface{}) (affected int64, err error) {
	return DB().QueryTable(GetTable(table)).Filter("Id__in", id...).Update(orm.Params{
		field: value,
	})
}

//根据指定的表和指定条件更新表字段，不支持批量更新
//@param            table                   需要更新的表
//@param            data                    需要更新的字段
//@param            filter                  过滤条件，如"Id__in"
//@param            filterValue             过滤条件的值
//@return           affected                影响的记录数
//@return           err                     错误
func UpdateByField(table string, data map[string]interface{}, filter string, filterValue ...interface{}) (affected int64, err error) {
	return DB().QueryTable(GetTable(table)).Filter(filter, filterValue...).Update(data)
}

//从单表中根据ID获取数据
//@param            table           需要查询的表
//@param            id              ID
//@param			field			查询字段
//@return           data          数据列表
//@return           err             错误
func GetById(table string, id int, field ...string) (data orm.Params, err error) {
	var params []orm.Params
	_, err = DB().QueryTable(GetTable(table)).Filter("id", id).Limit(1).Values(&params, field...)
	fmt.Println(params)
	if err == nil {
		for k, v := range params {
			if k == 0 {
				return v, err
			}
		}
	}
	return nil, err
}

//从单表中根据条件获取数据列表
//@param            table           需要查询的表
//@param            p               页码
//@param            listRows        每页显示记录数
//@param            condition       查询条件
//@param            orderby         排序
//@return           params          数据列表
//@return           rows            返回的记录数
//@return           err             错误
func GetList(table string, p, listRows int, condition *orm.Condition, orderby ...string) (params []orm.Params, rows int64, err error) {
	rows, err = DB().QueryTable(GetTable(table)).SetCond(condition).Limit(listRows).Offset((p - 1) * listRows).OrderBy(orderby...).Values(&params)
	return params, rows, err
}

//获取指定Strut的字段
//@param            tableObj        Strut结构对象，引用传递
//@return           fields          返回字段数组
func GetFields(tableObj interface{}) (fields []string) {
	elem := reflect.ValueOf(tableObj).Elem()
	for i := 0; i < elem.NumField(); i++ {
		fields = append(fields, elem.Type().Field(i).Name)
	}
	return fields
}

//左联合查询创建SQL语句
//@param                tables                  需要作为联合查询的数据表。注意：数据表的第一个表是主表
//@param                on                      联合查询的on查询条件，必须必表(tables)少一个。比如user表和user_info表做联合查询，那么on查询条件只有一个，必tables的数组元素少一个
//@param                fields                  需要查询的字段
//@param                p                       页码
//@param                listRows                每页查询记录数
//@param                orderBy                 排序条件，可以穿空数组
//@param                groupBy                 按组查询
//@param                condition               查询条件
//@param                conditionArgs           查询条件参数
//@return               sql                     返回生成的SQL语句
//@return               err                     错误。如果返回的错误不为nil，则SQL语句为空字符串
//使用示例：
//tables := []string{"document", "document_info info", "document_store store"}
//fields := map[string][]string{
//"document": {"Id Did", "Title", "Filename"},
//"info":     {"Vcnt", "Dcnt"},
//"store":    {"Md5", "Page"},
//}
//on := []map[string]string{
//{"document.Id": "info.Id"},
//{"info.DsId": "store.Id"},
//}
//orderby := []string{"doc.Id desc", "store.Page desc"}
//sql, err := LeftJoinSqlBuild(tables, on, fields, 1, 100, orderby, nil, "")
//fmt.Println(sql, err)
func LeftJoinSqlBuild(tables []string, on []map[string]string, fields map[string][]string, p, listRows int, orderBy []string, groupBy []string, condition string) (sql string, err error) {
	if len(tables) < 2 || len(tables)-1 != len(on) {
		err = errors.New("参数不规范：联合查询的数据表数量必须在2个或2个以上，同时表数量比on条件多一个")
		return
	}
	var (
		FieldSlice   []string
		StrOrderBy   string
		StrGroupBy   string
		StrCondition string
		joinKV       string
		join         = []string{tables[0]}
		usedTables   = []string{}
	)
	for table, field := range fields {
		for _, f := range field {
			FieldSlice = append(FieldSlice, strings.Trim(fmt.Sprintf("%v.%v", table, f), "."))
		}
	}
	for index, table := range tables {
		slice := strings.Split(strings.TrimSpace(table), " ")
		if len(slice) == 1 {
			slice = append(slice, slice[0])
		}
		usedTables = append(usedTables, slice[1])
		if index > 0 {
			on, joinKV = joinOn(slice[1], usedTables, on)
			join = append(join, "left join "+table+" on "+joinKV)
		}
	}
	if len(orderBy) > 0 {
		StrOrderBy = " order by " + strings.Join(orderBy, ",")
	}
	if len(condition) > 0 {
		StrCondition = " where " + condition
	}
	if len(groupBy) > 0 {
		StrGroupBy = " group by " + strings.Join(groupBy, ",")
	}

	sql = fmt.Sprintf("select %v from %v %v %v %v limit %v offset %v", strings.Join(FieldSlice, ","), strings.Join(join, " "), StrCondition, StrGroupBy, StrOrderBy, listRows, (p-1)*listRows)
	return
}

//只供LeftJoinSqlBuild创建SQL语句使用
//@param                table               需要左联查询的表
//@param                usedTables          已使用的表
//@param                on                  联合查询条件
//@return               newon               新的联合查询条件(返回未被使用的联合查询条件)
//@return               ret                 返回组装联合查询条件
func joinOn(table string, usedTables []string, on []map[string]string) (newon []map[string]string, ret string) {
	table = table + "."
	lenon := len(on)
	for index, v := range on {
		for key, val := range v {
			if strings.HasPrefix(key, table) || strings.HasPrefix(val, table) {
				for _, used := range usedTables {
					if strings.HasPrefix(key, used) || strings.HasPrefix(val, used) {
						ret = key + "=" + val
						if index > 0 {
							newon = append(newon, on[0:index]...)
						}
						if index+1 <= lenon {
							newon = append(newon, on[(index+1):]...)
						}
						return
					}
				}
			}
		}
	}
	return
}

//对单表记录进行统计查询
//@param            table           需要查询或者统计的表
//@param            cond            查询条件
//@return           cnt             统计的记录数
func Count(table string, cond *orm.Condition) (cnt int64) {
	cnt, _ = DB().QueryTable(GetTable(table)).SetCond(cond).Count()
	return
}

//检查数据库是否存在
//@param            host            数据库地址
//@param            port            端口
//@param            password        密码
//@param            database        数据库
//@return           err             错误
func CheckDatabaseIsExist(host string, port int, username, password, database string) (err error) {
	var (
		db      *sql.DB
		timeout = make(chan bool, 1)
	)
	go func() {
		if db, err = sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
			username, password, host, port, database,
		)); err == nil {
			err = db.Ping()
		}
		timeout <- false
	}()
	time.AfterFunc(3*time.Second, func() {
		timeout <- true
	})

	if t := <-timeout; t {
		err = errors.New("MySQL数据库连接失败，请检查数据库链接是否正确")
	}
	return
}
