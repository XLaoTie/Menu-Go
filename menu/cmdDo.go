package menu
import "fmt"

//指令操作相关
type Cmd_help struct {
}
func (this Cmd_help)do() {
	fmt.Println("Commands in Menu:")
	p := List.head
	for p != nil {
		fmt.Printf("%s : %s \n", p.cmd_s, p.desc)
		p = p.next
	}
}
type Cmd_quit struct {
}
func (this Cmd_quit)do() {
	
	panic("Quit！")
}
type Cmd_ls struct {
}
func (this Cmd_ls)do() {
	fmt.Println("file1, file2, file3.....")
}
