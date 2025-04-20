CREATE TABLE health_status (                                                   
    service TEXT,
    status_key TEXT, 
    unhealthy_started_at INTEGER NOT NULL,
    unhealthy_delay_in_seconds INTEGER NOT NULL,
    message TEXT NOT NULL,
    expires_at INTEGER NOT NULL,
    PRIMARY KEY (service, status_key)
);

CREATE INDEX index_status_key ON health_status (status_key);

