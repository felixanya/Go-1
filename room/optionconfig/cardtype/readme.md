# 番型选项配置说明


## 番型表

番型 ID | 说明
--------|--------------
0 | 平胡
1 | 清一色
2 | 七对
3 | 清七对
4 | 龙七对
5 | 清龙七对
6 | 碰碰胡
7 | 清碰
8 | 金钩钓
9 | 清金钩钓
10 | 十八罗汉
11 | 清十八罗汉
12 | 大四喜
13 | 大三元
14 | 九莲宝灯
15 | 大于五
16 | 小于五
17 | 大七星
18 | 连七对
19 | 四杠
20 | 小四喜
21 | 小三元
22 | 双龙会
23 | 字一色
24 | 四暗刻
24 | 四暗刻
25 | 四同顺
26 | 三元七对
27 | 四喜七对
28 | 28四连刻
29 | 四步高
30 | 混幺九
31 | 三杠
32 | 天听
33 | 四字刻
34 | 大三风
35 | 三同顺
36 | 三连刻
37 | 全花
38 | 三暗刻
39 | 清龙
40 | 三步高
41 | 双箭刻
42 | 双暗杠
43 | 小三风
44 | 混一色
45 | 全带幺
46 | 双明杠
47 | 报听
48 | 报听一发
49 | 春夏秋冬
50 | 梅兰竹菊
51 | 无花牌
52 | 门风刻
53 | 圈风刻
54 | 箭刻
55 | 四归一
56 | 断幺
57 | 双暗刻
58 | 暗杠
59 | 门前清
60 | 一般高
61 | 连六
62 | 老少副
63 | 花牌
64 | 明杠
65 | 边张
66 | 坎张
67 |单钓将
68 | 天胡
69 | 地胡
70 | 人胡
71 | 妙手回春
72 | 全求人
73 | 不求人
74 | 绝张
75 | 自摸
76 | 杠上开花
77 | 杠上海底捞
78 | 杠后炮
79 | 海底捞月
80 | 抢杠胡

## 函数 ID 

