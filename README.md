# interview

## 算法与数据结构

### 常见面试题
#### [拜托，面试别再问我TopK了！！！](https://mp.weixin.qq.com/s/FFsvWXiaZK96PtUg-mmtEw)
#### [小米三面面试题](https://github.com/TomorrowWu/golang-algorithms/blob/master/algorithms/uncategorized/%E5%B0%8F%E7%B1%B3%E4%B8%89%E9%9D%A2%E9%9D%A2%E8%AF%95%E9%A2%98/README.md)
#### [判断是否为2的幂](算法与数据结构/常见面试题/判断是否为2的幂.md)
#### [大数乘法](https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0043.multiply-strings/src/multiply-strings.go)
#### [三门问题](https://zh.wikipedia.org/wiki/%E8%92%99%E6%8F%90%E9%9C%8D%E7%88%BE%E5%95%8F%E9%A1%8C)

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

### 跳表
#### [为什么Redis要用跳表来实现有序集合，而不是红黑树](算法与数据结构/跳表/为什么Redis要用跳表来实现有序集合，而不是红黑树.md)

### 递归
#### [对递归的理解](算法与数据结构/递归/对递归的理解.md)

### 排序
#### [特定算法是依赖特定的数据结构的,排序算法一般都是基于数组实现的,如果数据存储在链表中,排序算法还能工作吗?如果能,那相应的时间复杂度、空间复杂度又是多少呢?](算法与数据结构/排序/基于链表存储,排序算法思考?.md)
#### [10个接口访问日志文件,每个文件约300MB,每个文件的日志都是按照时间戳从小到大排序的,将10个文件合并为1个文件,仍按照时间戳从小到大排列,如果处理任务的机器内存只有1GB,有什么好的解决思路,能"快速"地将这10个文件合并?]()
```
先构建十条io流，分别指向十个文件，每条io流读取对应文件的第一条数据，
然后比较时间戳，选择出时间戳最小的那条数据，将其写入一个新的文件，
然后指向该时间戳的io流读取下一行数据，然后继续刚才的操作，比较选出最小的时间戳数据，写入新文件，io流读取下一行数据，
以此类推，完成文件的合并， 
这种处理方式，日志文件有n个,数据就要比较n次，每次比较选出一条数据来写入，时间复杂度是O（n），空间复杂度是O（1）,几乎不占用内存，
```
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

### 二分查找
#### [如何快速定位IP对应的省份地址？](算法与数据结构/二分查找/如何快速定位IP对应的省份地址.md)

### 散列表（Hash表）
#### [如何设计一个工业级的散列表](算法与数据结构/散列表/如何设计一个工业级的散列表.md)

### 树
#### [说说树的概念](算法与数据结构/树/树的概念.md)
#### [给定一组数据，比如 1，3，5，6，9，10，可以构建出多少种不同的二叉树？]()
#### [如何实现按层遍历二叉树？]()
#### [散列表如此高效，为什么还需要二叉树？](算法与数据结构/树/散列表、二叉查找树对比.md)
#### 如何通过编程，求给定一棵二叉树的确切高度
```
- 第一种是深度优先思想的递归，分别求左右子树的高度。当前节点的高度就是左右子树中较大的那个+1
- 第二种可以采用层次遍历的方式，每一层记录都记录下当前队列的长度，这个是队尾，每一层队头从0开始。
然后每遍历一个元素，队头下标+1。直到队头下标等于队尾下标。这个时候表示当前层遍历完成。
每一层刚开始遍历的时候，树的高度+1。最后队列为空，就能得到树的高度。
```

### 哈希算法
#### [哈希算法有哪些应用场景](算法与数据结构/哈希算法/哈希算法的应用场景.md)

##   计算机网络
#### [如果一个局域网里面有多个交换机,ARP广播的模式会出现什么问题呢?](网络协议/一个局域网里面有多个交换机,会出现什么问题.md)
#### [ARP协议,已知IP地址求MAC;RARP协议,已知MAC求IP,可以用来做什么?](网络协议/RARP协议的作用.md)
#### [STP协议能够很好的解决环路问题,但是也有它的缺点,你能举几个例子吗?](网络协议/STP协议的缺点.md)
#### [在一个比较大的网络中,如果两台机器不通,你知道应该用什么方式调试吗?](网络协议/两台机器不通,如何调试?.md)
#### [你觉得基于RTMP的视频流传输的机制存在什么问题？如何进行优化？](网络协议/揭秘RTMP流媒体协议.md)
#### [把电影下载下来看，那么大，如何快速下载？](网络协议/快速下载.md)
#### 负载均衡为什么要分地址和运营商呢？
```
为了返回最优的IP，即离用户最近的IP，提高访问速度
```

## MySQL
### [什么是数据库事务]()
```
一次会话，多个命令要么都执行成功，要么都执行失败
```
### [ACID 四大特性]()
### [SQL注入攻击](https://lovecoding.club/2018/12/12/web%E5%AE%89%E5%85%A8-sql%E6%B3%A8%E5%85%A5.html)


