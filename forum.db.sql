BEGIN TRANSACTION;
DROP TABLE IF EXISTS "posts";
CREATE TABLE "posts" (
	"post_id"	TEXT NOT NULL,
	"user_id"	TEXT NOT NULL,
	"title"	TEXT NOT NULL CHECK(length('title') >= 3),
	"body"	TEXT NOT NULL CHECK(length('body') >= 3),
	"created_at"	INTEGER DEFAULT (strftime('%s', 'now')),
	PRIMARY KEY("post_id"),
	FOREIGN KEY("user_id") REFERENCES "users"("user_id")
);
DROP TABLE IF EXISTS "post_likes";
CREATE TABLE "post_likes" (
	"post_id"	TEXT NOT NULL,
	"user_id"	TEXT NOT NULL,
	"up"	INTEGER DEFAULT 1 CHECK("up" = 0 OR "up" = 1),
	FOREIGN KEY("post_id") REFERENCES "posts"("post_id"),
	FOREIGN KEY("user_id") REFERENCES "users"("user_id")
);
DROP TABLE IF EXISTS "comment_likes";
CREATE TABLE "comment_likes" (
	"comment_id"	TEXT NOT NULL,
	"user_id"	TEXT NOT NULL,
	"up"	INTEGER DEFAULT 1 CHECK("up" = 0 OR "up" = 1),
	FOREIGN KEY("comment_id") REFERENCES "comments"("comment_id"),
	FOREIGN KEY("user_id") REFERENCES "users"("user_id")
);
DROP TABLE IF EXISTS "comments";
CREATE TABLE "comments" (
	"comment_id"	TEXT NOT NULL UNIQUE,
	"user_id"	TEXT NOT NULL,
	"post_id"	TEXT NOT NULL,
	"body"	TEXT NOT NULL CHECK(length("body") >= 2),
	"created_at"	INTEGER DEFAULT (strftime('%s', 'now')),
	PRIMARY KEY("comment_id"),
	FOREIGN KEY("post_id") REFERENCES "posts"("post_id"),
	FOREIGN KEY("user_id") REFERENCES "users"("user_id")
);
DROP TABLE IF EXISTS "cats";
CREATE TABLE "cats" (
	"cat_id"	TEXT NOT NULL,
	"title"	TEXT NOT NULL CHECK(length('title') >= 3),
	"description"	TEXT CHECK(length('title') >= 3),
	PRIMARY KEY("cat_id")
);
DROP TABLE IF EXISTS "post_cats";
CREATE TABLE "post_cats" (
	"post_id"	TEXT NOT NULL,
	"cat_id"	TEXT NOT NULL,
	FOREIGN KEY("post_id") REFERENCES "posts"("post_id")
	FOREIGN KEY("cat_id") REFERENCES "cats"("cat_id")
);
DROP TABLE IF EXISTS "users";
CREATE TABLE "users" (
	"user_id"	TEXT NOT NULL,
	"username"	TEXT NOT NULL CHECK(length("username") >= 3) UNIQUE,
	"email"	TEXT NOT NULL CHECK("email" LIKE '%_@__%.__%') UNIQUE,
	"password"	TEXT NOT NULL CHECK(length('password') >= 3),
	"avatar"	TEXT DEFAULT '',
	"created_at"	INTEGER DEFAULT (strftime('%s', 'now')),
	PRIMARY KEY("user_id")
);
-- users
INSERT INTO "users" VALUES ('73fa04f7-18c1-4211-8e03-3e9357c49ee5','jonsnow','jonsnow@gmail.com','$2a$10$kUHK4BCc2u4tTLw3q75h6OV7tneonMD0SjI8gGNzzyIaq7MD6R8AW','',1611934585);
INSERT INTO "users" VALUES ('427d6f6c-4be9-4774-813a-6b90af3a50fd','ygritte','ygritte@gmail.com','123','',1611935626);
INSERT INTO "users" VALUES ('41fd1865-0176-41c9-9f60-792f77e12470','tyrion','tyrion@gmail.com','123','',1611941957);

-- cats
INSERT INTO "cats" VALUES ('6fa18344-7698-4b1d-a5d7-ac07fd83a068','💡Your suggestions/ideas💡','Share your opinion, experience, view, point of view, blah blah blah. Improvement suggestions');
INSERT INTO "cats" VALUES ('d92b1461-799d-450f-b741-958d79d33156','Music 🎵🎸','Mumble rap is bs. Change my mind');
INSERT INTO "cats" VALUES ('ac6115dc-ed29-4a31-8956-4fd6adf35844','Games 🎯🎮','Everything about games');
INSERT INTO "cats" VALUES ('97819604-6d06-4aaf-a3b7-7c800c80e645','Movies 🎥🎬','Game of Thrones and others');
INSERT INTO "cats" VALUES ('f97c0fbb-832d-4f93-9c13-126958ae3ae3','Memes 😆💬','You dont need the description. Just post your fav stuff');
INSERT INTO "cats" VALUES ('1a15719d-784f-479b-869f-bf1a1a6c5a91','Comics 💭🦹‍♀️','Batman v Spider-Man and all that stuff');
INSERT INTO "cats" VALUES ('c5c98e31-2812-41a7-9c63-46f1ca3270ef','Politics 🤵💩','Politics? why would you allow politics in forum??');
INSERT INTO "cats" VALUES ('ebddb84d-9d53-4717-a8c3-3c95790a1c2a','News 📰🗞','News. Just news.');
INSERT INTO "cats" VALUES ('ec8a9ee8-e514-4526-a8cf-727e28f69e23','Science 🔬👩‍🔬','Just dont mention corona vaccines, please');

