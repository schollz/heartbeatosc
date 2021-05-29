
```
cd <thisfolder>
go build -v
./heartbeatosc
```

goto https://app.hyperate.io/B8963

put this in the console:

```javascript
var t=setInterval(function(){fetch('http://localhost:8098/'+$(".heartrate").innerHTML);},1000);
```
