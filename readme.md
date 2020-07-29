## This is a test web application.

###### It has 2 methods:
	1. _http://yourserverurl:8081/get-phrase-hash_
	   it takes a single argument which will be treated as a string, then it returns a JSON with folowing structure
	   {"phrase":,"HashedPhrase":}, where phrase is the object you passed to the method and the HashedPhrase is sha256 
	   from the "phrase" trimed and converted to int64 type
	2. _http://yourserverurl:8081/get-detailed-hash_
	   Same as 1. but response JSON includes additional fields, such as **"HexedPhrase":"","HexedCuttedPhrase":""**
	   where HexedPhrase is sha256 sum in hex format, and HexedCuttedPhrase is trimmed HexedPhrase to fit int64
		

###### The repo has Dockerfile
which you can use to setup docker image with the app onboard.
For docker build command to be correctly fulfilled it is mandatory to executed with _--network=host_ option, 
since apk add git won't pass without it.
The reason is repo accessibilitie, so you may run into an error with text:
```bash 
WARNING: Ignoring http://dl-cdn.alpinelinux.org/alpine/v3.12/main/x86_64/APKINDEX.tar.gz: temporary error (try again later)
```
So the command to build image is:
```bash
 docker build . -t mega-test:v3 --network=host
 ```
