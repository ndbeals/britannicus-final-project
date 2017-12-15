DROP TABLE IF EXISTS tblPaymentMethod CASCADE;

CREATE TABLE tblPaymentMethod (
    payment_method_id   SERIAL PRIMARY KEY,
    payment_method      VARCHAR(25)
);

INSERT INTO tblPaymentMethod (payment_method_id, payment_method) VALUES (1, 'Cash');
INSERT INTO tblPaymentMethod (payment_method_id, payment_method) VALUES (2, 'Debit');
INSERT INTO tblPaymentMethod (payment_method_id, payment_method) VALUES (3, 'Visa');
INSERT INTO tblPaymentMethod (payment_method_id, payment_method) VALUES (4, 'Mastercard');
INSERT INTO tblPaymentMethod (payment_method_id, payment_method) VALUES (5, 'American Express');