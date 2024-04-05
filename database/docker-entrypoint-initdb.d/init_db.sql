-- auto-generated definition
CREATE TABLE Employee (
    EmployeeId INT PRIMARY KEY,
    FirstName VARCHAR(50),
    LastName VARCHAR(50),
    Position VARCHAR(50),
    Department VARCHAR(50),
    HireDate DATE,
    Salary INT,
    Email VARCHAR(100),
    PhoneNumber VARCHAR(20),
    Address VARCHAR(255)
);

INSERT INTO Employee (EmployeeId, FirstName, LastName, Position, Department, HireDate, Salary, Email, PhoneNumber, Address)
VALUES (1, 'John', 'Doe', 'Manager', 'HR', '2023-01-15', 60000, 'john.doe@example.com', '1234567890', '123 Main St');

INSERT INTO Employee (EmployeeId, FirstName, LastName, Position, Department, HireDate, Salary, Email, PhoneNumber, Address)
VALUES (2, 'Alice', 'Smith', 'Software Engineer', 'Engineering', '2022-11-20', 70000, 'alice.smith@example.com', '2345678901', '456 Elm St');

INSERT INTO Employee (EmployeeId, FirstName, LastName, Position, Department, HireDate, Salary, Email, PhoneNumber, Address)
VALUES (3, 'Bob', 'Johnson', 'Sales Representative', 'Sales', '2023-03-10', 50000, 'bob.johnson@example.com', '3456789012', '789 Oak St');

INSERT INTO Employee (EmployeeId, FirstName, LastName, Position, Department, HireDate, Salary, Email, PhoneNumber, Address)
VALUES (4, 'Emma', 'Williams', 'Marketing Specialist', 'Marketing', '2022-09-05', 55000, 'emma.williams@example.com', '4567890123', '1011 Pine St');

INSERT INTO Employee (EmployeeId, FirstName, LastName, Position, Department, HireDate, Salary, Email, PhoneNumber, Address)
VALUES (5, 'Michael', 'Brown', 'Financial Analyst', 'Finance', '2023-02-28', 65000, 'michael.brown@example.com', '5678901234', '1213 Cedar St');

INSERT INTO Employee (EmployeeId, FirstName, LastName, Position, Department, HireDate, Salary, Email, PhoneNumber, Address)
VALUES (6, 'Sophia', 'Taylor', 'Human Resources Coordinator', 'HR', '2022-12-12', 48000, 'sophia.taylor@example.com', '6789012345', '1415 Maple St');

INSERT INTO Employee (EmployeeId, FirstName, LastName, Position, Department, HireDate, Salary, Email, PhoneNumber, Address)
VALUES (7, 'Matthew', 'Martinez', 'Quality Assurance Specialist', 'Engineering', '2023-05-20', 60000, 'matthew.martinez@example.com', '7890123456', '1617 Walnut St');

INSERT INTO Employee (EmployeeId, FirstName, LastName, Position, Department, HireDate, Salary, Email, PhoneNumber, Address)
VALUES (8, 'Olivia', 'Garcia', 'Customer Service Representative', 'Customer Service', '2022-08-03', 45000, 'olivia.garcia@example.com', '8901234567', '1819 Cherry St');

INSERT INTO Employee (EmployeeId, FirstName, LastName, Position, Department, HireDate, Salary, Email, PhoneNumber, Address)
VALUES (9, 'Ethan', 'Martinez', 'Systems Administrator', 'IT', '2023-04-18', 70000, 'ethan.martinez@example.com', '9012345678', '2021 Oak St');

INSERT INTO Employee (EmployeeId, FirstName, LastName, Position, Department, HireDate, Salary, Email, PhoneNumber, Address)
VALUES (10, 'Ava', 'Hernandez', 'Graphic Designer', 'Marketing', '2022-10-30', 52000, 'ava.hernandez@example.com', '0123456789', '2223 Pine St');

