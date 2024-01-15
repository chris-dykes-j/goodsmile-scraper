CREATE OR REPLACE FUNCTION insert_nendoroid
(
    item_num int,
    name_en varchar(255),
    name_jp varchar(255),
    name_zh varchar(255),
    desc_en text,
    desc_jp text,
    desc_zh text,
    details_en JSONB,
    details_ja JSONB,
    details_zh JSONB,
    item_link_en text,
    item_link_ja text,
    item_link_zh text,
    blog_link_en text,
    blog_link_ja text, 
    blog_link_zh text
)
RETURNS VOID AS $$
DECLARE _id INTEGER;
BEGIN
    INSERT INTO nendoroid (item_number) VALUES (item_num) RETURNING id INTO _id;

    INSERT INTO nendoroid_name (nendoroid_id, language_code, text) VALUES (_id, 'en', name_en);
    INSERT INTO nendoroid_name (nendoroid_id, language_code, text) VALUES (_id, 'ja', name_jp);
    INSERT INTO nendoroid_name (nendoroid_id, language_code, text) VALUES (_id, 'zh', name_zh);

    INSERT INTO nendoroid_description (nendoroid_id, language_code, text) VALUES (_id, 'en', desc_en);
    INSERT INTO nendoroid_description (nendoroid_id, language_code, text) VALUES (_id, 'ja', desc_jp);
    INSERT INTO nendoroid_description (nendoroid_id, language_code, text) VALUES (_id, 'zh', desc_zh);

    INSERT INTO nendoroid_details (nendoroid_id, language_code, details) VALUES (_id, 'en', details_en);
    INSERT INTO nendoroid_details (nendoroid_id, language_code, details) VALUES (_id, 'ja', details_ja);
    INSERT INTO nendoroid_details (nendoroid_id, language_code, details) VALUES (_id, 'zh', details_zh);

    INSERT INTO nendoroid_link (nendoroid_id, language_code, text) VALUES (_id, 'en', item_link_en);
    INSERT INTO nendoroid_link (nendoroid_id, language_code, text) VALUES (_id, 'ja', item_link_ja);
    INSERT INTO nendoroid_link (nendoroid_id, language_code, text) VALUES (_id, 'zh', item_link_zh);

    INSERT INTO nendoroid_blog_link (nendoroid_id, language_code, text) VALUES (_id, 'en', blog_link_en);
    INSERT INTO nendoroid_blog_link (nendoroid_id, language_code, text) VALUES (_id, 'ja', blog_link_ja);
    INSERT INTO nendoroid_blog_link (nendoroid_id, language_code, text) VALUES (_id, 'zh', blog_link_zh); 
END;
$$ LANGUAGE plpgsql;
