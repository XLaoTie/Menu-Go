package menu
import "fmt"

//指令操作相关
type cmd_help struct {
}
func (this cmd_help)do() {
	fmt.Println("Commands in Menu:")
	p := List.head
	for p != nil {
		fmt.Printf("%s : %s \n", p.cmd_s, p.desc)
		p = p.next
	}
}
type cmd_quit struct {
}
func (this cmd_quit)do() {
	
	panic("Quit！")
}
type cmd_ls struct {
}
func (this cmd_ls)do() {
	fmt.Println("file1, file2, file3.....")
}
