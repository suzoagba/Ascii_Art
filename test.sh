go run . "hello" standard.txt | cat -e
go run . "HELLO" standard.txt | cat -e
go run . "HeLlo HuMaN" standard.txt | cat -e
go run . "1Hello 2There" standard.txt | cat -e
go run . "Hello\nThere" standard.txt | cat -e
go run . "Hello\n\nThere" standard.txt | cat -e
go run . "{Hello & There #}" standard.txt | cat -e
go run . 'hello There 1 to 2!' standard.txt | cat -e
go run . "MaD3IrA&LiSboN" standard.txt | cat -e
go run . "1a\"#FdwHywR&/()=" standard.txt | cat -e
go run . "{|}~" standard.txt | cat -e
go run . "[\]^_ 'a" standard.txt | cat -e
go run . "RGB" standard.txt | cat -e
go run . ":;<=>?@" standard.txt | cat -e
go run . '\!" #$%&'"'"'()*+,-./' standard.txt | cat -e
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" standard.txt | cat -e
go run . "abcdefghijklmnopqrstuvwxyz" standard.txt | cat -e