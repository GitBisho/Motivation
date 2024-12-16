package main

const ( EMPTY = iota
		wP = iota
		wN = iota
		wB = iota
		wQ = iota
		wK = iota
		bP = iota
		bN = iota
		bB = iota
		bQ = iota
		bK = iota 
)

const ( FILE_A = iota
		FILE_B = iota
		FILE_C = iota
		FILE_D = iota
		FILE_E = iota
		FILE_F = iota
		FILE_G = iota
		FILE_H = iota 
)

const ( RANK_1 = iota
		RANK_2 = iota
		RANK_3 = iota
		RANK_4 = iota
		RANK_5 = iota
		RANK_6 = iota
		RANK_7 = iota
		RANK_8 = iota

)
const ( WHITE = iota
		BLACK = iota
		BOTH = iota
)
const (A1 = iota 
B1 = iota 
C1 = iota 
D1 = iota 
E1 = iota 
F1 = iota 
G1 = iota 
H1 = iota
A2 = iota 
B2 = iota 
C2 = iota 
D2 = iota 
E2 = iota 
F2 = iota 
G2 = iota 
H2 = iota
A3 = iota 
B3 = iota 
C3 = iota 
D3 = iota 
E3 = iota 
F3 = iota 
G3 = iota 
H3 = iota
A4 = iota 
B4 = iota 
C4 = iota 
D4 = iota 
E4 = iota 
F4 = iota 
G4 = iota 
H4 = iota
A5 = iota 
B5 = iota 
C5 = iota 
D5 = iota 
E5 = iota 
F5 = iota 
G5 = iota 
H5 = iota
A6 = iota 
B6 = iota 
C6 = iota 
D6 = iota 
E6 = iota 
F6 = iota 
G6 = iota 
H6 = iota
A7 = iota 
B7 = iota 
C7 = iota 
D7 = iota 
E7 = iota 
F7 = iota 
G7 = iota 
H7 = iota
A8 = iota 
B8 = iota 
C8 = iota 
D8 = iota 
E8 = iota 
F8 = iota 
G8 = iota 
H8 = iota
)

var Sq120ToSq64 = [120]int32{}
var Sq64ToSq120 = [64]int32{}

func FileRankToSquare(file int32, rank int32) int32 {
	return (21 + file) + (rank * 10);
}