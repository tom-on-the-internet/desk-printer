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
	{normal, available, annoying, sick, available},
	{normal, available, available, sick, normal},
	{normal, available, normal, normal, annoying},
	{annoying, sick, available, available, normal},
	{normal, available, annoying, sick, available},
}
