package main

import (
  "fmt"
  "flag"
  "os"
  "bufio"
  "os/exec"
)

type selpg_args struct{
  start_pg int
  end_pg int
  page_len int
  page_type bool
  destination string
  srcName string
}

var progname string

func recieve_args() *selpg_args {
  args := new(selpg_args)
  flag.IntVar(&args.start_pg, "s", -1, "(mandatory)the starting page number")
  flag.IntVar(&args.end_pg, "e", -1, "(mandatory)the ending page number")
  flag.IntVar(&args.page_len, "l", 72, "the number of lines in a page, default is 72")
  flag.BoolVar(&args.page_type, "f", false, "whether use /f to seperate page")
  flag.StringVar(&args.destination, "d", "", "the destination to recieve output, default to stdout")
  flag.Parse()
  flag.Usage = usage
  if flag.NArg() > 0 {
    args.srcName = flag.Args()[0]
  }
  return args
}

func process_args(args *selpg_args) {
  if len(os.Args) < 3 {
    fmt.Fprintf(os.Stderr, "%s: not enough arguments\n", progname)
    flag.Usage()
    os.Exit(1)
  }
  if args.start_pg < 1 {
    fmt.Fprintf(os.Stderr, "%s: invalid start page %d", progname, args.start_pg)
    flag.Usage()
    os.Exit(2)
  }
  if args.end_pg < 1 || args.end_pg < args.start_pg {
    fmt.Fprintf(os.Stderr, "%s: invalid end page %d", progname, args.end_pg)
    flag.Usage()
    os.Exit(3)
  }
  if args.page_len < 1 {
    fmt.Fprintf(os.Stderr, "%s: invalid page length %d", progname, args.page_len)
    flag.Usage()
    os.Exit(4)
  }
  if args.page_type == true && args.page_len != 72 {
    fmt.Fprintf(os.Stderr, "%s: -f and -l can not be assigned at the same time")
    flag.Usage()
    os.Exit(5)
  }
}

func usage() {
  fmt.Fprintf(os.Stderr, "selpg Usage: selpg -s number -e number [-l number] [-f] [-d destination] [filename]\n")
  fmt.Printf("Options:\n")
  flag.PrintDefaults()
}

func process_input(args *selpg_args) {
  //set input source
  fin := os.Stdin
  var err error
  if args.srcName != "" {
    fin, err = os.Open(args.srcName)
    if err != nil {
      panic(err)
    }
  }
  cmd := exec.Command(args.destination)
  stdin, err := cmd.StdinPipe()
  if err != nil {
    panic(err)
  }
  cmd.Stdout = os.Stdout
  cur_page := 1
  //fix-length page method
  if !args.page_type {
    cur_line := 0
    bs := bufio.NewScanner(fin)
    for {
      if bs.Scan() {
        cur_line++
        if cur_line > args.page_len {
          cur_page++
          cur_line = 1
        }
        if cur_page >= args.start_pg && cur_page <= args.end_pg {
          if args.destination != "" {
            stdin.Write([]byte(bs.Text() + "\n"))
          } else {
            os.Stdout.Write([]byte(bs.Text() + "\n"))
          }
        }
      } else {
        break
      }
    }
  } else {
    //dynamic-length page method
    bs := bufio.NewScanner(fin)
    bs.Split(bufio.ScanBytes)
    for {
      if bs.Scan() {
        if bs.Text() == "\f" {
          cur_page++
        }
        if cur_page >= args.start_pg && cur_page <= args.end_pg {
          if args.destination != "" {
            stdin.Write([]byte(bs.Text()))
          } else {
            os.Stdout.Write([]byte(bs.Text()))
          }
        }
      } else {
        break
      }
    }
  }
  if args.destination != "" {
    err = cmd.Run()
    if err != nil{
      panic(err)
    }
  }

  if cur_page < args.start_pg {
    fmt.Fprintf(os.Stderr, "Start page greater than total page, no output")
  } else {
    if cur_page < args.end_pg {
      fmt.Fprintf(os.Stderr, "total page smaller than end page")
    }
  }

  fin.Close()
  stdin.Close()
  fmt.Fprintf(os.Stderr, "Done")

}

func main() {
  progname = os.Args[0];
  arguments := recieve_args()
  process_args(arguments)
  //fmt.Printf("start:%d end:%d length:%d type:%s dest:%s inname:%s\n", arguments.start_pg, arguments.end_pg, arguments.page_len, arguments.page_type, arguments.destination, arguments.srcName)
  process_input(arguments)
}
