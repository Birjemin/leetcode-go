## 问题
You are playing the following Bulls and Cows game with your friend: You write down a number and ask your friend to guess what the number is. Each time your friend makes a guess, you provide a hint that indicates how many digits in said guess match your secret number exactly in both digit and position (called "bulls") and how many digits match the secret number but locate in the wrong position (called "cows"). Your friend will use successive guesses and hints to eventually derive the secret number.

Write a function to return a hint according to the secret number and friend's guess, use A to indicate the bulls and B to indicate the cows. 

Please note that both secret number and friend's guess may contain duplicate digits.

Example 1:
```
Input: secret = "1807", guess = "7810"

Output: "1A3B"

Explanation: 1 bull and 3 cows. The bull is 8, the cows are 0, 1 and 7.
```

Example 2:
```
Input: secret = "1123", guess = "0111"

Output: "1A1B"

Explanation: The 1st 1 in friend's guess is a bull, the 2nd or 3rd 1 is a cow.
```

Note: You may assume that the secret number and your friend's guess only contain digits, and their lengths are always equal.

## 分析
你正在和你的朋友玩 猜数字（Bulls and Cows）游戏：你写下一个数字让你的朋友猜。每次他猜测后，你给他一个提示，告诉他有多少位数字和确切位置都猜对了（称为“Bulls”, 公牛），有多少位数字猜对了但是位置不对（称为“Cows”, 奶牛）。你的朋友将会根据提示继续猜，直到猜出秘密数字


## 最高分
```golang
func getHint(secret string, guess string) string {
    var countS, countG [10]int
    bulls, cows := 0,0
    for i := range secret{
        ns := int(secret[i] - '0')
        ng := int(guess[i] - '0')
        if ng == ns {
            bulls++
            continue
        }
        
        if countG[ns] > 0 {
            cows++
            countG[ns]--
        } else {
            countS[ns]++
        }
        
        if countS[ng] > 0 {
            cows++
            countS[ng]--
        } else {
            countG[ng]++
        }
    }
    
    return fmt.Sprintf("%dA%dB", bulls, cows)
}
```

## 实现
常规方式map实现
```golang
func getHint(secret string, guess string) string {
	var count, ret int
	length := len(secret)
	tmp := make(map[int32]int, length)

	for i, v := range secret {
		if v == int32(guess[i]) {
			count++
		}
		if _, ok := tmp[v]; ok {
			tmp[v]++
		} else {
			tmp[v] = 1
		}
	}
	for _, v := range guess {
		if _, ok := tmp[v]; ok {
			if tmp[v] > 0 {
				ret++
			}
			tmp[v]--
		}
	}
	return fmt.Sprintf("%dA%dB", count, ret-count)
}
```

## 改进
数组实现
```golang
func getHint(secret string, guess string) string {
	var arr1 [10]int
	var count, ret int
	length := len(secret)
	
	for i := 0; i < length; i++ {
		s1, s2 := secret[i]-'0', guess[i]-'0'
		if s1 == s2 {
			count++
		} else {
			arr1[s1]++
			arr1[s2]--
		}
	}

	for _, v := range arr1 {
		if v > 0 {
			ret += v
		}
	}
	return fmt.Sprintf("%dA%dB", count, length-count-ret)
}

```

## 反思

## 参考