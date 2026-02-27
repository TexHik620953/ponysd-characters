-- +goose Up
-- +goose StatementBegin

-- Локализации для body_type
INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Slim' FROM glossary WHERE type = 'body_type' AND value = 'slim';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Medium' FROM glossary WHERE type = 'body_type' AND value = 'medium';

-- Локализации для breast_type
INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Flat' FROM glossary WHERE type = 'breast_type' AND value = 'flat';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Small' FROM glossary WHERE type = 'breast_type' AND value = 'small';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Medium' FROM glossary WHERE type = 'breast_type' AND value = 'medium';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Large' FROM glossary WHERE type = 'breast_type' AND value = 'large';

-- Локализации для butt_type
INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Skinny' FROM glossary WHERE type = 'butt_type' AND value = 'skinny';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Athletic' FROM glossary WHERE type = 'butt_type' AND value = 'athletic';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Medium' FROM glossary WHERE type = 'butt_type' AND value = 'medium';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Large' FROM glossary WHERE type = 'butt_type' AND value = 'large';

-- Локализации для eyes_color
INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Brown' FROM glossary WHERE type = 'eyes_color' AND value = 'brown';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Blue' FROM glossary WHERE type = 'eyes_color' AND value = 'blue';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Hazel' FROM glossary WHERE type = 'eyes_color' AND value = 'hazel';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Gray' FROM glossary WHERE type = 'eyes_color' AND value = 'gray';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Amber' FROM glossary WHERE type = 'eyes_color' AND value = 'amber';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Red' FROM glossary WHERE type = 'eyes_color' AND value = 'red';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Violet' FROM glossary WHERE type = 'eyes_color' AND value = 'violet';

-- Локализации для hair_color
INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Black' FROM glossary WHERE type = 'hair_color' AND value = 'black';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Blue' FROM glossary WHERE type = 'hair_color' AND value = 'blue';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Brown' FROM glossary WHERE type = 'hair_color' AND value = 'brown';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Blonde' FROM glossary WHERE type = 'hair_color' AND value = 'blonde';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Red' FROM glossary WHERE type = 'hair_color' AND value = 'red';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Auburn' FROM glossary WHERE type = 'hair_color' AND value = 'auburn';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Chestnut' FROM glossary WHERE type = 'hair_color' AND value = 'chestnut';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Grey' FROM glossary WHERE type = 'hair_color' AND value = 'grey';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'White' FROM glossary WHERE type = 'hair_color' AND value = 'white';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Salt and Pepper' FROM glossary WHERE type = 'hair_color' AND value = 'salt and pepper';

