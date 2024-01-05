# context
context is a cli utility to add options to cli commands based on the context (dir, env) in which they are run

For example
create a context configuration file in the current directory
    $ echo .context << __END__
	[echo]
	$@ = $@ world
	__END__

use context to run the echo command
    $ context echo hello
	hello world

