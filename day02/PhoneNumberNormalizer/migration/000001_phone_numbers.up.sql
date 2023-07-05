CREATE TABLE IF NOT EXISTS phonebook (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    phonenumber VARCHAR(100)
);

INSERT INTO phonebook(phonenumber) VALUES ('1234567890');
INSERT INTO phonebook(phonenumber) VALUES ('123 456 7891');
INSERT INTO phonebook(phonenumber) VALUES ('(123) 456 7892');
INSERT INTO phonebook(phonenumber) VALUES ('(123) 456-7893');
INSERT INTO phonebook(phonenumber) VALUES ('123-456-7894');
INSERT INTO phonebook(phonenumber) VALUES ('123-456-7890');
INSERT INTO phonebook(phonenumber) VALUES ('1234567892');
INSERT INTO phonebook(phonenumber) VALUES ('(123)456-7892');
INSERT INTO phonebook(phonenumber) VALUES ('аа98741привет)');