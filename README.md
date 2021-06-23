# 基于Golang实现平滑加权轮询算法、接口方式实现

### go get -u github.com/JoeyZeYi/weight


##### 代码有测试用例可供参考
##抽奖代码逻辑
####配置三个礼物
```cassandraql
ID  权重
1   5
2   3
3   2
```

#### 抽奖之前Training
```cassandraql
id   Weight   CurrentWeight
1      5           0
2      3           0
3      2           0
```
####轮询十次  抽奖总是取CurrentWeight最大的  每次抽取时CurrentWeight+=Weight
第一次抽  抽完后Training为  id为1
```cassandraql
id   Weight   CurrentWeight
1      5           -5
2      3           3
3      2           2
```

第二次抽 取CurrentWeight最高的一条 然后设置CurrentWeight为CurrentWeight-10+Weight  奖品id为2 抽完后Training为  
```cassandraql
id   Weight   CurrentWeight
1      5           0
2      3           -4    
3      2           4
```
第三次抽 奖品id为3 抽完后Training为 
```cassandraql
id   Weight   CurrentWeight
1      5           5
2      3           -1    
3      2           -4
```

第四次抽 奖品id为 1 抽完后Training为
```cassandraql
id   Weight   CurrentWeight
1      5           0
2      3           2    
3      2           -2
```
第五次抽 奖品id为 1 抽完后Training为
```cassandraql
id   Weight   CurrentWeight
1      5           -5
2      3           5    
3      2           0

```
第六次抽 奖品id为 2 抽完后Training为
```cassandraql
id   Weight   CurrentWeight
1      5           0
2      3           -2    
3      2           2

```
第七次抽 奖品id为 1 抽完后Training为
```cassandraql
id   Weight   CurrentWeight
1      5           -5
2      3           1    
3      2           4
```
第八次抽 奖品id为 3 抽完后Training为
```cassandraql
id   Weight   CurrentWeight
1      5           0
2      3           4    
3      2           -4
```
第九次抽 奖品id为 2 抽完后Training为
```cassandraql
id   Weight   CurrentWeight
1      5           5
2      3           -3    
3      2           -2
```
第10次抽 奖品id为 1 抽完后Training为
```cassandraql
id   Weight   CurrentWeight
1      5           0
2      3           0    
3      2           0
```
总计id为1的出了5次  id为2的出了3次  id为3的出了两次、十次循环结束后、回到了最原始的数据。
