#### 作业题目
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码

#### 思路
1. DAO层的作用是相当于一个黑盒，对数据库类型不明感，那么对于sql.ErrNoRows这类error，其实在DAO层可以做一个转化，统一为NoFoundData
2. DAO层发生sql.ErrNoRows之后，应该wrap错误到上层，以便定位错误位置