-- posts
INSERT INTO "posts" VALUES ('48b00ce8-01a0-4a8f-bf5c-02d4d9e50278','41fd1865-0176-41c9-9f60-792f77e12470','Ghost👻 of Tsushima🎌 is great!','Ghost of Tsushima is a five star game, I think everyone should try it',1611942145);
INSERT INTO "posts" VALUES ('29da509b-d87f-4d30-8ac4-204a32f55180','73fa04f7-18c1-4211-8e03-3e9357c49ee5','Snowpiercer❄⛄ is a sequel to Willy Wonka!','Snowpiercer (2013) is a sequel to Willy Wonka and the Chocolate Factory (1971) some are saying. Really? Bong Joon-ho''s (who recently won an Oscar for Parasite) post-apocalyptic action movie is a sequel to the beloved children''s classic? Actually, it may not be as crazy as it sounds.',1611942360);
INSERT INTO "posts" VALUES ('fcb3378c-8600-4a09-be13-b6d48fc729d5','427d6f6c-4be9-4774-813a-6b90af3a50fd','Lil Wayne🐐 on his ting again','Fresh off securing a pardon in the closing hours of the Trump administration, Lil Wayne released a new song featuring Fousheé titled “Ain’t Got Time” where he details some of the circumstances that led to his arrest.',1611942532);
INSERT INTO "posts" VALUES ('87ea019f-0735-474c-bcf4-de220f05a081','73fa04f7-18c1-4211-8e03-3e9357c49ee5','Music to be Murdered by Side B is 👎🏽👎🏽 garbage','Eminem''s mtbmb side b is just awful. he should just quit music, he''s too old anyway...',1611942808);
INSERT INTO "posts" VALUES ('4212e160-3168-4259-98eb-abfe9ea55ff1','427d6f6c-4be9-4774-813a-6b90af3a50fd','Wanda🦸‍♀️🦸‍♂️Vision episode four: is Scarlet Witch the villain of her own story?','The efficiency of the storytelling is impressive, and all the better for us experiencing it as Monica does: we see her erased in Thanos’s “Snap” from Avengers: Infinity War, and returned five years later by the Hulk’s from Endgame. In her five-year absence, her mother (Maria, who we met in Captain Marvel) has died. It’s a fast-track to empathy and you’re already on Monica’s side by the time we reach SWORD.',1611942901);
INSERT INTO "posts" VALUES ('63e66f01-da14-4eee-98c0-90c2d64bfab9','41fd1865-0176-41c9-9f60-792f77e12470','Everything You Get in Disco Elysium🎮: The Final Cut Edition','Disco Elysium is a unique and stylish isometric RPG from indie studio ZA/UM. After garnering universal acclaim with its well-written story and clever approach to RPG elements, the title went on to win over a dozen awards. Now, just over a year since its launch, players will get an opportunity to return to the rain-damp streets of Disco Elysium''s Revachol.',1611942985);

-- post cats
INSERT INTO "post_cats" VALUES ('48b00ce8-01a0-4a8f-bf5c-02d4d9e50278','ac6115dc-ed29-4a31-8956-4fd6adf35844');
INSERT INTO "post_cats" VALUES ('29da509b-d87f-4d30-8ac4-204a32f55180','97819604-6d06-4aaf-a3b7-7c800c80e645');
INSERT INTO "post_cats" VALUES ('fcb3378c-8600-4a09-be13-b6d48fc729d5','d92b1461-799d-450f-b741-958d79d33156');
INSERT INTO "post_cats" VALUES ('87ea019f-0735-474c-bcf4-de220f05a081','d92b1461-799d-450f-b741-958d79d33156');
INSERT INTO "post_cats" VALUES ('4212e160-3168-4259-98eb-abfe9ea55ff1','97819604-6d06-4aaf-a3b7-7c800c80e645');
INSERT INTO "post_cats" VALUES ('63e66f01-da14-4eee-98c0-90c2d64bfab9','ac6115dc-ed29-4a31-8956-4fd6adf35844');

-- comments
INSERT INTO "comments" VALUES ('ea634f11-94db-4536-aabc-e43c9d748a55','427d6f6c-4be9-4774-813a-6b90af3a50fd','87ea019f-0735-474c-bcf4-de220f05a081','pfff, you know nothing',1611943246);
INSERT INTO "comments" VALUES ('2f9644b4-4bb7-4b16-a88e-d3668c9b3591','73fa04f7-18c1-4211-8e03-3e9357c49ee5','87ea019f-0735-474c-bcf4-de220f05a081','🤔 🤔  ithought you were dead 🤔🤔',1611943482);
INSERT INTO "comments" VALUES ('392917ad-c1a7-4700-a4db-c3c11c8cdb86','41fd1865-0176-41c9-9f60-792f77e12470','fcb3378c-8600-4a09-be13-b6d48fc729d5','wayne the 🐐🐐 !!!!!!',1611943925);

-- likes
INSERT INTO "post_likes" VALUES ('48b00ce8-01a0-4a8f-bf5c-02d4d9e50278','73fa04f7-18c1-4211-8e03-3e9357c49ee5',1);
INSERT INTO "post_likes" VALUES ('29da509b-d87f-4d30-8ac4-204a32f55180','41fd1865-0176-41c9-9f60-792f77e12470',1);
INSERT INTO "post_likes" VALUES ('87ea019f-0735-474c-bcf4-de220f05a081','41fd1865-0176-41c9-9f60-792f77e12470',0);

COMMIT;
