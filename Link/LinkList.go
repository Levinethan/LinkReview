package main

import (
	"fmt"
)

type Element int64
const  (ERROR = 10001)
//链表基本结构
type LinkNode struct {
	Data Element
	pNext *LinkNode
}


//对结构进行定义
type LinkNoder interface {
	Add (head *LinkNode,data Element,ishead bool)   //
	AddHead(head *LinkNode,data Element) //头插
	AddEnd (tail *LinkNode,data Element) //尾插
	Delete (head *LinkNode, index int)Element	//hard
	GetLength(head *LinkNode)int  //easy
	Search(head *LinkNode,data Element)bool   //easy
	Getdata(head *LinkNode, index int)Element   //easy
	Reverse()		//链表反转 hard
	Clear(phead *LinkNode)			//清空
 	Show(head *LinkNode)
	BubbleSort(head *LinkNode)   //冒泡排序   easy  三行代码搞定
	InserSort(head *LinkNode) *LinkNode		//插入排序hard     重要
	SelectSort(head *LinkNode)   //选择排序hard   性能最低  不太重要
	MergeSort(left *LinkNode , right *LinkNode) *LinkNode	//递归排序hard	重要
	QuickSort(head *LinkNode) *LinkNode		//快速排序hard	重要
}


//创建头部节点  不放任何数据
func NewLinkList() *LinkNode {
	return &LinkNode{Data:0,pNext:nil}
}
//链表有两个  风格  头部插入 尾部插入
func Add(head *LinkNode,data Element,ishead bool)  {
	if ishead {
		AddHead(head,data)
	}else {
		AddEnd(head,data)
	}
}

func AddHead(head *LinkNode,data Element)  {
	//首先新建一个节点
	node := &LinkNode{
		Data:  data,
		pNext: head.pNext,
	}
	head.pNext=node
}

func AddEnd(head *LinkNode,data Element)  {
	node := &LinkNode{
		Data:  data,
	}
	if head.pNext==nil{
		head.pNext=node
	}else {
		pcur := head
		for pcur.pNext!=nil{
			pcur=pcur.pNext
		}
		pcur.pNext=node  //循环到最后 尾部插入
	}
}

func Show(head *LinkNode)  {
	//一直循环到最后节点 nil
	pahead := head.pNext  //从头节点的下一个节点开始
	for pahead !=nil{
		fmt.Println(pahead.Data)
		pahead=pahead.pNext
	}
	fmt.Println("show over")
}

func GetLength(head *LinkNode)int  {
	pahead := head.pNext
	var length int =0
	for pahead !=nil{
		length++
		pahead=pahead.pNext
	}
	return length
}

func Search(head *LinkNode,data Element)bool  {
	pahead := head.pNext
	var isfind bool=false
	for pahead !=nil{
		if pahead.Data==data{
			isfind = true
			break
		}
		pahead= pahead.pNext

	}
	return isfind
}

func Getdata(head *LinkNode, index int)Element  {
	if index<=0 || index>GetLength(head){


		return  ERROR
	}else {
		pgo := head
		for i := 0;i<index;i++{
			pgo = pgo.pNext //一直循环
		}
		return pgo.Data
	}
}

func Delete (head *LinkNode, index int)Element {
	if index<=0 || index>GetLength(head){
		return  ERROR
	}else {
		pgo := head
		for i := 0;i<index;i++{
			pgo = pgo.pNext //一直循环
		}
		data := pgo.pNext.Data
		//pgo 是需要删除的位置
		pgo.pNext=pgo.pNext.pNext

		return data
	}
}

func Clear(phead *LinkNode)  {
	phead=nil
}

func Insert(head *LinkNode, index int,data Element)  {
	if index<=0 || index>GetLength(head){
		return
	}else {
		pgo := head
		for i := 0;i<index;i++{
			pgo = pgo.pNext //一直循环
		}
		var  node LinkNode //新节点
		node.Data=data   //插入操作
		node.pNext=pgo.pNext
		pgo.pNext=&node
		return
	}
}

func reverseNode(head *LinkNode) *LinkNode  {
	var preNode *LinkNode =nil //第一个节点
	var nextNode *LinkNode=nil //第二个节点
	for head!=nil{  //双指针循环
		nextNode = head.pNext //保存当前节点的下一个节点
		head.pNext=preNode //存储上一个节点
		preNode = head //更新前一个节点
		head = nextNode //更新
	}
	return preNode

}

func BubbleSort(head *LinkNode)  {
	for phead1 := head.pNext;phead1.pNext!=nil;phead1=phead1.pNext{
		for phead2 := head.pNext;phead2.pNext!=nil;phead2=phead2.pNext{
			if phead2.Data> phead2.pNext.Data{
				phead2.pNext.Data,phead2.Data= phead2.Data,phead2.pNext.Data
			}
		}
	}
	//下下策 两层for循环
}

