package main

func Example_main() {
	main()
	// Output:
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
