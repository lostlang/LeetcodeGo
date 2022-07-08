package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var outList *ListNode

	if list2 == nil && list1 == nil {
		return outList
	}

	if list1 == nil {
		outList = &ListNode{list2.Val, nil}
		list2 = list2.Next
	} else if list2 == nil {
		outList = &ListNode{list1.Val, nil}
		list1 = list1.Next
	} else if list1.Val < list2.Val {
		outList = &ListNode{list1.Val, nil}
		list1 = list1.Next
	} else {
		outList = &ListNode{list2.Val, nil}
		list2 = list2.Next
	}

	var curList = outList

	for list1 != nil || list2 != nil {
		if list1 == nil {
			curList.Next = &ListNode{list2.Val, nil}
			curList = curList.Next
			list2 = list2.Next
		} else if list2 == nil {
			curList.Next = &ListNode{list1.Val, nil}
			curList = curList.Next
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			curList.Next = &ListNode{list1.Val, nil}
			curList = curList.Next
			list1 = list1.Next
		} else {
			curList.Next = &ListNode{list2.Val, nil}
			curList = curList.Next
			list2 = list2.Next
		}
	}

	return outList
}