-- Локализации для hair_style
INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Asymmetrical Cut' FROM glossary WHERE type = 'hair_style' AND value = 'asymmetrical cut';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Blunt Cut' FROM glossary WHERE type = 'hair_style' AND value = 'blunt cut';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Bob Cut' FROM glossary WHERE type = 'hair_style' AND value = 'bob cut';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Braided Bob' FROM glossary WHERE type = 'hair_style' AND value = 'braided bob';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Buzz Cut' FROM glossary WHERE type = 'hair_style' AND value = 'buzz cut';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Choppy Cut' FROM glossary WHERE type = 'hair_style' AND value = 'choppy cut';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Curly Bob' FROM glossary WHERE type = 'hair_style' AND value = 'curly bob';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Curtain Bangs' FROM glossary WHERE type = 'hair_style' AND value = 'curtain bangs';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Faux Hawk' FROM glossary WHERE type = 'hair_style' AND value = 'faux hawk';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Feathered Cut' FROM glossary WHERE type = 'hair_style' AND value = 'feathered cut';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'French Bob' FROM glossary WHERE type = 'hair_style' AND value = 'french bob';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Layered Cut' FROM glossary WHERE type = 'hair_style' AND value = 'layered cut';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Long Bob' FROM glossary WHERE type = 'hair_style' AND value = 'long bob';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Mohawk' FROM glossary WHERE type = 'hair_style' AND value = 'mohawk';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Pixie Cut' FROM glossary WHERE type = 'hair_style' AND value = 'pixie cut';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Shag Cut' FROM glossary WHERE type = 'hair_style' AND value = 'shag cut';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Side-Swept Bangs' FROM glossary WHERE type = 'hair_style' AND value = 'side-swept bangs';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Textured Cut' FROM glossary WHERE type = 'hair_style' AND value = 'textured cut';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Undercut' FROM glossary WHERE type = 'hair_style' AND value = 'undercut';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Wavy Bob' FROM glossary WHERE type = 'hair_style' AND value = 'wavy bob';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Faux Hawk Short Pixie' FROM glossary WHERE type = 'hair_style' AND value = 'faux hawk short pixie';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Brave Short Haircut' FROM glossary WHERE type = 'hair_style' AND value = 'brave short haircut with shaved sides';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Tapered Haircut' FROM glossary WHERE type = 'hair_style' AND value = 'tapered haircut with shaved side';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Stacked Bob' FROM glossary WHERE type = 'hair_style' AND value = 'stacked bob';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Lemonade Braids' FROM glossary WHERE type = 'hair_style' AND value = 'lemonade braids';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Middle Part Ponytails' FROM glossary WHERE type = 'hair_style' AND value = 'middle part ponytails';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Deep Side Part' FROM glossary WHERE type = 'hair_style' AND value = 'deep side part';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'French Braids' FROM glossary WHERE type = 'hair_style' AND value = 'french braids';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Box Braids' FROM glossary WHERE type = 'hair_style' AND value = 'box braids';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Two Dutch Braids' FROM glossary WHERE type = 'hair_style' AND value = 'two dutch braids';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Wavy Cut Curtain Bangs' FROM glossary WHERE type = 'hair_style' AND value = 'wavy cut with curtains bangs';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Right Side Shaved' FROM glossary WHERE type = 'hair_style' AND value = 'right side shaved';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Sweeping Pixie' FROM glossary WHERE type = 'hair_style' AND value = 'sweeping pixie';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Smooth Lob' FROM glossary WHERE type = 'hair_style' AND value = 'smooth lob';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Long Pixie' FROM glossary WHERE type = 'hair_style' AND value = 'long pixie';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Sideswept Pixie' FROM glossary WHERE type = 'hair_style' AND value = 'sideswept pixie';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Italian Bob' FROM glossary WHERE type = 'hair_style' AND value = 'italian bob';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Shullet' FROM glossary WHERE type = 'hair_style' AND value = 'shullet';

-- Локализации для nationality
INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Asian' FROM glossary WHERE type = 'nationality' AND value = 'asian';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'White' FROM glossary WHERE type = 'nationality' AND value = 'white';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Latina' FROM glossary WHERE type = 'nationality' AND value = 'latina';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Arab' FROM glossary WHERE type = 'nationality' AND value = 'arab';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Indian' FROM glossary WHERE type = 'nationality' AND value = 'indian';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Japanese' FROM glossary WHERE type = 'nationality' AND value = 'japanese';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Elf' FROM glossary WHERE type = 'nationality' AND value = 'elf';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Demon' FROM glossary WHERE type = 'nationality' AND value = 'demon';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Angel' FROM glossary WHERE type = 'nationality' AND value = 'angel';

-- Локализации для action
INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Looking at viewer' FROM glossary WHERE type = 'action' AND value = 'looking at viewer';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Looking away' FROM glossary WHERE type = 'action' AND value = 'looking away';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Looking up' FROM glossary WHERE type = 'action' AND value = 'looking up';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Looking down' FROM glossary WHERE type = 'action' AND value = 'looking down';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Looking back' FROM glossary WHERE type = 'action' AND value = 'looking back';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Open mouth' FROM glossary WHERE type = 'action' AND value = 'open mouth';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Half-closed mouth' FROM glossary WHERE type = 'action' AND value = 'half-closed mouth';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Closed mouth' FROM glossary WHERE type = 'action' AND value = 'closed mouth';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Open eyes' FROM glossary WHERE type = 'action' AND value = 'open eyes';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Half-closed eyes' FROM glossary WHERE type = 'action' AND value = 'half-closed eyes';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Closed eyes' FROM glossary WHERE type = 'action' AND value = 'closed eyes';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Wink' FROM glossary WHERE type = 'action' AND value = 'wink';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Standing' FROM glossary WHERE type = 'action' AND value = 'standing';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Sitting' FROM glossary WHERE type = 'action' AND value = 'sitting';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Lying' FROM glossary WHERE type = 'action' AND value = 'lying';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Walking' FROM glossary WHERE type = 'action' AND value = 'walking';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Running' FROM glossary WHERE type = 'action' AND value = 'running';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Adjusting hair' FROM glossary WHERE type = 'action' AND value = 'adjusting hair';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Waving' FROM glossary WHERE type = 'action' AND value = 'waving';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Hand on hip' FROM glossary WHERE type = 'action' AND value = 'hand on hip';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Crossed arms' FROM glossary WHERE type = 'action' AND value = 'crossed arms';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Smile' FROM glossary WHERE type = 'action' AND value = 'smile';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Sad' FROM glossary WHERE type = 'action' AND value = 'sad';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Angry' FROM glossary WHERE type = 'action' AND value = 'angry';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Sleepy' FROM glossary WHERE type = 'action' AND value = 'sleepy';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Tired' FROM glossary WHERE type = 'action' AND value = 'tired';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Expressionless' FROM glossary WHERE type = 'action' AND value = 'expressionless';

