package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type InscripcionGrupos struct {
	Id                     int                   `orm:"column(id);pk"`
	FechaEspacioInscrito   time.Time             `orm:"column(fecha_espacio_inscrito);type(date)"`
	NotaDefinitiva         float64               `orm:"column(nota_definitiva);null"`
	GrupoAcademico         *GrupoAcademico       `orm:"column(grupo_academico);rel(fk)"`
	EstudiantePlanEstudios *EstudianteProyecto   `orm:"column(estudiante_plan_estudios);rel(fk)"`
	SesionOrganizacion     *SesionOrganizacion   `orm:"column(sesion_organizacion);rel(fk)"`
	TipoObservacionGrupo   *TipoObservacionGrupo `orm:"column(tipo_observacion_grupo);rel(fk)"`
}

func (t *InscripcionGrupos) TableName() string {
	return "inscripcion_grupos"
}

func init() {
	orm.RegisterModel(new(InscripcionGrupos))
}

// AddInscripcionGrupos insert a new InscripcionGrupos into database and returns
// last inserted Id on success.
func AddInscripcionGrupos(m *InscripcionGrupos) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInscripcionGruposById retrieves InscripcionGrupos by Id. Returns error if
// Id doesn't exist
func GetInscripcionGruposById(id int) (v *InscripcionGrupos, err error) {
	o := orm.NewOrm()
	v = &InscripcionGrupos{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInscripcionGrupos retrieves all InscripcionGrupos matches certain condition. Returns empty list if
// no records exist
func GetAllInscripcionGrupos(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(InscripcionGrupos))
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

	var l []InscripcionGrupos
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

// UpdateInscripcionGrupos updates InscripcionGrupos by Id and returns error if
// the record to be updated doesn't exist
func UpdateInscripcionGruposById(m *InscripcionGrupos) (err error) {
	o := orm.NewOrm()
	v := InscripcionGrupos{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInscripcionGrupos deletes InscripcionGrupos by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInscripcionGrupos(id int) (err error) {
	o := orm.NewOrm()
	v := InscripcionGrupos{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&InscripcionGrupos{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
