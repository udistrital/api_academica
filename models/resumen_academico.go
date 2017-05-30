package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ResumenAcademico struct {
	Id                     int                     `orm:"column(id);pk"`
	Promedio               float64                 `orm:"column(promedio);null"`
	CreditosOtorgados      float64                 `orm:"column(creditos_otorgados);null"`
	CantidadCursos         float64                 `orm:"column(cantidad_cursos);null"`
	FechaDiploma           string                  `orm:"column(fecha_diploma);null"`
	CategoriaCertificado   *CategoriaCertificado   `orm:"column(categoria_certificado);rel(fk)"`
	OrganizacionPersonaRol *OrganizacionPersonaRol `orm:"column(organizacion_persona_rol);rel(fk)"`
}

func (t *ResumenAcademico) TableName() string {
	return "resumen_academico"
}

func init() {
	orm.RegisterModel(new(ResumenAcademico))
}

// AddResumenAcademico insert a new ResumenAcademico into database and returns
// last inserted Id on success.
func AddResumenAcademico(m *ResumenAcademico) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetResumenAcademicoById retrieves ResumenAcademico by Id. Returns error if
// Id doesn't exist
func GetResumenAcademicoById(id int) (v *ResumenAcademico, err error) {
	o := orm.NewOrm()
	v = &ResumenAcademico{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllResumenAcademico retrieves all ResumenAcademico matches certain condition. Returns empty list if
// no records exist
func GetAllResumenAcademico(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ResumenAcademico))
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

	var l []ResumenAcademico
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

// UpdateResumenAcademico updates ResumenAcademico by Id and returns error if
// the record to be updated doesn't exist
func UpdateResumenAcademicoById(m *ResumenAcademico) (err error) {
	o := orm.NewOrm()
	v := ResumenAcademico{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteResumenAcademico deletes ResumenAcademico by Id and returns error if
// the record to be deleted doesn't exist
func DeleteResumenAcademico(id int) (err error) {
	o := orm.NewOrm()
	v := ResumenAcademico{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ResumenAcademico{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
