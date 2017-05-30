package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ProyectoCurricular struct {
	Id                       int                       `orm:"column(id);pk"`
	Descripcion              string                    `orm:"column(descripcion)"`
	CodigoAbreviacion        string                    `orm:"column(codigo_abreviacion)"`
	FechaCreacion            time.Time                 `orm:"column(fecha_creacion);type(date)"`
	FechaModificacion        time.Time                 `orm:"column(fecha_modificacion);type(date)"`
	Organizacion             *Organizacion             `orm:"column(organizacion);rel(fk)"`
	NucleoBasicoConocimiento *NucleoBasicoConocimiento `orm:"column(nucleo_basico_conocimiento);rel(fk)"`
	NivelPrograma            *NivelPrograma            `orm:"column(nivel_programa);rel(fk)"`
	UnidadDuracion           *UnidadTiempo             `orm:"column(unidad_duracion);rel(fk)"`
}

func (t *ProyectoCurricular) TableName() string {
	return "proyecto_curricular"
}

func init() {
	orm.RegisterModel(new(ProyectoCurricular))
}

// AddProyectoCurricular insert a new ProyectoCurricular into database and returns
// last inserted Id on success.
func AddProyectoCurricular(m *ProyectoCurricular) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProyectoCurricularById retrieves ProyectoCurricular by Id. Returns error if
// Id doesn't exist
func GetProyectoCurricularById(id int) (v *ProyectoCurricular, err error) {
	o := orm.NewOrm()
	v = &ProyectoCurricular{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProyectoCurricular retrieves all ProyectoCurricular matches certain condition. Returns empty list if
// no records exist
func GetAllProyectoCurricular(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProyectoCurricular))
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

	var l []ProyectoCurricular
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

// UpdateProyectoCurricular updates ProyectoCurricular by Id and returns error if
// the record to be updated doesn't exist
func UpdateProyectoCurricularById(m *ProyectoCurricular) (err error) {
	o := orm.NewOrm()
	v := ProyectoCurricular{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProyectoCurricular deletes ProyectoCurricular by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProyectoCurricular(id int) (err error) {
	o := orm.NewOrm()
	v := ProyectoCurricular{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProyectoCurricular{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
