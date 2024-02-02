package resolver

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/graphql-go/graphql"
	command "github.con/AndyGo/go-ddd/internal/application/command/ALLNote/create"
	query "github.con/AndyGo/go-ddd/internal/application/command/ALLNote/query"
	"github.con/AndyGo/go-ddd/internal/application/services"
	"github.con/AndyGo/go-ddd/internal/domain/entities"
	mssql "github.con/AndyGo/go-ddd/internal/infrastructure/db/mssql/respoitories"
)

func GetAllNotesGQL(p graphql.ResolveParams) (interface{}, error) {
	allnoteRepo := mssql.NewALLNoteRepository()
	allnoteService := services.NewALLNoteService(allnoteRepo)

	datasourceRequest := p.Args["dataSourceRequest"].(string)
	var dsRequest entities.DataSourceRequest
	if err := json.Unmarshal([]byte(datasourceRequest), &dsRequest); err != nil {
		return nil, err
	}

	data, err := allnoteService.GetAllALLNote(query.ListALLNoteCommand{
		DatasourceRequest: dsRequest,
	})

	if err != nil {
		return nil, err
	}

	allNoteList := make([]entities.ALLNote, len(data.Data))
	for i, validatedNote := range data.Data {
		allNoteList[i] = validatedNote.ALLNote
	}
	return map[string]interface{}{
		"Data":  allNoteList,
		"Total": data.Total,
	}, nil

}

// func CreateAllNotesGQL(p graphql.ResolveParams) (interface{}, error) {
// 	allnoteRepo := mssql.NewALLNoteRepository()
// 	allnoteService := services.NewALLNoteService(allnoteRepo)

// 	text_note := p.Args["note"].(string)

// 	data, err := allnoteService.Create(command.CreateALLNoteCommand{
// 		ALLNote: entities.ALLNote{
// 			Note: text_note,
// 		},
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return map[string]interface{}{
// 		"Uid": data.Result,
// 	}, nil

// }
func CreateAllNotesGQL(p graphql.ResolveParams) (interface{}, error) {
	allnoteRepo := mssql.NewALLNoteRepository()
	allnoteService := services.NewALLNoteService(allnoteRepo)

	// Obtener el input del argumento
	allNoteInput, ok := p.Args["note"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Error al obtener el input de la nota")
	}

	// Crear un nuevo objeto ALLNote con los datos proporcionados
	allNote := entities.ALLNote{}

	for key, value := range allNoteInput {
		field := reflect.ValueOf(&allNote).Elem().FieldByName(key)
		if field.IsValid() && field.CanSet() {
			// Verificar si el campo es de tipo *time.Time
			if field.Type().Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Struct && field.Type().Elem().Name() == "Time" {
				// Convertir el valor a time.Time
				if stringValue, ok := value.(string); ok {
					if stringValue == "" {
						// Tratar un valor nulo para *time.Time
						field.Set(reflect.Zero(field.Type()))
					} else {
						parsedTime, err := time.Parse(time.RFC3339, stringValue)
						if err != nil {
							return nil, err
						}
						field.Set(reflect.ValueOf(&parsedTime))
					}
				} else {
					return nil, fmt.Errorf("El valor para el campo %s no es de tipo string", key)
				}
			} else if field.Type().Kind() == reflect.Ptr {
				// Si el campo es un puntero, asignar el valor como un puntero
				if value == nil {
					// Tratar un valor nulo para punteros
					field.Set(reflect.Zero(field.Type()))
				} else {
					switch field.Type().Elem().Kind() {
					case reflect.String:
						stringValue, ok := value.(string)
						if !ok {
							return nil, fmt.Errorf("El valor para el campo %s no es de tipo string", key)
						}
						field.Set(reflect.ValueOf(&stringValue))
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						// Aquí puedes agregar más casos según sea necesario
						intValue, ok := value.(int)
						if !ok {
							return nil, fmt.Errorf("El valor para el campo %s no es de tipo int", key)
						}
						field.Set(reflect.ValueOf(&intValue))
					case reflect.Bool:
						boolValue, ok := value.(bool)
						if !ok {
							return nil, fmt.Errorf("El valor para el campo %s no es de tipo bool", key)
						}
						field.Set(reflect.ValueOf(&boolValue))
					default:
						return nil, fmt.Errorf("No se pudo manejar el tipo de campo %s", key)
					}
				}
			} else if field.Type().Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.Uint8 {
				// Si el campo es de tipo []uint8, convertir el valor a []byte
				stringValue, ok := value.(string)
				if !ok {
					return nil, fmt.Errorf("El valor para el campo %s no es de tipo string", key)
				}
				byteValue := []byte(stringValue)
				field.Set(reflect.ValueOf(byteValue))
			} else {
				// Para otros tipos, asignar directamente
				field.Set(reflect.ValueOf(value))
			}
		}
	}

	data, err := allnoteService.Create(command.CreateALLNoteCommand{
		ALLNote: allNote,
		// Otros campos según sea necesario
	})

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"Uid": data.Result,
	}, nil
}
