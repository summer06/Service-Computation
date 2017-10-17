# Go Program - Selpg
---
##### Requirment: Implement a CLI program named Selpg
##### for more detail: https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html

##### Introduction: selpg is a useful cli program that can select pages from a big file. It is useful when we want to print something. For example, if we want to print a specific chapter of a book, we can use this program to select the page we want and print them.

##### Develop Tool: Atom 1.21.0

----

### Data Structure
```
//declare a struct named selpg_args
type selpg_args struct{
  start_pg int  // number of start page
  end_pg int  //number of end page
  page_len int  //number of lines in a page
  page_type bool  //whether the page use '/f' as page seperator
  destination string  //the destination to print
  srcName string  //source file name
}
```
All the variables we need to get from the command line are capsuled in this struct named selpg_args.

### Function Declaration
```
//get arguments from command line
func recieve_args() *selpg_args
//check weather the arguments is valid
unc process_args(args *selpg_args)
//process input
func process_input(args *selpg_args)
```

----
### Implement Detail
#### Use Flag

Command-line flags are a common way to specify options for command-line programs. For example, in wc -l the -l is a command-line flag.

Go provides a flag package supporting basic command-line flag parsing. I use this package to implement my command-line program.
```
//bind the variables with the command args
flag.IntVar(&args.start_pg, "s", -1, "(mandatory)the starting page number")
flag.IntVar(&args.end_pg, "e", -1, "(mandatory)the ending page number")
flag.IntVar(&args.page_len, "l", 72, "the number of lines in a page, default is 72")
flag.BoolVar(&args.page_type, "f", false, "whether use /f to seperate page")
flag.StringVar(&args.destination, "d", "", "the destination to recieve output, default to stdout")
flag.Parse()
```
use flag.IntVar, flag.BoolVar, flag.StringVar to bind the command-ling flag with the variables we declare. It is convenient because we don't need to analysis the command-line flag by ourselves.

**remeber: after define all the flags, use flag.Parse() to make it work!**

know more clearly about flag, you can access https://gobyexample.com/command-line-flags

#### Use Exec
Package exec runs external commands. It wraps os.StartProcess to make it easier to remap stdin and stdout, connect I/O with pipes, and do other adjustments.

Given that I don't have a printer,so I implement it as 'cat' command.

```
cmd := exec.Command("cat", "-n")
stdin, err := cmd.StdinPipe()
if err != nil {
  panic(err)
}
cmd.Stdout = os.Stdout
```
know more about Package exec, go to
https://golang.org/pkg/os/exec/

---
### Test Result
**The input file is from https://github.com/smallGum/service-computing/tree/master/selpg**

1.select specific pages

```
$ ./selpg -s 1 -e 1 input_file
No. 1 line of No.1 page
No. 2 line of No.1 page
No. 3 line of No.1 page
No. 4 line of No.1 page
No. 5 line of No.1 page
No. 6 line of No.1 page
No. 7 line of No.1 page
No. 8 line of No.1 page
No. 9 line of No.1 page
...
No. 70 line of No.1 page
No. 71 line of No.1 page
No. 72 line of No.1 page

```

2.redirect input from file
```
$ ./selpg -s 1 -e 1 < input_file
No. 1 line of No.1 page
No. 2 line of No.1 page
No. 3 line of No.1 page
No. 4 line of No.1 page
No. 5 line of No.1 page
No. 6 line of No.1 page
No. 7 line of No.1 page
No. 8 line of No.1 page
No. 9 line of No.1 page
...
No. 70 line of No.1 page
No. 71 line of No.1 page
No. 72 line of No.1 page
```

3.other command's output as selpg's input
```
$ cat input_file | ./selpg -s 10 -e 20

No. 1 line of No.10 page
No. 2 line of No.10 page
No. 3 line of No.10 page
No. 4 line of No.10 page
No. 5 line of No.10 page
No. 6 line of No.10 page
...
No. 50 line of No.16 page
No. 51 line of No.16 page
No. 52 line of No.16 page
...
No. 70 line of No.20 page
No. 71 line of No.20 page
No. 72 line of No.20 page
```

4.result redirect to a file
```
$ ./selpg -s 10 -e 20 input_file >output_file
$
```
you can refer to the output_file to see the result

5.use the output as other command's input
```
Donedengxiajuns-MacBook-Pro:selpg dengxiajun$ ./selpg -s 10 -e 20 input_file | cat -n
     1
        No. 1 line of No.10 page
     2	No. 2 line of No.10 page
     3	No. 3 line of No.10 page
     4	No. 4 line of No.10 page
     5	No. 5 line of No.10 page
     ...
   427	No. 67 line of No.15 page
   428	No. 68 line of No.15 page
   429	No. 69 line of No.15 page
   430	No. 70 line of No.15 page
   ...
   790	No. 70 line of No.20 page
   791	No. 71 line of No.20 page
   792	No. 72 line of No.20 page
```
in command 'cat', flag '-n' can show the number of each line, thus the result is output by 'cat' command

6.use pipe to command ‘cat’
```
$ ./selpg -s 1 -e 1 -d input_file input_file
     1	No. 1 line of No.1 page
     2	No. 2 line of No.1 page
     3	No. 3 line of No.1 page
     4	No. 4 line of No.1 page
     5	No. 5 line of No.1 page
     6	No. 6 line of No.1 page
     7	No. 7 line of No.1 page
     ...
    70	No. 70 line of No.1 page
    71	No. 71 line of No.1 page
    72	No. 72 line of No.1 page
```
because the output result have line number, so it is output by 'cat' command.

---
### Other Useful links
- https://studygolang.com/articles/4367  --use of bufio package
- https://studygolang.com/articles/5024  --use of os package
