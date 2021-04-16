module hisheng.com/hello

go 1.14

replace hisheng.com/greetings => ../greetings

require (
	han.com/age v0.0.0-00010101000000-000000000000
	hisheng.com/greetings v0.0.0-00010101000000-000000000000
)

replace han.com/age => ../age
