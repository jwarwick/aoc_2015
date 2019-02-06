%% Advent of Code 2015, Day 24

array[int] of 0.0..125.0: weights;
int: cnt = length(weights);

array[1..cnt] of var bool: g1;
array[1..cnt] of var bool: g2;
array[1..cnt] of var bool: g3;
array[1..cnt] of var bool: g4;

var float: g1_weight = sum(i in 1..cnt)(weights[i] * g1[i]);
var float: g2_weight = sum(i in 1..cnt)(weights[i] * g2[i]);
var float: g3_weight = sum(i in 1..cnt)(weights[i] * g3[i]);
var float: g4_weight = sum(i in 1..cnt)(weights[i] * g4[i]);

var 0..cnt: g1_cnt = sum(g1);
var 0..cnt: g2_cnt = sum(g2);
var 0..cnt: g3_cnt = sum(g3);
var 0..cnt: g4_cnt = sum(g4);

array[1..cnt] of var float: logs;
constraint forall (i in 1..cnt)(logs[i] = ln(weights[i]));

array[1..cnt] of var 0.0..15.0: g1_logs;
constraint forall (i in 1..cnt)(g1_logs[i] = if g1[i] then logs[i] else 0 endif);

var float: qe = sum(g1_logs);

constraint g1_weight = g2_weight /\ g2_weight = g3_weight /\ g3_weight = g4_weight;
constraint forall (i in 1..cnt) (g1[i] \/ g2[i] \/ g3[i] \/ g4[i]);

constraint g1_cnt <= g2_cnt /\ g1_cnt <= g3_cnt /\ g1_cnt <= g4_cnt;

constraint cnt = g1_cnt + g2_cnt + g3_cnt + g4_cnt;

%% after determining the number of packages in the first group
%% fix that value and minimize the product of those packages
/*constraint g1_cnt = 6;*/

%% wrong
/*constraint g1 = [true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, true, true, true];*/
%% correct
/*constraint g1 = [true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, true, true, true];*/

/*solve minimize g1_cnt;*/
/*solve minimize sum(g1_logs);*/
solve minimize qe;

output ["g1 = \(g1), sum = \(g1_weight), cnt = \(g1_cnt)\n"];
output ["g2 = \(g2), sum = \(g2_weight), cnt = \(g2_cnt)\n"];
output ["g3 = \(g3), sum = \(g3_weight), cnt = \(g3_cnt)\n"];
output ["g4 = \(g4), sum = \(g4_weight), cnt = \(g4_cnt)\n"];
output ["logs = \(logs)\n"];
output ["g1_logs = \(g1_logs)\n"];
output ["ln(qe) = \(qe), qe = \(exp(qe))\n"];
output [if fix(g1[i]) then show(weights[i]) ++ " * " else "" endif | i in 1..cnt] ++ ["\n"];
