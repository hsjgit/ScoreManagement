# ScoreManagement

假设我们为老师提供一个成绩统计系统，该系统支持批量多份上传成绩结果。
成绩文件为多行，每行内容为：班级、姓名、分数、科目。
上传后的成绩统计可根据不同的字段进行排序和查询某一个用户的科目分数信息。


```mermaid
graph TB
	main(入口) --> Bootstrap(程序加载)
    Bootstrap -->|连接数据库| mysql[数据库]
    Bootstrap -->|注册路由| route[路由]
    mysql -->|连接失败| End(结束)
    route --> upload[上传文件]
    route --> select[查询成绩]
    upload --> save[解析文件保存到数据库]
    upload --> back[返回上传结果]
    select --> result[获取查询结果]
    back --> err[统一错误处理]
    result -->err[统一错误处理]
```

