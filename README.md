# hashid

Software to identify the different types of hashes
Go implimentation of https://github.com/psypanda/hashID

Identify the different types of hashes used to encrypt data and especially passwords.
hashid is a tool written in Go which supports the identification of over 220 unique hash types using regular expressions.
hashid is written as a Go library so that it can easily be inlcuded in other tools. 
It is able to identify a single hash. hashid is also capable of including the corresponding hashcat mode and/or JohnTheRipper format in its output.
Note: When identifying a hash on *nix operating systems use single quotes to prevent interpolation.