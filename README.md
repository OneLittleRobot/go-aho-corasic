# go-aho-corasic


A go version of [Aho Corasic](https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm), used for searching for phrases within text.

It's pretty fast, should be much faster than any built in string matching.

## How to use 
```go
import (
	"go-aho-corasic/aho"
)

func Search() {
    search := aho.NewSearch([]string{"something","another thing"})
    search.Build()
    result := search.Exec("'Lorem ipsum dolor something, consectetur adipiscing elit. In sem felis, tincidunt vitae orci et, ornare malesuada ante. Cras ultrices interdum leo id imperdiet. Lorem ipsum dolor sit amet, consectetur adipiscing elit.'")
    //[{index:18, text:'something'}]
}

```