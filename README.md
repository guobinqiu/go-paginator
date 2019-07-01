DB Paginator
===

### 1.获得db连接
```
db, err := sql.Open("mysql", <YOUR DSN>)
```

### 2.创建分页对象，传入db连接
```
paginator := paginator.New(db)
```

### 3.创建查询对象，传入sql查询语句及查询条件
```
query := paginator.CreateQuery(<YOUR SQL>, ...CONDS)
```  

### 4.输入页号及页大小返回分页结果对象
```
//返回分页结果对象
pagination, err := paginator.Paginate(query, <PAGE>, <PAGE_SIZE>)

//获取分页信息
page := pagination.Page //第几页
pageSize := pagination.PageSize //每页大小
pageCount := pagination.PageCount //总页数
rowCount := pagnination.RowCount //总行数

//获取所有行数据
rows = pagination.Rows

//获取第一行数据
row := pagination.RowIndex(0)

//以指定类型返回行内列的值
v, err := row.String(<COLUMN_NAME>) //as string
v, err := row.Int(<COLUMN_NAME>) //as int
v, err := row.Float(<COLUMN_NAME>) //as float64
v, err := row.Time(<COLUMN_NAME>) //as time.Time
```
