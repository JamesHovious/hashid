# hashid [![GoDoc](https://godoc.org/github.com/JamesHovious/go-badge?status.svg)](https://godoc.org/github.com/JamesHovious/hashid)

Given a string determine the possible hashing algorithms used to produce that string. 

hashid is implemented as a Go library so that it can be easily incorporated into other projects. 

A CLI wrapper for this library is also available in the releases tab. 

## Command line usage:

```
Given a string determine the possible hashing algorithms used to produce that string. 
hashid can also give the corresponding hashcat mode and/or JohnTheRipper format in its output.

Note: When identifying a hash on *nix operating systems use single quotes to prevent interpolation.

Usage: hashid [options]... [hash]

Options:
  -e, --extended         list all hash algorithms including salted passwords
  -j, --john             show corresponding JohnTheRipper format in output
  -m, --mode             show corresponding hashcat mode in output
  -o, --outfile string   write output to file (default: STDOUT)
  -v, --version          show program's version number and exit

Sample:

$ ./hashid 8743b52063cd84097a65d1633f5c74f5
MD2
MD5
MD4
Double MD5
LM
RIPEMD-128
Haval-128
Tiger-128
Skein-256(128)
Skein-512(128)
Lotus Notes/Domino 5
Skype
Snefru-128
NTLM
Domain Cached Credentials
Domain Cached Credentials 2
DNSSEC(NSEC3)
RAdmin v2.x
```




### Credits:

This is a Go implementation of https://github.com/psypanda/hashID
