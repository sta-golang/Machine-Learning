## 朴素贝叶斯
![](https://img.shields.io/badge/author-TheSevenSky-blue) ![](https://img.shields.io/badge/build-passing-yellow) ![](https://img.shields.io/badge/Release-Development-red)

### 贝叶斯算法
贝叶斯算法是一类算法的总称，这些算法均以贝叶斯定理为基础

### 贝叶斯定理

公式如下<br/>

在事件B已经发生的情况下，事件A发生的概率

![](https://img-blog.csdn.net/20180504155102233?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ppY2h1bnc=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

同理可得在事件A发生的情况下 事件B发生的概率

![](https://img-blog.csdn.net/20180504155218557?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ppY2h1bnc=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

那么很容易推导得

![](https://img-blog.csdn.net/20180504155332794?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ppY2h1bnc=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

### 朴素贝叶斯

特征独立性假设是指特征之间是没有联系的
具体的知识点自行百度


### 举例: 基于朴素贝叶斯算法进行账户异常检测

在电商网站中，往往会存在一些异常用户，包括恶意刷单用户。
爬虫爬取数据用户等。这些用户的数据就毫无意义需要屏蔽，

#### 数据准备

具体可以查看 data.go

| 注册天数 | 活跃天数 | 购物次数 | 点击商品个数 | 是否为异常用户 |
 | :------: | :------: | :------: | :------: | :------: | 
 | 320 | 204 | 198 | 265 | 是 |
 | 253 | 53 | 15 | 2243 | 否 |
 | 53 | 32 | 5 | 325 | 否 |
 | 63 | 50 | 42 | 98 | 是 |
 | 1302 | 523 | 202 | 5430 | 否 |
 | 32 | 22 | 5 | 143 | 否 |
 | 105 | 85 | 70 | 322 | 是 |
 | 872 | 730 | 840 | 2762 | 是 |
 | 16 | 15 | 13 | 52 | 是 |
 | 92 | 70 | 21 | 693 | 否 |
 
 然后预测一个用户 [134,84,235,349] 是否为异常用户
 
 该用户注册了134天 活跃 84天 总共下了235个订单 一个用户下了这么多订单 按道理来说应该会浏览上千或者上万的商品。所以这应该是一个异常用户