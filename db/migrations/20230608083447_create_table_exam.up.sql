CREATE TABLE exam 
(
  id VARCHAR(10) NOT NULL,
  employee_id VARCHAR(10) NOT NULL,
  exam_result INT DEFAULT 0,
  exam_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT fk_exam_employee
    FOREIGN KEY (employee_id) REFERENCES employee (id) ON DELETE CASCADE
) ENGINE = InnoDB;