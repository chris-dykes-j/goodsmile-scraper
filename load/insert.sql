    INSERT INTO nendoroid (nendoroid_number) VALUES ($1) RETURNING id;

    INSERT INTO nendoroid_name (nendoroid_id, language_code, text) VALUES (id, 'en', $2);
    INSERT INTO nendoroid_name (nendoroid_id, language_code, text) VALUES (id, 'ja', $3);
    INSERT INTO nendoroid_name (nendoroid_id, language_code, text) VALUES (id, 'zh', $4);

    INSERT INTO nendoroid_description (nendoroid_id, language_code, text) VALUES (id, 'en', $5);
    INSERT INTO nendoroid_description (nendoroid_id, language_code, text) VALUES (id, 'ja', $6);
    INSERT INTO nendoroid_description (nendoroid_id, language_code, text) VALUES (id, 'zh', $7);

    INSERT INTO nendoroid_details (nendoroid_id, language_code, details) VALUES (id, 'en', $8);
    INSERT INTO nendoroid_details (nendoroid_id, language_code, details) VALUES (id, 'ja', $9);
    INSERT INTO nendoroid_details (nendoroid_id, language_code, details) VALUES (id, 'zh', $10);

    INSERT INTO nendoroid_link (nendoroid_id, language_code, text) VALUES (id, 'en', $11);
    INSERT INTO nendoroid_link (nendoroid_id, language_code, text) VALUES (id, 'ja', $12);
    INSERT INTO nendoroid_link (nendoroid_id, language_code, text) VALUES (id, 'zh', $13);

    INSERT INTO nendoroid_blog_link (nendoroid_id, language_code, text) VALUES (id, 'en', $14);
    INSERT INTO nendoroid_blog_link (nendoroid_id, language_code, text) VALUES (id, 'ja', $15);
    INSERT INTO nendoroid_blog_link (nendoroid_id, language_code, text) VALUES (id, 'zh', $16);
