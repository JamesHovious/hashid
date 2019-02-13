# hashid [![GoDoc](https://godoc.org/github.com/JamesHovious/go-badge?status.svg)](https://godoc.org/github.com/JamesHovious/hashid)

Given a string determine the possible hashing algorithms used to produce that string. 

hashid is impliemnted as a Go library so that it can be easily incorporated into other projects. 

A CLI wrapper for this library is also available in the releases tab. 

## Command line usage:

```
Given a string determine the possible hashing algorithms used to produce that string. 
hashid can also give the corresponding hashcat mode and/or JohnTheRipper format in its output.

Note: When identifying a hash on *nix operating systems use single quotes to prevent interpolation.

Usage:
  -e, --extended         list all hash algorithms including salted passwords
  -j, --john             show corresponding JohnTheRipper format in output
  -m, --mode             show corresponding hashcat mode in output
  -o, --outfile string   write output to file (default: STDOUT)
  -v, --version          show program's version number and exit
```




### Credits:

This is a Go implimentation of https://github.com/psypanda/hashID
