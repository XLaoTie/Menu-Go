package menu
import "fmt"
import "errors"
import "sync"
var once sync.Once
var List *CmdList
var MAX_CMD int = 10
func init()  {
	once.Do(
		func(){
			List = new(CmdList)
		})
	tmp1 := &CmdNode{
		cmd_s : "help",
		desc : "show all command and descirbe!",
		cmd_do : new(cmd_help),
		next : nil,
	}
	tmp2 := &CmdNode{
		cmd_s : "quit",
		desc : "quit the Menu!",
		cmd_do : new(cmd_quit),
		next : nil,
	}
	tmp3 := &CmdNode{
		cmd_s : "ls",
		desc : "show all file list",
		cmd_do : new(cmd_ls),
		next : nil,
	}
	List.Add_CmdNode(tmp1)
	List.Add_CmdNode(tmp2)
	List.Add_CmdNode(tmp3)
}

//指令列表
type CmdList struct{
	head *CmdNode
	tail *CmdNode
	cnt int 
	lock sync.Mutex
}
//节点列表相关
func (this *CmdList)Add_CmdNode(newC *CmdNode) error{
	this.lock.Lock()
	if this.cnt >= MAX_CMD {
		return errors.New("full cmd!")
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
	this.lock.Unlock()
	return nil
}
func (this *CmdList)Process(cmd string) error{
	p := this.head
	for p != nil {
		if p.cmd_s == cmd {
			p.Do()
			return nil;
		}
		p = p.next
	}
	return errors.New("error cmd")
}
func ExecuteMenu() {
	var cmd string
	for{
		fmt.Printf("%s :", "Input a cmd number >")
		fmt.Scanln(&cmd)
		err := List.Process(cmd)
		if err != nil {
			fmt.Println(err)
		}
	}
}






