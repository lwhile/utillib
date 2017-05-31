# strs

- #### 通配符匹配
  
  "?" 匹配单个字符, "*" 可以匹配0个或多个字符

  Example:

        import "github.com/lwhile/utillib/strs"

        strs.WildcardMatch("abc", "*") // true
        strs.WildcardMatch("abc", "*?") // true
        strs.WildcardMatch("abc", "*d") // false
        strs.WildcardMatch("abc", "?") // false