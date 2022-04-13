/*
* @Author: lishuang
* @Date:   2022/4/13 11:31
 */

package main

import "fmt"

/*
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。

示例 1：
输入：n = 3
输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

题解连接 : https://leetcode-cn.com/problems/unique-binary-search-trees-ii/solution/golang-shuang-bai-jie-fa-cong-di-gui-dao-cmcz/

这种题直接完蛋。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	dp := make([][][]*TreeNode, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]*TreeNode, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][i] = []*TreeNode{
			&TreeNode{i, nil, nil},
		}
	}

	for k := 2; k <= n; k++ {
		p := 1
		q := k
		for p <= n && q <= n {
			dp[p][q] = []*TreeNode{}
			for e := p; e <= q; e++ {
				fmt.Println(k, p, q, e)
				var leftTreeSet []*TreeNode
				if e-1 < 1 {
					leftTreeSet = []*TreeNode{}
				} else {
					leftTreeSet = dp[p][e-1]
				}
				var rightTreeSet []*TreeNode
				if e+1 > n {
					rightTreeSet = []*TreeNode{}
				} else {
					rightTreeSet = dp[e+1][q]
				}

				if leftTreeSet == nil || len(leftTreeSet) == 0 {
					for i := 0; i < len(rightTreeSet); i++ {
						root := &TreeNode{
							Val:   e,
							Left:  nil,
							Right: nil,
						}
						root.Right = rightTreeSet[i]
						dp[p][q] = append(dp[p][q], root)
					}
				} else if rightTreeSet == nil || len(rightTreeSet) == 0 {
					for i := 0; i < len(leftTreeSet); i++ {
						root := &TreeNode{
							Val:   e,
							Left:  nil,
							Right: nil,
						}
						root.Left = leftTreeSet[i]
						dp[p][q] = append(dp[p][q], root)
					}
				} else {
					for i := 0; i < len(leftTreeSet); i++ {
						for j := 0; j < len(rightTreeSet); j++ {
							root := &TreeNode{
								Val:   e,
								Left:  nil,
								Right: nil,
							}
							root.Left = leftTreeSet[i]
							root.Right = rightTreeSet[j]
							dp[p][q] = append(dp[p][q], root)
						}
					}
				}
			}
			p++
			q++
		}
	}

	return dp[1][n]
}
