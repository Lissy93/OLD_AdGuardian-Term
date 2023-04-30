package pains

import (
	"fmt"

	"github.com/gizak/termui/v3/widgets"
	"github.com/lissy93/adguardian-term/common"
	"github.com/lissy93/adguardian-term/fetch"
)

type NodeValue string

func (nv NodeValue) String() string {
	return string(nv)
}

func formatLabelValuePair(label string, value string) string {
	return fmt.Sprintf("[%s](fg:cyan,mod:bold): [%s](fg:white)", label, value)
}

func createTreeNodes(queryLog fetch.AdGuardQueryLog) []*widgets.TreeNode {
	nodes := make([]*widgets.TreeNode, 0, len(queryLog.Data))

	for _, entry := range queryLog.Data {
		question := entry.Question
		// question.FilteredBlackList

		var blockColor = "bg:green,fg:blue,mod:bold"
		if entry.Reason == "FilteredBlackList" {
			blockColor = "bg:red,fg:white,mod:bold"
		}

		queryStr := fmt.Sprintf("[IN](%s) [[%s]](fg:blue) %s", blockColor, question.Type, question.Name)

		dnsNode := &widgets.TreeNode{
			Value: NodeValue(queryStr),
			Nodes: []*widgets.TreeNode{
				{Value: NodeValue(formatLabelValuePair("Time", entry.Time))},
				{Value: NodeValue(formatLabelValuePair("Client", entry.Client))},
				{Value: NodeValue(formatLabelValuePair("Upstream DNS", entry.Upstream))},
				{Value: NodeValue(formatLabelValuePair("Time Taken", entry.ElapsedMs+" ms"))},
				{Value: NodeValue(formatLabelValuePair("Status", entry.Status))},
			},
		}

		for _, answer := range entry.Answer {
			answerNode := &widgets.TreeNode{
				Value: NodeValue(formatLabelValuePair("Server", answer.Value)),
			}
			dnsNode.Nodes = append(dnsNode.Nodes, answerNode)
		}
		nodes = append(nodes, dnsNode)
	}

	return nodes
}

func QueryTree(queryLog fetch.AdGuardQueryLog) *widgets.Tree {

	nodes := createTreeNodes(queryLog)

	l := widgets.NewTree()
	l.Title = "Query Log"
	common.SetCommonStyles(l)
	l.SetNodes(nodes)

	return l
}
