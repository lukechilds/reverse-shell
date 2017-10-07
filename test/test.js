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
