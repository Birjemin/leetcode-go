## 问题
You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which will return whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.

Example:
```
Given n = 5, and version = 4 is the first bad version.

call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true

Then 4 is the first bad version. 
```

## 分析
你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。

假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。

你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。

提供了一个`isBadVersion`方法，写一个函数`firstBadVersion`方法，实现查找第一个错误版本，应该尽量减少对这个 `API` 的调用次数

## 最高分
```golang
func firstBadVersion(n int) int {
    l, r := 0, n
    for l < r{
        mid := l + (r-l)/2
        if isBadVersion(mid){
            r = mid
        } else{
            l = mid+1
        }
    }
    
    return l
}
```

## 实现
二分法这是最基础的方式
```golang
func firstBadVersion(n int) int {
	return cal(1, n)
}

func cal(min, max int) int {
	if min == max {
		return min
	}
	i := (min + max) / 2
	if isBadVersion(i) {
		return cal(min, i)
	} else {
		return cal(i+1, max)
	}
}
```

## 改进
```golang
func firstBadVersion(n int) int {

	for l := 1; l != n; {
		mid := (l + n) / 2
		if isBadVersion(mid) {
			n = mid
		} else {
			l = mid + 1
		}
	}
	return n
}
```

## 反思

## 参考