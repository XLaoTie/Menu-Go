package menu
//指令结点
type CmdNode struct{
	cmd_s string
	desc string
	cmd_do CmdDo
	next *CmdNode
}
//指令操作
type CmdDo interface{
	do()
}
//节点相关 
func(this *CmdNode)Do() {
	this.cmd_do.do()
}
func(this *CmdNode)CreateCmd(name string, desc string, cmd_do CmdDo) *CmdNode{
	return &CmdNode{
		cmd_s : name,
		desc : desc,
		cmd_do : cmd_do,
	}
}






