// A custom argument parser struct.
// 
// Any core arguments used to perform standard operations 
// are stored in the coreArgs slice.
// 
// Any extra flags used for finer control of these standard 
// operations are stored in the flags slice.
//
package main

import "strings"

type Parser struct {
	coreArgs	[]string
	flags		[]string
}
	
// parse the command line arguments
func (p *Parser) parse(args []string) {

	for _, arg := range(args) {
		if (strings.HasPrefix(arg, "-")) {
			p.flags = append(p.flags, arg)
		} else {
			p.coreArgs = append(p.coreArgs, arg)
		}
	}
}

// get the core argument at an index
func (p *Parser) get(index int) string {

	return p.coreArgs[index]
}

// get all core arguments after an index
func (p *Parser) getFrom(index int) []string {

	return p.coreArgs[index:]
}

// check if coreArgs contains a string
func (p *Parser) containsOne(arg0 string) bool {

	args := []string{ arg0 }

	return p.contains(args)
}

// check if coreArgs contains two strings in order
func (p *Parser) containsTwo(arg0 string, arg1 string) bool {

	args := []string{ arg0, arg1 }

	return p.contains(args)
}

// check if coreArgs contains three strings in order
func (p *Parser) containsThree(arg0 string, arg1 string, arg2 string) bool {

	args := []string{ arg0, arg1, arg2 }

	return p.contains(args)
}

// check if coreArgs contains a slice of strings in order
func (p *Parser) contains(args []string) bool {

	for index, x := range(args) {
		if p.coreArgs[index] != x {
			return false;
		} 
	}

	return true
}

// check if parser has a flag
func (p *Parser) hasFlag(flag string) bool {

	for _, f := range(p.flags) {
		if f == flag {
			return true;
		}
	}

	return false
}
