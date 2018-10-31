# interview

## 算法与数据结构


#### [拜托，面试别再问我TopK了！！！](https://mp.weixin.qq.com/s/FFsvWXiaZK96PtUg-mmtEw)
#### [小米三面面试题](https://github.com/TomorrowWu/golang-algorithms/blob/master/algorithms/uncategorized/%E5%B0%8F%E7%B1%B3%E4%B8%89%E9%9D%A2%E9%9D%A2%E8%AF%95%E9%A2%98/README.md)

### 数组
#### [实现Array](https://github.com/TomorrowWu/golang-algorithms/blob/master/data-structures/array/array.go)

### 链表
#### [实现linked-list及相关操作,链表反转,判断循环链表,合并两个有序链表,删除倒数第n个节点](https://github.com/TomorrowWu/golang-algorithms/blob/master/data-structures/linked-list/singlelinkedlist.go)
#### [单向链表判断是否回文](https://github.com/TomorrowWu/golang-algorithms/blob/master/data-structures/linked-list/palindrome.go)

### 栈
#### [为什么函数调用要用“栈”来保存临时变量]()
#### [JVM中的"栈"和数据结构中的"栈"是一回事吗?](算法与数据结构/栈/JVM中的"栈"和数据结构中的"栈"是一回事吗?.md)
#### [栈在表达式求值中应用(带小括号)]()
#### [基于数组实现栈数据结构]()
#### [基于单链表实现栈数据结构]()
#### [基于栈结构实现浏览器前进和后退功能]()

### 队列
#### [哪些场景中会用到队列的排队请求?](算法与数据结构/队列/哪些场景中会用到队列的排队请求?.md)
#### [实现无锁并发队列]()

### 递归
#### [对递归的理解](算法与数据结构/递归/对递归的理解.md)

### 排序
#### [特定算法是依赖特定的数据结构的,排序算法一般都是基于数组实现的,如果数据存储在链表中,排序算法还能工作吗?如果能,那相应的时间复杂度、空间复杂度又是多少呢?](算法与数据结构/排序/基于链表存储,排序算法思考?.md)
#### [10个接口访问日志文件,每个文件约300MB,每个文件的日志都是按照时间戳从小到大排序的,将10个文件合并为1个文件,仍按照时间戳从小到大排列,如果处理任务的机器内存只有1GB,有什么好的解决思路,能"快速"地将这10个文件合并?]()
桶排序,合并多个桶
#### [在O(n)内查找第K大元素,代码实现](https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0215.kth-largest-element-in-an-array/src/kth-largest-element-in-an-array.go)
#### 根据年龄,给100万用户排序
计数排序
#### [拜托，面试别再问我计数排序了！！！](https://mp.weixin.qq.com/s/KU-AUGOnLeRtE_hivl2uSA)
#### [拜托，面试别再问我基数排序了！！！](https://mp.weixin.qq.com/s/Z8gU9QLpMnA-zoMc9ZeR2w)
#### 如何在1000万个整数中快速查找某个整数?(内存限制是100MB,每个数据大小是8字节)
```
- 解决: 将数据存储在数组中,内存占用80MB,从小到大排序,再二分查找算法
- 散列表和二叉树等支持快速查找的动态数据结构,大部分情况下,二分查找可以解决的问题,用上述两种数据结构都可以解决
- 但是,上述两种结构需要比较多的额外的内存空间,100MB内存存不下
- 二分查找底层依赖的是数组,除了数据本身以外,不需要额外存储其他信息,最省内存
```
#### 如何编程实现“求一个数的平方根”？要求精确到小数点后 6 位。

#### 如果数据使用链表存储，二分查找的时间复杂就会变得很高，那查找的时间复杂度究竟是多少呢？


## 网络协议
#### [如果一个局域网里面有多个交换机,ARP广播的模式会出现什么问题呢?](网络协议/一个局域网里面有多个交换机,会出现什么问题.md)
#### [ARP协议,已知IP地址求MAC;RARP协议,已知MAC求IP,可以用来做什么?](网络协议/RARP协议的作用.md)
#### [STP协议能够很好的解决环路问题,但是也有它的缺点,你能举几个例子吗?](网络协议/STP协议的缺点.md)
#### [在一个比较大的网络中,如果两台机器不通,你知道应该用什么方式调试吗?](网络协议/两台机器不通,如何调试?.md)
#### [你觉得基于RTMP的视频流传输的机制存在什么问题？如何进行优化？](网络协议/揭秘RTMP流媒体协议.md)
#### [把电影下载下来看，那么大，如何快速下载？](网络协议/快速下载.md)

## MySQL


## 网络协议
#### [TCP和UDP的区别](网络协议/TCP和UDP的区别.md)
#### [UDP的特点及使用场景](网络协议/UDP的特点及使用场景.md)


## 系统设计
#### [基于Redis做IP限速](https://redis.io/commands/incr)
