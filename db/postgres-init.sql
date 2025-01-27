CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    timestamp TIMESTAMP NOT NULL,
    region VARCHAR(50) NOT NULL,
    amount DECIMAL(20, 2) NOT NULL,
    status VARCHAR(20),
    flagged BOOLEAN DEFAULT FALSE
);
