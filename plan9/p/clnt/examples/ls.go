package main

import "flag"
import "fmt"
import "log"
import "os"
import "plan9/p"
import "plan9/p/clnt"

var addr = flag.String("addr", "127.0.0.1:5640", "network address")

func main() {
	var user p.User;
	var err *p.Error;
	var c *clnt.Clnt;
	var file *clnt.File;
	var d []*p.Dir;

	flag.Parse();
	user = p.OsUsers.Uid2User(os.Geteuid());
	c, err = clnt.Mount("tcp", *addr, "", user);
	if err != nil {
		goto error
	}

	if flag.NArg() != 1 {
		log.Stderr("invalid arguments");
		return;
	}

	file, err = c.FOpen(flag.Arg(0), p.OREAD);
	if err != nil {
		goto error
	}

	for {
		d, err = file.Readdir(0);
		if err != nil {
			goto error
		}

		if d == nil || len(d) == 0 {
			break
		}

		for i := 0; i < len(d); i++ {
			os.Stdout.WriteString(d[i].Name + "\n")
		}
	}

	file.Close();
	return;

error:
	log.Stderr(fmt.Sprintf("Error: %s %d", err.Error, err.Errornum));
}