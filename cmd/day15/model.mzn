enum INGREDIENTS;

array[INGREDIENTS] of int: capacity;
array[INGREDIENTS] of int: durability;
array[INGREDIENTS] of int: flavor;
array[INGREDIENTS] of int: texture;
array[INGREDIENTS] of int: calories;

array[INGREDIENTS] of var 0..100: teaspoons;

constraint sum(teaspoons) = 100;

% For Part 2
constraint calorie_score = 500;

% These don't need to be floats, but this version of MiniZinc seems to make an
% intermediate variable that is larger than MAX_INT
var float: capacity_score = max(0, sum([teaspoons[i] * capacity[i] | i in INGREDIENTS]));
var float: durability_score = max(0, sum([teaspoons[i] * durability[i] | i in INGREDIENTS]));
var float: flavor_score = max(0, sum([teaspoons[i] * flavor[i] | i in INGREDIENTS]));
var float: texture_score = max(0, sum([teaspoons[i] * texture[i] | i in INGREDIENTS]));
var float: calorie_score = max(0, sum([teaspoons[i] * calories[i] | i in INGREDIENTS]));

var float: score;
score = capacity_score * durability_score * flavor_score * texture_score;


solve maximize score;

output ["Teaspoons = \(teaspoons);\nTotal Score = \(score);\n"];
