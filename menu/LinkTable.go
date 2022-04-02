package menu
import "errors"
import "sync"
var once sync.Once
var List *LinkTable
var MAX_CMD int = 10
func init()  {
	//单例模式
	once.Do(
		func(){
			List = new(LinkTable)	
		})
}
//指令列表
type LinkTable struct{
	head *LinkTableNode
	tail *LinkTableNode
	cnt int 
	
}
//删除所有命令销毁
//GO自动回收销毁
// func (this *LinkTable)Del_LinkTable(){
// 	if this.head == nil {
// 		return
// 	}
// 	for this.head != nil {
// 		// var p *LinkTableNode = this.head
// 		this.head = this.head.next
// 	}
// }
//节点列表相关
func (this *LinkTable)Add_LinkTableNode(newC *LinkTableNode) error{
	p := this.FindCmd(newC.cmd_s)
	if p != nil {
		return errors.New("Exit same command")
	}
	if this.cnt >= MAX_CMD {
		return errors.New("Full cmd!")
	}
	if this.head == nil {
		this.head = newC
		this.tail = newC
		this.cnt = 1
		return nil
	}
	this.cnt++;
	this.tail.next = newC
	this.tail = this.tail.next
	return nil
}
func (this *LinkTable)FindCmd(cmd string) *LinkTableNode {
	for p := this.head; p!= nil; p = p.next {
		if p.cmd_s == cmd {
			return p
		}
	}
	return nil
}
func (this *LinkTable)Del_LinkTableNode1(cmd string) error{
	if this.head == nil {
		return errors.New("Empty")
	}
	if this.head.cmd_s == cmd {
		this.head = this.head.next
		if this.head == nil {
			this.tail = this.head
		}
		return nil
	}
	pre := this.head
	for p:=pre.next; p != nil;{
		if p.cmd_s == cmd {
			pre.next = p.next
			if this.tail == p {
				this.tail = pre
			}
			return nil
		}
		pre = p
		p = p.next 
	}
	return errors.New("Doesn't exit")
}
// func (this *LinkTable)Del_LinkTableNode2(newC *LinkTableNode) {
	
// }
func (this *LinkTable)GetLinkTableHead() *LinkTableNode{
	if this.head == nil {
		return nil
	}
	return this.head
}

func (this *LinkTable)Process(cmd string) error{
	// p := this.head
	// for p != nil {
	// 	if p.cmd_s == cmd {
	// 		p.Do()
	// 		return nil;
	// 	}
	// 	p = p.next
	// }
	p := this.FindCmd(cmd)
	if p!= nil {
		p.Do()
		return nil
	}
	return errors.New("error cmd")
}







