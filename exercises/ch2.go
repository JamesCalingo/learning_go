package exercises

import "fmt"

func printIAndF1() {
	var i int = 20
	var f float32 = float32(i)
	fmt.Println(i, f)
	//We get 20 and 20 (Well, technically the second one is 20.0, but the trailing decimals are left off for some reason)
}

func printIAndF2() {
	const value = 20
	var i int = value
	var f float32 = value
	fmt.Println(i, f)
	//Just like the last exercise, we get 20 and 20.0 shortened to 20
}

func funWithOverflow() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	fmt.Println(b + 1, smallI + 1, bigI + 1)
	/*We get 0, -2147483648, and 0 thanks to a little property known as overflow.

	For those who don't know, overflow is a computer science property which dictates what happens if you try to go over a certain integer's range. Instead of trying to expand or change type (because few programming languages will actually do that without being asked), the value will reset to the lowest value in the range (i.e. 0 for unsigned ints and the negative number for signed ints).

	This has had some very interesting real world implications - most notably in the classic video game Pac Man. The game's level counter is stored as an 8-bit integer (i.e. b in this example), so there are, in practice, 255 levels. However, players have been able to beat level 255 and reach "level 256" - a number the hardware has no idea how to represent (remember, most languages start counting from 0). What you end up with is a glitched out screen as the game attempts to create a "level 256", but ends up resetting to "level 0" (though it's a bit more complex than this due to other underlying mechanics). This wiki article goes into more detail, as well as showing what this glitch looks like:
	https://pacman.fandom.com/wiki/Map_256_Glitch

	Also, this has some worried about 2038, when the 32 bit time system a lot of the computing world has been using since 1970 will overflow. A solution to this has been to switch to 64 bit time systems, but not everyone has adapted this yet, so no one knows what devices stuck in the 32 bit time system will do...
	
	(P.S.: There's also underflow, which is basically the same concept but going the other way)
	*/
}

func Ch2() {
	printIAndF1()
	printIAndF2()
	funWithOverflow()
}