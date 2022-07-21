package main

// nolint
var emptyDesks = desks{
	{available, available, available, available, available},
	{available, available, available, available, available},
	{available, available, available, available, available},
	{available, available, available, available, available},
	{available, available, available, available, available},
}

// nolint
var normalDesks = desks{
	{normal, available, annoying, dirty, available},
	{normal, available, available, dirty, normal},
	{normal, available, normal, normal, annoying},
	{annoying, dirty, available, available, normal},
	{normal, available, annoying, dirty, available},
}
