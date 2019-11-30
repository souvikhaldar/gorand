# gorand
A simple golang package to generate random string 

## Installation 
`go get github.com/souvikhaldar/gorand`  

## Usage
First import it.  
```
import "github.com/souvikhaldar/gorand"
```

```
//random string on given set of characters "abcdef" of length l
randString := gorand.RandStrWithCharset(l,"abcdef")

//random string of given length l
randString := gorand.RandStr(l)
```
## Reference
https://www.calhoun.io/creating-random-strings-in-go/
