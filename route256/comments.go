package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Node struct {
	id     int
	parent int
	text   string
}

func buildBranchesForLine(branches []bool) string {
	var result []string
	for i := range branches {
		if branches[i] {
			result = append(result, "|")
		} else {
			result = append(result, " ")
		}
	}

	return strings.Join(result, "  ")
}

func formatSubTree(node Node, result *strings.Builder, depth, childCount int, branches []bool, nodeChilds map[int][]Node) {
	if depth != 1 {
		result.WriteString(buildBranchesForLine(branches))
		result.WriteString("--")
	}
	result.WriteString(node.text + "\n")
	nodes := nodeChilds[node.id]
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].id < nodes[j].id
	})
	branches = append(branches, true)
	for i := range nodes {
		if i == childCount-1 {
			branches[depth-2] = false
		}
		result.WriteString(buildBranchesForLine(branches) + "\n")
		formatSubTree(nodes[i], result, depth+1, len(nodes), branches, nodeChilds)
	}
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var l int
	fmt.Fscan(in, &l)

	for i := 0; i < l; i++ {
		var n int
		nodeChilds := map[int][]Node{}
		var heads []Node
		var formatted strings.Builder
		fmt.Fscan(in, &n)
		in.ReadString('\n')
		for j := 0; j < n; j++ {
			var node Node
			fmt.Fscanf(in, "%d %d", &node.id, &node.parent)
			node.text, _ = in.ReadString('\n')
			node.text = strings.TrimRight(node.text, "\n")
			node.text = strings.TrimLeft(node.text, " ")
			if node.parent == -1 {
				heads = append(heads, node)
			}
			if _, ok := nodeChilds[node.id]; ok == false {
				nodeChilds[node.id] = []Node{}
			}
			nodeChilds[node.parent] = append(nodeChilds[node.parent], node)
		}
		sort.Slice(heads, func(i, j int) bool {
			return heads[i].id < heads[j].id
		})

		for k := range heads {
			formatSubTree(heads[k], &formatted, 1, 0, []bool{}, nodeChilds)
			formatted.WriteString("\n")
		}
		fmt.Fprintln(out, formatted.String())
	}
}