计算函数 ID | 说明 
-----------|-----------------------
0 | 平胡，所有牌型均满足平胡
1 | 清一色，手牌、碰牌、杠牌、胡牌、吃牌所有花色均相同
2 | 七对，手牌+胡牌数量为14张，每张牌的数量为2的整数倍
3 | 清七对，满足七对和清一色条件
4 | 龙七对，手牌+胡牌数量为14张，每张牌的数量为2的整数倍，且至少有一张牌的数量不小于4
5 | 清龙七对，满足龙七对和清一色条件
6 | 碰碰胡，手牌+胡牌的顺的数量为0，吃牌数为0
7 | 清碰，满足清一色和碰碰胡的条件
8 | 金钩钓，手牌+胡牌数量为2且相同
9 | 清金钩钓，满足金钩钓和清一色条件
10 | 十八罗汉，满足金钩钓的条件，吃、碰数为0
11 | 清十八罗汉，满足清一色和十八罗汉的条件
12 | 大四喜,胡牌时，含有 4 副风刻或杠（明刻、暗刻、明杠、暗杠）
13 | 大三元,胡牌时，含有“中发白”3 副刻子（明刻、暗刻）
14 | 九莲宝灯,由一种花色序数牌子组成的特定牌型（1112345678999），见同花色任何 1 张序数牌即成胡牌；
15 | 大于五,由序数牌 6-9 的顺子、刻子、将牌组成的胡牌
16 | 小于五,由序数牌 1-4 的顺子、刻子、将牌组成的胡牌
17 | 大七星，胡牌为七对,并且由“东南西北中发白”其中的字牌构成
18 | 连七对，由一种花色序数牌组成序数相连的 7 个对子组成的胡牌
19 | 四杠，胡牌时,含有 4 个杠(明杠、暗杠)
20 | 小四喜,胡牌时，含有风牌的 3 副刻子（明刻、暗刻）及将牌
21 | 小三元,胡牌时，含有箭牌的 2 副刻子（明刻、暗刻）及将牌
22 | 双龙会，由一种花色的 2 个老少副,5 为将牌组成的胡牌
23 | 字一色，由字牌组成的胡牌
24 | 四暗刻，胡牌时,含有 4 个暗刻(或暗杠)
24 | 四暗刻,胡牌时，含有 4 个暗刻（或暗杠）
25 | 四同顺,胡牌时，含有一种花色 4 副相同的顺子
26 | 三元七对,胡牌为七对，并且包含“中发白
27 | 四喜七对，胡牌为七对,并且包含“东南西北”
28 | 28四连刻，胡牌时,含有一种花色 4 副依次递增一位数的刻子
29 | 四步高,胡牌时，含有一种花色 4 副依次递增一位数或依次递增两位数的顺子
30 | 混幺九,由序数牌 1、9 和字牌的刻子、将牌组成的胡牌
31 | 三杠，胡牌时,含有 3 个杠(明杠、暗杠)
32 | 天听，分为以下 3 种情况: 庄家打出第一张牌前报听称为天听,发完牌后闲家便报听也称为天听; 若发完牌后有玩家补花,补花之后报听也算天听; 若庄家在发完牌后有暗杠,则庄家不算天听,但算报听
33 | 四字刻,胡牌时，含有 4 个字牌的刻或杠（明刻、暗刻、明杠、暗杠
34 | 大三风,胡牌时，含有 3 个风刻（明刻、暗刻）
35 | 三同顺，胡牌时,含有一种花色 3 副序数相同的顺子;
36 | 三连刻，胡牌时,含有一种花色 3 副依次递增一位数字的刻子
37 | 全花，摸到全部 8 张花牌;
38 | 三暗刻，胡牌时,含有 3 个暗刻
39 | 清龙,胡牌时，含有一种花色 1-9 相连接的序数牌
40 | 三步高,胡牌时，含有一种花色 3 副依次递增一位或依次递增两位数字的顺子
41 | 双箭刻,胡牌时，含有 2 副箭刻或杠（明刻、暗刻、明杠、暗杠）
42 | 双暗杠，胡牌时,含有 2 个暗杠
43 | 小三风,胡牌时，含有 2 个风牌的刻或杠（明刻、暗刻、明杠、暗杠），以及 1 对风牌作为将眼
44 | 混一色，由一种花色序数牌及字牌组成的胡牌
45 | 全带幺,胡牌时，每副牌、将牌都有幺九牌
46 | 双明杠,胡牌时，含有 2 个明杠
47 | 报听,主动选择听牌,自动摸牌打牌后胡牌(报听后可补杠、暗杠,不可直杠)
48 | 报听一发,报听后紧接着就胡牌(包括正常的点炮或自摸,以及补花后的点炮或自摸)
49 | 春夏秋冬,摸到”春夏秋冬”4 张花牌,计花牌
50 | 梅兰竹菊,摸到”梅兰竹菊”4 张花牌,计花牌
51 | 无花牌
52 | 门风刻,与本门风相同的风刻（明刻、暗刻）
53 | 圈风刻,与圈风刻相同的风刻（明刻、暗刻）
54 | 箭刻,由“中发白”三张相同的牌组成的刻子（明刻、暗刻）
55 | 四归一,4 张相同的牌归于一家的顺、刻子、对、将牌中（不包括杠牌）
56 | 断幺,胡牌时,没有一、九及字牌
57 | 双暗刻,胡牌时,含有 2 个暗刻
58 | 暗杠,自抓 4 张相同的牌开杠,暗杠对本家亮明一张,对对家不亮明
59 | 门前清,没有吃、碰、杠,胡别人打出的牌
60 | 一般高,由一种花色 2 副相同的顺子组成的胡牌
61 | 连六,胡牌时，含有一种花色 6 张相连的序数牌
62 | 老少副,胡牌时，含有一种花色牌的 123、789 两幅顺子
63 | 花牌，即春夏秋冬梅兰竹菊，每花计 1 番
64 | 明杠,自己有暗刻,碰别人打出的一张相同的牌开杠;或自己有明刻,自己抓进一张相同的牌开杠
65 | 边张,单胡 123 的 3 及 789 的 7 或 1233 胡 3、77879 胡 7 都为张;手中有 12345胡 3,56789 胡 6 不算边
66 | 坎张,胡 2 张牌之间的牌,4556 胡 5 也为坎张,手中有 45567 胡 6 不算坎张
67 |单钓将,钓单张牌作将成胡,1112 胡 2 算单钓将,1234 胡 1、4 不算单钓将
68 | 天胡，天胡，庄家在发完手牌后就胡牌,此为天胡,若庄家有补花,在补完花后就胡牌也算;若庄家在发完牌后有暗杠杠出,那么不算天胡
69 | 地胡，地胡，闲家摸到第一张牌就胡牌,此为地胡,若闲家抓的第一张牌是花牌,那么补花之后胡牌也算地胡;若闲家抓牌前有人吃碰杠(包括暗杠),那么不算地胡
70 | 人胡，庄家打出的第一张牌闲家就胡牌,此为人胡,若庄家出牌前有暗杠,那么不算人胡
71 | 妙手回春,自摸牌墙上最后一张牌胡牌
72 | 全求人，全靠吃牌、碰牌、单钓别人打出的牌胡牌
73 | 不求人,（4 副牌及将牌中）没有吃牌、碰牌（包括明杠）的自摸胡牌
74 | 绝张,胡牌池、桌面已亮明的 3 张牌所剩的第 4 张牌（抢杠和不计和绝张
75 | 自摸，自己抓进牌成胡牌
76 | 杠上开花，开杠补抓的牌成胡牌(补花之后也算)
77 | 杠上海底捞，满足杠上开花，海底捞条件
78 | 杠后炮,玩家杠后，出牌，被点炮
79 | 海底捞月,胡打出的最后一张牌
80 | 抢杠胡，胡别人自抓开补杠的牌
