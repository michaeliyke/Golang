module example.com/types

go 1.16

replace example.com/log => ../log

require (
	example.com/log v0.0.0-00010101000000-000000000000
	golang.org/x/tour v0.0.0-20210526031051-3891a3eb15c0
)
