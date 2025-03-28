package infraestructure

import (
	"examc1/core"
	"examc1/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()

	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) CreatePerson(person entities.Person) error {

	query := "INSERT INTO person (name, sex, gender) VALUES (?,?,?)"

	_, err := mysql.conn.ExecutePreparedQuery(query, person.Name, person.Sex, person.Gender)

	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err.Error())
	}

	return nil

}

func (mysql *MySQL) GetAllPersons() ([]entities.Person, error) {

	var persons []entities.Person
	var person entities.Person

	query := "SELECT * FROM person"

	rows := mysql.conn.FetchRows(query)

	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&person.Id, &person.Name, &person.Sex, &person.Gender); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %v", err.Error())
		}

		persons = append(persons, person)

	}

	return persons, nil

}

func (mysql *MySQL) CountGender() ([]int, error) {

	var countH, countF int
	var counts []int

	query1 := "SELECT COUNT(id_person) FROM person WHERE gender = 'Masculino'"
	query2 := "SELECT COUNT(id_person) FROM person WHERE gender = 'Femenino'"

	// Query para la parte del conteo de la cantidad de hombres
	rows1 := mysql.conn.FetchRows(query1)

	defer rows1.Close()

	for rows1.Next() {

		if err := rows1.Scan(&countH); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %v", err.Error())
		}

		counts = append(counts, countH)

	}

	// Query para la parte del conteo de la cantidad de mujeres
	rows2 := mysql.conn.FetchRows(query2)

	defer rows2.Close()

	for rows2.Next() {

		if err := rows2.Scan(&countF); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %v", err.Error())
		}

		counts = append(counts, countF)

	}

	return counts, nil

}
