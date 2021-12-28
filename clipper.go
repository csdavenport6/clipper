package main

import (
	"bufio"
	"flag"
	"os"
)

var url = flag.String("u", "", "url of thing to clip")
var name = flag.String("n", "", "name of clipping (both file and field in file)")
var author = flag.String("a", "", "author of thing (optional)")
var category = flag.String("c", "Article", "category of thing (default is Article)")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// TODO: add error handling for missing name

func main() {
	flag.Parse()
	homedir, err := os.UserHomeDir()
	check(err)
	homedir += "/secondbrain/3 Resources/Unsorted/"
	f, err := os.Create(homedir + *name + ".md")
	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString("## Name\n" + *name + "\n")
	check(err)
	_, err = w.WriteString("## Author\n" + *author + "\n")
	check(err)
	_, err = w.WriteString("## Category\n" + *category + "\n")
	check(err)
	_, err = w.WriteString("## URL\n" + *url + "\n")
	check(err)
	_, err = w.WriteString("## Tags\n#unread\n")
	check(err)
	_, err = w.WriteString("## Ratings/Comments\n\n")
	check(err)

	w.Flush()
}
