package main

import cmd "tracker/cmd/web"

func init() {
	cmd.AllBands()
	cmd.AllRelation()
}

func main() {
	cmd.Handler()
}
