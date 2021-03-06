package main

import (
	"testing"
	"fmt"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := BuildTree(preorder, inorder)
	fmt.Print(root)
	//InOrderPrint(root)
	//LevelOrderPrint(root)
}

func TestBuildTreeFromInAndPostOrder(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := BuildTree_2(inorder, postorder)
	fmt.Print(root)
	//InOrderPrint(root)
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := BuildLinkList([]int{2, 4, 8, 7, 4, 7})
	l2 := BuildLinkList([]int{5, 6, 4, 6})
	link := AddTwoNumbers(l1, l2)
	link.Print()
}

func TestAdd(t *testing.T) {
	array := []int{3, 22, 5, 45, 23, 8}

	head := BuildLinkList(array)

	link := head
	fmt.Println(link.Find(100))
	fmt.Println(link.FindIndex(3))
	//link.Print()

}

func TestRemoveNthFromEnd(t *testing.T) {
	array := []int{3, 22, 5, 45, 23, 8}

	head := BuildLinkList(array)

	link := RemoveNthFromEnd(head, 6)
	link.Print()
}

func TestRemoveNthFromEnd2(t *testing.T) {
	array := []int{1, 2}

	head := BuildLinkList(array)

	link := RemoveNthFromEnd2(head, 2)
	link.Print()
}

func TestMajorityElement(t *testing.T) {
	inorder := []int{1, 1, 2, 3, 1}
	fmt.Print(MajorityElement(inorder))
}

func TestLongestCommonStr(t *testing.T) {
	A := "134756778216"
	B := "357486782"
	fmt.Println(LongestCommonSubLen(A, B))
}

func TestClimbStairs(t *testing.T) {
	n := 4
	fmt.Println(ClimbStairs(n))
}

func TestLongestIncreaseList(t *testing.T) {
	array := []int{3, 22, 2, 45, 23, 8}
	fmt.Println(LongestIncreaseList(array))
}

func TestLongestConsecutive(t *testing.T) {
	array := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	fmt.Println(LongestConsecutive(array))
}

func TestTwoSum(t *testing.T) {
	array := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	fmt.Println(TwoSum(array, 6))
}
