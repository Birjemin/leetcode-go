package simplify_path

import (
    "strings"
)

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
