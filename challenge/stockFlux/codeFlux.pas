//https://app.codesignal.com/challenge/jiLk3Wj4oiiEDHfgt
//reverse challenge
//
//Array record representation is the following:
//  type
//    Array<type> = array of <fpc type>;
//    ArrayArray<type> = array of array of <type>;
//    ArrayArrayArray<type> = array of array of array of <type>;
//
//  where <type> is one of (Integer, Int64, Boolean, Extended, Char, String).
//  For example, array of array of integers will be represented in the following way:
//  type
//    ArrayArrayInteger = array of array of integer;
var 
    tmp: extended;
    i: longint;
    
function stockFlux(p: ArrayExtended): extended;
begin
    tmp := 100.0;
    for i := 0 to length(p)-1 do
    begin
        tmp := tmp + (tmp * p[i] * 0.01);
    end;
    result := tmp - 100.0;
end;
