CREATE TABLE employee 
(
  id VARCHAR(10) NOT NULL,
  name VARCHAR(50) NOT NULL,
  ktp VARCHAR(20) NOT NULL,
  status BOOLEAN DEFAULT 0,
  PRIMARY KEY (id)
) ENGINE = InnoDB;