-- Локализации для clothes
INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Underwear' FROM glossary WHERE type = 'clothes' AND value = 'underwear';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Clothed' FROM glossary WHERE type = 'clothes' AND value = 'clothed';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Casual' FROM glossary WHERE type = 'clothes' AND value = 'casual';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Dress' FROM glossary WHERE type = 'clothes' AND value = 'dress';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Swimsuit' FROM glossary WHERE type = 'clothes' AND value = 'swimsuit';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Uniform' FROM glossary WHERE type = 'clothes' AND value = 'uniform';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Bikini' FROM glossary WHERE type = 'clothes' AND value = 'bikini';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'One-piece swimsuit' FROM glossary WHERE type = 'clothes' AND value = 'one-piece swimsuit';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Shirt' FROM glossary WHERE type = 'clothes' AND value = 'shirt';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Blouse' FROM glossary WHERE type = 'clothes' AND value = 'blouse';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Sweater' FROM glossary WHERE type = 'clothes' AND value = 'sweater';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Hoodie' FROM glossary WHERE type = 'clothes' AND value = 'hoodie';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Jeans' FROM glossary WHERE type = 'clothes' AND value = 'jeans';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Pants' FROM glossary WHERE type = 'clothes' AND value = 'pants';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Shorts' FROM glossary WHERE type = 'clothes' AND value = 'shorts';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Skirt' FROM glossary WHERE type = 'clothes' AND value = 'skirt';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Vest' FROM glossary WHERE type = 'clothes' AND value = 'vest';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Coat' FROM glossary WHERE type = 'clothes' AND value = 'coat';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Trench coat' FROM glossary WHERE type = 'clothes' AND value = 'trenchoat';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Jacket' FROM glossary WHERE type = 'clothes' AND value = 'jacket';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Short dress' FROM glossary WHERE type = 'clothes' AND value = 'short dress';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Long dress' FROM glossary WHERE type = 'clothes' AND value = 'long dress';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Off-shoulder' FROM glossary WHERE type = 'clothes' AND value = 'off-shoulder';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Backless' FROM glossary WHERE type = 'clothes' AND value = 'backless';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Hairbow' FROM glossary WHERE type = 'clothes' AND value = 'hairbow';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Hair ribbon' FROM glossary WHERE type = 'clothes' AND value = 'hair ribbon';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Hair tie' FROM glossary WHERE type = 'clothes' AND value = 'hair tie';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Hairband' FROM glossary WHERE type = 'clothes' AND value = 'hairband';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Cap' FROM glossary WHERE type = 'clothes' AND value = 'cap';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Beanie' FROM glossary WHERE type = 'clothes' AND value = 'beanie';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Bucket hat' FROM glossary WHERE type = 'clothes' AND value = 'bucket hat';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Sun hat' FROM glossary WHERE type = 'clothes' AND value = 'sun hat';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Straw hat' FROM glossary WHERE type = 'clothes' AND value = 'straw hat';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Rice hat' FROM glossary WHERE type = 'clothes' AND value = 'rice hat';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Witch hat' FROM glossary WHERE type = 'clothes' AND value = 'witch hat';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Crown' FROM glossary WHERE type = 'clothes' AND value = 'crown';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Chain necklace' FROM glossary WHERE type = 'clothes' AND value = 'chain necklace';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Tooth necklace' FROM glossary WHERE type = 'clothes' AND value = 'tooth necklace';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Choker' FROM glossary WHERE type = 'clothes' AND value = 'choker';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Pendant' FROM glossary WHERE type = 'clothes' AND value = 'pendant';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Bracelet' FROM glossary WHERE type = 'clothes' AND value = 'bracelet';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Watch' FROM glossary WHERE type = 'clothes' AND value = 'watch';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Ring' FROM glossary WHERE type = 'clothes' AND value = 'ring';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Earring' FROM glossary WHERE type = 'clothes' AND value = 'earring';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Anklet' FROM glossary WHERE type = 'clothes' AND value = 'anklet';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Belt' FROM glossary WHERE type = 'clothes' AND value = 'belt';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Scarf' FROM glossary WHERE type = 'clothes' AND value = 'scarf';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Gloves' FROM glossary WHERE type = 'clothes' AND value = 'gloves';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Mittens' FROM glossary WHERE type = 'clothes' AND value = 'mittens';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Socks' FROM glossary WHERE type = 'clothes' AND value = 'socks';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Stockings' FROM glossary WHERE type = 'clothes' AND value = 'stockings';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Tights' FROM glossary WHERE type = 'clothes' AND value = 'tights';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Leggings' FROM glossary WHERE type = 'clothes' AND value = 'leggings';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Boots' FROM glossary WHERE type = 'clothes' AND value = 'boots';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Sneakers' FROM glossary WHERE type = 'clothes' AND value = 'sneakers';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Heels' FROM glossary WHERE type = 'clothes' AND value = 'heels';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Sandals' FROM glossary WHERE type = 'clothes' AND value = 'sandals';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Flip-flops' FROM glossary WHERE type = 'clothes' AND value = 'flip-flops';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Slippers' FROM glossary WHERE type = 'clothes' AND value = 'slippers';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Loafers' FROM glossary WHERE type = 'clothes' AND value = 'loafers';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Oxfords' FROM glossary WHERE type = 'clothes' AND value = 'oxfords';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Brogues' FROM glossary WHERE type = 'clothes' AND value = 'brogues';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Derbies' FROM glossary WHERE type = 'clothes' AND value = 'derbies';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Monk shoes' FROM glossary WHERE type = 'clothes' AND value = 'monk shoes';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Chelsea boots' FROM glossary WHERE type = 'clothes' AND value = 'chelsea boots';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Combat boots' FROM glossary WHERE type = 'clothes' AND value = 'combat boots';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Riding boots' FROM glossary WHERE type = 'clothes' AND value = 'riding boots';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Rain boots' FROM glossary WHERE type = 'clothes' AND value = 'rain boots';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Wedge heels' FROM glossary WHERE type = 'clothes' AND value = 'wedge heels';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Platform heels' FROM glossary WHERE type = 'clothes' AND value = 'platform heels';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Stilettos' FROM glossary WHERE type = 'clothes' AND value = 'stilettos';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Block heels' FROM glossary WHERE type = 'clothes' AND value = 'block heels';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Kitten heels' FROM glossary WHERE type = 'clothes' AND value = 'kitten heels';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Moccasins' FROM glossary WHERE type = 'clothes' AND value = 'moccasins';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Espadrilles' FROM glossary WHERE type = 'clothes' AND value = 'espadrilles';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Pumps' FROM glossary WHERE type = 'clothes' AND value = 'pumps';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Flats' FROM glossary WHERE type = 'clothes' AND value = 'flats';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Ballet flats' FROM glossary WHERE type = 'clothes' AND value = 'ballet flats';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Mary janes' FROM glossary WHERE type = 'clothes' AND value = 'mary janes';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Slingbacks' FROM glossary WHERE type = 'clothes' AND value = 'slingbacks';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Peep-toe' FROM glossary WHERE type = 'clothes' AND value = 'peep-toe';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Mule sandals' FROM glossary WHERE type = 'clothes' AND value = 'mule sandals';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Gladiator sandals' FROM glossary WHERE type = 'clothes' AND value = 'gladiator sandals';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Thong sandals' FROM glossary WHERE type = 'clothes' AND value = 'thong sandals';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Slide sandals' FROM glossary WHERE type = 'clothes' AND value = 'slide sandals';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Espadrille sandals' FROM glossary WHERE type = 'clothes' AND value = 'espadrille sandals';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Wedge sandals' FROM glossary WHERE type = 'clothes' AND value = 'wedge sandals';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Platform sandals' FROM glossary WHERE type = 'clothes' AND value = 'platform sandals';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Ankle boots' FROM glossary WHERE type = 'clothes' AND value = 'ankle boots';

