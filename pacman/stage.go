package pacman

type stage struct {
	matrix []string
}

var stage1 = stage{
	[]string{
		"3888888884888888885",
		"gqrrrrrrrgrrrrrrrqg",
		"gr02r789ror789r02rg",
		"gracrrrrrrrrrrracrg",
		"grikr6r78489r6rikrg",
		"grrrrgrrrgrrrgrrrrg",
		"gr02r789ror789r02rg",
		"grikrgrrrrrrrgrikrg",
		"grrrror39v75rorrrrg",
		"gr39rrrguwtgrrr75rg",
		"grgrr6rl888nr6rrgrg",
		"gror7nrrrsrrrl9rorg",
		"grrrrrr6r6r6rrrrrrg",
		"gr012r3nrorl5r012rg",
		"grijkrorrprrorijkrg",
		"gqrrrrrr385rrrrrrqg",
		"l888888m8m88888888n",
	},
}
var defaultStage = &stage1
