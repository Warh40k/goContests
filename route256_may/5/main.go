package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//var m int
//var result []string

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)

	var result = make([]string, t)

	for m := 0; m < t; m++ {
		var n int
		fmt.Fscan(in, &n)
		input := bytes.Buffer{}

		for i := 0; i <= n; i++ {
			s, _ := in.ReadString('\n')
			input.WriteString(s)
		}
		var data interface{}
		json.Unmarshal(input.Bytes(), &data)
		res, _ := prettify(data)
		b, _ := json.Marshal(res)
		result[m] = string(b)
	}
	fmt.Fprintln(out, "["+strings.Join(result, ",")+"]")
}

func prettify(data interface{}) (interface{}, bool) {
	var empty bool
	switch v := data.(type) {
	case map[string]interface{}:
		for key := range v {
			v[key], empty = prettify(v[key])
			if empty {
				delete(v, key)
			}
		}
		if len(v) == 0 {
			return nil, true
		}
		return v, false
	case []interface{}:
		for i := 0; i < len(v); i++ {
			v[i], empty = prettify(v[i])
			if empty {
				v = append(v[:i], v[i+1:]...)
				i--
			}
		}
		if len(v) == 0 {
			return nil, true
		}
		return v, false
	default:
		return v, false
	}
}
