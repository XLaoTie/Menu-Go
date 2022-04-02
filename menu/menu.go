package menu
import "fmt"
func init() {
	go MenuConfig("help", "show all command and descirbe!", new(Cmd_help))
	go MenuConfig("quit", "quit the Menu!", new(Cmd_quit))
	go MenuConfig("ls", "show all file list", new(Cmd_ls))
}
func MenuConfig(cmd_s string, desc string, cmd_do CmdDo) {
	p := CreateCmd(cmd_s, desc, cmd_do)
	err := List.Add_LinkTableNode(p)
	if err != nil {
		fmt.Println(err)
	}
}
func ExecuteMenu() {
	var cmd string
	//go List.Del_LinkTableNode1("ls")
	for{
		fmt.Printf("%s :", "Input a cmd number >")
		fmt.Scanln(&cmd)
		err := List.Process(cmd)
		if err != nil {
			fmt.Println(err)
		}
	}
}