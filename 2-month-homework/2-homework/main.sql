CREATE TABLE cars (
    car_id INT PRIMARY KEY AUTO_INCREMENT,
    car_name VARCHAR(255) NOT NULL,
    model_year INT NOT NULL
);

CREATE TABLE users (
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);

CREATE TABLE user_car (
    car_id INT,
    user_id INT,
    PRIMARY KEY (car_id, user_id),
    FOREIGN KEY (car_id) REFERENCES cars(car_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

INSERT INTO cars (car_name, model_year) VALUES ('Toyota Camry', 2020);
INSERT INTO cars (car_name, model_year) VALUES ('Honda Accord', 2019);

INSERT INTO users (username, email) VALUES ('john_doe', 'john@example.com');
INSERT INTO users (username, email) VALUES ('jane_smith', 'jane@example.com');


INSERT INTO yasang (car_id, user_id) VALUES (1, 1); 
INSERT INTO yasang (car_id, user_id) VALUES (2, 1); 
INSERT INTO yasang (car_id, user_id) VALUES (1, 2); 


SELECT cars.car_name, cars.model_year
FROM cars
JOIN yasang ON cars.car_id = yasang.car_id
WHERE yasang.user_id = 1;
