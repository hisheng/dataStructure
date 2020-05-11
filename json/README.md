 https://colobu.com/2017/06/21/json-tricks-in-Go/
 
 一
 用字符串传递数字
 ```sql
 type TestObject struct {
     Field1 int    `json:",string"`
 }
```

 这个对应的json是 {"Field1": "100"}
 如果json是 {"Field1": 100} 则会报错

二 多个的化，比如 Master 这样也有利于理解，直接写到一起
```sql
type MysqlConfig2 struct {
	MaxIdle     int `json:"max_idle"`
	MaxOpen     int `json:"max_open"`
	MaxLifetime int `json:"max_lifetime"`
	Master      struct {
		Dsn string `json:"dsn"`
	} `json:"master"`
	Slaves []struct {
		Dsn string `json:"dsn"`
	} `json:"slaves"`
}
```