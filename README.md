# go_bitWord
go位向量的实现.


搞明白位向量是个什么东西了.他就是 bitmap

之所以搞了一个数组是因为,只有64个数字是不够表示的.

比如010101 最低位为1表示1这个数字在bitmap中. 所以bitmap中元素是1,3,5

但是只用一个长度为64的unint来表示,显然不够因为才64个数.
所以建立一个数组来表示64多少倍的数.

比如64*100那么就能表示6400个数

比如我要表示128那么我就对应计算为 找bitmap里面第二个index所对应的word里面的最后1位即可.看他是不是1.他是1就表示128这个整数在我的set里面.














