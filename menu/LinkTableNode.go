package menu
//指令结点
type LinkTableNode struct{
	cmd_s string
	desc string
	cmd_do CmdDo
	next *LinkTableNode
}
//指令操作
type CmdDo interface{
	do()
}
//节点相关 
func(this *LinkTableNode)Do() {
	this.cmd_do.do()
}
func CreateCmd(name string, desc string, cmd_do CmdDo) *LinkTableNode{
	return &LinkTableNode{
		cmd_s : name,
		desc : desc,
		cmd_do : cmd_do,
		next : nil,
	}
}






