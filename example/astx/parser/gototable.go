// generated by gocc; DO NOT EDIT.

package parser

const numNTSymbols = 3

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		-1, // S'
		1,  // StmtList
		2,  // Stmt

	},
	gotoRow{ // S1
		-1, // S'
		-1, // StmtList
		4,  // Stmt

	},
	gotoRow{ // S2
		-1, // S'
		-1, // StmtList
		-1, // Stmt

	},
	gotoRow{ // S3
		-1, // S'
		-1, // StmtList
		-1, // Stmt

	},
	gotoRow{ // S4
		-1, // S'
		-1, // StmtList
		-1, // Stmt

	},
}