INSERT INTO glossary_localization (glossary_id, local, name) 
SELECT id, 'en', 'Knee-high boots' FROM glossary WHERE type = 'clothes' AND value = 'knee-high boots';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Over-the-knee boots' FROM glossary WHERE type = 'clothes' AND value = 'over-the-knee boots';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Thigh-high boots' FROM glossary WHERE type = 'clothes' AND value = 'thigh-high boots';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Wellington boots' FROM glossary WHERE type = 'clothes' AND value = 'wellington boots';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Chukka boots' FROM glossary WHERE type = 'clothes' AND value = 'chukka boots';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Desert boots' FROM glossary WHERE type = 'clothes' AND value = 'desert boots';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Hiking boots' FROM glossary WHERE type = 'clothes' AND value = 'hiking boots';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Work boots' FROM glossary WHERE type = 'clothes' AND value = 'work boots';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Snow boots' FROM glossary WHERE type = 'clothes' AND value = 'snow boots';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Cowboy boots' FROM glossary WHERE type = 'clothes' AND value = 'cowboy boots';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Biker boots' FROM glossary WHERE type = 'clothes' AND value = 'biker boots';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Duck boots' FROM glossary WHERE type = 'clothes' AND value = 'duck boots';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Military boots' FROM glossary WHERE type = 'clothes' AND value = 'military boots';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Western boots' FROM glossary WHERE type = 'clothes' AND value = 'western boots';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Ankle strap heels' FROM glossary WHERE type = 'clothes' AND value = 'ankle strap heels';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Chunky heels' FROM glossary WHERE type = 'clothes' AND value = 'chunky heels';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Cone heels' FROM glossary WHERE type = 'clothes' AND value = 'cone heels';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Slingback heels' FROM glossary WHERE type = 'clothes' AND value = 'slingback heels';

