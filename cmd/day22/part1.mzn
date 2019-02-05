%% Advent of Code 2015, Day 22
%% Minimize spending mana while winning the battle

int: base_hit_points;
int: base_mana_points;

int: boss_hit_points;
int: boss_damage;

int: per_turn_damage;

enum SPELL;
array[SPELL] of int: spell_cost;
array[SPELL] of int: damage;

int: MAX_ROUNDS = 22;
var 1..MAX_ROUNDS: last_round;

array[-6..MAX_ROUNDS] of var SPELL: spells;

constraint forall(i in -6..0) (spells[i] = NONE);
constraint forall(i in 1..MAX_ROUNDS)
  (i mod 2 = 0 <-> spells[i] = BOSS_TURN);
constraint forall(i in 1..MAX_ROUNDS)
  (if i <= last_round then
   spells[i] != NONE
   else
   spells[i] = NONE \/ spells[i] = BOSS_TURN
   endif);

constraint forall(i in 1..MAX_ROUNDS)
  (player_mana[i-1] + (101 * recharge_active[i]) >= spell_cost[spells[i]]);


array[0..MAX_ROUNDS] of var int: player_health;
array[0..MAX_ROUNDS] of var int: player_mana;
array[1..MAX_ROUNDS] of var int: mana_spent;
array[1..MAX_ROUNDS] of var bool: poison_active;
array[1..MAX_ROUNDS] of var bool: recharge_active;
array[1..MAX_ROUNDS] of var bool: shield_active;

constraint player_health[0] = base_hit_points;
constraint forall (i in 1..MAX_ROUNDS)
  (if i mod 2 = 0 /\ boss_health[i] > 0 then
   player_health[i] = player_health[i-1] - max(1, boss_damage - (7 * shield_active[i]))
   else
   player_health[i] = (player_health[i-1] - per_turn_damage) + (2 * (spells[i] = DRAIN))
   endif);

constraint player_mana[0] = base_mana_points;
constraint forall (i in 1..MAX_ROUNDS)
  (mana_spent[i] = spell_cost[spells[i]]);
constraint forall (i in 1..MAX_ROUNDS)
  (player_mana[i] = player_mana[i-1] + (101 * recharge_active[i]) - mana_spent[i]);

constraint forall (i in 1..MAX_ROUNDS)
  (poison_active[i] = true <-> exists(x in i-6..i-1) (spells[x] = POISON));
constraint forall (i in 1..MAX_ROUNDS-6)
  ((spells[i] = POISON) -> (forall (x in i+1..i+5)(spells[x] != POISON)));

constraint forall (i in 1..MAX_ROUNDS)
  (recharge_active[i] = true <-> exists(x in i-5..i-1) (spells[x] = RECHARGE));
constraint forall (i in 1..MAX_ROUNDS-5)
  ((spells[i] = RECHARGE) -> (forall (x in i+1..i+4)(spells[x] != RECHARGE)));

constraint forall (i in 1..MAX_ROUNDS)
  (shield_active[i] = true <-> exists(x in i-6..i-1) (spells[x] = SHIELD));
constraint forall (i in 1..MAX_ROUNDS-6)
  ((spells[i] = SHIELD) -> (forall (x in i+1..i+5)(spells[x] != SHIELD)));

array[0..MAX_ROUNDS] of var int: boss_health;
constraint boss_health[0] = boss_hit_points;
constraint forall (i in 1..MAX_ROUNDS)
  (boss_health[i] = boss_health[i-1] - (damage[spells[i]] + (3 * poison_active[i])));

constraint exists(i in 1..MAX_ROUNDS) (boss_health[i] <= 0 /\ player_health[i-1] > 0 /\ last_round = i);
constraint forall(i in 1..last_round) (player_health[i-1] > per_turn_damage);
constraint forall(i in 1..last_round) (boss_health[i-1] >= 0);

var int: total_mana_spent;
constraint total_mana_spent = sum(mana_spent);

solve minimize total_mana_spent;

%% Example 1
/*constraint spells[1] = POISON;*/
/*constraint spells[3] = MAGIC_MISSLE;*/

%% Example 2
/*constraint spells[1] = RECHARGE;*/
/*constraint spells[3] = SHIELD;*/
/*constraint spells[5] = DRAIN;*/

output [if i mod 2 = 0 then "-- Boss Turn \(i)\n" else "-- Player Turn \(i)\n" endif
++ "- Player has \(player_health[i]) hit points, \(7 * shield_active[i]) armor, \(player_mana[i]) mana\n"
++ "- Boss has \(boss_health[i]) hit points\n"
++ if fix(poison_active[i]) then "Poison deals 3 damage\n" else "" endif
++ if fix(recharge_active[i]) then "Recharge provides 101 mana\n" else "" endif
++ if i mod 2 = 0 then
"Boss attacks for \(boss_damage) - \(shield_active[i] * 7)\n"
else
"Player casts \(spells[i])\n"
endif
++ "\n"
| i in 1..fix(last_round)];
output [if i mod 2 != 0 then show(spells[i]) else " -> " endif | i in 1..fix(last_round)] ++ ["\n"];
output ["Total Mana Spent: \(total_mana_spent)"];
