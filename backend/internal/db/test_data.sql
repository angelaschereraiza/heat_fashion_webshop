INSERT INTO categories (id, name) VALUES ('45c614ce-76a5-48cf-8a11-3ee00b48efc4', 'TEST CATEGORY 1');
INSERT INTO categories (id, name) VALUES ('30f997a1-3ed2-4fbc-9b2f-975b3ae0d249', 'TEST CATEGORY 2');
INSERT INTO categories (id, name) VALUES ('44ce29c7-ef1c-40db-9899-19397df590a9', 'TEST CATEGORY 3');

INSERT INTO products (id, category_id, ali_express_id, name, price) VALUES ('ea420fd8-139e-4578-807a-6493caba9231', '45c614ce-76a5-48cf-8a11-3ee00b48efc4', 1, 'TEST', 100);
INSERT INTO products (id, category_id, ali_express_id, name, price) VALUES ('277cb7db-0b3e-461b-834a-465519db28b7', '45c614ce-76a5-48cf-8a11-3ee00b48efc4', 2, 'TEST 2', 50);
INSERT INTO products (id, category_id, ali_express_id, name, price) VALUES ('b962f906-7790-482a-9153-fca90a359300', '30f997a1-3ed2-4fbc-9b2f-975b3ae0d249', 3, 'TEST 3', 30);
INSERT INTO products (id, category_id, ali_express_id, name, price) VALUES ('a8393aee-ddb2-422a-9a0a-845a020ab33a', '30f997a1-3ed2-4fbc-9b2f-975b3ae0d249', 4, 'TEST 4', 50);
INSERT INTO products (id, category_id, ali_express_id, name, price) VALUES ('3a70f69b-8a3d-4737-886c-05e6c272dcbe', '44ce29c7-ef1c-40db-9899-19397df590a9', 5, 'TEST 5', 100);
INSERT INTO products (id, category_id, ali_express_id, name, price) VALUES ('1c527032-6000-4f8d-9350-326a53852034', '44ce29c7-ef1c-40db-9899-19397df590a9', 6, 'TEST 6', 100);