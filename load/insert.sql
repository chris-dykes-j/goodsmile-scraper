CREATE OR REPLACE FUNCTION insert_nendo
(
    item_num int,
    name_en varchar(255),
    name_ja varchar(255),
    name_zh varchar(255),
    desc_en text,
    desc_ja text,
    desc_zh text,
    details_en JSONB,
    details_ja JSONB,
    details_zh JSONB,
    link_en text,
    link_ja text,
    link_zh text,
    blog_en text,
    blog_ja text, 
    blog_zh text
)
RETURNS VOID AS $$
DECLARE _id INTEGER;
BEGIN
    INSERT INTO nendoroid (item_number) VALUES (item_num) RETURNING id INTO _id;

    INSERT INTO nendoroid_data (nendoroid_id, language_code, name, description, item_link, blog_link, details)
    VALUES (_id, 'en', name_en, desc_en, link_en, blog_en, details_en);
    
    INSERT INTO nendoroid_data (nendoroid_id, language_code, name, description, item_link, blog_link, details)
    VALUES (_id, 'ja', name_ja, desc_ja, link_ja, blog_ja, details_ja);
    
    INSERT INTO nendoroid_data (nendoroid_id, language_code, name, description, item_link, blog_link, details)
    VALUES (_id, 'zh', name_zh, desc_zh, link_zh, blog_zh, details_zh);

END;
$$ LANGUAGE plpgsql;
