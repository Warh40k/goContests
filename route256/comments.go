package main

import (
	"bufio"
	"fmt"
	"os"
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

func formatSubTree(node Node, result *strings.Builder, depth int, delBranch bool, branches []bool, nodeChilds map[int][]Node) {
	if depth != 0 {
		result.WriteString(buildBranchesForLine(branches))
		result.WriteString("--")
	}
	result.WriteString(node.text + "\n")
	nodes := nodeChilds[node.id]
	//sort.SliceStable(nodes, func(i, j int) bool {
	//	return nodes[i].id < nodes[j].id
	//})
	nodes = bubbleSort(nodes, len(nodes))
	if delBranch {
		branches[depth-1] = false
		delBranch = false
	}

	branches = append(branches, true)

	for i := range nodes {
		if i == len(nodes)-1 {
			delBranch = true
		}
		result.WriteString(buildBranchesForLine(branches) + "\n")
		formatSubTree(nodes[i], result, depth+1, delBranch, branches, nodeChilds)
	}
}

func bubbleSort(arr []Node, size int) []Node {
	if size <= 1 {
		return arr
	}

	var i = 0
	for i < size-1 {
		if arr[i].id > arr[i+1].id {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}

		i++
	}

	bubbleSort(arr, size-1)

	return arr
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	sc := bufio.NewScanner(os.Stdin)
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
			//fmt.Fscanf(in, "%d %d ", &node.id, &node.parent)
			sc.Scan()
			parts := strings.SplitN(sc.Text(), " ", 3)
			fmt.Sscan(parts[0], &node.id)
			fmt.Sscan(parts[1], &node.parent)
			node.text = parts[2]
			if node.parent == -1 {
				heads = append(heads, node)
			}
			if _, ok := nodeChilds[node.id]; ok == false {
				nodeChilds[node.id] = []Node{}
			}
			nodeChilds[node.parent] = append(nodeChilds[node.parent], node)
		}
		heads = bubbleSort(heads, len(heads))

		for k := range heads {
			formatSubTree(heads[k], &formatted, 0, false, []bool{}, nodeChilds)
			formatted.WriteString("\n")
		}
		fmt.Fprint(out, formatted.String())
	}
}
