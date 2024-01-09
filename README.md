# GolangPattern
 
1. SINGLETON
```
    var once sync.Once
    once.Do
```
```
    func init() {
	    instance = &singleton{count: 100}
    }
```