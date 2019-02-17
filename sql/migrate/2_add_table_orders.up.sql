CREATE TABLE IF NOT EXISTS orders (
    id INT AUTO_INCREMENT,
    start_lat DECIMAL(10, 8) NOT NULL,
    start_lon DECIMAL(11, 8) NOT NULL,
    end_lat DECIMAL(10, 8) NOT NULL,
    end_lon DECIMAL(11, 8) NOT NULL,
    distance INT NOT NULL,
    status VARCHAR(32) NOT NULL DEFAULT 'UNASSIGNED',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY status (status),
    KEY created_at (created_at),
    KEY updated_at (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=ascii;