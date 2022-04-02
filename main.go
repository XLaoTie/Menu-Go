package main
/*********************************************************************************
  *Copyright(C),2022-2022,USTC
  *FileName:  menu.go
  *Author:  xutiezhong
  *Version:  1.0
  *Date:  2022.03.22
  *Description: 实现简单的类似Shell的菜单程序
  *Others:
  *Function List:  //主要函数列表，每条记录应包含函数名及功能简要说明
     1.func(LinkTableNode)Do() :指令结点调用指令操作函数
     2.func (*LinkTable)p_add(*LinkTableNode) : 向治理列表添加指令结点
	 3.func (*LinkTable)Process(string) : 从指令列表查找指令执行
	 4.func (cmd_help)do() : help显示指令操作
	 5.func (cmd_quit)do() : quit退出指令操作
	 6.func (cmd_ls)do() : ls显示文件列表指令操作 
  *History:  //修改历史记录列表，每条修改记录应包含修改日期、修改者及修改内容简介
     1.Date:2022.03.22
       Author:XTZ
       Modification:创建菜单程序的初始框架
	 2.Data:2022.03.26
	   Author:XTZ
	   Modification:将菜单程序打包成menu包，在包中分三模块：指令结点模块、指令操作模块、菜单模块
	   			菜单使用单例模式生成菜单列表，添加互斥锁，同步使用。
**********************************************************************************/
import . "test/menu_project/menu"
func main()  {
	ExecuteMenu()
}