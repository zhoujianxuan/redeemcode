# redeemcode

[借鉴](https://blog.csdn.net/qq_61692791/article/details/131493798?spm=1001.2014.3001.5502)

做了一个go语言版本，生成的量级注意不要超过`uint32`的最大值

`num`可以采用redis的`incr`自增，不过这样生成的速度将受到制约

如果需要重兑校验可以
>  基于BitMap：兑换或没兑换就是两个状态，对应0和1，而兑换码使用的是自增id.我们如果每一个自增id对应一个bit位，用每一个bit位的状态表示兑换状态，是不是完美解决问题。而这种算法恰好就是BitMap的底层实现，而且Redis中的BitMap刚好能支持2^32个bit位。