-- Локализации для environment
INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Sunshine from window' FROM glossary WHERE type = 'environment' AND value = 'sunshine from window';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Neon night, city' FROM glossary WHERE type = 'environment' AND value = 'neon night, city';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Sunset over sea' FROM glossary WHERE type = 'environment' AND value = 'sunset over sea';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Golden time' FROM glossary WHERE type = 'environment' AND value = 'golden time';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Sci-fi RGB glowing, cyberpunk' FROM glossary WHERE type = 'environment' AND value = 'sci-fi RGB glowing, cyberpunk';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Natural lighting' FROM glossary WHERE type = 'environment' AND value = 'natural lighting';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Warm atmosphere, at home, bedroom' FROM glossary WHERE type = 'environment' AND value = 'warm atmosphere, at home, bedroom';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Magic lit' FROM glossary WHERE type = 'environment' AND value = 'magic lit';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Evil, gothic, in a cave' FROM glossary WHERE type = 'environment' AND value = 'evil, gothic, in a cave';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Light and shadow' FROM glossary WHERE type = 'environment' AND value = 'light and shadow';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Shadow from window' FROM glossary WHERE type = 'environment' AND value = 'shadow from window';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Soft studio lighting' FROM glossary WHERE type = 'environment' AND value = 'soft studio lighting';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Home atmosphere, cozy bedroom illumination' FROM glossary WHERE type = 'environment' AND value = 'home atmosphere, cozy bedroom illumination';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Neon, Wong Kar-wai, warm' FROM glossary WHERE type = 'environment' AND value = 'neon, Wong Kar-wai, warm';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Moonlight through curtains' FROM glossary WHERE type = 'environment' AND value = 'moonlight through curtains';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Stormy sky lighting' FROM glossary WHERE type = 'environment' AND value = 'stormy sky lighting';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Foggy forest at dawn' FROM glossary WHERE type = 'environment' AND value = 'foggy forest at dawn';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Golden hour in a meadow' FROM glossary WHERE type = 'environment' AND value = 'golden hour in a meadow';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Rainbow reflections, neon' FROM glossary WHERE type = 'environment' AND value = 'rainbow reflections, neon';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Cozy candlelight' FROM glossary WHERE type = 'environment' AND value = 'cozy candlelight';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Apocalyptic, smoky atmosphere' FROM glossary WHERE type = 'environment' AND value = 'apocalyptic, smoky atmosphere';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Red glow, emergency lights' FROM glossary WHERE type = 'environment' AND value = 'red glow, emergency lights';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Mystical glow, enchanted forest' FROM glossary WHERE type = 'environment' AND value = 'mystical glow, enchanted forest';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Campfire light' FROM glossary WHERE type = 'environment' AND value = 'campfire light';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Harsh, industrial lighting' FROM glossary WHERE type = 'environment' AND value = 'harsh, industrial lighting';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Sunrise in the mountains' FROM glossary WHERE type = 'environment' AND value = 'sunrise in the mountains';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Evening glow in the desert' FROM glossary WHERE type = 'environment' AND value = 'evening glow in the desert';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Moonlight in a dark alley' FROM glossary WHERE type = 'environment' AND value = 'moonlight in a dark alley';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Golden glow at a fairground' FROM glossary WHERE type = 'environment' AND value = 'golden glow at a fairground';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Midnight in the forest' FROM glossary WHERE type = 'environment' AND value = 'midnight in the forest';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Purple and pink hues at twilight' FROM glossary WHERE type = 'environment' AND value = 'purple and pink hues at twilight';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Foggy morning, muted light' FROM glossary WHERE type = 'environment' AND value = 'foggy morning, muted light';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Candle-lit room, rustic vibe' FROM glossary WHERE type = 'environment' AND value = 'candle-lit room, rustic vibe';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Fluorescent office lighting' FROM glossary WHERE type = 'environment' AND value = 'fluorescent office lighting';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Lightning flash in storm' FROM glossary WHERE type = 'environment' AND value = 'lightning flash in storm';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Night, cozy warm light from fireplace' FROM glossary WHERE type = 'environment' AND value = 'night, cozy warm light from fireplace';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Ethereal glow, magical forest' FROM glossary WHERE type = 'environment' AND value = 'ethereal glow, magical forest';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Dusky evening on a beach' FROM glossary WHERE type = 'environment' AND value = 'dusky evening on a beach';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Afternoon light filtering through trees' FROM glossary WHERE type = 'environment' AND value = 'afternoon light filtering through trees';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Blue neon light, urban street' FROM glossary WHERE type = 'environment' AND value = 'blue neon light, urban street';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Red and blue police lights in rain' FROM glossary WHERE type = 'environment' AND value = 'red and blue police lights in rain';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Aurora borealis glow, arctic landscape' FROM glossary WHERE type = 'environment' AND value = 'aurora borealis glow, arctic landscape';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Sunrise through foggy mountains' FROM glossary WHERE type = 'environment' AND value = 'sunrise through foggy mountains';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Golden hour on a city skyline' FROM glossary WHERE type = 'environment' AND value = 'golden hour on a city skyline';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Mysterious twilight, heavy mist' FROM glossary WHERE type = 'environment' AND value = 'mysterious twilight, heavy mist';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Early morning rays, forest clearing' FROM glossary WHERE type = 'environment' AND value = 'early morning rays, forest clearing';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Colorful lantern light at festival' FROM glossary WHERE type = 'environment' AND value = 'colorful lantern light at festival';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Soft glow through stained glass' FROM glossary WHERE type = 'environment' AND value = 'soft glow through stained glass';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Harsh spotlight in dark room' FROM glossary WHERE type = 'environment' AND value = 'harsh spotlight in dark room';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Mellow evening glow on a lake' FROM glossary WHERE type = 'environment' AND value = 'mellow evening glow on a lake';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Crystal reflections in a cave' FROM glossary WHERE type = 'environment' AND value = 'crystal reflections in a cave';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Vibrant autumn lighting in a forest' FROM glossary WHERE type = 'environment' AND value = 'vibrant autumn lighting in a forest';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Gentle snowfall at dusk' FROM glossary WHERE type = 'environment' AND value = 'gentle snowfall at dusk';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Hazy light of a winter morning' FROM glossary WHERE type = 'environment' AND value = 'hazy light of a winter morning';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Soft, diffused foggy glow' FROM glossary WHERE type = 'environment' AND value = 'soft, diffused foggy glow';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Underwater luminescence' FROM glossary WHERE type = 'environment' AND value = 'underwater luminescence';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Rain-soaked reflections in city lights' FROM glossary WHERE type = 'environment' AND value = 'rain-soaked reflections in city lights';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Golden sunlight streaming through trees' FROM glossary WHERE type = 'environment' AND value = 'golden sunlight streaming through trees';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Fireflies lighting up a summer night' FROM glossary WHERE type = 'environment' AND value = 'fireflies lighting up a summer night';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Glowing embers from a forge' FROM glossary WHERE type = 'environment' AND value = 'glowing embers from a forge';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Dim candlelight in a gothic castle' FROM glossary WHERE type = 'environment' AND value = 'dim candlelight in a gothic castle';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Midnight sky with bright starlight' FROM glossary WHERE type = 'environment' AND value = 'midnight sky with bright starlight';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Warm sunset in a rural village' FROM glossary WHERE type = 'environment' AND value = 'warm sunset in a rural village';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Flickering light in a haunted house' FROM glossary WHERE type = 'environment' AND value = 'flickering light in a haunted house';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Desert sunset with mirage-like glow' FROM glossary WHERE type = 'environment' AND value = 'desert sunset with mirage-like glow';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Golden beams piercing through storm clouds' FROM glossary WHERE type = 'environment' AND value = 'golden beams piercing through storm clouds';

