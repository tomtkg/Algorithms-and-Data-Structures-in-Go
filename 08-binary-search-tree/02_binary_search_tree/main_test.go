package main

import "mylib"

func Example_main() {
	names := []string{"apple", "lemon", "cranberry", "raspberry"}
	mylib.SetStdin(names)
	main()
	// Output:
	// -- Generated binary search tree --
	// 			{apple 120}
	// 		{banana 30}
	// 	{blueberry 320}
	// 		{durian 980}
	// {kiwi 150}
	// 			{lemon 70}
	// 		{litchi 50}
	// 				{mango 180}
	// 			{melon 780}
	// 	{orange 80}
	// 		{papaya 200}
	// 				{pear 90}
	// 			{persimmon 100}
	// 					{pineapple 270}
	// 				{raspberry 350}
	//
	// apple: 120 yen
	// lemon: 70 yen
	// cranberry does not exists
	// raspberry: 350 yen
}
