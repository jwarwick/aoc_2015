%% Advent of Code, Day 21
%% Minimize spending while surviving the battle

int: base_hit_points;

int: boss_hit_points;
int: boss_damage;
int: boss_armor;

enum WEAPON;
array[WEAPON] of int: weapon_cost;
array[WEAPON] of int: weapon_damage;

enum ARMOR;
array[ARMOR] of int: armor_cost;
array[ARMOR] of int: armor_armor;

enum RING;
array[RING] of int: ring_cost;
array[RING] of int: ring_damage;
array[RING] of int: ring_armor;

var WEAPON: weapon;
var ARMOR: armor;
var RING: ring1;
var RING: ring2;

constraint ring1 != ring2;

var int: damage_dealt;
damage_dealt = max(1, (weapon_damage[weapon] + ring_damage[ring1] + ring_damage[ring2]) - boss_armor);
var int: damage_taken;
damage_taken = max(1, boss_damage - (armor_armor[armor] + ring_armor[ring1] + ring_armor[ring2]));

var 0.0..1000.0: self_rounds;
self_rounds = base_hit_points / damage_taken;
var 0.0..1000.0: boss_rounds;
boss_rounds = boss_hit_points / damage_dealt;

var int: self_ceil;
var int: boss_ceil;
constraint self_ceil >= self_rounds;
constraint self_ceil < self_rounds + 1;
constraint boss_ceil >= boss_rounds;
constraint boss_ceil < boss_rounds + 1;

constraint self_ceil >= boss_ceil;

var int: cost;
cost = weapon_cost[weapon] + armor_cost[armor] + ring_cost[ring1] + ring_cost[ring2];

solve minimize cost;
