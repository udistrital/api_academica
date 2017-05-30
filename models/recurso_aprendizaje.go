package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type RecursoAprendizaje struct {
	Id                      int                     `orm:"column(id);pk"`
	Titulo                  string                  `orm:"column(titulo)"`
	Descripcion             string                  `orm:"column(descripcion);null"`
	Url                     string                  `orm:"column(url);null"`
	Importancia             string                  `orm:"column(importancia);null"`
	Tema                    string                  `orm:"column(tema);null"`
	CodigoTema              string                  `orm:"column(codigo_tema);null"`
	SistemaAsignacionTema   string                  `orm:"column(sistema_asignacion_tema);null"`
	FechaCreacion           time.Time               `orm:"column(fecha_creacion);type(date)"`
	FechaModificacion       time.Time               `orm:"column(fecha_modificacion);type(date)"`
	FechaPublicacion        time.Time               `orm:"column(fecha_publicacion);type(date);null"`
	Version                 string                  `orm:"column(version);null"`
	AutorNombre             string                  `orm:"column(autor_nombre);null"`
	AutorEmail              string                  `orm:"column(autor_email);null"`
	AutorUrl                string                  `orm:"column(autor_url);null"`
	EditorNombre            string                  `orm:"column(editor_nombre);null"`
	EditorEmail             string                  `orm:"column(editor_email);null"`
	EditorUrl               string                  `orm:"column(editor_url);null"`
	TitularDerechosAutor    string                  `orm:"column(titular_derechos_autor);null"`
	AnoDerechosAutor        string                  `orm:"column(ano_derechos_autor);null"`
	UrlDerechosAutor        string                  `orm:"column(url_derechos_autor);null"`
	TiempoRecurso           string                  `orm:"column(tiempo_recurso);null"`
	EdadMinima              float64                 `orm:"column(edad_minima);null"`
	EdadMaxima              float64                 `orm:"column(edad_maxima);null"`
	ValorComplejidadTexto   string                  `orm:"column(valor_complejidad_texto);null"`
	SistemaComplejidadTexto string                  `orm:"column(sistema_complejidad_texto);null"`
	TipoRecursoAprendizaje  *TipoRecursoAprendizaje `orm:"column(tipo_recurso_aprendizaje);rel(fk)"`
	UsuarioFinal            *UsuarioFinal           `orm:"column(usuario_final);rel(fk)"`
	Idioma                  *Idioma                 `orm:"column(idioma);rel(fk)"`
	PropositoRecurso        *PropositoRecurso       `orm:"column(proposito_recurso);rel(fk)"`
	TipoInteractividad      *TipoInteractividad     `orm:"column(tipo_interactividad);rel(fk)"`
	TipoApiAcceso           *TipoApiAcceso          `orm:"column(tipo_api_acceso);rel(fk)"`
	TipoRiesgoRecurso       *TipoRiesgoRecurso      `orm:"column(tipo_riesgo_recurso);rel(fk)"`
	ModoAcceso              *ModoAcceso             `orm:"column(modo_acceso);rel(fk)"`
	FormatoLibro            *FormatoLibro           `orm:"column(formato_libro);rel(fk)"`
	MetodoEntradaRecurso    *MetodoEntradaRecurso   `orm:"column(metodo_entrada_recurso);rel(fk)"`
	MedioDigital            *MedioDigital           `orm:"column(medio_digital);rel(fk)"`
	CondicionDerechosAutor  *CondicionDerechosAutor `orm:"column(condicion_derechos_autor);rel(fk)"`
	TipoAutor               *TipoAutor              `orm:"column(tipo_autor);rel(fk)"`
}

func (t *RecursoAprendizaje) TableName() string {
	return "recurso_aprendizaje"
}

func init() {
	orm.RegisterModel(new(RecursoAprendizaje))
}

// AddRecursoAprendizaje insert a new RecursoAprendizaje into database and returns
// last inserted Id on success.
func AddRecursoAprendizaje(m *RecursoAprendizaje) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRecursoAprendizajeById retrieves RecursoAprendizaje by Id. Returns error if
// Id doesn't exist
func GetRecursoAprendizajeById(id int) (v *RecursoAprendizaje, err error) {
	o := orm.NewOrm()
	v = &RecursoAprendizaje{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRecursoAprendizaje retrieves all RecursoAprendizaje matches certain condition. Returns empty list if
// no records exist
func GetAllRecursoAprendizaje(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RecursoAprendizaje))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []RecursoAprendizaje
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateRecursoAprendizaje updates RecursoAprendizaje by Id and returns error if
// the record to be updated doesn't exist
func UpdateRecursoAprendizajeById(m *RecursoAprendizaje) (err error) {
	o := orm.NewOrm()
	v := RecursoAprendizaje{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRecursoAprendizaje deletes RecursoAprendizaje by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRecursoAprendizaje(id int) (err error) {
	o := orm.NewOrm()
	v := RecursoAprendizaje{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RecursoAprendizaje{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
