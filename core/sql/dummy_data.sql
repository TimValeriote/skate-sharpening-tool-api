INSERT INTO users(first_name, last_name, email, phone_number, uuid, is_staff) VALUES
('Tim','Valeriote','tv91@rogers.com','tim-uuid','5195900620',1),
('Clayton','Peterek','c.peterek93@gmail.com','clay-uuid','5199986136',1);

INSERT INTO staff(user_id) VALUES
((SELECT id FROM users WHERE email = 'tv91@rogers.com')),
((SELECT id FROM users WHERE email = 'c.peterek93@gmail.com'));

INSERT INTO brands(name, short_name, is_skate, is_steel, is_holder) VALUES
('CCM', 'CCM', 1, 1, 1),
('Bauer', 'Bauer', 1, 1, 0),
('Graf', 'Graf', 1, 0, 1),
('Step Steel', 'Step Steel', 0, 1, 0),
('Tuuk', 'Tuuk', 0, 0, 1);

INSERT INTO model(name, alias) VALUES
('Tacks', 'AS-V'),
('Tacks', 'AS 590'),
('Jetspeed', 'FT4'),
('Jetspeed', 'FT485'),
('Ribcor', '100K'),
('Ribcor', '86K'),
('Vapor', 'Hyperlite'),
('Vapor', '3X'),
('Supreme', 'MACH'),
('Supreme', 'M5'),
('Bauer X', 'X-LP'),
('PeakSpeed', 'PK7900'),
('PeakSpeed', 'PK7700');

INSERT INTO fits(name) VALUES('Tapered'),('Flexible'),('Wide'),('FIT 1'),('FIT 2');

INSERT INTO skates(model_id, brand_id) VALUES
((SELECT id FROM model WHERE name = 'Tacks' AND alias = 'AS-V'),(SELECT id FROM brands WHERE short_name = 'CCM')),
((SELECT id FROM model WHERE name = 'Tacks' AND alias = 'AS 590'),(SELECT id FROM brands WHERE short_name = 'CCM')),
((SELECT id FROM model WHERE name = 'Jetspeed' AND alias = 'FT4'),(SELECT id FROM brands WHERE short_name = 'CCM')),
((SELECT id FROM model WHERE name = 'Ribcor' AND alias = '100K'),(SELECT id FROM brands WHERE short_name = 'CCM')),
((SELECT id FROM model WHERE name = 'Ribcor' AND alias = '86K'),(SELECT id FROM brands WHERE short_name = 'CCM')),
((SELECT id FROM model WHERE name = 'Vapor' AND alias = 'Hyperlite'),(SELECT id FROM brands WHERE short_name = 'Bauer')),
((SELECT id FROM model WHERE name = 'Vapor' AND alias = '3X'),(SELECT id FROM brands WHERE short_name = 'Bauer')),
((SELECT id FROM model WHERE name = 'Bauer X' AND alias = 'X-LP'),(SELECT id FROM brands WHERE short_name = 'Bauer')),
((SELECT id FROM model WHERE name = 'Supreme' AND alias = 'M5'),(SELECT id FROM brands WHERE short_name = 'Bauer')),
((SELECT id FROM model WHERE name = 'PeakSpeed' AND alias = 'PK7900'),(SELECT id FROM brands WHERE short_name = 'Graf')),
((SELECT id FROM model WHERE name = 'PeakSpeed' AND alias = 'PK7700'),(SELECT id FROM brands WHERE short_name = 'Graf'));

INSERT INTO colour(colour) VALUES ('red'),('green'),('black'),('white');

INSERT INTO store(name, address, city, country, phone_number, store_number) VALUES
('Kitchener','123 Kitchener Street N2N 2N2','Kitchener','Canada','5195555555', 6044),
('Cambridge','456 Cambridge Street N3N 3N3','Cambridge','Canada','5196666666', 42069),
('Waterloo','789 Waterloo Street N4N 4N4','Waterloo','Canada','5197777777', 69699);

INSERT INTO user_skates(user_id, skate_id, holder_brand_id, holder_size, skate_size, lace_colour_id, has_steel, steel_id, has_guards, guard_colour_id, preferred_radius, fit_id) VALUES
(1, (SELECT id FROM skates WHERE id = 1), (SELECT id FROM brands WHERE short_name = 'CCM'), 9, 8.5, (SELECT id FROM colour WHERE colour = 'white'), 1, (SELECT id FROM brands WHERE short_name = 'CCM'), 0, null, '7/8', 1),
(1, (SELECT id FROM skates WHERE id = 2), (SELECT id FROM brands WHERE short_name = 'Tuuk'), 7, 8, (SELECT id FROM colour WHERE colour = 'black'), 1, (SELECT id FROM brands WHERE short_name = 'CCM'), 1, (SELECT id FROM colour WHERE colour = 'white'), '5/8', 2),
(1, (SELECT id FROM skates WHERE id = 10), (SELECT id FROM brands WHERE short_name = 'Graf'), 6.5, 6.5, (SELECT id FROM colour WHERE colour = 'white'), 1, (SELECT id FROM brands WHERE short_name = 'Step Steel'), 0, null, '1/3', 3),
(1, (SELECT id FROM skates WHERE id = 4), (SELECT id FROM brands WHERE short_name = 'CCM'), 8, 8.5, (SELECT id FROM colour WHERE colour = 'black'), 1, (SELECT id FROM brands WHERE short_name = 'Step Steel'), 0, null, '1 1/4', 1),
(2, (SELECT id FROM skates WHERE id = 1), (SELECT id FROM brands WHERE short_name = 'CCM'), 9, 9.5, (SELECT id FROM colour WHERE colour = 'green'), 1, (SELECT id FROM brands WHERE short_name = 'Step Steel'), 1, (SELECT id FROM colour WHERE colour = 'green'), '1', 3),
(2, (SELECT id FROM skates WHERE id = 5), (SELECT id FROM brands WHERE short_name = 'Tuuk'), 10, 10.5, (SELECT id FROM colour WHERE colour = 'black'), 1, (SELECT id FROM brands WHERE short_name = 'Step Steel'), 0, null, '3/4', 5),
(2, (SELECT id FROM skates WHERE id = 6), (SELECT id FROM brands WHERE short_name = 'Tuuk'), 4, 3.5, (SELECT id FROM colour WHERE colour = 'white'), 1, (SELECT id FROM brands WHERE short_name = 'Bauer'), 0, null, '1/2', 4),
(2, (SELECT id FROM skates WHERE id = 7), (SELECT id FROM brands WHERE short_name = 'Tuuk'), 5, 5, (SELECT id FROM colour WHERE colour = 'green'), 1, (SELECT id FROM brands WHERE short_name = 'Bauer'), 0, null, '7/8', 1);


