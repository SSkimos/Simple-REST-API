CREATE OR REPLACE FUNCTION increase_salaries()
RETURNS VOID AS $$
DECLARE
    emp_record RECORD;
BEGIN
    FOR emp_record IN SELECT id, salary FROM employees
    LOOP
        UPDATE employees
        SET salary = emp_record.salary * 1.1
        WHERE id = emp_record.id;
    END LOOP;
END;
$$ LANGUAGE plpgsql;