## Redis
### [RDB与AOF两种持久化](http://redisdoc.com/topic/persistence.html)


## 网络协议
#### [TCP和UDP的区别](网络协议/TCP和UDP的区别.md)
#### [UDP的特点及使用场景](网络协议/UDP的特点及使用场景.md)
#### [用websocket连接服务器时，为什么要自己做心跳机制，tcp不是有心跳机制？](网络协议/TCP和UDP的区别.md)
#### [HTTP状态码](网络协议/HTTP状态码.md)


## 分布式系统
#### [PoW和Paxos/Raft对比](分布式系统/PoW,Paxos,Raft对比.md)

## 区块链
#### [比特币中PoW的优劣](区块链/比特币中PoW的优劣.md)
#### Base58编码转换过程
```
相当于将一个二进制数字，转换为58进制
```

## 系统设计
#### [基于Redis做IP限速](https://redis.io/commands/incr)
```
counter = current_time # for example 15:03
count = INCR counter
EXPIRE counter 60 # just to make sure redis doesn't store it forever

if count > 5
  print "Exceeded the limit"
```
#### [秒杀](https://mp.weixin.qq.com/s?__biz=MzU2OTUyNzk1NQ==&mid=2247491172&idx=1&sn=1b3c5ca9b043bafd23687f9a55946492&source=41#wechat_redirect)

## 操作系统
#### [面试必备之乐观锁与悲观锁](https://lovecoding.club/2018/11/20/%E4%B9%90%E8%A7%82%E9%94%81%E4%B8%8E%E6%82%B2%E8%A7%82%E9%94%81.html)
#### [多版本并发控制(MVCC)在分布式系统中的应用](https://coolshell.cn/articles/6790.html)
#### [进程、线程、协程区别]()
```
- 进程是系统进行资源分配和调度的一个独立单位。每个进程都有自己的独立内存空间，不同进程通过进程间通信来通信。由于进程比较重量，占据独立的内存，所以上下文进程间的切换开销（栈、寄存器、虚拟内存、文件句柄等）比较大，但相对比较稳定安全。
- 线程是进程的一个实体,是CPU调度和分派的基本单位，只拥有一点在运行中必不可少的资源(如程序计数器,一组寄存器和栈)，线程间通信主要通过共享内存，上下文切换很快，资源开销较少，但相比进程不够稳定容易丢失数据
- 协程是一种用户态的轻量级线程，协程的调度完全由用户控制。协程拥有自己的寄存器上下文和栈                                                     

```

## Linux
### 你常用的5个Linux命令
```
1. grep: 文本搜索   ps -ef|grep java:查找指定进程
2. cp: 复制文件     cp file1 file2 file3 dir ：把文件file1、file2、file3复制到目录dir中
3. mv: 移动文件、目录或更名   mv file1 file2 file3 dir : 把文件file1、file2、file3移动到目录dir中   mv file1 file2 : 把文件file1重命名为file2  
4. rm
5. ps: 列出系统中当前运行的进程
6. kill: 杀死进程
7. tar: 压缩，解压
8. chmod: 改变文件权限
9. tail:  tail -f 20160921.logs ：查看正在改变的日志文件   tail -3000 catalina.out：查看倒数前3000行的数据  tail -3000 catalina.out | grep 'AA'：查看倒数前3000行包含字母'AA'的数据
10. lsof -i:8080：根据端口查看进程pid
11. ps -ef|grep ganache-cli | grep -v grep|sort -nrk2|head -n 1|awk '{print $1,$4}'
-v : 屏蔽 
wc chap1: 显示文件中的行数，字数，和字节数
sort -nrk2: r: 从大到小 k2:第几列
head -n 1: 头部的 1行
awk '{print $1,$4}' log.txt: 输出文本中的1、4项
12. top:查看当前系统的运行情况
13. tar: 解压
14. du -sh * | sort -n 
查看当前目录大小，并按文件大小排序
```
###  查看进程
```
1，ps -ef|grep mysql
2. netstat -anp|grep 3306
3. lsof -i:3306    (list of file )
```
### Linux系统的/proc目录的作用
```
在GUN/Linux操作系统中，/proc是一个位于内存中的伪文件系统(in-memory pseudo-file system)。
该目录下保存的不是真正的文件和目录，而是一些“运行时”信息，如系统内存、磁盘io、设备挂载信息和硬件配置信息等。
proc目录是一个控制中心，用户可以通过更改其中某些文件来改变内核的运行状态。
proc目录也是内核提供给我们的查询中心，我们可以通过这些文件查看有关系统硬件及当前正在运行进程的信息
```

## 分布式
### 分布式系统、集群、负载均衡
```
不同机器部署不同的服务，对外暴露一个系统
不同机器相同的服务，对外暴露一个服务
平衡的将任务分给集群中的每一个服务
```

## 智力题
### [有9个球,其中一个的质量与其他的不同,有一个天平,通过最多几次可以找出那个质量不一样的球](https://www.nowcoder.com/questionTerminal/2b1ac592b56745b2969672c1c4e0d2c5?source=relative)
