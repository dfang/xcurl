# README

make `curl -LO -# -w <file>.tar.gz https://github.com/<file>.tar.gz` faster or work in your server that don't have VPN or proxy support.


## install

```
curl -sf https://gobinaries.com/dfang/xcurl | sh
```


```
alias curl='xcurl'
```

## test 

`time xcurl -LO -# -w 'OpenJDK8U-jdk_x64_linux_hotspot_8u372b07.tar.gz\n' https://github.com/adoptium/temurin8-binaries/releases/download/jdk8u372-b07/OpenJDK8U-jdk_x64_linux_hotspot_8u372b07.tar.gz`

## thanks

https://ghproxy.com