-- Локализации для background
INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A cozy bed and a lamp' FROM glossary WHERE type = 'background' AND value = 'a cozy bed and a lamp';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A forest clearing with mist' FROM glossary WHERE type = 'background' AND value = 'a forest clearing with mist';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A bustling marketplace' FROM glossary WHERE type = 'background' AND value = 'a bustling marketplace';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A quiet beach at dusk' FROM glossary WHERE type = 'background' AND value = 'a quiet beach at dusk';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'An old, cobblestone street' FROM glossary WHERE type = 'background' AND value = 'an old, cobblestone street';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A futuristic cityscape' FROM glossary WHERE type = 'background' AND value = 'a futuristic cityscape';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A tranquil lake with mountains' FROM glossary WHERE type = 'background' AND value = 'a tranquil lake with mountains';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A mysterious cave entrance' FROM glossary WHERE type = 'background' AND value = 'a mysterious cave entrance';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Bookshelves and plants in the background' FROM glossary WHERE type = 'background' AND value = 'bookshelves and plants in the background';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'An ancient temple in ruins' FROM glossary WHERE type = 'background' AND value = 'an ancient temple in ruins';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Tall skyscrapers and neon signs' FROM glossary WHERE type = 'background' AND value = 'tall skyscrapers and neon signs';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A starry sky over a desert' FROM glossary WHERE type = 'background' AND value = 'a starry sky over a desert';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A bustling café' FROM glossary WHERE type = 'background' AND value = 'a bustling café';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Rolling hills and farmland' FROM glossary WHERE type = 'background' AND value = 'rolling hills and farmland';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A modern living room with a fireplace' FROM glossary WHERE type = 'background' AND value = 'a modern living room with a fireplace';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'An abandoned warehouse' FROM glossary WHERE type = 'background' AND value = 'an abandoned warehouse';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A picturesque mountain range' FROM glossary WHERE type = 'background' AND value = 'a picturesque mountain range';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A starry night sky' FROM glossary WHERE type = 'background' AND value = 'a starry night sky';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'The interior of a futuristic spaceship' FROM glossary WHERE type = 'background' AND value = 'the interior of a futuristic spaceship';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'The cluttered workshop of an inventor' FROM glossary WHERE type = 'background' AND value = 'the cluttered workshop of an inventor';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'The glowing embers of a bonfire' FROM glossary WHERE type = 'background' AND value = 'the glowing embers of a bonfire';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A misty lake surrounded by trees' FROM glossary WHERE type = 'background' AND value = 'a misty lake surrounded by trees';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'An ornate palace hall' FROM glossary WHERE type = 'background' AND value = 'an ornate palace hall';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A busy street market' FROM glossary WHERE type = 'background' AND value = 'a busy street market';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A vast desert landscape' FROM glossary WHERE type = 'background' AND value = 'a vast desert landscape';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A peaceful library corner' FROM glossary WHERE type = 'background' AND value = 'a peaceful library corner';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Bustling train station' FROM glossary WHERE type = 'background' AND value = 'bustling train station';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A mystical, enchanted forest' FROM glossary WHERE type = 'background' AND value = 'a mystical, enchanted forest';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'An underwater reef with colorful fish' FROM glossary WHERE type = 'background' AND value = 'an underwater reef with colorful fish';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A quiet rural village' FROM glossary WHERE type = 'background' AND value = 'a quiet rural village';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A sandy beach with palm trees' FROM glossary WHERE type = 'background' AND value = 'a sandy beach with palm trees';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A vibrant coral reef, teeming with life' FROM glossary WHERE type = 'background' AND value = 'a vibrant coral reef, teeming with life';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Snow-capped mountains in distance' FROM glossary WHERE type = 'background' AND value = 'snow-capped mountains in distance';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A stormy ocean, waves crashing' FROM glossary WHERE type = 'background' AND value = 'a stormy ocean, waves crashing';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A rustic barn in open fields' FROM glossary WHERE type = 'background' AND value = 'a rustic barn in open fields';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A futuristic lab with glowing screens' FROM glossary WHERE type = 'background' AND value = 'a futuristic lab with glowing screens';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A dark, abandoned castle' FROM glossary WHERE type = 'background' AND value = 'a dark, abandoned castle';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'The ruins of an ancient civilization' FROM glossary WHERE type = 'background' AND value = 'the ruins of an ancient civilization';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A bustling urban street in rain' FROM glossary WHERE type = 'background' AND value = 'a bustling urban street in rain';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'An elegant grand ballroom' FROM glossary WHERE type = 'background' AND value = 'an elegant grand ballroom';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A sprawling field of wildflowers' FROM glossary WHERE type = 'background' AND value = 'a sprawling field of wildflowers';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A dense jungle with sunlight filtering through' FROM glossary WHERE type = 'background' AND value = 'a dense jungle with sunlight filtering through';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A dimly lit, vintage bar' FROM glossary WHERE type = 'background' AND value = 'a dimly lit, vintage bar';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'An ice cave with sparkling crystals' FROM glossary WHERE type = 'background' AND value = 'an ice cave with sparkling crystals';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A serene riverbank at sunset' FROM glossary WHERE type = 'background' AND value = 'a serene riverbank at sunset';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A narrow alley with graffiti walls' FROM glossary WHERE type = 'background' AND value = 'a narrow alley with graffiti walls';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A peaceful zen garden with koi pond' FROM glossary WHERE type = 'background' AND value = 'a peaceful zen garden with koi pond';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A high-tech control room' FROM glossary WHERE type = 'background' AND value = 'a high-tech control room';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A quiet mountain village at dawn' FROM glossary WHERE type = 'background' AND value = 'a quiet mountain village at dawn';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A lighthouse on a rocky coast' FROM glossary WHERE type = 'background' AND value = 'a lighthouse on a rocky coast';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A rainy street with flickering lights' FROM glossary WHERE type = 'background' AND value = 'a rainy street with flickering lights';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A frozen lake with ice formations' FROM glossary WHERE type = 'background' AND value = 'a frozen lake with ice formations';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'An abandoned theme park' FROM glossary WHERE type = 'background' AND value = 'an abandoned theme park';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A small fishing village on a pier' FROM glossary WHERE type = 'background' AND value = 'a small fishing village on a pier';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Rolling sand dunes in a desert' FROM glossary WHERE type = 'background' AND value = 'rolling sand dunes in a desert';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A dense forest with towering redwoods' FROM glossary WHERE type = 'background' AND value = 'a dense forest with towering redwoods';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A snowy cabin in the mountains' FROM glossary WHERE type = 'background' AND value = 'a snowy cabin in the mountains';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A mystical cave with bioluminescent plants' FROM glossary WHERE type = 'background' AND value = 'a mystical cave with bioluminescent plants';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A castle courtyard under moonlight' FROM glossary WHERE type = 'background' AND value = 'a castle courtyard under moonlight';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A bustling open-air night market' FROM glossary WHERE type = 'background' AND value = 'a bustling open-air night market';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'An old train station with steam' FROM glossary WHERE type = 'background' AND value = 'an old train station with steam';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A tranquil waterfall surrounded by trees' FROM glossary WHERE type = 'background' AND value = 'a tranquil waterfall surrounded by trees';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A vineyard in the countryside' FROM glossary WHERE type = 'background' AND value = 'a vineyard in the countryside';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A quaint medieval village' FROM glossary WHERE type = 'background' AND value = 'a quaint medieval village';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A bustling harbor with boats' FROM glossary WHERE type = 'background' AND value = 'a bustling harbor with boats';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A high-tech futuristic mall' FROM glossary WHERE type = 'background' AND value = 'a high-tech futuristic mall';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'A lush tropical rainforest' FROM glossary WHERE type = 'background' AND value = 'a lush tropical rainforest';

