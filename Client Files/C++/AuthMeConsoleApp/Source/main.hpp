// main.hpp : This file declares the functions that will used in main.cpp

#include <Windows.h>
#include <iostream>

// Console : This namespace will hold all example functions and variables used.
namespace Console
{
	// Cout : Shorthand std::cout into a function for quicker coding.
	void Cout(const char* text, const char* endl)
	{
		std::cout << text << endl;
	}

	// Printf : Shorthand printf() into a function for quicker coding.
	void Printf(const char* text, auto format)
	{
		printf(text, format);
	}

	// Cls : Shorthand system("cls") into a function for quicker coding.
	void Cls()
	{
		system("cls");
	}

	// Pause : Shorthand system("pause") into a function for quicker coding.
	void Pause()
	{
		system("pause");
	}
}