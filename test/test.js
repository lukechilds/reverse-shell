import test from 'ava';
import reverseShell from 'this';

test('reverseShell is a function', t => {
	t.is(typeof reverseShell, 'function');
});

test('reverseShell returns shell code with /host:port variables', t => {
	const req = {
		url: '/foo:bar'
	};
	const returnValue = reverseShell(req);
	t.true(returnValue.indexOf('("foo",bar)') > -1);
});

test('reverseShell returns usage if host and port aren\'t set', t => {
	const req = {
		url: ''
	};
	const returnValue = reverseShell(req);
	t.true(returnValue.indexOf('# Reverse Shell as a Service') === 0);
});
