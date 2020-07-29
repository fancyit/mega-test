This is a test web application.
it has 2 methods:
	1. http://yourserverurl:8081/get-phrase-hash
	   it takes a single argument which will be treated as a string, then it returns a JSON with folowing structure
	   {"phrase":,"HashedPhrase":}, where phrase is the object you passed to the method and the HashedPhrase is sha256 
	   from the "phrase" trimed and converted to int64 type
	2. http://yourserverurl:8081/get-detailed-hash
	   Same as 1. but response JSON includes additional fields, such as "HexedPhrase":"","HexedCuttedPhrase":""
	   where HexedPhrase is sha256 sum in hex format, and HexedCuttedPhrase is trimmed HexedPhrase to fit int64
		
