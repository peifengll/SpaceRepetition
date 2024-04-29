function reviewCards(decks) {
    for (var i_1 = 0; i_1 < decks.length; i_1++) {
        var deck = decks[i_1];
        while (deck.cards.length > 0) {
            var aa = deck.cards.shift();
            console.log(aa);
            // 还要复习的时候，插入回去操作
            if (shouldReview(aa)) {
                insertCard(deck, aa);
            }
        }
        // 复习完了之后换牌组
        if (deck.cards.length === 0) {
            decks.splice(i_1, 1);
            i_1--;
        }
        console.log("现在牌组的长度:", decks.length);
    }
}
var i = 5;
function shouldReview(card) {
    i++;
    if (i % 2 == 0) {
        return false;
    }
    else {
        return true;
    }
    // 根据需要实现复习逻辑，返回true表示需要复习，false表示不需要复习
    return Math.random() < 0.5;
}
function insertCard(deck, card) {
    var randomIndex = Math.floor(Math.random() * 5) + 1;
    deck.cards.splice(randomIndex, 0, card);
}
var deccc = [
    {
        "id": 10,
        "name": "测试牌组号",
        "num": 2,
        "cards": [
            {
                "id": 14,
                "record_id": 46,
                "font": "二号二号的测试卡片",
                "originfont": "二号二号的测试卡片###这是二号",
                "back": "这是二号",
                "deckid": 10,
                "typeof": 1
            },
            {
                "id": 15,
                "record_id": 47,
                "font": "二号测试",
                "originfont": "二号测试###这个的目的是为了验证跳转问题",
                "back": "这个的目的是为了验证跳转问题",
                "deckid": 10,
                "typeof": 1
            }
        ]
    },
    {
        "id": 10,
        "name": "测试牌组号",
        "num": 2,
        "cards": [
            {
                "id": 14,
                "record_id": 46,
                "font": "二号二号的测试卡片",
                "originfont": "二号二号的测试卡片###这是二号",
                "back": "这是二号",
                "deckid": 10,
                "typeof": 1
            },
            {
                "id": 15,
                "record_id": 47,
                "font": "二号测试",
                "originfont": "二号测试###这个的目的是为了验证跳转问题",
                "back": "这个的目的是为了验证跳转问题",
                "deckid": 10,
                "typeof": 1
            }
        ]
    }
];
reviewCards(deccc);
