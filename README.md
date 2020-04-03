# LinkReview
  ## 单链表结构 
  ### 添加
  ### 头部插入
  ### 尾部插入
  ### 删除
  ### 长度
  ### 查找
  ### 获取数据
  ### 链表反转
  ### 清空
  ### 冒泡排序
  ### 插入排序
  ### 选择排序
  ### 递归排序
  ### 快速排序
`Add (head *LinkNode,data Element,ishead bool)   //
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
 `
