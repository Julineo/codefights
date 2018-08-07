/*
https://app.codesignal.com/arcade/intro/level-9/6M57rMTFB9MeDeSWo
Given the positions of a white bishop and a black pawn on the standard chess board, determine whether the bishop can capture the pawn in one move.

The bishop has no restrictions in distance for each move, but is limited to diagonal movement. Check out the example below to see how it can move:


Example

For bishop = "a1" and pawn = "c3", the output should be
bishopAndPawn(bishop, pawn) = true.



For bishop = "h1" and pawn = "h3", the output should be
bishopAndPawn(bishop, pawn) = false.
*/

bool bishopAndPawn(string bishop, string pawn) {
    bishop = bishop.replace("a", "8");
    bishop = bishop.replace("b", "7");
    bishop = bishop.replace("c", "6");
    bishop = bishop.replace("d", "5");
    bishop = bishop.replace("e", "4");
    bishop = bishop.replace("f", "3");
    bishop = bishop.replace("g", "2");
    bishop = bishop.replace("h", "1");

    pawn = pawn.replace("a", "8");
    pawn = pawn.replace("b", "7");
    pawn = pawn.replace("c", "6");
    pawn = pawn.replace("d", "5");
    pawn = pawn.replace("e", "4");
    pawn = pawn.replace("f", "3");
    pawn = pawn.replace("g", "2");
    pawn = pawn.replace("h", "1");

    int a1;
    int b1;
    int a2;
    int b2;

    a1 = to!int(bishop[0..1]);
    b1 = to!int(bishop[1..2]);
    a2 = to!int(pawn[0..1]);
    b2 = to!int(pawn[1..2]);

    return abs(a1-a2) == abs(b1-b2);
}
