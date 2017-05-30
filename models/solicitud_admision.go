package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SolicitudAdmision struct {
	Id                       int                     `orm:"column(id);pk"`
	FechaCreacion            time.Time               `orm:"column(fecha_creacion);type(date)"`
	FechaModificacion        time.Time               `orm:"column(fecha_modificacion);type(date)"`
	RequisitosCumplidos      bool                    `orm:"column(requisitos_cumplidos);null"`
	PromedioAcademico        float64                 `orm:"column(promedio_academico);null"`
	PercentilSecundaria      float64                 `orm:"column(percentil_secundaria);null"`
	PosicionSecundaria       float64                 `orm:"column(posicion_secundaria);null"`
	EstudiantesSecundaria    float64                 `orm:"column(estudiantes_secundaria);null"`
	TipoAdmision             *TipoAdmision           `orm:"column(tipo_admision);rel(fk)"`
	OrganizacionPersonaRol   *OrganizacionPersonaRol `orm:"column(organizacion_persona_rol);rel(fk)"`
	DominioPromedioAcademico *NivelEducativo         `orm:"column(dominio_promedio_academico);rel(fk)"`
}

func (t *SolicitudAdmision) TableName() string {
	return "solicitud_admision"
}

func init() {
	orm.RegisterModel(new(SolicitudAdmision))
}

// AddSolicitudAdmision insert a new SolicitudAdmision into database and returns
// last inserted Id on success.
func AddSolicitudAdmision(m *SolicitudAdmision) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSolicitudAdmisionById retrieves SolicitudAdmision by Id. Returns error if
// Id doesn't exist
func GetSolicitudAdmisionById(id int) (v *SolicitudAdmision, err error) {
	o := orm.NewOrm()
	v = &SolicitudAdmision{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSolicitudAdmision retrieves all SolicitudAdmision matches certain condition. Returns empty list if
// no records exist
func GetAllSolicitudAdmision(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SolicitudAdmision))
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

	var l []SolicitudAdmision
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

// UpdateSolicitudAdmision updates SolicitudAdmision by Id and returns error if
// the record to be updated doesn't exist
func UpdateSolicitudAdmisionById(m *SolicitudAdmision) (err error) {
	o := orm.NewOrm()
	v := SolicitudAdmision{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSolicitudAdmision deletes SolicitudAdmision by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSolicitudAdmision(id int) (err error) {
	o := orm.NewOrm()
	v := SolicitudAdmision{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SolicitudAdmision{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
