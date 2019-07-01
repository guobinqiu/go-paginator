DB Paginator
===

### 获得一个db连接
```
db, err := sql.Open("mysql", <YOUR DSN>)
```

### 初始化一个分页对象，传入刚才创建得db连接
```
paginator := paginator.New(db)
```

### 用分页对象创建查询一个对象，传入sql查询语句及不定参查询条件
```
query := paginator.CreateQuery(<YOUR SQL>, ...CONDS)
```  

### 返回分页结果对象
```
pagination, err := paginator.Paginate(query, <CURRENT_PAGE>, <PAGE_SIZE>)

// 获取分页信息
page := pagination.Page
pageSize := pagination.PageSize
pageCount := pagination.PageCount
rowCount := pagnination.RowCount
```

### 获取所有行
```
rows = pagination.Rows
```

### 获取单行
```
row := pagination.RowIndex(0)
```

### 以指定类型返回单行里某列的值
```
v, err := row.String(<COLUMN_NAME>)
v, err := row.Int(<COLUMN_NAME>)
v, err := row.Float(<COLUMN_NAME>)
v, err := row.Time(<COLUMN_NAME>)
