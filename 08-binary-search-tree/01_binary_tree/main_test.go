package main

func Example_a() {
	a()
	// Output:
	// 問8.1
	// -- Binary tree --
	// 	{Steven 1994}
	// {Gary 1997}
	// 		{Upton 1988}
	// 	{Tyler 1985}
	// 		{Martin 1973}
}

func Example_b() {
	b()
	// Output:
	// 8.3節
	// -- Binary search tree --
	// 		{Bridget 1982}
	// 	{Cindy 1983}
	// 			{Jack 1995}
	// 		{Luther 1989}
	// 			{Maroon 1998}
	// {Minks 1985}
	// 		{Minuet 1978}
	// 	{Rady 1998}
	// 		{Ruby 1995}
	// 			{Tiki 1982}
	//
	// {Minks 1985}
	// {Cindy 1983}
	// {Luther 1989}
	// {Maroon 1998}
	//
	// -- Binary search tree --
	// 		{Bridget 1982}
	// 	{Cindy 1983}
	// 				{Elpe 1982}
	// 			{Jack 1995}
	// 		{Luther 1989}
	// 			{Maroon 1998}
	// {Minks 1985}
	// 		{Minuet 1978}
	// 	{Rady 1998}
	// 		{Ruby 1995}
	// 			{Tiki 1982}
}

func Example_c() {
	c()
	// Output:
	// 8.4節
	// -- Generated binary search tree --
	// 			{Dali 1904}
	// 		{Hokusai 1760}
	// 			{Kandinsky 1866}
	// 				{Matisse 1869}
	// 	{Michelangelo 1475}
	// 			{Picasso 1881}
	// 		{Raphael 1483}
	// 			{Utrillo 1883}
	// {da Vinci 1452}
	// 	{van Gogh 1853}
	//
	// -- preorder trversal --
	// {da Vinci 1452}
	// {Michelangelo 1475}
	// {Hokusai 1760}
	// {Dali 1904}
	// {Kandinsky 1866}
	// {Matisse 1869}
	// {Raphael 1483}
	// {Picasso 1881}
	// {Utrillo 1883}
	// {van Gogh 1853}
	//
	// -- inorder trversal --
	// {Dali 1904}
	// {Hokusai 1760}
	// {Kandinsky 1866}
	// {Matisse 1869}
	// {Michelangelo 1475}
	// {Picasso 1881}
	// {Raphael 1483}
	// {Utrillo 1883}
	// {da Vinci 1452}
	// {van Gogh 1853}
	//
	// -- postorder trversal --
	// {Dali 1904}
	// {Matisse 1869}
	// {Kandinsky 1866}
	// {Hokusai 1760}
	// {Picasso 1881}
	// {Utrillo 1883}
	// {Raphael 1483}
	// {Michelangelo 1475}
	// {van Gogh 1853}
	// {da Vinci 1452}
}
