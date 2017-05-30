package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type PlanEstudios struct {
	Id                 int                 `orm:"column(id);pk"`
	Descripcion        string              `orm:"column(descripcion)"`
	CodigoAbreviacion  string              `orm:"column(codigo_abreviacion)"`
	FechaCreacion      time.Time           `orm:"column(fecha_creacion);type(date)"`
	FechaModificacion  time.Time           `orm:"column(fecha_modificacion);type(date)"`
	FechaAprobacion    time.Time           `orm:"column(fecha_aprobacion);type(date);null"`
	TotalDedicacion    float64             `orm:"column(total_dedicacion);null"`
	TiempoDuracion     float64             `orm:"column(tiempo_duracion);null"`
	UnidadTiempo       *UnidadTiempo       `orm:"column(unidad_tiempo);rel(fk)"`
	ProyectoCurricular *ProyectoCurricular `orm:"column(proyecto_curricular);rel(fk)"`
	TipoJornada        *TipoJornada        `orm:"column(tipo_jornada);rel(fk)"`
	UnidadDedicacion   *UnidadDedicacion   `orm:"column(unidad_dedicacion);rel(fk)"`
	PerfilTitulacion   *PerfilTitulacion   `orm:"column(perfil_titulacion);rel(fk)"`
	ModalidadPrograma  *ModalidadPrograma  `orm:"column(modalidad_programa);rel(fk)"`
}

func (t *PlanEstudios) TableName() string {
	return "plan_estudios"
}

func init() {
	orm.RegisterModel(new(PlanEstudios))
}

// AddPlanEstudios insert a new PlanEstudios into database and returns
// last inserted Id on success.
func AddPlanEstudios(m *PlanEstudios) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPlanEstudiosById retrieves PlanEstudios by Id. Returns error if
// Id doesn't exist
func GetPlanEstudiosById(id int) (v *PlanEstudios, err error) {
	o := orm.NewOrm()
	v = &PlanEstudios{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPlanEstudios retrieves all PlanEstudios matches certain condition. Returns empty list if
// no records exist
func GetAllPlanEstudios(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PlanEstudios))
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

	var l []PlanEstudios
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

// UpdatePlanEstudios updates PlanEstudios by Id and returns error if
// the record to be updated doesn't exist
func UpdatePlanEstudiosById(m *PlanEstudios) (err error) {
	o := orm.NewOrm()
	v := PlanEstudios{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePlanEstudios deletes PlanEstudios by Id and returns error if
// the record to be deleted doesn't exist
func DeletePlanEstudios(id int) (err error) {
	o := orm.NewOrm()
	v := PlanEstudios{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PlanEstudios{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
