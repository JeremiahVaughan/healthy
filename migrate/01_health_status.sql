CREATE TABLE health_status (                                                   
    status_key TEXT PRIMARY KEY, 
    service TEXT NOT NULL,
    unhealthy INTEGER NOT NULL,
    unhealthy_at INTEGER NOT NULL,
    unhealthy_delay_in_seconds INTEGER NOT NULL,
    message TEXT NOT NULL
);

CREATE UNIQUE INDEX health_status_sort on health_status(service, status_key);
