DROP TABLE IF EXISTS tblPaymentMethod CASCADE;

CREATE TABLE tblPaymentMethod (
    payment_method_id   SERIAL PRIMARY KEY,
    payment_method      VARCHAR(25)
);

INSERT INTO tblPaymentMethod (payment_method) VALUES ('Cash');
INSERT INTO tblPaymentMethod (payment_method) VALUES ('Debit');
INSERT INTO tblPaymentMethod (payment_method) VALUES ('Visa');
INSERT INTO tblPaymentMethod (payment_method) VALUES ('Mastercard');
INSERT INTO tblPaymentMethod (payment_method) VALUES ('American Express');