//选择排序  我们需要用到前一个指针   在单链表里面为什么需要使用到双指针  为了方便删除
func SelectSort(head *LinkNode) *LinkNode {
	if head==nil|| head.pNext == nil{
		return head
	}else {
		//首先 我们需要双指针
		for newhead:=head.pNext;newhead != nil;newhead=newhead.pNext{
			pahead := newhead
			maxnode := newhead
			for pahead!=nil{
				if pahead.Data>maxnode.Data{
					maxnode = pahead
				}
				pahead=pahead.pNext
			}
			maxnode.Data,newhead.Data=newhead.Data,maxnode.Data
		}
		return head
	}
}

func InserSort(head *LinkNode) *LinkNode {
	if head == nil || head.pNext==nil{
		return head
	}else {
		var pstart  *LinkNode = new(LinkNode)
		var pend *LinkNode = head
		var p*LinkNode=head.pNext
		pstart.pNext=head //存储头部位置

		for p!=nil{
			var ptmp *LinkNode=pstart.pNext
			var pre *LinkNode=pstart //记录前一个位置 方便删除与插入
			for  ptmp!=p && p.Data>=ptmp.Data{ //找到要插入的位置
				ptmp = ptmp.pNext
				pre = pre.pNext
			}
			if ptmp == p{  //插入
				pend = p
			}else {
				pend.pNext = p.pNext
				p.pNext = ptmp
				pre.pNext = p
			}
			p=pend.pNext
		}
		head = pstart.pNext
		return head
	}
}

func Merge(left *LinkNode , right *LinkNode) *LinkNode  {
	//归并操作
	if left == nil{
		return right
	}
	if right==nil{
		return left
	}
	var  pres *LinkNode= new(LinkNode)  //新建链表
	var ptmp *LinkNode = pres   //开辟 并且 备份节点

	for left !=nil && right!=nil{
		if left.Data<right.Data{
			ptmp.pNext=left
			ptmp=ptmp.pNext
			left = left.pNext

		}else {
			ptmp.pNext=right
			ptmp=ptmp.pNext
			right = right.pNext
		}
	}

	if left != nil{
		ptmp.pNext=left
	}else {
		ptmp.pNext=right
	}

	ptmp=pres.pNext  //回到最开始的位置
	return ptmp

}

func MergeSort (head *LinkNode ) *LinkNode {
	if head == nil || head.pNext ==nil{
		return  head
	}else {
		//拆分
		var phead *LinkNode=head
		var qhead *LinkNode=head
		//双指针循环
		var pre *LinkNode=nil  //上一个指针
		//归并位置
		//

		for qhead!=nil && qhead.pNext!=nil{
			//归并排序  找到中间位置   切分步长  一个人走2步 一个人走一步
			//q 在这里走两步
			qhead = qhead.pNext.pNext
			pre=phead  //记录中间位置的上一步
			phead=phead.pNext  //中间  双指针错位
		}
		pre.pNext=nil
		var left *LinkNode = MergeSort(head)  //继续拆分左边
		var right *LinkNode = MergeSort(phead)  //继续拆分右边

		return  Merge(left,right)
	}
}

func  QuickSort(head *LinkNode) *LinkNode{
	if head == nil || head.pNext == nil{
		return head
	}else {
		var tmphead LinkNode = LinkNode{
			Data:  0,
			pNext: head,
		}

		qSortlist(&tmphead,head,nil)

		return  head
	}
}

func qSortlist(headpre,head,tail *LinkNode)  {
	if head != tail && head.pNext!=tail{
		var mid *LinkNode = partition(headpre,head,tail)
		qSortlist(headpre,headpre.pNext,mid)
		qSortlist(mid,mid.pNext,tail)
	}
}
//3    8 9 2 1
//3    8 2 9 1
//3    1 2  8 9
//指针交换

func partition(lowpre,low,high *LinkNode) *LinkNode  {
	key := low.Data //取得中间数据

	var node1 LinkNode = LinkNode{
		Data:  0,
		pNext: nil,
	}

	var node2 LinkNode = LinkNode{
		Data:  0,
		pNext: nil,
	}

	var small ,big *LinkNode=&node1,&node2   //用来做指针替换

	for i:= low.pNext;i!=high;i=i.pNext{
		if i.Data<key{
			small.pNext=i
			small=i

		}else {
			big.pNext=i
			big=i
		}
	}
	big.pNext=high
	small.pNext=low
	low.pNext=node2.pNext
	lowpre.pNext=node1.pNext
	return low


}

func main()  {
	var phead *LinkNode = NewLinkList()


	AddEnd(phead,11)
	AddEnd(phead,2)
	AddEnd(phead,32)
	AddEnd(phead,95)
	AddEnd(phead,25)
	AddEnd(phead,60)

	Show(phead)
	QuickSort(phead)
	Show(phead)

}
