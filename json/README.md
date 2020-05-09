 https://colobu.com/2017/06/21/json-tricks-in-Go/
 
 用字符串传递数字
 type TestObject struct {
     Field1 int    `json:",string"`
 }
 这个对应的json是 {"Field1": "100"}
 
 如果json是 {"Field1": 100} 则会报错
