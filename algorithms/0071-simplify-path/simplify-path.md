## 问题
Given an absolute path for a file (Unix-style), simplify it. Or in other words, convert it to the canonical path.

In a UNIX-style file system, a period . refers to the current directory. Furthermore, a double period .. moves the directory up a level. For more information, see: Absolute path vs relative path in Linux/Unix

Note that the returned canonical path must always begin with a slash /, and there must be only a single slash / between two directory names. The last directory name (if it exists) must not end with a trailing /. Also, the canonical path must be the shortest string representing the absolute path.

Example 1:
```
Input: "/home/"
Output: "/home"
Explanation: Note that there is no trailing slash after the last directory name.
```

Example 2:
```
Input: "/../"
Output: "/"
Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
```

Example 3:
```
Input: "/home//foo/"
Output: "/home/foo"
Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
```

Example 4:
```
Input: "/a/./b/../../c/"
Output: "/c"
```

Example 5:
```
Input: "/a/../../b/../c//.//"
Output: "/c"
```

Example 6:
```
Input: "/a//b////c/d//././/.."
Output: "/a/b/c"
```

## 分析
标准化Unix路径

## 最高分
类似的解法，先分割，然后入栈和出栈
```golang
func simplifyPath(path string) string {
	absDirs := []string{}
	dirs := strings.Split(path, "/")
	for _, d := range dirs {
		if d == "" {
			continue
		}
		if d == ".." {
			if len(absDirs) > 0 {
				absDirs = absDirs[:len(absDirs)-1]
			}
		} else if d == "." {

		} else {
			absDirs = append(absDirs, d)
		}
	}
	return "/" + strings.Join(absDirs, "/")
}
```

## 实现
考察栈的使用
```golang
func simplifyPath(path string) string {
    tmp, stack, length := strings.Split(path, "/"), make([]string, 0), 0
    for _, v := range tmp {
        v = strings.TrimSpace(v)
        if v == "" || v == "." {
            continue
        }
        switch v {
        case "..":
            if length != 0 {
                stack, length = stack[:length-1], length-1
            }
        default:
            stack, length = append(stack, v), length+1
        }
    }
    return "/" + strings.Join(stack, "/")
}
```

## 改进
```golang

```

## 反思

## 参考