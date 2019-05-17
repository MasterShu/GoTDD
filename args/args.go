package args

import (
	"strconv"
	"strings"
)

type Args struct {
	l bool
	p int
	d string
}

func ArgsParse(argument string) Args {
	var args Args
	getNumber(argument, &args)

	getBoolean(argument, &args)

	getString(argument, &args)
	return args
}

func getString(argument string, args *Args) {
	if strings.Index(argument, "-d") > -1 {
		temp := argument[3:]
		args.d = temp
	}
}

func getBoolean(argument string, args *Args) {
	if strings.Index(argument, "-l") > -1 {
		args.l = true
	}
}

func getNumber(argument string, args *Args) {
	if strings.Index(argument, "-p ") > -1 {
		temp := argument[3:]
		args.p, _ = strconv.Atoi(temp)
	}
}