-- Локализации для nsfw
INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Nude' FROM glossary WHERE type = 'nsfw' AND value = 'nude';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Breast' FROM glossary WHERE type = 'nsfw' AND value = 'breast';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Clothes lift' FROM glossary WHERE type = 'nsfw' AND value = 'clothes lift';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Pussy juice trail' FROM glossary WHERE type = 'nsfw' AND value = 'pussy juice trail';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Pussy juice puddle' FROM glossary WHERE type = 'nsfw' AND value = 'pussy juice puddle';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Cum on body' FROM glossary WHERE type = 'nsfw' AND value = 'cum on body';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Cum inside' FROM glossary WHERE type = 'nsfw' AND value = 'cum inside';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Fingering' FROM glossary WHERE type = 'nsfw' AND value = 'fingering';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Doggystyle' FROM glossary WHERE type = 'nsfw' AND value = 'doggystyle';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Cowgirl' FROM glossary WHERE type = 'nsfw' AND value = 'cowgirl';

INSERT INTO glossary_localization (glossary_id, local, name)
SELECT id, 'en', 'Reversed cowgirl' FROM glossary WHERE type = 'nsfw' AND value = 'reversed cowgirl';

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
-- Удаление данных в обратном порядке (сначала локализации, затем глоссарий)
DELETE FROM glossary_localization;
DELETE FROM glossary;
-- +goose StatementEnd