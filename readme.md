This is a test web application.
it has 2 methods:
	1. http://yourserverurl:8081/get-phrase-hash
	   it takes a single argument which will be treated as a string, then it returns a JSON with folowing structure
	   {"phrase":,"HashedPhrase":}, where phrase is the object you passed to the method and the HashedPhrase is sha256 
	   from the "phrase" trimed and converted to int64 type
	2. http://yourserverurl:8081/get-detailed-hash
	   Same as 1. but response JSON includes additional fields, such as "HexedPhrase":"","HexedCuttedPhrase":""
	   where HexedPhrase is sha256 sum in hex format, and HexedCuttedPhrase is trimmed HexedPhrase to fit int64
		

The repo has Dockerfile which you can you to setap docker image with this app onboard,
to build pass correctly it is mandatory to run docker build with --network=host option, since apk add git won't pass without this option, caused by issue 
with repo, and you'll run into an error with text:
WARNING: Ignoring http://dl-cdn.alpinelinux.org/alpine/v3.12/main/x86_64/APKINDEX.tar.gz: temporary error (try again later)
