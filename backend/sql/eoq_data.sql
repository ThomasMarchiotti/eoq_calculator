CREATE TABLE eoq_data (
    id INTEGER PRIMARY KEY,
    demand REAL,
    setup_cost REAL,
    holding_cost REAL
);

INSERT INTO eoq_data (demand, setup_cost, holding_cost) VALUES (1000, 50, 5);
INSERT INTO eoq_data (demand, setup_cost, holding_cost) VALUES (2000, 60, 6);
INSERT INTO eoq_data (demand, setup_cost, holding_cost) VALUES (3000, 70, 7);