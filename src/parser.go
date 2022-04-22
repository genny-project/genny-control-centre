package main

import "strings"

type Parser struct {
	coreArgs	[]string
	flags		[]string
}
	
func (p *Parser) parse(args []string) {

	for _, arg := range(args) {
		if (strings.Contains(arg, "-")) {
			p.flags = append(p.flags, arg)
		} else {
			p.coreArgs = append(p.coreArgs, arg)
		}
	}
}

func (p *Parser) get(index int) string {

	return p.coreArgs[index]
}

func (p *Parser) getFrom(index int) []string {

	return p.coreArgs[index:]
}

func (p *Parser) containsOne(arg0 string) bool {

	args := []string{ arg0 }

	return p.contains(args)
}

func (p *Parser) containsTwo(arg0 string, arg1 string) bool {

	args := []string{ arg0, arg1 }

	return p.contains(args)
}

func (p *Parser) containsThree(arg0 string, arg1 string, arg2 string) bool {

	args := []string{ arg0, arg1, arg2 }

	return p.contains(args)
}

func (p *Parser) contains(args []string) bool {

	for index, x := range(args) {
		if p.coreArgs[index] != x {
			return false;
		} 
	}

	return true
}

func (p *Parser) hasFlag(flag string) bool {

	for _, f := range(p.flags) {
		if f == flag {
			return true;
		}
	}

	return false
}
